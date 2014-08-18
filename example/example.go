package main

import (
	"fmt"

	vw "github.com/travisbrady/go-vw"
)

func main() {
	h := vw.Initialize("-q st --noconstant --quiet")
	//example := ReadExample(vw, "1 |s p^the_man w^the w^man |t p^un_homme w^un w^homme")
	example := vw.ReadExample(h, "1 |x a b")
	score := vw.Learn(h, example)
	ex2 := vw.ReadExample(h, "1 |x a b |y c")
	vw.Learn(h, ex2)
	fmt.Printf("[1 Learns] Pred = %v\n", vw.GetPrediction(ex2))
	vw.Learn(h, ex2)
	fmt.Printf("[2 Learns] Pred = %v\n", vw.GetPrediction(ex2))
	vw.Learn(h, ex2)
	fmt.Printf("[3 Learns] Pred = %v\n", vw.GetPrediction(ex2))
	vw.Learn(h, ex2)
	fmt.Printf("[4 Learns] Pred = %v\n", vw.GetPrediction(ex2))
	pred := vw.GetPrediction(ex2)
	imp := vw.GetImportance(ex2)
	featHash := vw.HashFeature(h, "a", 42)
	fmt.Printf("Num Features: %v\n", vw.GetNumFeatures(ex2))
	fmt.Printf("Loss: %v\n", vw.GetLoss(ex2))
	fmt.Printf("GetSimpleLabelPrediction: %v\n", vw.GetSimpleLabelPrediction(ex2))
	fmt.Printf("GetSimpleLabelLabel: %v\n", vw.GetSimpleLabelLabel(ex2))
	fmt.Printf("GetUpdatedPrediction: %v\n", vw.GetUpdatedPrediction(ex2))
	fmt.Printf("GetSumLoss: %v\n", vw.GetSumLoss(h))
	vw.Finish(h)
	fmt.Printf("Score = %v Pred: %v Importance: %v FeatHash: %v\n", score, pred, imp, featHash)
}
