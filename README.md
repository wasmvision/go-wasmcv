# wasmCV

Generated Go bindings for wasmCV WebAssembly interfaces to computer vision systems.

See https://github.com/wasmvision/wasmcv for information about wasmCV.

## How to use

This example Go module exports a `process()` function to the WASM host application. When the host calls the processor, it passes in the wasmCV image `Mat` to be processed. The wasmCV module then calls functions on that `Mat` which are handled by the host application, by calling OpenCV to actually perform the operations.

```go
package main

import (
	"github.com/hybridgroup/mechanoid/convert"
	"wasmcv.org/wasm/cv/mat"
)

//go:wasmimport hosted println
func println(ptr, size uint32)

//export process
func process(image mat.Mat) mat.Mat {
	println(convert.StringToWasmPtr("Cols: " +
		convert.IntToString(int(image.Cols())) +
		" Rows: " +
		convert.IntToString(int(image.Rows())) +
		" Type: " +
		convert.IntToString(int(image.Type()))))

	return image
}

func main() {}
```

Install the `wasmcv` package into your Go package:

```shell
go get wasmcv.org/wasm/cv
```

You can then compile this module using the TinyGo compiler.

```shell
tinygo build -o processor.wasm -target=wasm-unknown processor.go
```

Note that the `wasm-unknown` target can be used with wasmCV to produce very lightweight guest modules. The example above compiles to around 31k, including debug information.

```shell
-rwxrwxr-x 1 ron ron 31248 sep 11 11:00 processor.wasm
```
