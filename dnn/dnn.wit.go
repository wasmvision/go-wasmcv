// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package dnn represents the imported interface "wasm:cv/dnn".
package dnn

import (
	"go.bytecodealliance.org/cm"
	"wasmcv.org/wasm/cv/mat"
	"wasmcv.org/wasm/cv/types"
)

// Mat represents the imported type alias "wasm:cv/dnn#mat".
//
// See [mat.Mat] for more information.
type Mat = mat.Mat

// Size represents the type alias "wasm:cv/dnn#size".
//
// See [types.Size] for more information.
type Size = types.Size

// Scalar represents the type alias "wasm:cv/dnn#scalar".
//
// See [types.Scalar] for more information.
type Scalar = types.Scalar

// Rect represents the type alias "wasm:cv/dnn#rect".
//
// See [types.Rect] for more information.
type Rect = types.Rect

// BlobParams represents the type alias "wasm:cv/dnn#blob-params".
//
// See [types.BlobParams] for more information.
type BlobParams = types.BlobParams

// DataLayoutType represents the type alias "wasm:cv/dnn#data-layout-type".
//
// See [types.DataLayoutType] for more information.
type DataLayoutType = types.DataLayoutType

// PaddingModeType represents the type alias "wasm:cv/dnn#padding-mode-type".
//
// See [types.PaddingModeType] for more information.
type PaddingModeType = types.PaddingModeType

// NetBackendType represents the enum "wasm:cv/dnn#net-backend-type".
//
//	enum net-backend-type {
//		net-backend-default,
//		net-backend-halide,
//		net-backend-openvino,
//		net-backend-opencv,
//		net-backend-vkcom,
//		net-backend-cuda
//	}
type NetBackendType uint8

const (
	NetBackendTypeNetBackendDefault NetBackendType = iota
	NetBackendTypeNetBackendHalide
	NetBackendTypeNetBackendOpenvino
	NetBackendTypeNetBackendOpencv
	NetBackendTypeNetBackendVkcom
	NetBackendTypeNetBackendCuda
)

var stringsNetBackendType = [6]string{
	"net-backend-default",
	"net-backend-halide",
	"net-backend-openvino",
	"net-backend-opencv",
	"net-backend-vkcom",
	"net-backend-cuda",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e NetBackendType) String() string {
	return stringsNetBackendType[e]
}

// NetTargetType represents the enum "wasm:cv/dnn#net-target-type".
//
//	enum net-target-type {
//		net-target-cpu,
//		net-target-fp32,
//		net-target-fp16,
//		net-target-vpu,
//		net-target-vulkan,
//		net-target-fpga,
//		net-target-cuda,
//		net-target-cuda-fp16
//	}
type NetTargetType uint8

const (
	NetTargetTypeNetTargetCPU NetTargetType = iota
	NetTargetTypeNetTargetFp32
	NetTargetTypeNetTargetFp16
	NetTargetTypeNetTargetVpu
	NetTargetTypeNetTargetVulkan
	NetTargetTypeNetTargetFpga
	NetTargetTypeNetTargetCuda
	NetTargetTypeNetTargetCudaFp16
)

var stringsNetTargetType = [8]string{
	"net-target-cpu",
	"net-target-fp32",
	"net-target-fp16",
	"net-target-vpu",
	"net-target-vulkan",
	"net-target-fpga",
	"net-target-cuda",
	"net-target-cuda-fp16",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e NetTargetType) String() string {
	return stringsNetTargetType[e]
}

// Layer represents the imported resource "wasm:cv/dnn#layer".
//
//	resource layer
type Layer cm.Resource

// ResourceDrop represents the imported resource-drop for resource "layer".
//
// Drops a resource handle.
//
//go:nosplit
func (self Layer) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_LayerResourceDrop((uint32)(self0))
	return
}

// NewLayer represents the imported constructor for resource "layer".
//
//	constructor()
//
//go:nosplit
func NewLayer() (result Layer) {
	result0 := wasmimport_NewLayer()
	result = cm.Reinterpret[Layer]((uint32)(result0))
	return
}

// GetName represents the imported method "get-name".
//
// GetName returns the name of the layer.
//
//	get-name: func() -> string
//
//go:nosplit
func (self Layer) GetName() (result string) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_LayerGetName((uint32)(self0), &result)
	return
}

