// Code generated by wit-bindgen-go. DO NOT EDIT.

package cv

// This file contains wasmimport and wasmexport declarations for "wasm:cv".

//go:wasmimport wasm:cv/cv arrowed-line
//go:noescape
func wasmimport_ArrowedLine(img0 uint32, point10 uint32, point11 uint32, point20 uint32, point21 uint32, c0 uint32, c1 uint32, c2 uint32, c3 uint32, thickness0 uint32)

//go:wasmimport wasm:cv/cv rectangle
//go:noescape
func wasmimport_Rectangle(img0 uint32, r0 uint32, r1 uint32, r2 uint32, r3 uint32, c0 uint32, c1 uint32, c2 uint32, c3 uint32, thickness0 uint32)

//go:wasmimport wasm:cv/cv circle
//go:noescape
func wasmimport_Circle(img0 uint32, center0 uint32, center1 uint32, radius0 uint32, c0 uint32, c1 uint32, c2 uint32, c3 uint32, thickness0 uint32)

//go:wasmimport wasm:cv/cv line
//go:noescape
func wasmimport_Line(img0 uint32, point10 uint32, point11 uint32, point20 uint32, point21 uint32, c0 uint32, c1 uint32, c2 uint32, c3 uint32, thickness0 uint32)

//go:wasmimport wasm:cv/cv put-text
//go:noescape
func wasmimport_PutText(img0 uint32, text0 *uint8, text1 uint32, org0 uint32, org1 uint32, fontFace0 uint32, fontScale0 float64, c0 uint32, c1 uint32, c2 uint32, c3 uint32, thickness0 uint32)

//go:wasmimport wasm:cv/cv adaptive-threshold
//go:noescape
func wasmimport_AdaptiveThreshold(src0 uint32, maxValue0 float32, adaptiveType0 uint32, thresholdType0 uint32, blockSize0 uint32, c0 float32) (result0 uint32)

//go:wasmimport wasm:cv/cv blur
//go:noescape
func wasmimport_Blur(src0 uint32, kSize0 uint32, kSize1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv box-filter
//go:noescape
func wasmimport_BoxFilter(src0 uint32, depth0 uint32, kSize0 uint32, kSize1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv canny
//go:noescape
func wasmimport_Canny(src0 uint32, threshold10 float32, threshold20 float32) (result0 uint32)

//go:wasmimport wasm:cv/cv cvt-color
//go:noescape
func wasmimport_CvtColor(src0 uint32, code0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv dilate
//go:noescape
func wasmimport_Dilate(src0 uint32, kernel0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv erode
//go:noescape
func wasmimport_Erode(src0 uint32, kernel0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv equalize-hist
//go:noescape
func wasmimport_EqualizeHist(src0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv gaussian-blur
//go:noescape
func wasmimport_GaussianBlur(src0 uint32, size0 uint32, size1 uint32, sigmaX0 float32, sigmaY0 float32, border0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv get-structuring-element
//go:noescape
func wasmimport_GetStructuringElement(shape0 uint32, size0 uint32, size1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv hough-lines
//go:noescape
func wasmimport_HoughLines(src0 uint32, rho0 float64, theta0 float64, threshold0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv hough-lines-p
//go:noescape
func wasmimport_HoughLinesP(src0 uint32, rho0 float64, theta0 float64, threshold0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv median-blur
//go:noescape
func wasmimport_MedianBlur(src0 uint32, kSize0 uint32, kSize1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv resize
//go:noescape
func wasmimport_Resize(src0 uint32, size0 uint32, size1 uint32, fx0 float32, fy0 float32, interp0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv threshold
//go:noescape
func wasmimport_Threshold(src0 uint32, thresh0 float32, maxValue0 float32, thresholdType0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv transpose-ND
//go:noescape
func wasmimport_TransposeND(src0 uint32, order0 *int32, order1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv estimate-affine2d
//go:noescape
func wasmimport_EstimateAffine2d(frm0 uint32, to0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv warp-affine
//go:noescape
func wasmimport_WarpAffine(src0 uint32, m0 uint32, size0 uint32, size1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv get-rotation-matrix2d
//go:noescape
func wasmimport_GetRotationMatrix2d(center0 uint32, center1 uint32, angle0 float64, scale0 float64) (result0 uint32)

//go:wasmimport wasm:cv/cv add
//go:noescape
func wasmimport_Add(src10 uint32, src20 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv add-weighted
//go:noescape
func wasmimport_AddWeighted(src10 uint32, alpha0 float64, src20 uint32, beta0 float64, gamma0 float64) (result0 uint32)

//go:wasmimport wasm:cv/cv exp
//go:noescape
func wasmimport_Exp(src0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv hconcat
//go:noescape
func wasmimport_Hconcat(src10 uint32, src20 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv vconcat
//go:noescape
func wasmimport_Vconcat(src10 uint32, src20 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv lut
//go:noescape
func wasmimport_Lut(src0 uint32, wblut0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/cv reduce-arg-max
//go:noescape
func wasmimport_ReduceArgMax(src0 uint32, axis0 uint32, lastIndex0 uint32) (result0 uint32)
