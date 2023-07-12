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

![image](https://github.com/themreza/ghz-test-sfixed64/assets/5465057/7e7528d3-9501-4f80-8b4d-7dc9582f0fd3)

![image](https://github.com/themreza/ghz-test-sfixed64/assets/5465057/702c9cbc-632e-46f9-afad-08a1d3772816)
