// Code generated by wit-bindgen-go. DO NOT EDIT.

package objdetect

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasm:cv".

//go:wasmimport wasm:cv/objdetect [resource-drop]cascade-classifier
//go:noescape
func wasmimport_CascadeClassifierResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/objdetect [constructor]cascade-classifier
//go:noescape
func wasmimport_NewCascadeClassifier(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.close
//go:noescape
func wasmimport_CascadeClassifierClose(self0 uint32)

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.detect-multi-scale
//go:noescape
func wasmimport_CascadeClassifierDetectMultiScale(self0 uint32, image0 uint32, result *cm.List[Rect])

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.detect-multi-scale-with-params
//go:noescape
func wasmimport_CascadeClassifierDetectMultiScaleWithParams(self0 uint32, image0 uint32, scale0 float64, minNeighbors0 uint32, flags0 uint32, minSize0 uint32, minSize1 uint32, maxSize0 uint32, maxSize1 uint32, result *cm.List[Rect])

//go:wasmimport wasm:cv/objdetect [method]cascade-classifier.load
//go:noescape
func wasmimport_CascadeClassifierLoad(self0 uint32, file0 *uint8, file1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [resource-drop]HOG-descriptor
//go:noescape
func wasmimport_HOGDescriptorResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/objdetect [constructor]HOG-descriptor
//go:noescape
func wasmimport_NewHOGDescriptor(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]HOG-descriptor.close
//go:noescape
func wasmimport_HOGDescriptorClose(self0 uint32)

//go:wasmimport wasm:cv/objdetect [method]HOG-descriptor.detect-multi-scale
//go:noescape
func wasmimport_HOGDescriptorDetectMultiScale(self0 uint32, image0 uint32, result *cm.List[Rect])

//go:wasmimport wasm:cv/objdetect [method]HOG-descriptor.detect-multi-scale-with-params
//go:noescape
func wasmimport_HOGDescriptorDetectMultiScaleWithParams(self0 uint32, image0 uint32, hitThreshold0 float64, winStride0 uint32, winStride1 uint32, padding0 uint32, padding1 uint32, scale0 float64, finalThreshold0 float64, useMeanshiftGrouping0 uint32, result *cm.List[Rect])

//go:wasmimport wasm:cv/objdetect [resource-drop]face-detector-YN
//go:noescape
func wasmimport_FaceDetectorYNResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/objdetect [constructor]face-detector-YN
//go:noescape
func wasmimport_NewFaceDetectorYN(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32, inputSize0 uint32, inputSize1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [static]face-detector-YN.new-with-params
//go:noescape
func wasmimport_FaceDetectorYNNewWithParams(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32, inputSize0 uint32, inputSize1 uint32, scoreThreshold0 float32, nmsThreshold0 float32, topK0 uint32, backendId0 uint32, targetId0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.close
//go:noescape
func wasmimport_FaceDetectorYNClose(self0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.detect
//go:noescape
func wasmimport_FaceDetectorYNDetect(self0 uint32, input0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-input-size
//go:noescape
func wasmimport_FaceDetectorYNGetInputSize(self0 uint32, result *Size)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-nms-threshold
//go:noescape
func wasmimport_FaceDetectorYNGetNmsThreshold(self0 uint32) (result0 float32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-score-threshold
//go:noescape
func wasmimport_FaceDetectorYNGetScoreThreshold(self0 uint32) (result0 float32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.get-topk
//go:noescape
func wasmimport_FaceDetectorYNGetTopk(self0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-input-size
//go:noescape
func wasmimport_FaceDetectorYNSetInputSize(self0 uint32, size0 uint32, size1 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-nms-threshold
//go:noescape
func wasmimport_FaceDetectorYNSetNmsThreshold(self0 uint32, threshold0 float32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-score-threshold
//go:noescape
func wasmimport_FaceDetectorYNSetScoreThreshold(self0 uint32, threshold0 float32)

//go:wasmimport wasm:cv/objdetect [method]face-detector-YN.set-topk
//go:noescape
func wasmimport_FaceDetectorYNSetTopk(self0 uint32, topk0 uint32)

//go:wasmimport wasm:cv/objdetect [resource-drop]face-recognizer-SF
//go:noescape
func wasmimport_FaceRecognizerSFResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/objdetect [constructor]face-recognizer-SF
//go:noescape
func wasmimport_NewFaceRecognizerSF(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [static]face-recognizer-SF.new-with-params
//go:noescape
func wasmimport_FaceRecognizerSFNewWithParams(model0 *uint8, model1 uint32, config0 *uint8, config1 uint32, backendId0 uint32, targetId0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.align-crop
//go:noescape
func wasmimport_FaceRecognizerSFAlignCrop(self0 uint32, src0 uint32, faceBox0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.close
//go:noescape
func wasmimport_FaceRecognizerSFClose(self0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.feature
//go:noescape
func wasmimport_FaceRecognizerSFFeature(self0 uint32, aligned0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.match
//go:noescape
func wasmimport_FaceRecognizerSFMatch(self0 uint32, face10 uint32, face20 uint32) (result0 float32)

//go:wasmimport wasm:cv/objdetect [method]face-recognizer-SF.match-with-params
//go:noescape
func wasmimport_FaceRecognizerSFMatchWithParams(self0 uint32, face10 uint32, face20 uint32, distance0 uint32) (result0 float32)