package upload

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/KenethSandoval/xopopu/pkg"
)

var (
	err    error
	data   []byte
	result bool
)

// UploadCSV Receive the formdata file and make sure it is a valid .csv
func UploadCSV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	defer file.Close()

	mimeType := handler.Header.Get("Content-Type")

	switch mimeType {
	case "image/png":
		// validar error
		fmt.Println("Solo se permite csv")
	default:
		// subir archivo
		saveFile(w, file, handler)
	}
}

func saveFile(w http.ResponseWriter, file multipart.File, handler *multipart.FileHeader) {
	data, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	result, err = pkg.CreateDir("./files")
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	if result {
		filename := handler.Filename

		err = ioutil.WriteFile("./files/"+filename, data, 0666)

		err = pkg.CSVToJson(filename)

		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		}
	}

	jsonResponse(w, 200, "File uploaded successfully!")
}

func jsonResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
