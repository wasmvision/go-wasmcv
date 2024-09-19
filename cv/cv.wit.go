// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package cv represents the world "wasm:cv/cv".
//
// wasmCV is a WebAssembly guest module interface for computer vision based on OpenCV.
package cv

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"wasmcv.org/wasm/cv/mat"
	"wasmcv.org/wasm/cv/types"
)

// Rectangle represents the imported function "rectangle".
//
// drawing functions
// Rectangle draws a simple, thick, or filled up-right rectangle.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d6/d6e/group__imgproc__draw.html#ga07d2f74cadcf8e305e810ce8f3d1e1b7
//
//	rectangle: func(img: mat, r: rect, c: rgba, thickness: u8)
//
//go:nosplit
func Rectangle(img mat.Mat, r types.Rect, c types.Rgba, thickness uint8) {
	img0 := cm.Reinterpret[uint32](img)
	r0, r1, r2, r3 := lower_Rect(r)
	c0, c1, c2, c3 := lower_Rgba(c)
	thickness0 := (uint32)(thickness)
	wasmimport_Rectangle((uint32)(img0), (uint32)(r0), (uint32)(r1), (uint32)(r2), (uint32)(r3), (uint32)(c0), (uint32)(c1), (uint32)(c2), (uint32)(c3), (uint32)(thickness0))
	return
}

//go:wasmimport wasm:cv/cv rectangle
//go:noescape
func wasmimport_Rectangle(img0 uint32, r0 uint32, r1 uint32, r2 uint32, r3 uint32, c0 uint32, c1 uint32, c2 uint32, c3 uint32, thickness0 uint32)

// PutText represents the imported function "put-text".
//
// PutText draws a text string.
// It renders the specified text string into the img Mat at the location
// passed in the "org" param, using the desired font face, font scale,
// color, and line thinkness.
//
// For further details, please see:
// http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576
//
//	put-text: func(img: mat, text: string, org: size, font-face: hershey-font-type,
//	font-scale: f64, c: rgba, thickness: s32)
//
//go:nosplit
func PutText(img mat.Mat, text string, org types.Size, fontFace types.HersheyFontType, fontScale float64, c types.Rgba, thickness int32) {
	img0 := cm.Reinterpret[uint32](img)
	text0, text1 := cm.LowerString(text)
	org0, org1 := lower_Size(org)
	fontFace0 := (uint32)(fontFace)
	fontScale0 := (float64)(fontScale)
	c0, c1, c2, c3 := lower_Rgba(c)
	thickness0 := (uint32)(thickness)
	wasmimport_PutText((uint32)(img0), (*uint8)(text0), (uint32)(text1), (uint32)(org0), (uint32)(org1), (uint32)(fontFace0), (float64)(fontScale0), (uint32)(c0), (uint32)(c1), (uint32)(c2), (uint32)(c3), (uint32)(thickness0))
	return
}

//go:wasmimport wasm:cv/cv put-text
//go:noescape
func wasmimport_PutText(img0 uint32, text0 *uint8, text1 uint32, org0 uint32, org1 uint32, fontFace0 uint32, fontScale0 float64, c0 uint32, c1 uint32, c2 uint32, c3 uint32, thickness0 uint32)

