go-webservice
=============

A web service library and stub generator for go/protobuf.

Goals:
* Given a protobuf definition file, generate stubs for webservice and several web service clients including but not limited to html/js web service client.
* Make webservices with clean and clear API's easier to deploy and use (ad hoc interfaces are quick, but costly in the long run)
* One backend (go), one transport (http/protobuf), several clients (go & js)

Antigoals:
* JSON & interoperability with existing web service standards, it's fully possible to do something similar to this system using json schemas to make it usable with legacy services. It's not the focus of this project however. If you are looking for a general web service client stub generator you need to keep looking. This is supposed to be an end to end solution.

Current status:
A simple framework to enable web service clients/servers for go
A simple library to enable js web service clients (still embedded in the sum example)
A hand coded examples
A generator with ability to create a server library and some example proto/main files using this library

Next steps:
* Extract go client library and examples
* Extract js client library and examples
