package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	functions "ascii-art/Functions"
)

type TemplateData struct {
	Input    string   // User input text
	AsciiArt []string // ASCII art lines
    Fontfile string   // Selected font file
}

func main() {
    mux:=http.NewServeMux()
	mux.HandleFunc("/", webhandler)
	mux.HandleFunc("/submit", HandleRequest)

	fmt.Println("Server running on http://localhost:5500/")

	err := http.ListenAndServe(":5500", mux)
	log.Fatal(err)
}

func webhandler(write http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("index.html"))
	temp.Execute(write, nil)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		
		return
	} else if r.Method == "POST" {
		r.ParseForm()
		
		input := r.Form.Get("text")
		fontFile := r.Form.Get("File")

		asciiArt, err := functions.Ascii(input, fontFile)
		if err != nil {
			log.Println("Error generating Ascii art:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}       

		data := TemplateData{
			Input:    input,
			AsciiArt: asciiArt,
			Fontfile: fontFile,
		}

		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, data)

	}
}
