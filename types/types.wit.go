// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package types represents the imported interface "wasm:cv/types".
package types

import (
	"go.bytecodealliance.org/cm"
)

// Size represents the record "wasm:cv/types#size".
//
// size is a 2-element integer vector.
// It represents a width and height.
//
//	record size {
//		x: s32,
//		y: s32,
//	}
type Size struct {
	_ cm.HostLayout
	X int32
	Y int32
}

// Point represents the type alias "wasm:cv/types#point".
//
// See [Size] for more information.
type Point = Size

// Scalar represents the record "wasm:cv/types#scalar".
//
// scalar is a 4-element floating point vector.
//
//	record scalar {
//		val1: f32,
//		val2: f32,
//		val3: f32,
//		val4: f32,
//	}
type Scalar struct {
	_    cm.HostLayout
	Val1 float32
	Val2 float32
	Val3 float32
	Val4 float32
}

// Rect represents the record "wasm:cv/types#rect".
//
// rect is a rectangle with integer coordinates.
// It is represented by the top-left corner and the bottom-right corner.
//
//	record rect {
//		min: size,
//		max: size,
//	}
type Rect struct {
	_   cm.HostLayout
	Min Size
	Max Size
}

// RGBA represents the record "wasm:cv/types#RGBA".
//
// RGBA is a color with red, green, blue, and alpha channels.
//
//	record RGBA {
//		r: u8,
//		g: u8,
//		b: u8,
//		a: u8,
//	}
type RGBA struct {
	_ cm.HostLayout
	R uint8
	G uint8
	B uint8
	A uint8
}

// BorderType represents the enum "wasm:cv/types#border-type".
//
// border-type is a type of border to add to an image.
//
//	enum border-type {
//		border-constant,
//		border-replicate,
//		border-reflect,
//		border-wrap,
//		border-reflect101,
//		border-transparent,
//		border-default,
//		border-isolated
//	}
type BorderType uint8

const (
	BorderTypeBorderConstant BorderType = iota
	BorderTypeBorderReplicate
	BorderTypeBorderReflect
	BorderTypeBorderWrap
	BorderTypeBorderReflect101
	BorderTypeBorderTransparent
	BorderTypeBorderDefault
	BorderTypeBorderIsolated
)

var stringsBorderType = [8]string{
	"border-constant",
	"border-replicate",
	"border-reflect",
	"border-wrap",
	"border-reflect101",
	"border-transparent",
	"border-default",
	"border-isolated",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e BorderType) String() string {
	return stringsBorderType[e]
}

// AdaptiveThresholdType represents the enum "wasm:cv/types#adaptive-threshold-type".
//
// adaptive-threshold-type is a type of adaptive thresholding.
//
//	enum adaptive-threshold-type {
//		adaptive-threshold-mean,
//		adaptive-threshold-gaussian
//	}
type AdaptiveThresholdType uint8

const (
	AdaptiveThresholdTypeAdaptiveThresholdMean AdaptiveThresholdType = iota
	AdaptiveThresholdTypeAdaptiveThresholdGaussian
)

var stringsAdaptiveThresholdType = [2]string{
	"adaptive-threshold-mean",
	"adaptive-threshold-gaussian",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e AdaptiveThresholdType) String() string {
	return stringsAdaptiveThresholdType[e]
}

// ThresholdType represents the enum "wasm:cv/types#threshold-type".
//
// threshold-type is a type of thresholding.
//
//	enum threshold-type {
//		threshold-binary,
//		threshold-binary-inv,
//		threshold-trunc,
//		threshold-to-zero,
//		threshold-to-zero-inv,
//		threshold-mask,
//		threshold-otsu,
//		tthreshold-triangle
//	}
type ThresholdType uint8

const (
	ThresholdTypeThresholdBinary ThresholdType = iota
	ThresholdTypeThresholdBinaryInv
	ThresholdTypeThresholdTrunc
	ThresholdTypeThresholdToZero
	ThresholdTypeThresholdToZeroInv
	ThresholdTypeThresholdMask
	ThresholdTypeThresholdOtsu
	ThresholdTypeTthresholdTriangle
)

var stringsThresholdType = [8]string{
	"threshold-binary",
	"threshold-binary-inv",
	"threshold-trunc",
	"threshold-to-zero",
	"threshold-to-zero-inv",
	"threshold-mask",
	"threshold-otsu",
	"tthreshold-triangle",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ThresholdType) String() string {
	return stringsThresholdType[e]
}

