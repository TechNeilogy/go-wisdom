package main

import (
	"github.com/TechNeilogy/go-wisdom/src/wisdom/errorw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/functionalw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/jsonw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/nongow"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/structw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/syncw"
)

func main() {

	errorw.RunCustomErrors(false)

	functionalw.RunFunctionalWisdom(false)

	jsonw.RunJsonAsInterface(false)
	jsonw.RunJsonMarshalling(false)
	jsonw.RunJsonMarshallingUnexported(false)

	nongow.RunScanset(true)

	structw.RunOopThisSemantics(false)
	structw.RunStringStructPenalty(false)

	syncw.RunFirstComeFirstServedGoroutinesVariadic(false)
	syncw.RunPool0(false)
	syncw.RunSyncWaitGroup(false)
}
