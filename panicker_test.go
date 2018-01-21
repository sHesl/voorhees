package voorhees

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanickerAdd(t *testing.T) {
	path := "added"
	value := "excellent"
	input := map[string]interface{}{
		"keepMe": "please",
	}

	expected := map[string]interface{}{
		"keepMe": "please",
		"added":  "excellent",
	}

	result := NewPanickerVoorhees(input).Add(path, value)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected .Add(%s, %s) to correctly add property '%s' to map with value: %s",
			path, value, finalPropertyOfPath(path), value)
	}

	if reflect.DeepEqual(result, input) {
		t.Errorf("Expected .Add() to not to modify input map")
	}
}

func TestPanickerChange(t *testing.T) {
	path := "layer1.changeMe"
	value := "changed"
	input := map[string]interface{}{
		"keepMe": "please",
		"layer1": map[string]interface{}{
			"changeMe": 123,
		},
	}

	expected := map[string]interface{}{
		"keepMe": "please",
		"layer1": map[string]interface{}{
			"changeMe": "changed",
		},
	}

	result := NewPanickerVoorhees(input).Change(path, value)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected .Change(%s, %s) to correctly change property '%s' to value: %s",
			path, value, finalPropertyOfPath(path), value)
	}

	if reflect.DeepEqual(result, input) {
		t.Errorf("Expected .Change() to not to modify input map")
	}
}

func TestPanickerDelete(t *testing.T) {
	path := "layer1.deleteMe"
	input := map[string]interface{}{
		"keepMe": "please",
		"layer1": map[string]interface{}{
			"deleteMe": 123,
		},
	}

	expected := map[string]interface{}{
		"keepMe": "please",
		"layer1": map[string]interface{}{},
	}

	result := NewPanickerVoorhees(input).Delete(path)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected .Delete(%s) to correctly delete property '%s'",
			path, finalPropertyOfPath(path))
	}

	if reflect.DeepEqual(result, input) {
		t.Errorf("Expected .Delete() to not to modify input map")
	}
}

func TestPanicsOnInvalidJSON(t *testing.T) {
	expected := "interface conversion: interface {} is nil, not map[string]interface {}"

	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			assert.Equal(t, expected, err.Error())
		}
	}()

	NewPanickerVoorhees(nil)
}

func TestPanickerArrayIsNotPresentInPath(t *testing.T) {
	expected := "[Voorhees]: Unable to navigate to notreal[0].uhoh. Failed to find node: notreal[0]"

	testCase := map[string]interface{}{
		"integer": 123,
	}

	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			assert.Equal(t, expected, err.Error())
		}
	}()

	pv := NewPanickerVoorhees(testCase)
	pv.Change("notreal[0].uhoh.added", "excellent")
}

func TestPanickerChangeNonexistantProperty(t *testing.T) {
	expected := "[Voorhees]: Unable to change changeMe because it doesn't exist at path layer1.layer2"

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe": "please",
			},
		},
	}

	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			assert.Equal(t, expected, err.Error())
		}
	}()

	pv := NewPanickerVoorhees(testCase)
	pv.Change("layer1.layer2.changeMe", "x")
}

func TestPanickerDeleteInvalidPath(t *testing.T) {
	expected := "[Voorhees]: Unable to navigate to layer1.layer2.uhoh. Failed to find node: uhoh"

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe":   "please",
				"deleteMe": "bye",
			},
		},
	}

	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			assert.Equal(t, expected, err.Error())
		}
	}()

	pv := NewPanickerVoorhees(testCase)
	pv.Delete("layer1.layer2.uhoh.deleteMe")
}
