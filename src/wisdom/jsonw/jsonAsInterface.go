package jsonw

import (
	"encoding/json"
	"fmt"
	"github.com/TechNeilogy/go-wisdom/src/util"
)

func RunJsonAsInterface() {

	util.PrintHeader("JSON as Raw Interfaces")

	var v interface{}

	err := json.Unmarshal([]byte("[\"+\", \"b\", 32, [\"and\", \"c\", true]]"), &v)

	if err != nil {
		panic(err)
	}

	// A tree of interfaces and values.
	fmt.Printf("Raw Interfaces: %v\n", v)

	// Pull out a known value type.
	j := v.([]interface{})
	j3 := j[3].([]interface{})
	for i, z := range []interface{}{
		j3[0].(string),
		j3[1].(string),
		j3[2].(bool),
	} {
		fmt.Printf("json[3][%v].(%T): %v\n", i, z, z)
	}

	if j32, ok := j3[1].(string); ok {
		fmt.Printf("data, ok := json[3][1].(%T): %v, %v\n", j32, j32, ok)
	}

	if j32, ok := j3[1].(int); !ok {
		fmt.Printf("data, ok := json[3][1].(%T): %v, %v\n", j32, j32, ok)
	}

	// Switch on an unknown value type.
	for i, z := range j {
		switch x := z.(type) {
		case float64:
			fmt.Printf("switch json[%v] case %T: %v\n", i, x, x)
		case string:
			fmt.Printf("switch json[%v] case %T: %v\n", i, x, x)
		default:
			fmt.Printf("switch json[%v] default: Unmapped type: %T\n", i, x)
		}
	}
}
