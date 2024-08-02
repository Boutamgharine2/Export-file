package ascii

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Data struct {
	Tableau []string
	input   string
}

var (
	output Data
	v      string
)

func Handler_rout(w http.ResponseWriter, r *http.Request) {
	output.Extrait()

	if r.URL.Path != "/" { // handel if url was not valide
		http.Error(w, "page not found :)", http.StatusNotFound)
		return
	}
	
	if r.Method != http.MethodGet { // handel if url was not valide
		http.Error(w, "methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("tmplt/index.html")
	r.ParseForm()
	tmpl.Execute(w, output.Tableau)

	if err != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
	}
}

func Handler_asci_art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("", "")
	tmpl, err := template.ParseFiles("tmplt/my.html")
	if err != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
	}

	r.ParseForm()

	banner := r.Form.Get("style")

	input := r.Form.Get("user_input")
	if !Checkinput(input) {
		tmp, err := template.ParseFiles("tmplt/error.html")
		if err != nil {
			http.Error(w, "internal server error 500", http.StatusInternalServerError)
		}
		tmp.Execute(w, nil)
		return

	}

	if (banner == "" || input == "") || (banner != "standard" && banner != "shadow" && banner != "thinkertoy") {

		tmpl, err := template.ParseFiles("tmplt/bad.html")
		if err != nil {
			http.Error(w, "internal server error 500", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusBadRequest)
		// r.ParseForm()

		tmpl.Execute(w, nil)

		return
	}
	v = Printing(input, banner)

	tmpl.Execute(w, v)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "downloaded.txt"
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")

	w.Header().Set("Content-Length", strconv.Itoa(len(v)))

	w.Write([]byte(v))
}

func (r *Data) Extrait() {
	dir, err1 := os.Open("Fonts")
	if err1 != nil {
		fmt.Println(err1)
	}

	tr, err2 := dir.Readdirnames(-1)
	if err2 != nil {
		fmt.Println(err2)
	}
	r.Tableau = nil
	for _, v := range tr {
		r.Tableau = append(r.Tableau, v[:len(v)-4])
	}
}