// DataLayoutType represents the enum "wasm:cv/types#data-layout-type".
//
// data-layout-type is a type of data layout.
//
//	enum data-layout-type {
//		data-layout-unknown,
//		data-layout-nd,
//		data-layout-nchw,
//		data-layout-ncdhw,
//		data-layout-nhwc,
//		data-layout-ndhwc,
//		data-layout-planar
//	}
type DataLayoutType uint8

const (
	DataLayoutTypeDataLayoutUnknown DataLayoutType = iota
	DataLayoutTypeDataLayoutNd
	DataLayoutTypeDataLayoutNchw
	DataLayoutTypeDataLayoutNcdhw
	DataLayoutTypeDataLayoutNhwc
	DataLayoutTypeDataLayoutNdhwc
	DataLayoutTypeDataLayoutPlanar
)

var stringsDataLayoutType = [7]string{
	"data-layout-unknown",
	"data-layout-nd",
	"data-layout-nchw",
	"data-layout-ncdhw",
	"data-layout-nhwc",
	"data-layout-ndhwc",
	"data-layout-planar",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e DataLayoutType) String() string {
	return stringsDataLayoutType[e]
}

// PaddingModeType represents the enum "wasm:cv/types#padding-mode-type".
//
//	enum padding-mode-type {
//		padding-mode-null,
//		padding-mode-crop-center,
//		padding-mode-letterbox
//	}
type PaddingModeType uint8

const (
	PaddingModeTypePaddingModeNull PaddingModeType = iota
	PaddingModeTypePaddingModeCropCenter
	PaddingModeTypePaddingModeLetterbox
)

var stringsPaddingModeType = [3]string{
	"padding-mode-null",
	"padding-mode-crop-center",
	"padding-mode-letterbox",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e PaddingModeType) String() string {
	return stringsPaddingModeType[e]
}

// BlobParams represents the record "wasm:cv/types#blob-params".
//
//	record blob-params {
//		scale-factor: f32,
//		size: size,
//		mean: scalar,
//		swap-RB: bool,
//		ddepth: u8,
//		data-layout: data-layout-type,
//		padding-mode: padding-mode-type,
//		border: scalar,
//	}
type BlobParams struct {
	_           cm.HostLayout
	ScaleFactor float32
	Size        Size
	Mean        Scalar
	SwapRB      bool
	Ddepth      uint8
	DataLayout  DataLayoutType
	PaddingMode PaddingModeType
	Border      Scalar
}

// MixMaxLocResult represents the record "wasm:cv/types#mix-max-loc-result".
//
//	record mix-max-loc-result {
//		min-val: f32,
//		max-val: f32,
//		min-loc: size,
//		max-loc: size,
//	}
type MixMaxLocResult struct {
	_      cm.HostLayout
	MinVal float32
	MaxVal float32
	MinLoc Size
	MaxLoc Size
}

// HersheyFontType represents the enum "wasm:cv/types#hershey-font-type".
//
//	enum hershey-font-type {
//		hershey-font-simplex,
//		hershey-font-plain,
//		hershey-font-duplex,
//		hershey-font-complex,
//		hershey-font-triplex,
//		hershey-font-complex-small,
//		hershey-font-script-simplex,
//		hershey-font-script-complex,
//		hershey-font-italic
//	}
type HersheyFontType uint8

const (
	HersheyFontTypeHersheyFontSimplex HersheyFontType = iota
	HersheyFontTypeHersheyFontPlain
	HersheyFontTypeHersheyFontDuplex
	HersheyFontTypeHersheyFontComplex
	HersheyFontTypeHersheyFontTriplex
	HersheyFontTypeHersheyFontComplexSmall
	HersheyFontTypeHersheyFontScriptSimplex
	HersheyFontTypeHersheyFontScriptComplex
	HersheyFontTypeHersheyFontItalic
)

var stringsHersheyFontType = [9]string{
	"hershey-font-simplex",
	"hershey-font-plain",
	"hershey-font-duplex",
	"hershey-font-complex",
	"hershey-font-triplex",
	"hershey-font-complex-small",
	"hershey-font-script-simplex",
	"hershey-font-script-complex",
	"hershey-font-italic",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e HersheyFontType) String() string {
	return stringsHersheyFontType[e]
}

// InterpolationType represents the enum "wasm:cv/types#interpolation-type".
//
//	enum interpolation-type {
//		interpolation-nearest,
//		interpolation-linear,
//		interpolation-cubic,
//		interpolation-area,
//		interpolation-lanczos4
//	}
type InterpolationType uint8

const (
	InterpolationTypeInterpolationNearest InterpolationType = iota
	InterpolationTypeInterpolationLinear
	InterpolationTypeInterpolationCubic
	InterpolationTypeInterpolationArea
	InterpolationTypeInterpolationLanczos4
)

