<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">compute</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

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

MIT © fadymondy
