package agent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNewLLMAdapterFromConfig(t *testing.T) {
	cfg := LLMConfig{
		APIKey:      "test-key",
		BaseURL:     "https://api.example.com/v1",
		Model:       "test-model",
		MaxTokens:   2048,
		Temperature: 0.5,
		TimeoutMs:   60000,
	}
	a := NewLLMAdapterFromConfig(cfg)
	if a.APIKey != "test-key" {
		t.Errorf("APIKey = %q, want test-key", a.APIKey)
	}
	if a.Model != "test-model" {
		t.Errorf("Model = %q, want test-model", a.Model)
	}
	if a.MaxTokens != 2048 {
		t.Errorf("MaxTokens = %d, want 2048", a.MaxTokens)
	}
	if a.Temperature != 0.5 {
		t.Errorf("Temperature = %v, want 0.5", a.Temperature)
	}
	if a.Timeout != 60*time.Second {
		t.Errorf("Timeout = %v, want 60s", a.Timeout)
	}
}

func TestNewLLMAdapterFromConfigDefaults(t *testing.T) {
	cfg := LLMConfig{APIKey: "key", BaseURL: "url", Model: "model"}
	a := NewLLMAdapterFromConfig(cfg)
	if a.MaxTokens != 4096 {
		t.Errorf("default MaxTokens = %d, want 4096", a.MaxTokens)
	}
	if a.Temperature != 0.7 {
		t.Errorf("default Temperature = %v, want 0.7", a.Temperature)
	}
	if a.Timeout != 120*time.Second {
		t.Errorf("default Timeout = %v, want 120s", a.Timeout)
	}
}

func TestNewLLMAdapterFromEnv(t *testing.T) {
	t.Setenv("LLM_API_KEY", "env-key")
	t.Setenv("LLM_BASE_URL", "https://env.example.com/v1")
	t.Setenv("LLM_MODEL", "env-model")

	a := NewLLMAdapterFromEnv()
	if a.APIKey != "env-key" {
		t.Errorf("APIKey = %q, want env-key", a.APIKey)
	}
	if a.BaseURL != "https://env.example.com/v1" {
		t.Errorf("BaseURL = %q, want env URL", a.BaseURL)
	}
	if a.Model != "env-model" {
		t.Errorf("Model = %q, want env-model", a.Model)
	}
}

func TestNewLLMAdapterFromEnvDefaults(t *testing.T) {
	os.Clearenv()
	t.Setenv("LLM_API_KEY", "key")

	a := NewLLMAdapterFromEnv()
	if a.BaseURL != "https://api.openai.com/v1" {
		t.Errorf("default BaseURL = %q", a.BaseURL)
	}
	if a.Model != "gpt-4o" {
		t.Errorf("default Model = %q", a.Model)
	}
}

func setupMockServer(handler http.HandlerFunc) (*httptest.Server, *LLMAdapter) {
	server := httptest.NewServer(handler)
	adapter := &LLMAdapter{
		APIKey:      "test-key",
		BaseURL:     server.URL,
		Model:       "test-model",
		MaxTokens:   1024,
		Temperature: 0.7,
		Timeout:     10 * time.Second,
	}
	return server, adapter
}

func TestCallSuccess(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("Authorization header = %q, want Bearer test-key", r.Header.Get("Authorization"))
		}
		resp := ChatResponse{
			Choices: []ChatChoice{{Message: ChatMessage{Content: "Hello world"}}},
			Usage:   ChatUsage{TotalTokens: 50},
		}
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	content, tokens, err := adapter.Call("system prompt", "user message")
	if err != nil {
		t.Fatalf("Call failed: %v", err)
	}
	if content != "Hello world" {
		t.Errorf("content = %q, want Hello world", content)
	}
	if tokens != 50 {
		t.Errorf("tokens = %d, want 50", tokens)
	}
}

func TestCallWithHistory(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		var req ChatRequest
		json.NewDecoder(r.Body).Decode(&req)
		if len(req.Messages) != 2 {
			t.Errorf("Messages length = %d, want 2", len(req.Messages))
		}
		resp := ChatResponse{
			Choices: []ChatChoice{{Message: ChatMessage{Content: "response"}}},
			Usage:   ChatUsage{TotalTokens: 30},
		}
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	messages := []ChatMessage{
		{Role: "system", Content: "sys"},
		{Role: "user", Content: "usr"},
	}
	content, _, err := adapter.CallWithHistory(messages)
	if err != nil {
		t.Fatalf("CallWithHistory failed: %v", err)
	}
	if content != "response" {
		t.Errorf("content = %q, want response", content)
	}
}

func TestCallErrorStatus(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("rate limited"))
	})
	defer server.Close()

	_, _, err := adapter.Call("sys", "usr")
	if err == nil {
		t.Error("expected error for 429 status")
	}
}

