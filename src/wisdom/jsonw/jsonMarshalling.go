package jsonw

import "encoding/json"

type InnerMessage struct {
	Tag0 string `json:"taggymctagface"`
	Tag1 string
}

type OuterMessage struct {
	Name  string
	Body  string
	Time  int64
	Inner InnerMessage
}

func makeSample() OuterMessage {
	return OuterMessage{
		"Fred",
		"Flintstone",
		1294706395881547000,
		InnerMessage{
			"t0",
			"t1",
		},
	}
}

func RunJsonMarshalling(run bool) {
	if !run {
		return
	}

	sample0 := makeSample()

	b, err := json.Marshal(sample0)

	println(string(b), err)

	var sample1 OuterMessage

	err = json.Unmarshal(b, &sample1)

	println(sample1.Inner.Tag0, err)

	var sample2 interface{}

	err = json.Unmarshal(b, &sample2)

	println(sample2, err)
}
