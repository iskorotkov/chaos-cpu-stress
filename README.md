# Chaos CPU Stress

Simple app used in chaos testing as a potential target. It multiplies matrices on CPU and prints measured performance.

- [Chaos CPU Stress](#chaos-cpu-stress)
  - [Env vars](#env-vars)
  - [Development](#development)

## Env vars

Service requires several env vars set (example values are provided in parentheses):

- `DIM` — dimension of the matrices to generate and multiply (`32`).

## Development

To build project:

```shell
go build ./...
```