func TestCallEmptyChoices(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		resp := ChatResponse{Choices: []ChatChoice{}}
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	_, _, err := adapter.Call("sys", "usr")
	if err == nil {
		t.Error("expected error for empty choices")
	}
}

func TestCallForJSONSuccess(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		resp := ChatResponse{
			Choices: []ChatChoice{{Message: ChatMessage{Content: `{"title": "test", "priority": "high"}`}}},
			Usage:   ChatUsage{TotalTokens: 40},
		}
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	result, tokens, err := adapter.CallForJSON("sys", "usr")
	if err != nil {
		t.Fatalf("CallForJSON failed: %v", err)
	}
	if result["title"] != "test" {
		t.Errorf("title = %v, want test", result["title"])
	}
	if tokens != 40 {
		t.Errorf("tokens = %d, want 40", tokens)
	}
}

func TestCallForJSONMarkdownWrapped(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		content := "Here is the JSON:\n```json\n{\"title\": \"wrapped\"}\n```\nDone."
		resp := ChatResponse{
			Choices: []ChatChoice{{Message: ChatMessage{Content: content}}},
			Usage:   ChatUsage{TotalTokens: 60},
		}
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	result, _, err := adapter.CallForJSON("sys", "usr")
	if err != nil {
		t.Fatalf("CallForJSON with markdown failed: %v", err)
	}
	if result["title"] != "wrapped" {
		t.Errorf("title = %v, want wrapped", result["title"])
	}
}

func TestCallForJSONInvalidResponse(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		resp := ChatResponse{
			Choices: []ChatChoice{{Message: ChatMessage{Content: "This is not JSON at all"}}},
			Usage:   ChatUsage{TotalTokens: 20},
		}
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	_, _, err := adapter.CallForJSON("sys", "usr")
	if err == nil {
		t.Error("expected error for non-JSON response")
	}
}

func TestCallForJSONRawCurlyBraces(t *testing.T) {
	server, adapter := setupMockServer(func(w http.ResponseWriter, r *http.Request) {
		content := "Some text before {\"key\": \"value\"} some text after"
		resp := ChatResponse{
			Choices: []ChatChoice{{Message: ChatMessage{Content: content}}},
			Usage:   ChatUsage{TotalTokens: 30},
		}
		json.NewEncoder(w).Encode(resp)
	})
	defer server.Close()

	result, _, err := adapter.CallForJSON("sys", "usr")
	if err != nil {
		t.Fatalf("CallForJSON with raw braces failed: %v", err)
	}
	if result["key"] != "value" {
		t.Errorf("key = %v, want value", result["key"])
	}
}

func TestExtractJSONFromMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantNil bool
		wantKey string
		wantVal string
	}{
		{"json_block", "```json\n{\"a\":1}\n```", false, "a", ""},
		{"plain_block", "```\n{\"b\":2}\n```", false, "b", ""},
		{"raw_curly", "text {\"c\":3} text", false, "c", ""},
		{"no_json", "just plain text no json here", true, "", ""},
		{"nested_json", "```json\n{\"outer\":{\"inner\":4}}\n```", false, "outer", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractJSONFromMarkdown(tt.input)
			if tt.wantNil {
				if result != nil {
					t.Errorf("expected nil, got %s", string(result))
				}
				return
			}
			if result == nil {
				t.Fatalf("expected non-nil JSON extraction")
			}
			var m map[string]interface{}
			if err := json.Unmarshal(result, &m); err != nil {
				t.Fatalf("extracted bytes not valid JSON: %v", err)
			}
			if tt.wantKey != "" {
				if m[tt.wantKey] == nil {
					t.Errorf("m[%q] = nil, want non-nil", tt.wantKey)
				}
				if tt.wantVal != "" && fmt.Sprintf("%v", m[tt.wantKey]) != tt.wantVal {
					t.Errorf("m[%q] = %v, want %v", tt.wantKey, m[tt.wantKey], tt.wantVal)
				}
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	if indexOf("hello world", "world") != 6 {
		t.Error("indexOf should find 'world' at position 6")
	}
	if indexOf("hello", "xyz") != -1 {
		t.Error("indexOf should return -1 for missing substring")
	}
	if indexOf("aaa", "a") != 0 {
		t.Error("indexOf should find first occurrence at 0")
	}
}

func TestLastIndexOf(t *testing.T) {
	if lastIndexOf("hello world hello", "hello") != 12 {
		t.Error("lastIndexOf should find last occurrence")
	}
	if lastIndexOf("hello", "xyz") != -1 {
		t.Error("lastIndexOf should return -1 for missing substring")
	}
}
