# compute

See the [main README](../README.md). Full framework docs: https://github.com/togo-framework/togo/tree/main/docs

togo's **compute** capability: `compute.Compute.Submit(ctx, Job)`. Ships the
`local` default; engine backends (compute-beam/spark/flink/databricks) register
into the `compute` slot, selected via `togo provider:use compute <name>`.
