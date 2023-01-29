package main

import (
	"github.com/TechNeilogy/go-wisdom/src/wisdom/errorw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/functionalw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/jsonw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/structw"
)

func main() {

	structw.RunStringStructPenalty(false)
	structw.RunOopThisSemantics(true)
	errorw.RunCustomErrors(true)
	jsonw.RunJsonAsInterface(true)
	functionalw.RunFunctionalWisdom(true)
}
