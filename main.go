package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/signintech/gopdf"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	names, err := filepath.Glob(os.Args[1] + "/*")
	if err != nil {
		log.Fatal(err)
	}
	sort.SliceStable(names, func(i, j int) bool {
		return len(names[i]) < len(names[j])
	})

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

	fmt.Println("Press the Enter Key to terminate the console...")
	fmt.Scanln()
}
