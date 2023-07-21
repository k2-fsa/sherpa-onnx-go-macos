//go:build darwin && arm64 && !ios

package sherpa_onnx

// #cgo LDFLAGS: -L ${SRCDIR}/lib/aarch64-apple-darwin -lsherpa-onnx-c-api -lsherpa-onnx-core -lkaldi-native-fbank-core -lonnxruntime -Wl,-rpath,${SRCDIR}/lib/aarch64-apple-darwin
import "C"
