package main

import (
	"github.com/baxiang/go-note/minio/minio"
	"net/http"
	"log"
	"io/ioutil"
	"crypto/rand"
	"fmt"
	"mime"
	"path/filepath"
	"os"
)

const (
	uploadPath    = "./tmp"
	maxUploadSize = 20 * 1024 * 1014
	location      = "cn-east-1"
	contentType   = "application/octet-stream"
)

func main() {
	http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files", http.StripPrefix("/files", fs))
	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{file_name} for download files.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func uploadFileHandler() func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			log.Println(err)
			renderError(w, http.StatusBadRequest, "FILE_TOO_BIG")
			return
		}
		//fileType := r.PostFormValue("type")
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			log.Println(err)
			renderError(w, http.StatusBadRequest, "INVALID_FILE")
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, http.StatusBadRequest, "INVALID_FILE")
			return
		}
		fileType := http.DetectContentType(fileBytes)
		switch fileType {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			//renderError(w, http.StatusBadRequest, "INVALID_FILE_TYPE")
			//return
		}
		fileName := randToken(12)
		fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			renderError(w, http.StatusBadRequest, "CANT_READ_FILE_TYPE")
			return
		}
		fileName = fileName + fileEndings[0]
		newPath := filepath.Join(uploadPath, fileName)
		newFile, err := os.Create(newPath)
		if err != nil {
			log.Println(err)
			renderError(w, http.StatusBadRequest, "CANT_WRITE_FILE")
			return
		}
		defer newFile.Close()
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			log.Println(err)
			renderError(w, http.StatusBadRequest, "CANT_WRITE_FILE")
			return
		}
		objectPath := minio.FPutObject("upload", location, fileName, newPath, contentType)
		w.Write([]byte(objectPath))
	})
}
func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
func renderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
