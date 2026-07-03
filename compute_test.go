package compute

import (
	"context"
	"strings"
	"testing"

	"github.com/togo-framework/togo"
)

func TestLocalSucceeds(t *testing.T) {
	run, err := (&localCompute{}).Submit(context.Background(), Job{Name: "echo", Cmd: []string{"echo", "hi"}})
	if err != nil {
		t.Fatal(err)
	}
	if run.Status != "succeeded" || !strings.Contains(run.Output, "hi") || run.ID == "" {
		t.Fatalf("bad run: %+v", run)
	}
}

func TestLocalFails(t *testing.T) {
	run, _ := (&localCompute{}).Submit(context.Background(), Job{Name: "false", Cmd: []string{"false"}})
	if run.Status != "failed" {
		t.Fatalf("expected failed, got %+v", run)
	}
}

func TestLocalNoCmd(t *testing.T) {
	if _, err := (&localCompute{}).Submit(context.Background(), Job{Name: "x"}); err == nil {
		t.Fatal("expected error for empty Cmd")
	}
}

func TestFromKernelDefaultsToLocal(t *testing.T) {
	c := FromKernel(togo.New())
	if c == nil {
		t.Fatal("no compute backend registered")
	}
	if _, ok := c.(*localCompute); !ok {
		t.Fatalf("default should be localCompute, got %T", c)
	}
}
