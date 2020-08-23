package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	ctx := context.Background()

	runner := NewRunner()

	err := runner.StartAPI(ctx)
	fmt.Println(err)
}

func NewRunner() *Runner {
	return &Runner{
		out: NewPrefixWriter(os.Stderr, "[rna] "),
	}
}

type Runner struct {
	out io.Writer
}

func (r *Runner) StartAPI(ctx context.Context) error {
	bin, err := r.BuildAPI(ctx)
	if err != nil {
		return fmt.Errorf("Runner.StartAPI: %w", err)
	}

	if err := r.startAPI(ctx, bin); err != nil {
		return fmt.Errorf("Runner.StartAPI: %w", err)
	}

	return nil
}

func (r *Runner) log(format string, args ...interface{}) {
	format = strings.TrimSpace(format)
	format += "\n"
	fmt.Fprintf(r.out, format, args...)
}

func (r *Runner) startAPI(ctx context.Context, bin string) error {
	r.log("starting server...")

	cmd := exec.CommandContext(ctx, bin)
	cmd.Dir = filepath.Join(".", "api")
	cmd.Env = []string{
		// DATABASE_URL
	}

	pw := NewPrefixWriter(os.Stderr, "[api] ")
	cmd.Stdout = pw
	cmd.Stderr = pw

	return cmd.Run()
}

func (r *Runner) BuildAPI(ctx context.Context) (string, error) {
	r.log("building server...")

	dest := filepath.Join(".", "bin", "server")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return dest, fmt.Errorf("Runner.BuildAPI: %w", err)
	}

	cacheDir, err := ioutil.TempDir("", "rna")
	if err != nil {
		return dest, fmt.Errorf("Runner.BuildAPI: %w", err)
	}
	defer os.RemoveAll(cacheDir)

	args := []string{"build", "-o", dest, "-tags", "dev", "."}

	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = filepath.Join(".", "api")
	cmd.Env = []string{
		"HOME=" + homeDir,
		"GOCACHE=" + cacheDir,
		"GOOS=" + runtime.GOOS,
		"GOARCH=" + runtime.GOARCH,
		"CGO_ENABLED=0",
		"TMP=" + os.TempDir(),
		"GOPATH=" + os.Getenv("GOPATH"),
	}

	pw := NewPrefixWriter(os.Stderr, "[api] ")
	cmd.Stdout = pw
	cmd.Stderr = pw

	if err := cmd.Run(); err != nil {
		return dest, fmt.Errorf("Runner.BuildAPI: %w", err)
	}

	return dest, nil
}