var stringsInterpolationType = [5]string{
	"interpolation-nearest",
	"interpolation-linear",
	"interpolation-cubic",
	"interpolation-area",
	"interpolation-lanczos4",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e InterpolationType) String() string {
	return stringsInterpolationType[e]
}

// ColorCoversionType represents the enum "wasm:cv/types#color-coversion-type".
//
//	enum color-coversion-type {
//		color-BGR-to-BGRA,
//		color-RGB-to-RGBA,
//		color-BGRA-to-BGR,
//		color-RGBA-to-RGB,
//		color-BGR-to-RGBA,
//		color-RGB-to-BGRA,
//		color-RGBA-to-BGR,
//		color-BGRA-to-RGB,
//		color-BGR-to-RGB,
//		color-RGB-to-BGR,
//		color-BGRA-to-RGBA,
//		color-RGBA-to-BGRA,
//		color-BGR-to-gray,
//		color-RGB-to-gray,
//		color-gray-to-BGR,
//		color-gray-to-RGB,
//		color-gray-to-BGRA,
//		color-gray-to-RGBA,
//		color-BGRA-to-gray,
//		color-RGBA-to-gray
//	}
type ColorCoversionType uint8

const (
	ColorCoversionTypeColorBGRToBGRA ColorCoversionType = iota
	ColorCoversionTypeColorRGBToRGBA
	ColorCoversionTypeColorBGRAToBGR
	ColorCoversionTypeColorRGBAToRGB
	ColorCoversionTypeColorBGRToRGBA
	ColorCoversionTypeColorRGBToBGRA
	ColorCoversionTypeColorRGBAToBGR
	ColorCoversionTypeColorBGRAToRGB
	ColorCoversionTypeColorBGRToRGB
	ColorCoversionTypeColorRGBToBGR
	ColorCoversionTypeColorBGRAToRGBA
	ColorCoversionTypeColorRGBAToBGRA
	ColorCoversionTypeColorBGRToGray
	ColorCoversionTypeColorRGBToGray
	ColorCoversionTypeColorGrayToBGR
	ColorCoversionTypeColorGrayToRGB
	ColorCoversionTypeColorGrayToBGRA
	ColorCoversionTypeColorGrayToRGBA
	ColorCoversionTypeColorBGRAToGray
	ColorCoversionTypeColorRGBAToGray
)

var stringsColorCoversionType = [20]string{
	"color-BGR-to-BGRA",
	"color-RGB-to-RGBA",
	"color-BGRA-to-BGR",
	"color-RGBA-to-RGB",
	"color-BGR-to-RGBA",
	"color-RGB-to-BGRA",
	"color-RGBA-to-BGR",
	"color-BGRA-to-RGB",
	"color-BGR-to-RGB",
	"color-RGB-to-BGR",
	"color-BGRA-to-RGBA",
	"color-RGBA-to-BGRA",
	"color-BGR-to-gray",
	"color-RGB-to-gray",
	"color-gray-to-BGR",
	"color-gray-to-RGB",
	"color-gray-to-BGRA",
	"color-gray-to-RGBA",
	"color-BGRA-to-gray",
	"color-RGBA-to-gray",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ColorCoversionType) String() string {
	return stringsColorCoversionType[e]
}

// MorphShape represents the enum "wasm:cv/types#morph-shape".
//
//	enum morph-shape {
//		morph-rect,
//		morph-cross,
//		morph-ellipse
//	}
type MorphShape uint8

const (
	MorphShapeMorphRect MorphShape = iota
	MorphShapeMorphCross
	MorphShapeMorphEllipse
)

var stringsMorphShape = [3]string{
	"morph-rect",
	"morph-cross",
	"morph-ellipse",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e MorphShape) String() string {
	return stringsMorphShape[e]
}

// KeyPoint represents the record "wasm:cv/types#key-point".
//
//	record key-point {
//		x: f32,
//		y: f32,
//		size: f32,
//		angle: f32,
//		response: f32,
//		octave: s32,
//		class-id: s32,
//	}
type KeyPoint struct {
	_        cm.HostLayout
	X        float32
	Y        float32
	Size     float32
	Angle    float32
	Response float32
	Octave   int32
	ClassID  int32
}

// DMatch represents the record "wasm:cv/types#d-match".
//
// DMatch is data structure for matching keypoint descriptors.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d4/de0/classcv_1_1DMatch.html#a546ddb9a87898f06e510e015a6de596e
//
//	record d-match {
//		query-idx: u32,
//		train-idx: u32,
//		img-idx: u32,
//		distance: f64,
//	}
type DMatch struct {
	_        cm.HostLayout
	QueryIdx uint32
	TrainIdx uint32
	ImgIdx   uint32
	Distance float64
}
