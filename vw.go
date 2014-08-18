package vw

//#cgo CFLAGS: -I . -I/usr/local/include -lstdc++
//#cgo LDFLAGS: -L. -I . -I/usr/local/include -lvw -lstdc++
//#include "stubs.h"
import "C"
import "unsafe"

//func Initialize(args string) unsafe.Pointer {
func Initialize(args string) C.VW_HANDLE {
	cs := C.CString(args)
	defer C.free(unsafe.Pointer(cs))
	return C.VW_InitializeA(cs)
}

func ReadExample(handle C.VW_HANDLE, example string) C.VW_EXAMPLE {
	cs := C.CString(example)
	defer C.free(unsafe.Pointer(cs))
	return C.VW_ReadExampleA(handle, cs)
}

func FinishExample(handle C.VW_HANDLE, example C.VW_EXAMPLE) {
	C.VW_FinishExample(handle, example)
}

func Learn(handle C.VW_HANDLE, example C.VW_EXAMPLE) float64 {
	return float64(C.VW_Learn(handle, example))
}

func LearnString(handle C.VW_HANDLE, exStr string) float64 {
	ex := ReadExample(handle, exStr)
	C.VW_Learn(handle, ex)
	pp := float64(C.VW_GetPartialPrediction(ex))
	FinishExample(handle, ex)
	return pp
}

func GetPrediction(example C.VW_EXAMPLE) float64 {
	return float64(C.VW_GetPrediction(example))
}

func GetPartialPrediction(example C.VW_EXAMPLE) float64 {
	return float64(C.VW_GetPartialPrediction(example))
}

func GetSimpleLabelPrediction(example C.VW_EXAMPLE) float64 {
	return float64(C.VW_GetSimpleLabelPrediction(example))
}

func GetSimpleLabelLabel(example C.VW_EXAMPLE) float64 {
	return float64(C.VW_GetSimpleLabel(example))
}

func GetPredictionString(handle C.VW_HANDLE, exStr string) float64 {
	return GetPrediction(ReadExample(handle, exStr))
}

func GetImportance(example C.VW_EXAMPLE) float64 {
	return float64(C.VW_GetImportance(example))
}

func HashFeature(handle C.VW_HANDLE, s string, u int64) int {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return int(C.VW_HashFeatureA(handle, cs, C.ulong(u)))
}

func GetNumFeatures(example C.VW_EXAMPLE) int {
	return int(C.VW_GetNumFeatures(example))
}

func GetLoss(example C.VW_EXAMPLE) float64 {
	return float64(C.VW_GetLoss(example))
}

func GetUpdatedPrediction(example C.VW_EXAMPLE) float64 {
	return float64(C.VW_GetUpdatedPrediction(example))
}

func Finish(handle C.VW_HANDLE) {
	C.VW_Finish(handle)
}

func GetTopicPrediction(example C.VW_EXAMPLE, i int) float64 {
	return float64(C.VW_GetTopicPrediction(example, C.size_t(i)))
}

func GetSumLoss(handle C.VW_HANDLE) float64 {
	return float64(C.VW_GetSumLoss(handle))
}
