//go:generate pbwsgen -req=MirrorReq -resp=MirrorResp -proto=../mirror.proto -out_client_js=static/scripts/lib

package main

import "net/http"

func main() {
	http.ListenAndServe(":8081", http.FileServer(http.Dir("static/")))
}
