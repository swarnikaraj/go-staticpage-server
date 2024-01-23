package main

import (
	"fmt"
	"log"
	"net/http"
)




 func methodValidater(next http.HandlerFunc) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request) {
		if r.Method!="GET"{
			http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
			return
		}
		next(w,r)
	}
	
 }

 
 func pathValidator(next http.HandlerFunc, validPath string) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path!=validPath{
			http.Error(w,"Invalid Path",http.StatusNotFound)
			return
		}
		next(w,r)
	}
 }

 func calenderHandler(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"./static/calendar.html")
 }

 func formHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./static/forms.html")
 }
func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	 http.Handle("/",fileserver)
    

	 http.Handle("/forms",pathValidator(methodValidater(formHandler),"/forms"))
	 http.Handle("/calendar",pathValidator(methodValidater(calenderHandler),"/calender"))
	
    fmt.Print("Server is runnning on port 8080")
	if err:= http.ListenAndServe(":8080",nil); err!=nil {
		log.Fatal(err)
	}else{
		fmt.Print("Server runnning on port 8080")
	}
}