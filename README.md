Steps to reproduce the bug:

1. (Optional) You can run `./build.sh` to rebuild the protobufs
2. Run the server: `go run ./grpc-server`
3. Run `grpcurl -plaintext 127.0.0.1:8000 describe`

On grpcurl v1.8.7 you will get this output:

```
grpc.reflection.v1.ServerReflection is a service:
service ServerReflection {
  rpc ServerReflectionInfo ( stream .grpc.reflection.v1.ServerReflectionRequest ) returns ( stream .grpc.reflection.v1.ServerReflectionResponse );
}
grpc.reflection.v1alpha.ServerReflection is a service:
service ServerReflection {
  rpc ServerReflectionInfo ( stream .grpc.reflection.v1alpha.ServerReflectionRequest ) returns ( stream .grpc.reflection.v1alpha.ServerReflectionResponse );
}
Failed to resolve symbol "servicepb.TestService": file "service.proto" included an unresolvable reference to ".sharedpb.HelloWorldResponse"
```

With grpcurl v1.8.8 you will get:

```
grpc.reflection.v1.ServerReflection is a service:
service ServerReflection {
  rpc ServerReflectionInfo ( stream .grpc.reflection.v1.ServerReflectionRequest ) returns ( stream .grpc.reflection.v1.ServerReflectionResponse );
}
grpc.reflection.v1alpha.ServerReflection is a service:
service ServerReflection {
  rpc ServerReflectionInfo ( stream .grpc.reflection.v1alpha.ServerReflectionRequest ) returns ( stream .grpc.reflection.v1alpha.ServerReflectionResponse );
}
Failed to resolve symbol "servicepb.TestService": proto: invalid syntax: "<unknown:0>"
```

If you apply the patch `change-path.patch` to the repo using `git apply change-path.patch` and restart the gRPC server, both versions of `grpcurl` will start to work.
