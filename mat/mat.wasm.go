// Code generated by wit-bindgen-go. DO NOT EDIT.

package mat

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasm:cv".

//go:wasmimport wasm:cv/mat [resource-drop]mat
//go:noescape
func wasmimport_MatResourceDrop(self0 uint32)

//go:wasmimport wasm:cv/mat [constructor]mat
//go:noescape
func wasmimport_NewMat(id0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [static]mat.new-with-size
//go:noescape
func wasmimport_MatNewWithSize(cols0 uint32, rows0 uint32, mattype0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.clone
//go:noescape
func wasmimport_MatClone(self0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.close
//go:noescape
func wasmimport_MatClose(self0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.col-range
//go:noescape
func wasmimport_MatColRange(self0 uint32, start0 uint32, end0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.cols
//go:noescape
func wasmimport_MatCols(self0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.convert-to
//go:noescape
func wasmimport_MatConvertTo(self0 uint32, mattype0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.copy-to
//go:noescape
func wasmimport_MatCopyTo(self0 uint32, dst0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.empty
//go:noescape
func wasmimport_MatEmpty(self0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.get-float-at
//go:noescape
func wasmimport_MatGetFloatAt(self0 uint32, row0 uint32, col0 uint32) (result0 float32)

//go:wasmimport wasm:cv/mat [method]mat.get-float-at3
//go:noescape
func wasmimport_MatGetFloatAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32) (result0 float32)

//go:wasmimport wasm:cv/mat [method]mat.get-int-at
//go:noescape
func wasmimport_MatGetIntAt(self0 uint32, row0 uint32, col0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.get-int-at3
//go:noescape
func wasmimport_MatGetIntAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.get-uchar-at
//go:noescape
func wasmimport_MatGetUcharAt(self0 uint32, row0 uint32, col0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.get-uchar-at3
//go:noescape
func wasmimport_MatGetUcharAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.get-vecb-at
//go:noescape
func wasmimport_MatGetVecbAt(self0 uint32, row0 uint32, col0 uint32, result *cm.List[uint8])

//go:wasmimport wasm:cv/mat [method]mat.get-vecf-at
//go:noescape
func wasmimport_MatGetVecfAt(self0 uint32, row0 uint32, col0 uint32, result *cm.List[float32])

//go:wasmimport wasm:cv/mat [method]mat.get-veci-at
//go:noescape
func wasmimport_MatGetVeciAt(self0 uint32, row0 uint32, col0 uint32, result *cm.List[int32])

//go:wasmimport wasm:cv/mat [method]mat.mattype
//go:noescape
func wasmimport_MatMattype(self0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.min-max-loc
//go:noescape
func wasmimport_MatMinMaxLoc(self0 uint32, result *MixMaxLocResult)

//go:wasmimport wasm:cv/mat [method]mat.region
//go:noescape
func wasmimport_MatRegion(self0 uint32, rect0 uint32, rect1 uint32, rect2 uint32, rect3 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.reshape
//go:noescape
func wasmimport_MatReshape(self0 uint32, channels0 uint32, rows0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.row-range
//go:noescape
func wasmimport_MatRowRange(self0 uint32, start0 uint32, end0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.rows
//go:noescape
func wasmimport_MatRows(self0 uint32) (result0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.set-float-at
//go:noescape
func wasmimport_MatSetFloatAt(self0 uint32, row0 uint32, col0 uint32, val0 float32)

//go:wasmimport wasm:cv/mat [method]mat.set-float-at3
//go:noescape
func wasmimport_MatSetFloatAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32, val0 float32)

//go:wasmimport wasm:cv/mat [method]mat.set-int-at
//go:noescape
func wasmimport_MatSetIntAt(self0 uint32, row0 uint32, col0 uint32, val0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.set-int-at3
//go:noescape
func wasmimport_MatSetIntAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32, val0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.set-uchar-at
//go:noescape
func wasmimport_MatSetUcharAt(self0 uint32, row0 uint32, col0 uint32, val0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.set-uchar-at3
//go:noescape
func wasmimport_MatSetUcharAt3(self0 uint32, x0 uint32, y0 uint32, z0 uint32, val0 uint32)

//go:wasmimport wasm:cv/mat [method]mat.size
//go:noescape
func wasmimport_MatSize(self0 uint32, result *cm.List[uint32])
