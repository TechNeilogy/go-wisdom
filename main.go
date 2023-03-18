package main

import (
	"github.com/TechNeilogy/go-wisdom/src/wisdom/jsonw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/structw"
	"github.com/TechNeilogy/go-wisdom/src/wisdom/syncw"
)

func main() {

	jsonw.RunJsonAsInterface(false)
	jsonw.RunJsonMarshalling(false)
	jsonw.RunJsonMarshallingUnexported(false)

	structw.RunOopThisSemantics(false)
	structw.RunStringStructPenalty(true)

	syncw.RunFirstComeFirstServedGoroutinesVariadic(false)
	syncw.RunPool0(false)
	syncw.RunSyncWaitGroup(false)
}