// AdaptiveThreshold represents the imported function "adaptive-threshold".
//
// imgproc functions
// AdaptiveThreshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga72b913f352e4a1b1b397736707afcde3
//
//	adaptive-threshold: func(src: mat, max-value: f32, adaptive-type: adaptive-threshold-type,
//	threshold-type: threshold-type, block-size: u32, c: f32) -> mat
//
//go:nosplit
func AdaptiveThreshold(src mat.Mat, maxValue float32, adaptiveType types.AdaptiveThresholdType, thresholdType types.ThresholdType, blockSize uint32, c float32) (result mat.Mat) {
	src0 := cm.Reinterpret[uint32](src)
	maxValue0 := (float32)(maxValue)
	adaptiveType0 := (uint32)(adaptiveType)
	thresholdType0 := (uint32)(thresholdType)
	blockSize0 := (uint32)(blockSize)
	c0 := (float32)(c)
	result0 := wasmimport_AdaptiveThreshold((uint32)(src0), (float32)(maxValue0), (uint32)(adaptiveType0), (uint32)(thresholdType0), (uint32)(blockSize0), (float32)(c0))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv adaptive-threshold
//go:noescape
func wasmimport_AdaptiveThreshold(src0 uint32, maxValue0 float32, adaptiveType0 uint32, thresholdType0 uint32, blockSize0 uint32, c0 float32) (result0 uint32)

// Blur represents the imported function "blur".
//
// Blur blurs an image Mat using a normalized box filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga8c45db9afe636703801b0b2e440fce37
//
//	blur: func(src: mat, k-size: size) -> mat
//
//go:nosplit
func Blur(src mat.Mat, kSize types.Size) (result mat.Mat) {
	src0 := cm.Reinterpret[uint32](src)
	kSize0, kSize1 := lower_Size(kSize)
	result0 := wasmimport_Blur((uint32)(src0), (uint32)(kSize0), (uint32)(kSize1))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv blur
//go:noescape
func wasmimport_Blur(src0 uint32, kSize0 uint32, kSize1 uint32) (result0 uint32)

// BoxFilter represents the imported function "box-filter".
//
// BoxFilter blurs an image using the box filter.
//
// For further details, please see:
// https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gad533230ebf2d42509547d514f7d3fbc3
//
//	box-filter: func(src: mat, depth: u32, k-size: size) -> mat
//
//go:nosplit
func BoxFilter(src mat.Mat, depth uint32, kSize types.Size) (result mat.Mat) {
	src0 := cm.Reinterpret[uint32](src)
	depth0 := (uint32)(depth)
	kSize0, kSize1 := lower_Size(kSize)
	result0 := wasmimport_BoxFilter((uint32)(src0), (uint32)(depth0), (uint32)(kSize0), (uint32)(kSize1))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv box-filter
//go:noescape
func wasmimport_BoxFilter(src0 uint32, depth0 uint32, kSize0 uint32, kSize1 uint32) (result0 uint32)

// GaussianBlur represents the imported function "gaussian-blur".
//
// GaussianBlur blurs an image using a Gaussian filter.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d4/d86/group__imgproc__filter.html#gae8bdcd9154ed5ca3cbc1766d960f45c1
//
//	gaussian-blur: func(src: mat, size: size, sigma-x: f32, sigma-y: f32, border: border-type)
//	-> mat
//
//go:nosplit
func GaussianBlur(src mat.Mat, size types.Size, sigmaX float32, sigmaY float32, border types.BorderType) (result mat.Mat) {
	src0 := cm.Reinterpret[uint32](src)
	size0, size1 := lower_Size(size)
	sigmaX0 := (float32)(sigmaX)
	sigmaY0 := (float32)(sigmaY)
	border0 := (uint32)(border)
	result0 := wasmimport_GaussianBlur((uint32)(src0), (uint32)(size0), (uint32)(size1), (float32)(sigmaX0), (float32)(sigmaY0), (uint32)(border0))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv gaussian-blur
//go:noescape
func wasmimport_GaussianBlur(src0 uint32, size0 uint32, size1 uint32, sigmaX0 float32, sigmaY0 float32, border0 uint32) (result0 uint32)

// Threshold represents the imported function "threshold".
//
// Threshold applies a fixed-level threshold to each array element.
//
// For further details, please see:
// https://docs.opencv.org/3.3.0/d7/d1b/group__imgproc__misc.html#gae8a4a146d1ca78c626a53577199e9c57
//
//	threshold: func(src: mat, thresh: f32, max-value: f32, threshold-type: threshold-type)
//	-> mat
//
//go:nosplit
func Threshold(src mat.Mat, thresh float32, maxValue float32, thresholdType types.ThresholdType) (result mat.Mat) {
	src0 := cm.Reinterpret[uint32](src)
	thresh0 := (float32)(thresh)
	maxValue0 := (float32)(maxValue)
	thresholdType0 := (uint32)(thresholdType)
	result0 := wasmimport_Threshold((uint32)(src0), (float32)(thresh0), (float32)(maxValue0), (uint32)(thresholdType0))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv threshold
//go:noescape
func wasmimport_Threshold(src0 uint32, thresh0 float32, maxValue0 float32, thresholdType0 uint32) (result0 uint32)

// TransposeNd represents the imported function "transpose-nd".
//
// Transpose for n-dimensional matrices.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gab1b1274b4a563be34cdfa55b8919a4ec
//
//	transpose-nd: func(src: mat, order: list<s32>) -> mat
//
//go:nosplit
func TransposeNd(src mat.Mat, order cm.List[int32]) (result mat.Mat) {
	src0 := cm.Reinterpret[uint32](src)
	order0, order1 := cm.LowerList(order)
	result0 := wasmimport_TransposeNd((uint32)(src0), (*int32)(order0), (uint32)(order1))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv transpose-nd
//go:noescape
func wasmimport_TransposeNd(src0 uint32, order0 *int32, order1 uint32) (result0 uint32)

// BlobFromImage represents the imported function "blob-from-image".
//
// dnn functions
// BlobFromImage creates 4-dimensional blob from image. Optionally resizes and crops
// image from center,
// subtract mean values, scales values by scalefactor, swap Blue and Red channels.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga29f34df9376379a603acd8df581ac8d7
//
//	blob-from-image: func(image: mat, scale-factor: f32, size: size, mean: scalar,
//	swap-rb: bool, crop: bool) -> mat
//
//go:nosplit
func BlobFromImage(image mat.Mat, scaleFactor float32, size types.Size, mean types.Scalar, swapRb bool, crop bool) (result mat.Mat) {
	image0 := cm.Reinterpret[uint32](image)
	scaleFactor0 := (float32)(scaleFactor)
	size0, size1 := lower_Size(size)
	mean0, mean1, mean2, mean3 := lower_Scalar(mean)
	swapRb0 := cm.BoolToU32(swapRb)
	crop0 := cm.BoolToU32(crop)
	result0 := wasmimport_BlobFromImage((uint32)(image0), (float32)(scaleFactor0), (uint32)(size0), (uint32)(size1), (float32)(mean0), (float32)(mean1), (float32)(mean2), (float32)(mean3), (uint32)(swapRb0), (uint32)(crop0))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv blob-from-image
//go:noescape
func wasmimport_BlobFromImage(image0 uint32, scaleFactor0 float32, size0 uint32, size1 uint32, mean0 float32, mean1 float32, mean2 float32, mean3 float32, swapRb0 uint32, crop0 uint32) (result0 uint32)

// BlobFromImageWithParams represents the imported function "blob-from-image-with-params".
//
// BlobFromImageWithParams creates 4-dimensional blob from image. Optionally resizes
// and crops image from center,
// subtract mean values, scales values by scalefactor, swap Blue and Red channels.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga29f34df9376379a603acd8df581ac8d7
//
//	blob-from-image-with-params: func(image: mat, params: blob-params) -> mat
//
//go:nosplit
func BlobFromImageWithParams(image mat.Mat, params types.BlobParams) (result mat.Mat) {
	image0 := cm.Reinterpret[uint32](image)
	params0, params1, params2, params3, params4, params5, params6, params7, params8, params9, params10, params11, params12, params13, params14 := lower_BlobParams(params)
	result0 := wasmimport_BlobFromImageWithParams((uint32)(image0), (float32)(params0), (uint32)(params1), (uint32)(params2), (float32)(params3), (float32)(params4), (float32)(params5), (float32)(params6), (uint32)(params7), (uint32)(params8), (uint32)(params9), (uint32)(params10), (float32)(params11), (float32)(params12), (float32)(params13), (float32)(params14))
	result = cm.Reinterpret[mat.Mat]((uint32)(result0))
	return
}

//go:wasmimport wasm:cv/cv blob-from-image-with-params
//go:noescape
func wasmimport_BlobFromImageWithParams(image0 uint32, params0 float32, params1 uint32, params2 uint32, params3 float32, params4 float32, params5 float32, params6 float32, params7 uint32, params8 uint32, params9 uint32, params10 uint32, params11 float32, params12 float32, params13 float32, params14 float32) (result0 uint32)

// BlobRectsToImageRects represents the imported function "blob-rects-to-image-rects".
//
// BlobRectsToImageRects converts blob rects to image rects.
//
// For further details, please see:
// https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee
//
//	blob-rects-to-image-rects: func(params: blob-params, blob-rects: list<rect>, image-size:
//	size) -> list<rect>
//
//go:nosplit
func BlobRectsToImageRects(params types.BlobParams, blobRects cm.List[types.Rect], imageSize types.Size) (result cm.List[types.Rect]) {
	params_ := wasmimport_BlobRectsToImageRects_params{params, blobRects, imageSize}
	wasmimport_BlobRectsToImageRects(&params_, &result)
	return
}

//go:wasmimport wasm:cv/cv blob-rects-to-image-rects
//go:noescape
func wasmimport_BlobRectsToImageRects(params *wasmimport_BlobRectsToImageRects_params, result *cm.List[types.Rect])

// wasmimport_BlobRectsToImageRects_params represents the flattened function params for [wasmimport_BlobRectsToImageRects].
// See the Canonical ABI flattening rules for more information.
type wasmimport_BlobRectsToImageRects_params struct {
	params    types.BlobParams
	blobRects cm.List[types.Rect]
	imageSize types.Size
}

// NmsBoxes represents the imported function "nms-boxes".
//
// NMSBoxes performs non maximum suppression given boxes and corresponding scores.
//
// For futher details, please see:
// https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee
//
//	nms-boxes: func(bboxes: list<rect>, scores: list<f32>, score-threshold: f32, nms-threshold:
//	f32) -> list<s32>
//
//go:nosplit
func NmsBoxes(bboxes cm.List[types.Rect], scores cm.List[float32], scoreThreshold float32, nmsThreshold float32) (result cm.List[int32]) {
	bboxes0, bboxes1 := cm.LowerList(bboxes)
	scores0, scores1 := cm.LowerList(scores)
	scoreThreshold0 := (float32)(scoreThreshold)
	nmsThreshold0 := (float32)(nmsThreshold)
	wasmimport_NmsBoxes((*types.Rect)(bboxes0), (uint32)(bboxes1), (*float32)(scores0), (uint32)(scores1), (float32)(scoreThreshold0), (float32)(nmsThreshold0), &result)
	return
}

//go:wasmimport wasm:cv/cv nms-boxes
//go:noescape
func wasmimport_NmsBoxes(bboxes0 *types.Rect, bboxes1 uint32, scores0 *float32, scores1 uint32, scoreThreshold0 float32, nmsThreshold0 float32, result *cm.List[int32])
