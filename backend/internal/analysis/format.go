package analysis

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	"image/png"
	"math"
)

// GetFloat32Array converts a byte slice to a float32 array.
func GetFloat32Array(aBytes []byte) []float32 {
	aArr := make([]float32, len(aBytes)/4)
	for i := 0; i < len(aArr); i++ {
		aArr[i] = BytesToFloat32(aBytes[i*4 : (i+1)*4])
	}
	return aArr
}

// Float32ArrayToBytes converts a float32 array to a byte slice.
func Float32ArrayToBytes(aArr []float32) []byte {
	aBytes := make([]byte, len(aArr)*4)
	for i := 0; i < len(aArr); i++ {
		copy(aBytes[i*4:(i+1)*4], Float32ToBytes(aArr[i]))
	}
	return aBytes
}

// Float32ToBytes converts a float32 to a byte slice.
func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

// BytesToFloat32 converts a byte slice to a float32.
func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

// GetPngDimensions returns the width and height of a PNG image.
func GetPngDimensions(imgBytes []byte) (int, int, error) {
	reader := bytes.NewReader(imgBytes)
	img, err := png.Decode(reader)
	if err != nil {
		return 0, 0, err
	}
	return img.Bounds().Dx(), img.Bounds().Dy(), nil
}

// Map3DToPng converts a 3D map to a PNG image.
func Map3DToPng(m Map3D) ([]byte, error) {
	height := len(m)
	width := len(m[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.SetRGBA(j, i, color.RGBA{
				R: uint8((m[i][j][0] + 1) * 127),
				G: uint8((m[i][j][1] + 1) * 127),
				B: uint8((m[i][j][2] + 1) * 127),
				A: 255,
			})
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
