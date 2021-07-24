## Custom Plugin for OPA-Envoy

This repo contains an example of creating a [custom plugin](https://www.openpolicyagent.org/docs/latest/extensions/#custom-plugins-for-opa-runtime)
which embeds the Opa-Envoy plugin as a library.

### Steps

1. Build an OPA executable with the plugin
```
$ go build -o opa++ .
```

2. Run the executable with the provided config and policy file
```
$ ./opa++ run --server --config-file config.yaml --skip-version-check policy.rego
```

3. Exercise the plugin
```
$ grpcurl -plaintext -d '
  {
    "attributes": {
      "request": {
        "http": {
          "method": "GET",
          "path": "/api/v1/products"
        }
      }
    }
  }' localhost:9191 envoy.service.auth.v2.Authorization/Check
```

4. If everything worked you will see the string `Logged by Custom Logger!` followed by the Go struct representation of the
   decision log event written to stdout.