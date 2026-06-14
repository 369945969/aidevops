package agent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RegisterBuiltInTools adds the 3 default DevOps tools to a registry.
func RegisterBuiltInTools(registry *ToolRegistry, workspacePath string) {
	registry.Register(readRepoFileDef(workspacePath), readRepoFileExec(workspacePath))
	registry.Register(writeFileDef(workspacePath), writeFileExec(workspacePath))
	registry.Register(triggerCICDDef(), triggerCICDExec())
}

// ============================================================
// Tool 1: read_repo_file — read a file from the workspace repo
// ============================================================

func readRepoFileDef(workspacePath string) ToolDefinition {
	return ToolDefinition{
		Name:        "read_repo_file",
		Description: "Read a file from the workspace code repository. Returns the file contents.",
		Parameters: map[string]ParamSpec{
			"path": {
				Type:        "string",
				Description: "Relative file path within the workspace repo (e.g. 'src/main.go')",
				Required:    true,
			},
			"max_lines": {
				Type:        "integer",
				Description: "Maximum lines to return (0 = all)",
				Required:    false,
			},
		},
	}
}

func readRepoFileExec(workspacePath string) ToolExecutor {
	return func(params map[string]interface{}) ToolResult {
		relPath, _ := params["path"].(string)
		if relPath == "" {
			return ToolResult{Success: false, Error: "path parameter is required"}
		}

		absPath := filepath.Join(workspacePath, relPath)
		absPath = filepath.Clean(absPath)
		if !strings.HasPrefix(absPath, filepath.Clean(workspacePath)) {
			return ToolResult{Success: false, Error: "path escapes workspace directory"}
		}

		data, err := os.ReadFile(absPath)
		if err != nil {
			return ToolResult{Success: false, Error: fmt.Sprintf("failed to read file: %v", err)}
		}

		content := string(data)
		maxLines := 0
		if ml, ok := params["max_lines"].(float64); ok && ml > 0 {
			maxLines = int(ml)
			lines := strings.Split(content, "\n")
			if len(lines) > maxLines {
				content = strings.Join(lines[:maxLines], "\n") + "\n... (truncated)"
			}
		}

		return ToolResult{Success: true, Output: content}
	}
}

// ============================================================
// Tool 2: write_file — write content to a file in the workspace
// ============================================================

func writeFileDef(workspacePath string) ToolDefinition {
	return ToolDefinition{
		Name:        "write_file",
		Description: "Write content to a file in the workspace repo. Creates parent directories if needed.",
		Parameters: map[string]ParamSpec{
			"path": {
				Type:        "string",
				Description: "Relative file path within the workspace repo (e.g. 'src/handler.go')",
				Required:    true,
			},
			"content": {
				Type:        "string",
				Description: "The content to write to the file",
				Required:    true,
			},
		},
	}
}

func writeFileExec(workspacePath string) ToolExecutor {
	return func(params map[string]interface{}) ToolResult {
		relPath, _ := params["path"].(string)
		content, _ := params["content"].(string)
		if relPath == "" {
			return ToolResult{Success: false, Error: "path parameter is required"}
		}
		if content == "" {
			return ToolResult{Success: false, Error: "content parameter is required"}
		}

		absPath := filepath.Join(workspacePath, relPath)
		absPath = filepath.Clean(absPath)
		if !strings.HasPrefix(absPath, filepath.Clean(workspacePath)) {
			return ToolResult{Success: false, Error: "path escapes workspace directory"}
		}

		dir := filepath.Dir(absPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return ToolResult{Success: false, Error: fmt.Sprintf("failed to create directory: %v", err)}
		}

		if err := os.WriteFile(absPath, []byte(content), 0644); err != nil {
			return ToolResult{Success: false, Error: fmt.Sprintf("failed to write file: %v", err)}
		}

		return ToolResult{Success: true, Output: fmt.Sprintf("Written %d bytes to %s", len(content), relPath)}
	}
}

// ============================================================
// Tool 3: trigger_cicd — trigger a CI/CD pipeline run
// ============================================================

func triggerCICDDef() ToolDefinition {
	return ToolDefinition{
		Name:        "trigger_cicd",
		Description: "Trigger a CI/CD pipeline run for the specified service/branch.",
		Parameters: map[string]ParamSpec{
			"service": {
				Type:        "string",
				Description: "Service name to trigger CI/CD for (e.g. 'backend-api')",
				Required:    true,
			},
			"branch": {
				Type:        "string",
				Description: "Git branch to build from (default: 'main')",
				Required:    false,
			},
			"action": {
				Type:        "string",
				Description: "CI/CD action to perform",
				Required:    true,
				Enum:        []string{"build", "test", "deploy"},
			},
		},
	}
}

func triggerCICDExec() ToolExecutor {
	return func(params map[string]interface{}) ToolResult {
		service, _ := params["service"].(string)
		action, _ := params["action"].(string)
		branch, _ := params["branch"].(string)
		if branch == "" {
			branch = "main"
		}
		if service == "" || action == "" {
			return ToolResult{Success: false, Error: "service and action are required"}
		}

		// Stub: real implementation would call Jenkins/GitLab/GitHub Actions API
		return ToolResult{
			Success: true,
			Output:  fmt.Sprintf("CI/CD %s triggered for service=%s branch=%s (stub — integrate with real CI/CD provider)", action, service, branch),
		}
	}
}
