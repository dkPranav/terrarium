grpcurl -plaintext -d '{"api_key": "my_secret", "name": "module_name", "description":"empty", "maturity": 3, "source_url": "url"}' -proto ./Projects/cie-terrarium/terrarium-grpc-gateway/pkg/terrarium/module.proto  localhost:9443 terrarium.module.Publisher/Configure