package main

import (
	"fmt"

	vw "github.com/travisbrady/govw"
)

// Lots of goofing around
// Initially intended to match the Python test.py
// From https://github.com/JohnLangford/vowpal_wabbit/blob/master/python/test.py
func main() {
	h := vw.Initialize("-q st --noconstant --quiet")
	//example := ReadExample(vw, "1 |s p^the_man w^the w^man |t p^un_homme w^un w^homme")
	example := vw.ReadExample(h, "1 |x a b")
	score := vw.Learn(h, example)
	ex2 := vw.ReadExample(h, "1 |x a b |y c")
	vw.Learn(h, ex2)
	fmt.Printf("[1 Learns] Pred = %v\n", ex2.GetPrediction())
	vw.Learn(h, ex2)
	fmt.Printf("[2 Learns] Pred = %v\n", ex2.GetPrediction())
	vw.Learn(h, ex2)
	fmt.Printf("[3 Learns] Pred = %v\n", ex2.GetPrediction())
	vw.Learn(h, ex2)
	fmt.Printf("[4 Learns] Pred = %v\n", ex2.GetPrediction())
	pred := ex2.GetPrediction()
	imp := ex2.GetImportance()
	featHash := h.HashFeature("a", 42)
	fmt.Printf("Num Features: %v\n", ex2.GetNumFeatures())
	fmt.Printf("Loss: %v\n", ex2.GetLoss())
	fmt.Printf("GetSimpleLabelPrediction: %v\n", ex2.GetSimpleLabelPrediction())
	fmt.Printf("GetSimpleLabelLabel: %v\n", ex2.GetSimpleLabelLabel())
	//fmt.Printf("GetUpdatedPrediction: %v\n", ex2.GetUpdatedPrediction())
	fmt.Printf("GetSumLoss: %v\n", h.GetSumLoss())
	h.Finish()
	fmt.Printf("Score = %v Pred: %v Importance: %v FeatHash: %v\n", score, pred, imp, featHash)
}
