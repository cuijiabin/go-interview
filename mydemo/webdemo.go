package mydemo

import (
"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func RunWeb() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)

}
