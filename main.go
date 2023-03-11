package main

import (
	"github.com/TechNeilogy/go-wisdom/src/wisdom/errorw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/functionalw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/jsonw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/structw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/syncw"
)

func main() {

	structw.RunStringStructPenalty(false)
	structw.RunOopThisSemantics(false)

	errorw.RunCustomErrors(false)

	jsonw.RunJsonAsInterface(false)

	functionalw.RunFunctionalWisdom(false)

	syncw.RunSyncWaitGroup(false)
	syncw.RunFirstComeFirstServedGoroutinesVariadic(false)
	syncw.RunPool0(true)
}
