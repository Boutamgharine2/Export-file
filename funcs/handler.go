package ascii

import (
	"html/template"
	"net/http"
)

func Handler_rout(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
	}

	tmpl.Execute(w, nil)
}

func Handler_asci_art(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "methode not allowed", http.StatusMethodNotAllowed)
		return
	}
	

	tmpl, err := template.ParseFiles("tmplt/my.html")
	
	if err != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
	}
	
	r.ParseForm()
	

	banner := r.Form.Get("style")
	
	input := r.Form.Get("user_input")
	
	if (banner == "" || input == "") || (banner != "standard" && banner != "shadow" && banner != "thinkertoy") {
		http.Error(w, "Bad Request 400", http.StatusBadRequest)
		return
	}
	v := Printing(input,banner)
	
   
	tmpl.Execute(w, v)
}
