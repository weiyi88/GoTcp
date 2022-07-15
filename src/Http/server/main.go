package main

import "net/http"

func test(w http.ResponseWriter, r *http.Request) {
	msg := `<div>
			<h1> This is Aring</h1>
			<span> another line </span>
			</div>
  			`
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/aring/test", test)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
