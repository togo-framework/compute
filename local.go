package compute

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
)

// localCompute runs a job as a local subprocess — the default `compute` backend,
// so compute works out of the box with no engine deployed.
type localCompute struct{}

func (l *localCompute) Submit(ctx context.Context, job Job) (Run, error) {
	if len(job.Cmd) == 0 {
		return Run{}, fmt.Errorf("compute: job %q has no Cmd", job.Name)
	}
	cmd := exec.CommandContext(ctx, job.Cmd[0], job.Cmd[1:]...)
	cmd.Env = os.Environ()
	for k, v := range job.Env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}
	var buf bytes.Buffer
	cmd.Stdout, cmd.Stderr = &buf, &buf
	run := Run{ID: newID()}
	if err := cmd.Run(); err != nil {
		run.Status = "failed"
		run.Output = buf.String()
		return run, nil // the job ran; failure is reported in Run, not as an error
	}
	run.Status = "succeeded"
	run.Output = buf.String()
	return run, nil
}

func newID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
