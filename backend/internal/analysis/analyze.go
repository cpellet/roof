package analysis

import (
	blob "backend/pkg/blobstore"
	"context"
	"github.com/google/uuid"
)

// Processor is a struct that contains the necessary dependencies for processing an analysis request.
type Processor struct {
	Blobstore blob.Blobstore
}

const (
	BlobKeyImage = "image.png"
	BlobKeyMask  = "mask.png"
)

// ProcessRequest processes an analysis request, storing the image and mask in the blobstore.
// It returns the mask and the ID of the analysis request, which can be used to retrieve the image and mask.
func (m *Processor) ProcessRequest(ctx context.Context, imageBytes []byte, emapBytes []byte) ([]byte, string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, "", err
	}

	err = m.Blobstore.CreateObject(ctx, id.String(), BlobKeyImage, imageBytes, true, "image/png")
	if err != nil {
		return nil, "", err
	}

	mask, err := computeSouthernRoofMask(imageBytes, emapBytes)
	if err != nil {
		return nil, "", err
	}

	err = m.Blobstore.CreateObject(ctx, id.String(), BlobKeyMask, mask, false, "image/png")
	if err != nil {
		return nil, "", err
	}

	return mask, id.String(), nil
}

// computeSouthernRoofMask computes the southern roof mask for the given image and elevation map.
// It returns the mask as a byte slice.
func computeSouthernRoofMask(imageBytes []byte, emapBytes []byte) ([]byte, error) {
	elevations := GetFloat32Array(emapBytes)
	width, _, err := GetPngDimensions(imageBytes)
	if err != nil {
		return nil, err
	}

	eMatrix := MakeElevationMatrix(elevations, width)
	dx, dy := ComputeDerivative(eMatrix)
	normalMap := ComputeNormal(dx, dy)
	u, v, err := ComputeUV(normalMap)
	if err != nil {
		return nil, err
	}
	dir, err := ComputeAngleNormalized(u, v)
	if err != nil {
		return nil, err
	}
	mask, err := ComputeMaskedDirection(dir, 0.5, 0.1)
	if err != nil {
		return nil, err
	}

	labels, biggestComponent := ComputeConnectedComponents(mask)
	largestComponentMask := ComputeLargestComponentMask(labels, biggestComponent)
	blurredMask := ApplyGaussianBlur(largestComponentMask)

	return Map3DToPng(blurredMask)
}
