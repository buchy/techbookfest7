package main

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image"
	"image/jpeg"
	_ "image/png" // (2)
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.tpl"))

	if err := tpl.Execute(w, nil); err != nil {
		log.Printf("Template実行に失敗しました。: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "不正なHTTP Methodです。", http.StatusMethodNotAllowed)
		return
	}

	// (3)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// (4)
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// (5)
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("File Closeに失敗しました。: %v", err)
		}
	}()

	// (6)
	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("Image decodeに作成に失敗しました。: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	drawImage(w, img)
}

func drawImage(w http.ResponseWriter, image image.Image) {

	// (7)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, image, nil); err != nil {
		log.Printf("Image encodeに失敗しました。: %v", err)
	}

	// (8)
	tpl := template.Must(template.ParseFiles("index.tpl"))
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data := map[string]string{"Image": str}
	if err := tpl.Execute(w, data); err != nil {
		log.Printf("Template実行に失敗しました。: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// (1)
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/analyze", analyzeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
