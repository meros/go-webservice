go-webservice
=============

A web service using go/protobuf

This is a testbed for a bigger idea. Right now it's a simple rpc over http using protobuf for serialization. The bigger idea is to automatically generate server and client stubs where you can easily hook in custom code for:

Server side:
* go

Client side:
* go
* js
* c++

Current status:
A simple decoder/encoder for protocol buffers
A simple web service library

An manually coded example server (send 2 uint32, receive sum as result)
* go webservice
* go client
* html/js client

Next steps:
* Create templates and generate example libraries from templates
* More examples
* Simpler library design (less deps?)
* More clients
