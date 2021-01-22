package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"

	"github.com/signintech/gopdf"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	names, err := ls(os.Args[1])
	fp, err := os.Open(names[0])
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	img, _, err := image.Decode(fp)
	if err != nil {
		log.Fatal(err)
	}
	rect := gopdf.Rect{W: float64(img.Bounds().Dx()), H: float64(img.Bounds().Dy())}

	pdf := new(gopdf.GoPdf)
	pdf.Start(gopdf.Config{PageSize: rect})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range names {
		fmt.Println(v)
		pdf.AddPage()
		pdf.Image(v, 0, 0, &rect)
	}

	if err := pdf.WritePdf(fmt.Sprintf("%s.pdf", os.Args[1])); err != nil {
		log.Fatal(err)
	}
}

func ls(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filenames := make([]string, len(files))
	for i, file := range files {
		filenames[i] = dir + "/" + file.Name()
	}
	return filenames, nil
}
