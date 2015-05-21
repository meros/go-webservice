go-webservice
=============

A web service library and stub generator for protobuf over http.

Goals
---
* Given a protobuf definition file, generate stubs for webservice and several web service clients including but not limited to html/js web service client.
* Make webservices with clean and clear API's easier to deploy and use (ad hoc interfaces are quick, but costly in the long run)
* One backend (go), one transport (http/protobuf), several clients (go & js)

Overview
---

pbws/

Documentation: http://godoc.org/github.com/meros/go-webservice/pbws

Simple client/server protocol buffers over web service library

pbwsgen/

Documentation: http://godoc.org/github.com/meros/go-webservice/pbwsgen

Stub/lib generator as well as some examples

Status
---
A simple framework to enable web service clients/servers for go
A working generator for go client/server and js client
