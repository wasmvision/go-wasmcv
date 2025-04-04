// Code generated by wit-bindgen-go. DO NOT EDIT.

package features2d

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasm:cv".

//go:wasmimport wasm:cv/features2d [resource-drop]AKAZE-detector
//go:noescape
func wasmimport_AKAZEDetectorResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/features2d [constructor]AKAZE-detector
//go:noescape
func wasmimport_NewAKAZEDetector(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [method]AKAZE-detector.close
//go:noescape
func wasmimport_AKAZEDetectorClose(self0 uint32)

//go:wasmimport wasm:cv/features2d [method]AKAZE-detector.compute
//go:noescape
func wasmimport_AKAZEDetectorCompute(self0 uint32, src0 uint32, mask0 uint32, kps0 *KeyPoint, kps1 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [method]AKAZE-detector.detect
//go:noescape
func wasmimport_AKAZEDetectorDetect(self0 uint32, src0 uint32, result *cm.Result[cm.List[KeyPoint], cm.List[KeyPoint], ErrorResult])

//go:wasmimport wasm:cv/features2d [method]AKAZE-detector.detect-and-compute
//go:noescape
func wasmimport_AKAZEDetectorDetectAndCompute(self0 uint32, src0 uint32, mask0 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [resource-drop]BRISK-detector
//go:noescape
func wasmimport_BRISKDetectorResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/features2d [constructor]BRISK-detector
//go:noescape
func wasmimport_NewBRISKDetector(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [method]BRISK-detector.close
//go:noescape
func wasmimport_BRISKDetectorClose(self0 uint32)

//go:wasmimport wasm:cv/features2d [method]BRISK-detector.compute
//go:noescape
func wasmimport_BRISKDetectorCompute(self0 uint32, src0 uint32, mask0 uint32, kps0 *KeyPoint, kps1 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [method]BRISK-detector.detect
//go:noescape
func wasmimport_BRISKDetectorDetect(self0 uint32, src0 uint32, result *cm.Result[cm.List[KeyPoint], cm.List[KeyPoint], ErrorResult])

//go:wasmimport wasm:cv/features2d [method]BRISK-detector.detect-and-compute
//go:noescape
func wasmimport_BRISKDetectorDetectAndCompute(self0 uint32, src0 uint32, mask0 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [resource-drop]KAZE-detector
//go:noescape
func wasmimport_KAZEDetectorResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/features2d [constructor]KAZE-detector
//go:noescape
func wasmimport_NewKAZEDetector(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [method]KAZE-detector.close
//go:noescape
func wasmimport_KAZEDetectorClose(self0 uint32)

//go:wasmimport wasm:cv/features2d [method]KAZE-detector.compute
//go:noescape
func wasmimport_KAZEDetectorCompute(self0 uint32, src0 uint32, mask0 uint32, kps0 *KeyPoint, kps1 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [method]KAZE-detector.detect
//go:noescape
func wasmimport_KAZEDetectorDetect(self0 uint32, src0 uint32, result *cm.Result[cm.List[KeyPoint], cm.List[KeyPoint], ErrorResult])

//go:wasmimport wasm:cv/features2d [method]KAZE-detector.detect-and-compute
//go:noescape
func wasmimport_KAZEDetectorDetectAndCompute(self0 uint32, src0 uint32, mask0 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [resource-drop]ORB-detector
//go:noescape
func wasmimport_ORBDetectorResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/features2d [constructor]ORB-detector
//go:noescape
func wasmimport_NewORBDetector(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [static]ORB-detector.new-with-params
//go:noescape
func wasmimport_ORBDetectorNewWithParams(features0 uint32, scale0 float32, levels0 uint32, edgeThreshold0 uint32, first0 uint32, wtak0 uint32, scoreType0 uint32, patchSize0 uint32, fastThreshold0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [method]ORB-detector.close
//go:noescape
func wasmimport_ORBDetectorClose(self0 uint32)

//go:wasmimport wasm:cv/features2d [method]ORB-detector.compute
//go:noescape
func wasmimport_ORBDetectorCompute(self0 uint32, src0 uint32, mask0 uint32, kps0 *KeyPoint, kps1 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [method]ORB-detector.detect
//go:noescape
func wasmimport_ORBDetectorDetect(self0 uint32, src0 uint32, result *cm.Result[cm.List[KeyPoint], cm.List[KeyPoint], ErrorResult])

//go:wasmimport wasm:cv/features2d [method]ORB-detector.detect-and-compute
//go:noescape
func wasmimport_ORBDetectorDetectAndCompute(self0 uint32, src0 uint32, mask0 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [resource-drop]SIFT-detector
//go:noescape
func wasmimport_SIFTDetectorResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/features2d [constructor]SIFT-detector
//go:noescape
func wasmimport_NewSIFTDetector(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [method]SIFT-detector.close
//go:noescape
func wasmimport_SIFTDetectorClose(self0 uint32)

//go:wasmimport wasm:cv/features2d [method]SIFT-detector.compute
//go:noescape
func wasmimport_SIFTDetectorCompute(self0 uint32, src0 uint32, mask0 uint32, kps0 *KeyPoint, kps1 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [method]SIFT-detector.detect
//go:noescape
func wasmimport_SIFTDetectorDetect(self0 uint32, src0 uint32, result *cm.Result[cm.List[KeyPoint], cm.List[KeyPoint], ErrorResult])

//go:wasmimport wasm:cv/features2d [method]SIFT-detector.detect-and-compute
//go:noescape
func wasmimport_SIFTDetectorDetectAndCompute(self0 uint32, src0 uint32, mask0 uint32, result *cm.Result[DetectorResultShape, DetectorResult, ErrorResult])

//go:wasmimport wasm:cv/features2d [resource-drop]BF-matcher
//go:noescape
func wasmimport_BFMatcherResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/features2d [constructor]BF-matcher
//go:noescape
func wasmimport_NewBFMatcher(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [static]BF-matcher.new-with-params
//go:noescape
func wasmimport_BFMatcherNewWithParams(norm0 uint32, crossCheck0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [method]BF-matcher.KNN-match
//go:noescape
func wasmimport_BFMatcherKNNMatch(self0 uint32, query0 uint32, train0 uint32, k0 uint32, result *cm.Result[cm.List[cm.List[DMatch]], cm.List[cm.List[DMatch]], ErrorResult])

//go:wasmimport wasm:cv/features2d [method]BF-matcher.close
//go:noescape
func wasmimport_BFMatcherClose(self0 uint32)

//go:wasmimport wasm:cv/features2d [method]BF-matcher.match
//go:noescape
func wasmimport_BFMatcherMatch(self0 uint32, query0 uint32, train0 uint32, result *cm.Result[cm.List[DMatch], cm.List[DMatch], ErrorResult])

//go:wasmimport wasm:cv/features2d [resource-drop]flann-based-matcher
//go:noescape
func wasmimport_FlannBasedMatcherResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/features2d [constructor]flann-based-matcher
//go:noescape
func wasmimport_NewFlannBasedMatcher(name0 *uint8, name1 uint32) (result0 uint32)

//go:wasmimport wasm:cv/features2d [method]flann-based-matcher.KNN-match
//go:noescape
func wasmimport_FlannBasedMatcherKNNMatch(self0 uint32, query0 uint32, train0 uint32, k0 uint32, result *cm.Result[cm.List[cm.List[DMatch]], cm.List[cm.List[DMatch]], ErrorResult])

//go:wasmimport wasm:cv/features2d [method]flann-based-matcher.close
//go:noescape
func wasmimport_FlannBasedMatcherClose(self0 uint32)
