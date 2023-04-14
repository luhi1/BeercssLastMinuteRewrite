package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", loginHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tplExec(w, "teacherEvents.html")
}

func tplExec(w http.ResponseWriter, filename string) {
	temp := template.Must(template.ParseFiles(filename))

	err := temp.Execute(w, nil)
	if err != nil {
		return
	}
}
