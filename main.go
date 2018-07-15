package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

// Popular wisdom is that Twitter will leave a png file as a png (not
// convert to jpeg) if it has even a single non-opaque pixel.
//
// This utility loads a png, tweaks one pixel, and saves the file as
// foo-tweet.png.

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: tweetpng <png file>")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	outfile := strings.TrimSuffix(filename, ".png") + "-tweet.png"

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Reading png: %s\n", err)
		os.Exit(1)
	}

	img, err := png.Decode(bytes.NewReader(buf))
	if err != nil {
		fmt.Printf("Decoding png: %s\n", err)
		os.Exit(1)
	}

	dst := image.NewNRGBA(img.Bounds())
	draw.Draw(dst, dst.Bounds(), img, img.Bounds().Min, draw.Src)

	if dst.Opaque() {
		dst.Pix[3] = 0xfe
	}

	w := bytes.NewBuffer(nil)
	if err := png.Encode(w, dst); err != nil {
		fmt.Printf("Encoding png: %s\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(outfile, w.Bytes(), 0644); err != nil {
		fmt.Printf("Writing png: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Wrote %s\n", outfile)
}