// Net represents the imported resource "wasm:cv/dnn#net".
//
//	resource net
type Net cm.Resource

// ResourceDrop represents the imported resource-drop for resource "net".
//
// Drops a resource handle.
//
//go:nosplit
func (self Net) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetResourceDrop((uint32)(self0))
	return
}

// NetRead represents the imported static function "read".
//
// ReadNet read deep learning network represented in one of the supported formats.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga138439da76f26266fdefec9723f6c5cd
//
//	read: static func(model: string, config: string) -> net
//
//go:nosplit
func NetRead(model string, config string) (result Net) {
	model0, model1 := cm.LowerString(model)
	config0, config1 := cm.LowerString(config)
	result0 := wasmimport_NetRead((*uint8)(model0), (uint32)(model1), (*uint8)(config0), (uint32)(config1))
	result = cm.Reinterpret[Net]((uint32)(result0))
	return
}

// NetReadFromONNX represents the imported static function "read-from-ONNX".
//
// ReadNetFromONNX reads a network model stored in ONNX framework's format.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga9198ecaac7c32ddf0aa7a1bcbd359567
//
//	read-from-ONNX: static func(model: string) -> net
//
//go:nosplit
func NetReadFromONNX(model string) (result Net) {
	model0, model1 := cm.LowerString(model)
	result0 := wasmimport_NetReadFromONNX((*uint8)(model0), (uint32)(model1))
	result = cm.Reinterpret[Net]((uint32)(result0))
	return
}

// Close represents the imported method "close".
//
// Close the network
//
//	close: func()
//
//go:nosplit
func (self Net) Close() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetClose((uint32)(self0))
	return
}

// Empty represents the imported method "empty".
//
// Empty returns true if there are no layers in the network.
//
// For further details, please see:
// https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#a6a5778787d5b8770deab5eda6968e66c
//
//	empty: func() -> bool
//
//go:nosplit
func (self Net) Empty() (result bool) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_NetEmpty((uint32)(self0))
	result = cm.U32ToBool((uint32)(result0))
	return
}

