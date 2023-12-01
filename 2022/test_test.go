package _022

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type SomeStruct struct {
	Field1 int32
	Nested NestedStruct
	Field2 float32
}

type NestedStruct struct {
	Nester string
}

func (x *SomeStruct) Decode(d *json.Decoder, path string) ([]string, error) {
	token, err := d.Token()
	if err != nil {
		return nil, err
	}
	if delim, ok := token.(json.Delim); !ok || delim != '{' {
		return nil, fmt.Errorf("unexpected token expected \"{\" got %q", token)
	}

	field1Found := false
	nestedFound := false
	field2Found := false

	missingFields := make([]string, 0)
	earlyExit := false
	for {
		token, err = d.Token()
		if err != nil {
			return nil, err
		}

		if delim, ok := token.(json.Delim); ok && delim == '}' {
			earlyExit = true
			break
		}

		str, ok := token.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected token %T with value %v", token, token)
		}

		switch str {
		case "field1":
			if field1Found {
				return nil, fmt.Errorf("duplicate field: field1")
			}

			err = d.Decode(&x.Field1)
			if err != nil {
				return nil, err
			}

			field1Found = true
		case "nested":
			if nestedFound {
				return nil, fmt.Errorf("duplicate field: nested")
			}

			otherMissingFields, err := x.Nested.Decode(d, path+"/nested")
			if err != nil {
				return nil, err
			}

			missingFields = append(missingFields, otherMissingFields...)
			nestedFound = true
		case "field2":
			if field2Found {
				return nil, fmt.Errorf("duplicate field: field2")
			}

			err = d.Decode(&x.Field2)
			if err != nil {
				return nil, err
			}

			field2Found = true
		default:
			// extra field, skip
			err = skip(d)
			if err != nil {
				return nil, err
			}
		}
	}

	if !field1Found {
		missingFields = append(missingFields, path+"/"+"field1")
	}
	if !nestedFound {
		missingFields = append(missingFields, path+"/"+"nested")
	}
	if !field2Found {
		missingFields = append(missingFields, path+"/"+"field2")
	}

	if !earlyExit {
		token, err = d.Token()
		if err != nil {
			return nil, err
		}
	}

	return missingFields, nil
}

func (x *NestedStruct) Decode(d *json.Decoder, path string) ([]string, error) {
	token, err := d.Token()
	if err != nil {
		return nil, err
	}
	if delim, ok := token.(json.Delim); !ok || delim != '{' {
		return nil, fmt.Errorf("unexpected token expected \"{\" got %q", token)
	}

	nesterFound := false

	missingFields := make([]string, 0)
	earlyExit := false
	for {
		token, err = d.Token()
		if err != nil {
			return nil, err
		}

		if delim, ok := token.(json.Delim); ok && delim == '}' {
			earlyExit = true
			break
		}

		str, ok := token.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected token %T with value %v", token, token)
		}

		switch str {
		case "nester":
			if nesterFound {
				return nil, fmt.Errorf("duplicate field: nester")
			}

			err = d.Decode(&x.Nester)
			if err != nil {
				return nil, err
			}

			nesterFound = true
		}

	}

	if !nesterFound {
		missingFields = append(missingFields, path+"/"+"nester")
	}

	if !earlyExit {
		token, err = d.Token()
		if err != nil {
			return nil, err
		}
	}

	return missingFields, nil
}

func skip(d *json.Decoder) error {
	square := 0
	curly := 0

	for {
		token, err := d.Token()
		if err != nil {
			return nil
		}

		switch token.(type) {
		case json.Delim:
			delim := token.(json.Delim)
			switch delim {
			case '{':
				curly++
			case '}':
				curly--
			case '[':
				square++
			case ']':
				square--
			}
		}

		if square == 0 && curly == 0 {
			break
		}
	}

	return nil
}

func Test(t *testing.T) {
	x := SomeStruct{}

	d := json.NewDecoder(strings.NewReader(`{"field1": 1234, "nested": {"nester": "1234"}, "extra": {}, "field2": 321}`))
	missingFields, err := x.Decode(d, "")
	if err != nil {
		t.Errorf("error: %v", err.Error())
	}

	t.Logf("value: %+v", x)
	t.Logf("value: %+v", missingFields)
}
