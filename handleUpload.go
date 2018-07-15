package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func handler(writer http.ResponseWriter, request * http.Request)  {
	fmt.Println("new connection")
	request.ParseMultipartForm(0);

	if request.MultipartForm != nil && request.MultipartForm.File != nil {
		for key, fhs := range request.MultipartForm.File {
			fmt.Println("Filekey:" + key)
			for _, fh := range fhs {
				fmt.Println("\tFilename:" + fh.Filename)

				file, err := fh.Open()
				if err != nil {
					fmt.Println(err)
					continue
				}
				defer file.Close()

				dstFile, err := os.OpenFile("./" + fh.Filename, os.O_WRONLY | os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err)
					continue
				}
				defer dstFile.Close()

				io.Copy(dstFile, file)
			}
		}
	}


	for key, texts := range request.Form {
		fmt.Println("Textkey:" + key)
		for _, text := range texts {
			fmt.Println("\tText:" + text)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("start listenning")
	http.ListenAndServe(":3000", nil)
}