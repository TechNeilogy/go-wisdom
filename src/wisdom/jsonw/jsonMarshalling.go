package jsonw

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

// JsonMarshalling shows typed JSON marshalling.
func JsonMarshalling() {
	// See ExampleJsonMarshalling.
}
