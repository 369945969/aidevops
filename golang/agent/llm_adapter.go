package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// ============================================================
// LLMAdapter — OpenAI-compatible API caller
// ============================================================

// LLMAdapter handles LLM API calls via OpenAI-compatible endpoints.
// Mirrors LLMToolAdapter from daily_stock_analysis/src/agent/llm_adapter.py
// but simplified for the DevOps use case (no tool calling yet).
type LLMAdapter struct {
	APIKey      string
	BaseURL     string
	Model       string
	MaxTokens   int
	Temperature float64
	Timeout     time.Duration
}

// LLMConfig holds configuration for building an LLMAdapter.
type LLMConfig struct {
	APIKey      string
	BaseURL     string
	Model       string
	MaxTokens   int
	Temperature float64
	TimeoutMs   int
}

// NewLLMAdapterFromConfig creates an adapter from explicit config.
func NewLLMAdapterFromConfig(cfg LLMConfig) *LLMAdapter {
	timeout := 120 * time.Second
	if cfg.TimeoutMs > 0 {
		timeout = time.Duration(cfg.TimeoutMs) * time.Millisecond
	}
	maxTokens := cfg.MaxTokens
	if maxTokens <= 0 {
		maxTokens = 4096
	}
	temperature := cfg.Temperature
	if temperature <= 0 {
		temperature = 0.7
	}
	return &LLMAdapter{
		APIKey:      cfg.APIKey,
		BaseURL:     cfg.BaseURL,
		Model:       cfg.Model,
		MaxTokens:   maxTokens,
		Temperature: temperature,
		Timeout:     timeout,
	}
}

// NewLLMAdapterFromEnv creates an adapter from environment variables.
// Env vars: LLM_API_KEY, LLM_BASE_URL, LLM_MODEL, LLM_MAX_TOKENS, LLM_TIMEOUT_MS
func NewLLMAdapterFromEnv() *LLMAdapter {
	cfg := LLMConfig{
		APIKey:      os.Getenv("LLM_API_KEY"),
		BaseURL:     os.Getenv("LLM_BASE_URL"),
		Model:       os.Getenv("LLM_MODEL"),
		MaxTokens:   0, // will use default
		Temperature: 0,
		TimeoutMs:   0,
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.openai.com/v1"
	}
	if cfg.Model == "" {
		cfg.Model = "gpt-4o"
	}
	return NewLLMAdapterFromConfig(cfg)
}

// ChatRequest mirrors the OpenAI chat completion request format.
type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	Temperature float64       `json:"temperature,omitempty"`
}

// ChatMessage is a single message in the conversation.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatResponse mirrors the OpenAI chat completion response format.
type ChatResponse struct {
	ID      string         `json:"id"`
	Choices []ChatChoice   `json:"choices"`
	Usage   ChatUsage      `json:"usage"`
	Model   string         `json:"model"`
}

// ChatChoice is a single choice in the response.
type ChatChoice struct {
	Index   int         `json:"index"`
	Message ChatMessage `json:"message"`
	Finish  string      `json:"finish_reason"`
}

// ChatUsage tracks token usage.
type ChatUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Call sends a chat completion request and returns the response content.
// This is the primary method agents use to invoke the LLM.
func (a *LLMAdapter) Call(systemPrompt, userMessage string) (string, int, error) {
	messages := []ChatMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: userMessage},
	}
	return a.CallWithHistory(messages)
}

// CallWithHistory sends a full conversation history and returns the response.
func (a *LLMAdapter) CallWithHistory(messages []ChatMessage) (string, int, error) {
	reqBody := ChatRequest{
		Model:       a.Model,
		Messages:    messages,
		MaxTokens:   a.MaxTokens,
		Temperature: a.Temperature,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", 0, fmt.Errorf("marshal request: %w", err)
	}

	client := &http.Client{Timeout: a.Timeout}
	req, err := http.NewRequest("POST", a.BaseURL+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", 0, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.APIKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("LLM API call failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("[LLMAdapter] API error: status=%d body=%s", resp.StatusCode, string(body))
		return "", 0, fmt.Errorf("LLM API returned status %d: %s", resp.StatusCode, string(body))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", 0, fmt.Errorf("unmarshal response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", 0, fmt.Errorf("LLM returned no choices")
	}

	content := chatResp.Choices[0].Message.Content
	tokens := chatResp.Usage.TotalTokens

	return content, tokens, nil
}

// CallForJSON sends a chat request and parses the response as JSON.
// Agents that produce structured output use this method.
func (a *LLMAdapter) CallForJSON(systemPrompt, userMessage string) (map[string]interface{}, int, error) {
	content, tokens, err := a.Call(systemPrompt, userMessage)
	if err != nil {
		return nil, tokens, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(content), &result); err != nil {
		// Try to extract JSON from markdown code blocks
		extracted := extractJSONFromMarkdown(content)
		preview := content
		if len(preview) > 200 {
			preview = preview[:200]
		}
		if extracted != nil {
			if err2 := json.Unmarshal(extracted, &result); err2 != nil {
				return nil, tokens, fmt.Errorf("response is not valid JSON: %w (raw: %s)", err, preview)
			}
		} else {
			return nil, tokens, fmt.Errorf("response is not valid JSON: %w (raw: %s)", err, preview)
		}
	}

	return result, tokens, nil
}

// extractJSONFromMarkdown tries to extract a JSON object from markdown code blocks.
func extractJSONFromMarkdown(content string) []byte {
	// Look for ```json ... ``` blocks
	start := -1
	end := -1
	jsonTag := "```json"
	plainTag := "```"

	idx := indexOf(content, jsonTag)
	if idx >= 0 {
		start = idx + len(jsonTag)
		endIdx := indexOf(content[start:], plainTag)
		if endIdx >= 0 {
			end = start + endIdx
		}
	} else {
		// Try plain ``` blocks
		idx = indexOf(content, plainTag)
		if idx >= 0 {
			start = idx + len(plainTag)
			endIdx := indexOf(content[start:], plainTag)
			if endIdx >= 0 {
				end = start + endIdx
			}
		}
	}

	if start >= 0 && end > start {
		return []byte(content[start:end])
	}

	// Try to find raw JSON object { ... }
	objStart := indexOf(content, "{")
	objEnd := lastIndexOf(content, "}")
	if objStart >= 0 && objEnd > objStart {
		return []byte(content[objStart : objEnd+1])
	}

	return nil
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func lastIndexOf(s, substr string) int {
	for i := len(s) - len(substr); i >= 0; i-- {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
