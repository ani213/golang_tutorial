package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	// http.FileServer is used to serve files from the directory ./(same directory)
	// This allows users to access files in that directory via HTTP

	http.Handle("/projectfiles/", http.StripPrefix("/projectfiles", fs))
	// http.Handle registers the /projectfiles/ route, i.e. it serve the file present in above path when route /projectfiles is accessed.
	// http.StripPrefix("/projectfiles", fs) ensures that when a request is made to /projectfiles/file.txt, it actually serves /users/billb/demo/file.txt.
	// i.e. it will remove the endpoint /projectfiles from the file path beacuse this end point is not something which is actually present, it is just a end point
	// we are using to access the files.

	http.HandleFunc("/samplepdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./pdf/sample.pdf")
		// This explicitly serves the file ./pdf/sample.pdf when a request is made to /samplepdf end point.
	})

	http.ListenAndServe(":8080", nil)
}
