## Body Base Routing (BBR) from Gateway API Inference Extension

Trying out body-base-routing from Gateway API Inference Extension https://github.com/kubernetes-sigs/gateway-api-inference-extension/tree/main/pkg/bbr


### Run Env

This example can be run with [`docker compose`](https://docs.docker.com/compose/install/)
and has a matching Envoy configuration.

```sh
$ docker compose up
```

#### Run Traffic

Send HTTP request to `localhost:10000/hello`:

```sh
curl -v http://127.0.0.1:10000/mcp \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -d '{
    "model": "gpt-4.1"
  }'
```

```sh
curl -v http://127.0.0.1:10000/mcp \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -d '{
    "model": "llama"
  }'
```

#### Clean up

```sh
$ docker compose down
```
