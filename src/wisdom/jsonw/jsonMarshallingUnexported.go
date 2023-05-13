package jsonw

import "encoding/json"

// Source:
// https://stackoverflow.com/questions/11126793/json-and-dealing-with-unexported-fields

type jsonData struct {
	Field1 string `json:"some_field"`
	Field2 string `json:"some_other_field"`
}

type JsonData struct {
	jsonData
}

// Implement json.Unmarshaller
func (d *JsonData) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &d.jsonData)
}

// Getter
func (d *JsonData) Field1() string {
	return d.jsonData.Field1
}

type NoExportedFields struct {
	someField int
}

func (nef *NoExportedFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SomeField int `json:"some_field"`
	}{
		SomeField: nef.someField,
	})
}

func (nef *NoExportedFields) UnmarshalJSON(b []byte) error {
	var x map[string]interface{}
	err := json.Unmarshal(b, &x)
	if err == nil {
		nef.someField = int(x["some_field"].(float64))
	}
	return err
}

func RunJsonMarshallingUnexported(run bool) {
	if !run {
		return
	}

	sample0 := jsonData{
		"field1",
		"field2",
	}

	sample1 := JsonData{
		sample0,
	}

	b, err := json.Marshal(sample1)

	println(string(b), err)

	b = []byte(`{"some_field":"field1","some_other_field":"field22"}`)

	err = sample1.UnmarshalJSON(b)

	println(sample1.Field2)

	we := NoExportedFields{23}

	b, err = we.MarshalJSON()

	println(string(b), err)

	s := `{"some_field":1000} `

	we2 := NoExportedFields{2322}

	println(we2.someField)

	err = json.Unmarshal([]byte(s), &we2)

	println(we2.someField, err)

	err = we2.UnmarshalJSON([]byte(s))

	println(we2.someField, err)
}
