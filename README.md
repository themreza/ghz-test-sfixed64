# ghz test sfixed64

A simple gRPC service to test a bug in ghz resulting in sfixed64 values to be
rounded.

## Compile Protobuf

```shell
# Already compiled.
docker run -v $PWD:/defs namely/protoc-all -f schema/test.proto -l go -o .
```

## Run Server

```shell
go build -o server cmd/test-server/main.go
./server
```

## Test with ghz

```shell
ghz --version
# v0.117.0

ghz --insecure \
      --proto schema/test.proto \
      --call protos.TestService.Echo \
      -d '{ "id": 1000000000000000089 }' \
      -n 1 \
      0.0.0.0:9009
      
# Server log:
# Id:  1000000000000000100
```

`1000000000000000089` is rounded up to `1000000000000000100`.

`10000000000000089` is rounded down to `10000000000000088`.

## Test with Postman

![image](https://github.com/themreza/ghz-test-sfixed64/assets/5465057/ebb45b35-f48b-420c-a62f-a1363e80e38a)

![image](https://github.com/themreza/ghz-test-sfixed64/assets/5465057/3276a1ba-59f8-41d9-8986-8fd14e7a7485)

