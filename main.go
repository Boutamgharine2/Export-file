package main

import (
	"fmt"
	"net/http"

	ascii "ascii/funcs"
)

const port = ":5050"

func main() {
	http.HandleFunc("/", ascii.Handler_rout)
	http.HandleFunc("/ascii-art", ascii.Handler_asci_art)
	http.HandleFunc("/download", ascii.DownloadHandler)
	fs := http.StripPrefix("/style/", http.FileServer(http.Dir("style")))
	http.Handle("/style/", fs)
	fmt.Println("(http://localhost:5050/)server started on port", port)
	http.ListenAndServe(port, nil)
}
