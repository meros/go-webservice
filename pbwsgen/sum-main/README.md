To build:

1. go to pbwsgen folder
2. run go run main.go -proto="sum" -req="SumReq" -resp="SumResp"
3. go to sum-main folder
4. go run main.go

Server will listen on http/8080 and ws is exposed on /sum/ws
