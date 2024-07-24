package analysis

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

// ReadFileBytes reads the bytes of a file.
func ReadFileBytes(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(file)
	b := make([]byte, len(file))
	_, err = reader.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// WritePng writes an image to a file.
func WritePng(filename string, imgBytes []byte, width, height int) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	img.Pix = imgBytes
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}
