// Code generated by wit-bindgen-go. DO NOT EDIT.

package dnn

import (
	"go.bytecodealliance.org/cm"
	"wasmcv.org/wasm/cv/types"
)

func lower_Size(v types.Size) (f0 uint32, f1 uint32) {
	f0 = (uint32)(v.X)
	f1 = (uint32)(v.Y)
	return
}

func lower_Scalar(v types.Scalar) (f0 float32, f1 float32, f2 float32, f3 float32) {
	f0 = (float32)(v.Val1)
	f1 = (float32)(v.Val2)
	f2 = (float32)(v.Val3)
	f3 = (float32)(v.Val4)
	return
}

func lower_BlobParams(v types.BlobParams) (f0 float32, f1 uint32, f2 uint32, f3 float32, f4 float32, f5 float32, f6 float32, f7 uint32, f8 uint32, f9 uint32, f10 uint32, f11 float32, f12 float32, f13 float32, f14 float32) {
	f0 = (float32)(v.ScaleFactor)
	f1, f2 = lower_Size(v.Size)
	f3, f4, f5, f6 = lower_Scalar(v.Mean)
	f7 = cm.BoolToU32(v.SwapRB)
	f8 = (uint32)(v.Ddepth)
	f9 = (uint32)(v.DataLayout)
	f10 = (uint32)(v.PaddingMode)
	f11, f12, f13, f14 = lower_Scalar(v.Border)
	return
}
