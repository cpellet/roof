package analysis

import (
	"gocv.io/x/gocv"
	"image"
	"math"
)

// ComputeMaskedDirection computes the masked direction of a map.
func ComputeMaskedDirection(dir Map3D, m float32, v float32) (Map3D, error) {
	height := len(dir)
	width := len(dir[0])
	masked := make(Map3D, height)
	for i := 0; i < height; i++ {
		masked[i] = make([][]float32, width)
		for j := 0; j < width; j++ {
			masked[i][j] = make([]float32, 3)
			if math.Abs(float64(dir[i][j][0]-m)) < float64(v) {
				masked[i][j][0] = dir[i][j][0]
				masked[i][j][1] = dir[i][j][1]
				masked[i][j][2] = dir[i][j][2]
			} else {
				masked[i][j][0] = 0
				masked[i][j][1] = 0
				masked[i][j][2] = 0
			}
		}
	}
	return masked, nil
}

// ComputeUV computes the UV coordinates of a normal map.
func ComputeUV(normal Map3D) (Map3D, Map3D, error) {
	height := len(normal)
	width := len(normal[0])
	u := make(Map3D, height)
	v := make(Map3D, height)
	for i := 0; i < height; i++ {
		u[i] = make([][]float32, width)
		v[i] = make([][]float32, width)
		for j := 0; j < width; j++ {
			u[i][j] = make([]float32, 1)
			v[i][j] = make([]float32, 1)
			u[i][j][0] = 0.5 + float32(math.Atan2(float64(normal[i][j][1]), float64(normal[i][j][0])))/(2*math.Pi)
			v[i][j][0] = 0.5 - float32(math.Asin(float64(normal[i][j][2])))/math.Pi
		}
	}
	return u, v, nil
}

// ComputeAngleNormalized computes the normalized angle between two maps.
func ComputeAngleNormalized(u Map3D, v Map3D) (Map3D, error) {
	height := len(u)
	width := len(u[0])
	dir := make(Map3D, height)
	for i := 0; i < height; i++ {
		dir[i] = make([][]float32, width)
		for j := 0; j < width; j++ {
			dir[i][j] = make([]float32, 3)
			dir[i][j][0] = (u[i][j][0] - 0) / (1 - 0)
			dir[i][j][1] = (v[i][j][0] - 0) / (1 - 0)
			dir[i][j][2] = 0
		}
	}
	return dir, nil
}

// ComputeNormal computes the normal of an elevation matrix.
func ComputeNormal(dx Map2D, dy Map2D) Map3D {
	height := len(dx)
	width := len(dx[0])
	normalMap := make(Map3D, height)
	for i := range normalMap {
		normalMap[i] = make([][]float32, width)
		for j := range normalMap[i] {
			normalMap[i][j] = make([]float32, 3)
		}
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nx := -dx[i][j]
			ny := -dy[i][j]
			nz := float32(1.0)

			length := float32(math.Sqrt(float64(nx*nx + ny*ny + nz*nz)))
			normalMap[i][j][0] = nx / length
			normalMap[i][j][1] = ny / length
			normalMap[i][j][2] = nz / length
		}
	}
	return normalMap
}

// ComputeDerivative computes the derivative of an elevation matrix in the x and y directions.
func ComputeDerivative(elevationMatrix Map2D) (Map2D, Map2D) {
	height := len(elevationMatrix)
	width := len(elevationMatrix[0])
	dx := make(Map2D, height)
	dy := make(Map2D, height)
	for i := 0; i < height; i++ {
		dx[i] = make([]float32, width)
		dy[i] = make([]float32, width)
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 {
				dx[i][j] = 0
			} else {
				dx[i][j] = (elevationMatrix[i+1][j] - elevationMatrix[i-1][j]) / 2
			}
			if j == 0 || j == width-1 {
				dy[i][j] = 0
			} else {
				dy[i][j] = (elevationMatrix[i][j+1] - elevationMatrix[i][j-1]) / 2
			}
		}
	}
	return dx, dy
}

// MakeElevationMatrix converts an array of elevations to a 2D matrix, given the width of the matrix.
func MakeElevationMatrix(elevations []float32, width int) Map2D {
	height := len(elevations) / width
	matrix := make([][]float32, height)
	for i := 0; i < height; i++ {
		matrix[i] = elevations[i*width : (i+1)*width]
	}
	return matrix
}

// ComputeConnectedComponents computes the connected components of a mask.
// It returns the labeled image and the number of components.
func ComputeConnectedComponents(mask Map3D) (gocv.Mat, int) {
	height := len(mask)
	width := len(mask[0])
	img := gocv.NewMatWithSize(height, width, gocv.MatTypeCV8U)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if mask[i][j][0] != 0 || mask[i][j][1] != 0 || mask[i][j][2] != 0 {
				img.SetUCharAt(i, j, 1)
			} else {
				img.SetUCharAt(i, j, 0)
			}
		}
	}
	labels := gocv.NewMat()
	numLabels := gocv.ConnectedComponents(img, &labels)
	// Find the largest component
	labelCounts := make([]int, numLabels)
	for i := 0; i < labels.Rows(); i++ {
		for j := 0; j < labels.Cols(); j++ {
			label := labels.GetIntAt(i, j)
			if label > 0 {
				labelCounts[label]++
			}
		}
	}
	biggestComponent := 1
	for i := 2; i < numLabels; i++ {
		if labelCounts[i] > labelCounts[biggestComponent] {
			biggestComponent = i
		}
	}
	return labels, biggestComponent
}

// ComputeLargestComponentMask computes the mask of the largest continuous component in a labeled image.
func ComputeLargestComponentMask(labels gocv.Mat, biggestComponent int) Map3D {
	height := labels.Rows()
	width := labels.Cols()
	mask := make(Map3D, height)
	for i := 0; i < height; i++ {
		mask[i] = make([][]float32, width)
		for j := 0; j < width; j++ {
			mask[i][j] = make([]float32, 3)
			if labels.GetIntAt(i, j) == int32(biggestComponent) {
				mask[i][j][0] = 1
				mask[i][j][1] = 1
				mask[i][j][2] = 1
			} else {
				mask[i][j][0] = 0
				mask[i][j][1] = 0
				mask[i][j][2] = 0
			}
		}
	}
	return mask
}

// ApplyGaussianBlur applies a Gaussian blur to a mask.
// The kernel size is 5x5.
func ApplyGaussianBlur(mask Map3D) Map3D {
	height := len(mask)
	width := len(mask[0])
	img := gocv.NewMatWithSize(height, width, gocv.MatTypeCV32F)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.SetFloatAt(i, j, mask[i][j][0])
		}
	}
	blurred := gocv.NewMat()
	gocv.GaussianBlur(img, &blurred, image.Point{X: 5, Y: 5}, 0, 0, gocv.BorderDefault)
	blurredMask := make(Map3D, height)
	for i := 0; i < height; i++ {
		blurredMask[i] = make([][]float32, width)
		for j := 0; j < width; j++ {
			blurredMask[i][j] = make([]float32, 3)
			value := blurred.GetFloatAt(i, j)
			blurredMask[i][j][0] = value
			blurredMask[i][j][1] = value
			blurredMask[i][j][2] = value
		}
	}
	return blurredMask
}
