package jsonw

import (
	"encoding/json"
	"fmt"
)

func ExampleJsonAsInterface() {

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

	// output: Raw Interfaces: [+ b 32 [and c true]]
	//json[3][0].(string): and
	//json[3][1].(string): c
	//json[3][2].(bool): true
	//data, ok := json[3][1].(string): c, true
	//data, ok := json[3][1].(int): 0, false
	//switch json[0] case string: +
	//switch json[1] case string: b
	//switch json[2] case float64: 32
	//switch json[3] default: Unmapped type: []interface {}
}

func ExampleJsonMarshalling() {

	sample0 := makeSample()

	b, err := json.Marshal(sample0)

	fmt.Println(string(b), err)

	var sample1 OuterMessage

	err = json.Unmarshal(b, &sample1)

	fmt.Println(sample1.Inner.Tag0, err)

	var sample2 interface{}

	err = json.Unmarshal(b, &sample2)

	fmt.Println(sample2, err)

	// output: {"Name":"Fred","Body":"Flintstone","Time":1294706395881547000,"Inner":{"taggymctagface":"t0","Tag1":"t1"}} <nil>
	//t0 <nil>
	//map[Body:Flintstone Inner:map[Tag1:t1 taggymctagface:t0] Name:Fred Time:1.294706395881547e+18] <nil>
}
