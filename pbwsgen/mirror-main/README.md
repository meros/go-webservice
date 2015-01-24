To build:

1. go to pbwsgen folder
2. run go run main.go -proto="mirror" -req="MirrorReq" -resp="MirrorResp"
3. go to mirror-main folder
4. go run main.go

Server will listen on http/8080 and ws is exposed on /mirror/ws
