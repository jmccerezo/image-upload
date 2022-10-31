package views

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ImageUpload(r *http.Request, name string) (string, error) {
	//* this function returns the filename(to save in database) of the saved file or an error if it occurs

	r.ParseMultipartForm(32 << 20)
	//* ParseMultipartForm parses a request body as multipart/form-data

	file, handler, err := r.FormFile(name) //* retrieve the file from form data
	//* replace file with the key your sent your image with
	if err != nil {
		return "", err
	}
	defer file.Close() //* close the file when we finish

	//* this is path which  we want to store the file
	path := "../image-upload/"
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)

	//* here we save our file to our path
	filepath := path + handler.Filename
	fmt.Println("FILE SUCCESSFULLY UPLOADED")
	return filepath, err
}
