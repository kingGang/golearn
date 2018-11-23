package main

import "net/http"

func main()  {
	http.HandleFunc("/hello",hello)
	http.Handle("/say",http.HandlerFunc(say))
	http.ListenAndServe(":8080",nil)
}

func hello(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Hello"))

}

func say(w http.ResponseWriter,req *http.Request)  {
	w.Write([]byte("I say,you are very good"))
}