package main

import (
	u "goprac/utils"
	"log"
	"mime/multipart"
	"net/http"
)

func validateFile(multipartFileHeader *multipart.FileHeader, allowedFiles []string) (res bool) {
	fileHeader := make([]byte, 512)
	file, _ := multipartFileHeader.Open()

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		log.Print(err)
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		log.Print(err)
	}

	if u.Contains(allowedFiles, http.DetectContentType(fileHeader)) {
		res = true
	}
	return
}
