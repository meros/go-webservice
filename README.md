go-webservice
=============

A web service generator for go/protobuf.

Goals:
* Given a protobuf definition file, generate stubs for webservice and several web service clients including but not limited to html/js web service client.
* Make designed and well thought throught API's easier to deploy and use (to avoid using JSON only because of ease of use)

Antigoals:
* JSON
* Interoperability with existing web service standards

Current status:
A simple framework to enable web service clients/servers for go
A hand coded example application, sum, consisting of a protobuf definition go server and go/js clients

Next steps:
* Complete go client library (move generic code from example application to library)
* Extract js library (extract generic code from example html application to library)
* Extract templates from the non-generic parts of the example application
* Create generator application and generate sum example from templates