// Forward represents the imported method "forward".
//
// Forward runs forward pass to compute output of layer with name outputName.
//
// For further details, please see:
// https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a98ed94cb6ef7063d3697259566da310b
//
//	forward: func(output-name: string) -> mat
//
//go:nosplit
func (self Net) Forward(outputName string) (result Mat) {
	self0 := cm.Reinterpret[uint32](self)
	outputName0, outputName1 := cm.LowerString(outputName)
	result0 := wasmimport_NetForward((uint32)(self0), (*uint8)(outputName0), (uint32)(outputName1))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

// ForwardLayers represents the imported method "forward-layers".
//
// ForwardLayers forward pass to compute outputs of layers listed in outBlobNames.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#afe22e099b60a2883e220645391f68d4c
//
//	forward-layers: func(output-names: list<string>) -> list<mat>
//
//go:nosplit
func (self Net) ForwardLayers(outputNames cm.List[string]) (result cm.List[Mat]) {
	self0 := cm.Reinterpret[uint32](self)
	outputNames0, outputNames1 := cm.LowerList(outputNames)
	wasmimport_NetForwardLayers((uint32)(self0), (*string)(outputNames0), (uint32)(outputNames1), &result)
	return
}

// GetLayer represents the imported method "get-layer".
//
// GetLayer returns layer with specified id.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#ac944d7f2d3ead5ef9b1b2fa3885f3ff1
//
//	get-layer: func(id: u32) -> layer
//
//go:nosplit
func (self Net) GetLayer(id uint32) (result Layer) {
	self0 := cm.Reinterpret[uint32](self)
	id0 := (uint32)(id)
	result0 := wasmimport_NetGetLayer((uint32)(self0), (uint32)(id0))
	result = cm.Reinterpret[Layer]((uint32)(result0))
	return
}

// GetLayerNames represents the imported method "get-layer-names".
//
// GetLayerNames returns names of layers in the network.
//
// For further details, please see:
// hhttps://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#a38e67098ae4ae5906bf8d8ea72199c2e
//
//	get-layer-names: func() -> list<string>
//
//go:nosplit
func (self Net) GetLayerNames() (result cm.List[string]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetGetLayerNames((uint32)(self0), &result)
	return
}

// GetUnconnectedOutLayers represents the imported method "get-unconnected-out-layers".
//
// GetUnconnectedOutLayers returns indexes of layers with unconnected outputs.
//
// For further details, please see:
// https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#ae26f0c29b3733d15d0482098ef9053e3
//
//	get-unconnected-out-layers: func() -> list<u32>
//
//go:nosplit
func (self Net) GetUnconnectedOutLayers() (result cm.List[uint32]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_NetGetUnconnectedOutLayers((uint32)(self0), &result)
	return
}

// SetInput represents the imported method "set-input".
//
// SetInput sets the new input value for the network.
//
// For further details, please see:
// https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a672a08ae76444d75d05d7bfea3e4a328
//
//	set-input: func(input: mat, name: string)
//
//go:nosplit
func (self Net) SetInput(input Mat, name string) {
	self0 := cm.Reinterpret[uint32](self)
	input0 := cm.Reinterpret[uint32](input)
	name0, name1 := cm.LowerString(name)
	wasmimport_NetSetInput((uint32)(self0), (uint32)(input0), (*uint8)(name0), (uint32)(name1))
	return
}

// BlobFromImage represents the imported function "blob-from-image".
//
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
func BlobFromImage(image Mat, scaleFactor float32, size Size, mean Scalar, swapRb bool, crop bool) (result Mat) {
	image0 := cm.Reinterpret[uint32](image)
	scaleFactor0 := (float32)(scaleFactor)
	size0, size1 := lower_Size(size)
	mean0, mean1, mean2, mean3 := lower_Scalar(mean)
	swapRb0 := cm.BoolToU32(swapRb)
	crop0 := cm.BoolToU32(crop)
	result0 := wasmimport_BlobFromImage((uint32)(image0), (float32)(scaleFactor0), (uint32)(size0), (uint32)(size1), (float32)(mean0), (float32)(mean1), (float32)(mean2), (float32)(mean3), (uint32)(swapRb0), (uint32)(crop0))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

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
func BlobFromImageWithParams(image Mat, params BlobParams) (result Mat) {
	image0 := cm.Reinterpret[uint32](image)
	params0, params1, params2, params3, params4, params5, params6, params7, params8, params9, params10, params11, params12, params13, params14 := lower_BlobParams(params)
	result0 := wasmimport_BlobFromImageWithParams((uint32)(image0), (float32)(params0), (uint32)(params1), (uint32)(params2), (float32)(params3), (float32)(params4), (float32)(params5), (float32)(params6), (uint32)(params7), (uint32)(params8), (uint32)(params9), (uint32)(params10), (float32)(params11), (float32)(params12), (float32)(params13), (float32)(params14))
	result = cm.Reinterpret[Mat]((uint32)(result0))
	return
}

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
func BlobRectsToImageRects(params BlobParams, blobRects cm.List[Rect], imageSize Size) (result cm.List[Rect]) {
	params_ := wasmimport_BlobRectsToImageRects_params{params: params, blobRects: blobRects, imageSize: imageSize}
	wasmimport_BlobRectsToImageRects(&params_, &result)
	return
}

// wasmimport_BlobRectsToImageRects_params represents the flattened function params for [wasmimport_BlobRectsToImageRects].
// See the Canonical ABI flattening rules for more information.
type wasmimport_BlobRectsToImageRects_params struct {
	_         cm.HostLayout
	params    BlobParams
	blobRects cm.List[Rect]
	imageSize Size
}

// NMSBoxes represents the imported function "NMS-boxes".
//
// NMSBoxes performs non maximum suppression given boxes and corresponding scores.
//
// For futher details, please see:
// https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee
//
//	NMS-boxes: func(bboxes: list<rect>, scores: list<f32>, score-threshold: f32, nms-threshold:
//	f32) -> list<s32>
//
//go:nosplit
func NMSBoxes(bboxes cm.List[Rect], scores cm.List[float32], scoreThreshold float32, nmsThreshold float32) (result cm.List[int32]) {
	bboxes0, bboxes1 := cm.LowerList(bboxes)
	scores0, scores1 := cm.LowerList(scores)
	scoreThreshold0 := (float32)(scoreThreshold)
	nmsThreshold0 := (float32)(nmsThreshold)
	wasmimport_NMSBoxes((*Rect)(bboxes0), (uint32)(bboxes1), (*float32)(scores0), (uint32)(scores1), (float32)(scoreThreshold0), (float32)(nmsThreshold0), &result)
	return
}
