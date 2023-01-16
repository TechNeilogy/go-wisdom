package jsonarray

import (
	"encoding/json"
	"fmt"
)

func Run() {

	var v interface{}

	err := json.Unmarshal([]byte("[\"a\", \"b\", 32, [\"c\", true]]"), &v)

	if err != nil {
		panic(err)
	}

	// A tree of interfaces and values.
	fmt.Printf("%v\n", v)

	// Pull out a known value type.
	z0 := v.([]interface{})
	z1 := z0[3].([]interface{})
	z2 := z1[0].(string)
	fmt.Printf("%v\n", z2)

	// Switch on an unknown value type.
	switch x := z0[2].(type) {
	case float64:
		fmt.Printf("float64: %v\n", x)
	default:
		fmt.Printf("Unexpected type: %T\n", x)
	}
}
