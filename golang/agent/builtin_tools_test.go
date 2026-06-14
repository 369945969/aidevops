package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRegisterBuiltInTools(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	defs := r.AllDefinitions()
	if len(defs) != 3 {
		t.Fatalf("expected 3 built-in tools, got %d", len(defs))
	}
	names := map[string]bool{}
	for _, d := range defs {
		names[d.Name] = true
	}
	for _, name := range []string{"read_repo_file", "write_file", "trigger_cicd"} {
		if !names[name] {
			t.Errorf("missing built-in tool: %s", name)
		}
	}
}

func TestReadRepoFileSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "src", "main.go")
	os.MkdirAll(filepath.Dir(testFile), 0755)
	os.WriteFile(testFile, []byte("package main\nfunc main() {}\n"), 0644)

	r := NewToolRegistry()
	RegisterBuiltInTools(r, tmpDir)

	result := r.Execute("read_repo_file", `{"path":"src/main.go"}`)
	if !result.Success {
		t.Fatalf("read_repo_file failed: %s", result.Error)
	}
	if result.Output != "package main\nfunc main() {}\n" {
		t.Errorf("Output mismatch: %q", result.Output)
	}
}

func TestReadRepoFileMissingParam(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	result := r.Execute("read_repo_file", `{"path":""}`)
	if result.Success {
		t.Error("empty path should fail")
	}

	result2 := r.Execute("read_repo_file", `{}`)
	if result2.Success {
		t.Error("missing path param should fail")
	}
}

func TestReadRepoFilePathTraversal(t *testing.T) {
	tmpDir := t.TempDir()
	r := NewToolRegistry()
	RegisterBuiltInTools(r, tmpDir)

	result := r.Execute("read_repo_file", `{"path":"../../etc/passwd"}`)
	if result.Success {
		t.Error("path traversal should be blocked")
	}
	if result.Error != "path escapes workspace directory" {
		t.Errorf("Error = %q, want path escapes message", result.Error)
	}
}

func TestReadRepoFileNonexistent(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	result := r.Execute("read_repo_file", `{"path":"nonexistent.go"}`)
	if result.Success {
		t.Error("reading nonexistent file should fail")
	}
}

func TestReadRepoFileMaxLines(t *testing.T) {
	tmpDir := t.TempDir()
	content := ""
	for i := 0; i < 100; i++ {
		content += "line " + string(rune('A'+i%26)) + "\n"
	}
	testFile := filepath.Join(tmpDir, "big.txt")
	os.WriteFile(testFile, []byte(content), 0644)

	r := NewToolRegistry()
	RegisterBuiltInTools(r, tmpDir)

	result := r.Execute("read_repo_file", `{"path":"big.txt","max_lines":5}`)
	if !result.Success {
		t.Fatalf("read with max_lines failed: %s", result.Error)
	}
	if !contains(result.Output, "... (truncated)") {
		t.Error("expected truncation marker in output")
	}
}

func TestWriteFileSuccess(t *testing.T) {
	tmpDir := t.TempDir()
	r := NewToolRegistry()
	RegisterBuiltInTools(r, tmpDir)

	result := r.Execute("write_file", `{"path":"src/handler.go","content":"package handler"}`)
	if !result.Success {
		t.Fatalf("write_file failed: %s", result.Error)
	}

	writtenFile := filepath.Join(tmpDir, "src", "handler.go")
	data, err := os.ReadFile(writtenFile)
	if err != nil {
		t.Fatalf("failed to read written file: %v", err)
	}
	if string(data) != "package handler" {
		t.Errorf("written content = %q, want package handler", string(data))
	}
}

func TestWriteFileCreatesParentDirs(t *testing.T) {
	tmpDir := t.TempDir()
	r := NewToolRegistry()
	RegisterBuiltInTools(r, tmpDir)

	result := r.Execute("write_file", `{"path":"deep/nested/dir/file.txt","content":"nested"}`)
	if !result.Success {
		t.Fatalf("write_file with nested dirs failed: %s", result.Error)
	}
}

func TestWriteFileMissingParams(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	result := r.Execute("write_file", `{"path":"file.txt","content":""}`)
	if result.Success {
		t.Error("empty content should fail")
	}

	result2 := r.Execute("write_file", `{"path":"","content":"data"}`)
	if result2.Success {
		t.Error("empty path should fail")
	}
}

func TestWriteFilePathTraversal(t *testing.T) {
	tmpDir := t.TempDir()
	r := NewToolRegistry()
	RegisterBuiltInTools(r, tmpDir)

	result := r.Execute("write_file", `{"path":"../../tmp/malicious.sh","content":"evil"}`)
	if result.Success {
		t.Error("path traversal should be blocked")
	}
}

func TestTriggerCICDSuccess(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	result := r.Execute("trigger_cicd", `{"service":"backend-api","action":"deploy"}`)
	if !result.Success {
		t.Fatalf("trigger_cicd failed: %s", result.Error)
	}
	if !contains(result.Output, "backend-api") {
		t.Errorf("output should mention service name: %q", result.Output)
	}
}

func TestTriggerCICDDefaultBranch(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	result := r.Execute("trigger_cicd", `{"service":"svc","action":"build"}`)
	if !contains(result.Output, "branch=main") {
		t.Errorf("default branch should be main: %q", result.Output)
	}
}

func TestTriggerCICDCustomBranch(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	result := r.Execute("trigger_cicd", `{"service":"svc","action":"build","branch":"develop"}`)
	if !contains(result.Output, "branch=develop") {
		t.Errorf("custom branch should be develop: %q", result.Output)
	}
}

func TestTriggerCICDMissingParams(t *testing.T) {
	r := NewToolRegistry()
	RegisterBuiltInTools(r, t.TempDir())

	result := r.Execute("trigger_cicd", `{"service":""}`)
	if result.Success {
		t.Error("missing service should fail")
	}

	result2 := r.Execute("trigger_cicd", `{"action":"deploy"}`)
	if result2.Success {
		t.Error("missing action should fail")
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
