// Package compute is togo's compute capability — submitting batch/stream jobs to
// a pluggable engine. It is distinct from togo's `worker` (supervised background
// goroutine workers) and `queue` (app job dispatch): compute is for data/compute
// engines (Apache Beam, Spark, Flink, Databricks), selectable at runtime.
//
// This repo is the MAIN plugin: it defines the Compute contract and ships the
// built-in `local` backend (runs the job as a local process). Engine backends
// live in their own repos (compute-beam, compute-spark, compute-flink,
// compute-databricks) and register into the same slot; pick one with
// `togo provider:use compute <name>` (or TOGO_COMPUTE_PROVIDER).
package compute

import (
	"context"

	"github.com/togo-framework/providers"
	"github.com/togo-framework/togo"
)

// Job is a unit of compute work. Cmd is the command (entrypoint + args) the
// engine runs — a script/binary locally, or a jar/module submitted to Spark/Beam.
type Job struct {
	Name string            `json:"name"`
	Cmd  []string          `json:"cmd"`
	Env  map[string]string `json:"env,omitempty"`
	Args map[string]string `json:"args,omitempty"` // engine-specific options
}

// Run is the result/handle of a submitted job.
type Run struct {
	ID     string `json:"id"`
	Status string `json:"status"` // succeeded | failed | running
	Output string `json:"output,omitempty"`
}

// Compute submits jobs to an engine. The built-in `local` backend runs them as a
// local process; engine backends submit to Spark/Beam/Flink/Databricks.
type Compute interface {
	Submit(ctx context.Context, job Job) (Run, error)
}

func init() {
	togo.RegisterProviderFunc("compute", togo.PriorityService, func(k *togo.Kernel) error {
		providers.Use(k, providers.CapCompute, "local", &localCompute{}, true)
		if k.Log != nil {
			k.Log.Info("plugin active", "plugin", "compute")
		}
		return nil
	})
}

// FromKernel returns the active Compute backend (or nil if none registered).
func FromKernel(k *togo.Kernel) Compute {
	if v, ok := k.Get(providers.CapCompute); ok {
		if c, ok := v.(Compute); ok {
			return c
		}
	}
	return nil
}
