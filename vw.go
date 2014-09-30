package vw

//#cgo CFLAGS: -I . -I/usr/local/include -lstdc++ -O2
//#cgo LDFLAGS: -L. -I . -I/usr/local/include -lvw -lstdc++
//#include "stubs.h"
import "C"
import "unsafe"

type Handle struct {
	c_handle C.VW_HANDLE
}

type Example struct {
	c_example C.VW_EXAMPLE
}

func Initialize(args string) *Handle {
	cs := C.CString(args)
	defer C.free(unsafe.Pointer(cs))
	c_h := C.VW_InitializeA(cs)
	h := Handle{c_h}
	return &h
}

func ReadExample(handle *Handle, example string) *Example {
	cs := C.CString(example)
	defer C.free(unsafe.Pointer(cs))
	c_ex := C.VW_ReadExampleA(handle.c_handle, cs)
	ex := Example{c_ex}
	return &ex
}

func EmptyExample(handle *Handle) *Example {
	c_ex := C.VW_EmptyExample(handle.c_handle)
	ex := Example{c_ex}
	return &ex
}

func (example *Example) ExamplePushFeature(ns int, fid int, v float64) {
	C.VW_ExamplePushFeature(example.c_example, C.uchar(ns), C.uint32_t(fid), C.float(v))
}

func (example *Example) FinishExample(handle *Handle) {
	C.VW_FinishExample(handle.c_handle, example.c_example)
}

func Learn(handle *Handle, example *Example) float64 {
	return float64(C.VW_Learn(handle.c_handle, example.c_example))
}

func (handle *Handle) LearnString(exStr string) float64 {
	ex := ReadExample(handle, exStr)
	C.VW_Learn(handle.c_handle, ex.c_example)
	pp := float64(C.VW_GetPartialPrediction(ex.c_example))
	ex.FinishExample(handle)
	return pp
}

func (example *Example) GetPrediction() float64 {
	return float64(C.VW_GetPrediction(example.c_example))
}

func (example *Example) GetPartialPrediction() float64 {
	return float64(C.VW_GetPartialPrediction(example.c_example))
}

func (example *Example) GetSimpleLabelPrediction() float64 {
	return float64(C.VW_GetSimpleLabelPrediction(example.c_example))
}

func (example *Example) GetSimpleLabelLabel() float64 {
	return float64(C.VW_GetSimpleLabel(example.c_example))
}

func (handle *Handle) GetPredictionString(exStr string) float64 {
	ex := ReadExample(handle, exStr)
	//return GetPrediction(ReadExample(handle, exStr))
	return ex.GetPrediction()
}

func (example *Example) GetImportance() float64 {
	return float64(C.VW_GetImportance(example.c_example))
}

func (handle *Handle) HashFeature(s string, u int64) int {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return int(C.VW_HashFeatureA(handle.c_handle, cs, C.ulong(u)))
}

func (example *Example) GetNumFeatures() int {
	return int(C.VW_GetNumFeatures(example.c_example))
}

func (example *Example) GetLoss() float64 {
	return float64(C.VW_GetLoss(example.c_example))
}

func (handle *Handle) Finish() {
	C.VW_Finish(handle.c_handle)
}

/*
func GetTopicPrediction(example C.VW_EXAMPLE, i int) float64 {
	return float64(C.VW_GetTopicPrediction(example, C.size_t(i)))
}
*/

func (handle *Handle) GetSumLoss() float64 {
	return float64(C.VW_GetSumLoss(handle.c_handle))
}
