package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"fmt"
	"path/filepath"
	"encoding/json"
	"lab4/models"
)

func main() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/check", checkOk)

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Payload struct {
	Proto string

}

func checkOk(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

    for k, v := range r.Header {
        fmt.Fprintf(w, "%q: %q\n\n", k, v)
	}

	p := &models.Payload{
		Method: r.Method,
		Addr: r.RemoteAddr,
		URL: r.URL.String(),
		Proto: r.Proto,
		Host: r.Host,
		StatusCode: 200,
		ContentLength: 200,
		ContentType: "application/json",
	}
	json.NewEncoder(w).Encode(p)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}