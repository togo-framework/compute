<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/compute</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/compute"><img src="https://pkg.go.dev/badge/github.com/togo-framework/compute.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/compute
```

<!-- /togo-header -->

togo's **compute** capability — submit batch/stream jobs to a pluggable engine.
Distinct from [`worker`](https://github.com/togo-framework/worker) (supervised
background goroutine workers) and [`queue`](https://github.com/togo-framework/queue)
(app job dispatch): `compute` is for **data/compute engines**.

Main plugin: defines `Compute` + ships the built-in **`local`** backend. Engine
backends register into the same slot:

| Backend | Repo |
|---|---|
| local *(default)* | built in |
| Apache Beam | `togo-framework/compute-beam` |
| Apache Spark | `togo-framework/compute-spark` |
| Apache Flink | `togo-framework/compute-flink` |
| Databricks | `togo-framework/compute-databricks` |

```bash
togo install togo-framework/compute-spark
togo provider:use compute spark
```

```go
c := compute.FromKernel(k)
run, _ := c.Submit(ctx, compute.Job{Name: "nightly-etl", Cmd: []string{"python", "etl.py"}})
```
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
