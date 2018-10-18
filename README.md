# ProtobufDemo
This repo contains:
* a [Go App](protobufDemo.go) that can write/read protobuf encoded files
* a [C# app](cSharpProtobufDemo) that can read protobuf encoded files

## Generating Protobuf dependencies

In order to use Protobuf with go and C# you first need to create a .proto file like [addressbook.proto](addressbook.proto),
then you use `protoc` to generate .go and .cs files using the following commands:

Generate protobuf files in go:
```bash
protoc --go_out=protogen addressbook.proto
```

Generate protobuf files in C#:
```bash
protoc --csharp_out=cSharpProtobufDemo\protobufDemo\protogen addressbook.proto
```

Now that the class/type files have been generated, they can be used within code just like normal classes/types in C#/Go. 
See [protobufDemo.go](protobufDemo.go) and [Program.cs](cSharpProtobufDemo/protobufDemo/Program.cs) for examples of this.