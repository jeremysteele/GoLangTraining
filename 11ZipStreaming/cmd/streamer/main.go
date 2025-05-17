package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func populateTmpFile(file string, mbSize int) (string, error) {
	result, err := exec.Command("dd", "if=/dev/random", "of="+file, "bs=1M", fmt.Sprintf("count=%d", mbSize)).CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(result), nil
}

// Sets up needed tmp files
func setupTmpFiles() ([]string, error) {
	files := []string{}

	_, err := populateTmpFile("/tmp/testfile", 100)

	if err != nil {
		return nil, err
	}

	files = append(files, "/tmp/testfile")

	_, err = populateTmpFile("/tmp/testfile2", 200)

	if err != nil {
		return nil, err
	}

	files = append(files, "/tmp/testfile2")

	return files, nil
}

func writeToZip(zw *zip.Writer, file string) error {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	cf, err := zw.Create(f.Name())
	if err != nil {
		return err
	}

	_, err = io.Copy(cf, f)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	tmpFiles, err := setupTmpFiles()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%+v", tmpFiles)
	}

	http.HandleFunc("/doit", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", "files"))
		w.Header().Set("Content-Transfer-Encoding", "binary")
		zw := zip.NewWriter(w)

		for _, file := range tmpFiles {
			err = writeToZip(zw, file)
			if err != nil {
				log.Fatal(err)
			}
		}

		// close the zip Writer to flush the contents to the ResponseWriter
		err = zw.Close()
		if err != nil {
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
