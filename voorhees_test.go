package voorhees

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	type testCase struct {
		path     string
		value    interface{}
		input    map[string]interface{}
		expected map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			"added",
			"excellent",
			map[string]interface{}{
				"keepMe": "please",
			},
			map[string]interface{}{
				"keepMe": "please",
				"added":  "excellent",
			},
		},

		testCase{
			".added",
			"excellent",
			map[string]interface{}{
				"keepMe": "please",
			},
			map[string]interface{}{
				"keepMe": "please",
				"added":  "excellent",
			},
		},

		testCase{
			"layer1.added",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"added":  "excellent",
				},
			},
		},

		testCase{
			"layer1.layer2.added",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"layer2": map[string]interface{}{
						"keepMe": "please",
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"layer2": map[string]interface{}{
						"keepMe": "please",
						"added":  "excellent",
					},
				},
			},
		},

		testCase{
			"layer1.array[0].added",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []map[string]interface{}{
						map[string]interface{}{
							"keepMe": "please",
						},
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []interface{}{
						map[string]interface{}{
							"keepMe": "please",
							"added":  "excellent",
						},
					},
				},
			},
		},

		testCase{
			"layer1.array[0].layer2.added",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []map[string]interface{}{
						map[string]interface{}{
							"keepMe": "please",
							"layer2": map[string]interface{}{
								"keepMe": "please",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []interface{}{
						map[string]interface{}{
							"keepMe": "please",
							"layer2": map[string]interface{}{
								"keepMe": "please",
								"added":  "excellent",
							},
						},
					},
				},
			},
		},

		testCase{
			"layer1.addLayer.added",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"addLayer": map[string]interface{}{
						"added": "excellent",
					},
				},
			},
		},

		testCase{
			"addedArray[0].added",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{},
				"addedArray": []interface{}{
					map[string]interface{}{
						"added": "excellent",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		result := NewVoorhees(testCase.input).Add(testCase.path, testCase.value)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected .Add(%s, %s) to correctly add property '%s' to map with value: %s",
				testCase.path, testCase.value, finalPropertyOfPath(testCase.path), testCase.value)
		}

		if reflect.DeepEqual(result, testCase) {
			t.Errorf("Expected .Add() to not to modify input map")
		}
	}
}

func Test_AddThen(t *testing.T) {
	input := map[string]interface{}{
		"keepMe": "please",
	}

	a := NewVoorhees(input).Add("addMe", "excellent")
	b := NewVoorhees(input).AddThen("addMe", "excellent").JSON

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected .Add() and .AddThen() to produce the same result with the same input.")
	}
}

func Test_Change(t *testing.T) {
	type testCase struct {
		path     string
		value    interface{}
		input    map[string]interface{}
		expected map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			"changeMe",
			"excellent",
			map[string]interface{}{
				"keepMe":   "please",
				"changeMe": "please",
			},
			map[string]interface{}{
				"keepMe":   "please",
				"changeMe": "excellent",
			},
		},

		testCase{
			"layer1.changeMe",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe":   "please",
					"changeMe": "please",
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe":   "please",
					"changeMe": "excellent",
				},
			},
		},

		testCase{
			"layer1.layer2.changeMe",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"layer2": map[string]interface{}{
						"keepMe":   "please",
						"changeMe": "please",
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"layer2": map[string]interface{}{
						"keepMe":   "please",
						"changeMe": "excellent",
					},
				},
			},
		},

		testCase{
			"layer1.array[0].changeMe",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []map[string]interface{}{
						map[string]interface{}{
							"keepMe":   "please",
							"changeMe": "please",
						},
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []interface{}{
						map[string]interface{}{
							"keepMe":   "please",
							"changeMe": "excellent",
						},
					},
				},
			},
		},
		testCase{
			"layer1.array[0].layer2.changeMe",
			"excellent",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []map[string]interface{}{
						map[string]interface{}{
							"keepMe": "please",
							"layer2": map[string]interface{}{
								"keepMe":   "please",
								"changeMe": "please",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []interface{}{
						map[string]interface{}{
							"keepMe": "please",
							"layer2": map[string]interface{}{
								"keepMe":   "please",
								"changeMe": "excellent",
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		result := NewVoorhees(testCase.input).Change(testCase.path, testCase.value)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected .Change(%s, %s) to correctly change property '%s' to value: %s",
				testCase.path, testCase.value, finalPropertyOfPath(testCase.path), testCase.value)
		}

		if reflect.DeepEqual(result, testCase) {
			t.Errorf("Expected .Change() to not to modify input map")
		}
	}
}

func Test_ChangeThen(t *testing.T) {
	input := map[string]interface{}{
		"changeMe": "please",
	}

	a := NewVoorhees(input).Change("changeMe", "excellent")
	b := NewVoorhees(input).ChangeThen("changeMe", "excellent").JSON

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected .Change() and .ChangeThen() to produce the same result with the same input.")
	}
}

func Test_Delete(t *testing.T) {
	type testCase struct {
		path     string
		input    map[string]interface{}
		expected map[string]interface{}
	}

	testCases := []testCase{
		testCase{
			"deleteMe",
			map[string]interface{}{
				"keepMe":   "please",
				"deleteMe": "please",
			},
			map[string]interface{}{
				"keepMe": "please",
			},
		},

		testCase{
			"layer1.deleteMe",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe":   "please",
					"deleteMe": "please",
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
				},
			},
		},

		testCase{
			"layer1.layer2.deleteMe",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"layer2": map[string]interface{}{
						"keepMe":   "please",
						"deleteMe": "please",
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"layer2": map[string]interface{}{
						"keepMe": "please",
					},
				},
			},
		},

		testCase{
			"layer1.array[0].deleteMe",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []map[string]interface{}{
						map[string]interface{}{
							"keepMe":   "please",
							"deleteMe": "please",
						},
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []interface{}{
						map[string]interface{}{
							"keepMe": "please",
						},
					},
				},
			},
		},
		testCase{
			"layer1.array[0].layer2.deleteMe",
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []map[string]interface{}{
						map[string]interface{}{
							"keepMe": "please",
							"layer2": map[string]interface{}{
								"keepMe":   "please",
								"deleteMe": "please",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"layer1": map[string]interface{}{
					"keepMe": "please",
					"array": []interface{}{
						map[string]interface{}{
							"keepMe": "please",
							"layer2": map[string]interface{}{
								"keepMe": "please",
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		result := NewVoorhees(testCase.input).Delete(testCase.path)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected .Delete(%s) to correctly delete property '%s'",
				testCase.path, finalPropertyOfPath(testCase.path))
		}

		if reflect.DeepEqual(result, testCase) {
			t.Errorf("Expected .Delete() to not to modify input map")
		}
	}
}

func Test_DeleteThen(t *testing.T) {
	input := map[string]interface{}{
		"deleteMe": "please",
	}

	a := NewVoorhees(input).Delete("deleteMe")
	b := NewVoorhees(input).DeleteThen("deleteMe").JSON

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected .Delete() and .DeleteThen() to produce the same result with the same input.")
	}
}

func Test_PropertyInPathIsNotAMap(t *testing.T) {
	expectedPanicMessage := "Voorhees: Add | Unable to navigate to integer.uhoh. Node: integer was not a map[string]interface{}"

	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			if expectedPanicMessage != s {
				t.Errorf("Expected any panic occuring inside Add() to produce message %s. Got: %s",
					expectedPanicMessage, s)
			}
		} else {
			t.Errorf("Test_PropertyInPathIsNotAMap panicked but did could not recover. Test setup is broken!")
		}
	}()

	input := map[string]interface{}{
		"integer": 123,
	}

	NewVoorhees(input).Add("integer.uhoh.added", "excellent")
}

func Test_ArrayIsNotPresentInPath(t *testing.T) {
	expectedPanicMessage := "Voorhees: Change | Unable to navigate to notreal[0].uhoh. Failed to find node: notreal[0]"

	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			if expectedPanicMessage != s {
				t.Errorf("Expected any panic occuring inside Change() to produce message %s. Got: %s",
					expectedPanicMessage, s)
			}
		} else {
			t.Errorf("Test_ArrayIsNotPresentInPath panicked but did could not recover. Test setup is broken!")
		}
	}()

	input := map[string]interface{}{
		"integer": 123,
	}

	NewVoorhees(input).Change("notreal[0].uhoh.added", "excellent")
}

func Test_Change_InvalidPath(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			if !strings.HasPrefix(s, "Voorhees: Change |") {
				t.Errorf("Expected any panic occuring inside Change() to be proper Voorhees formatted panics")
			}
		}
	}()

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe": "please",
			},
		},
	}

	NewVoorhees(testCase).Change("layer1.uhoh.layer2", "x")
}

func Test_Change_NonexistantProperty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			if !strings.HasPrefix(s, "Voorhees: Change |") {
				t.Errorf("Expected any panic occuring inside Change() to be proper Voorhees formatted panics")
			}
		}
	}()

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe": "please",
			},
		},
	}

	NewVoorhees(testCase).Change("layer1.layer2.changeMe", "x")
}

func Test_Delete_InvalidPath(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			s := r.(string)
			if !strings.HasPrefix(s, "Voorhees: Delete |") {
				t.Errorf("Expected any panic occuring inside Delete() to be proper Voorhees formatted panics")
			}
		}
	}()

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe":   "please",
				"deleteMe": "bye",
			},
		},
	}

	NewVoorhees(testCase).Delete("layer1.layer2.uhoh.deleteMe")
}

func Test_finalPropertyOfPath(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}
	testCases := []testCase{
		testCase{"prop", "prop"},
		testCase{"layer1.prop", "prop"},
		testCase{"layer1.layer2.prop", "prop"},
		testCase{"layer1.array[0].prop", "prop"},
		testCase{"array[0].prop", "prop"},
	}

	for _, testCase := range testCases {
		result := finalPropertyOfPath(testCase.input)

		if result != testCase.expected {
			t.Errorf("Expected finalPropertyOfPath(\"%s\") to return \"%s\". Got: %s",
				testCase.input, testCase.expected, result)
		}
	}
}

func Test_trimLastPropertyFromPath(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}
	testCases := []testCase{
		testCase{"prop", "prop"},
		testCase{"layer1.prop", "layer1"},
		testCase{"layer1.layer2.prop", "layer1.layer2"},
		testCase{"layer1.array[0].prop", "layer1.array[0]"},
		testCase{"array[0].prop", "array[0]"},
	}

	for _, testCase := range testCases {
		result := trimLastPropertyFromPath(testCase.input)

		if result != testCase.expected {
			t.Errorf("Expected trimLastPropertyFromPath(\"%s\") to return \"%s\". Got: %s",
				testCase.input, testCase.expected, result)
		}
	}
}

func Test_deconstructArrayPath(t *testing.T) {
	type testCase struct {
		s             string
		expectedName  string
		expectedIndex int
	}

	testCases := []testCase{
		testCase{"array[0]", "array", 0},
		testCase{"array[4]", "array", 4},
		testCase{"array[14]", "array", 14},
	}

	for _, testCase := range testCases {
		name, index := deconstructArrayPath(testCase.s)

		if name != testCase.expectedName {
			t.Errorf("Expected deconstructArrayPath(\"%s\") to return name \"%s\". Got: %s",
				testCase.s, testCase.expectedName, name)
		}

		if index != testCase.expectedIndex {
			t.Errorf("Expected deconstructArrayPath(\"%s\") to return index \"%d\". Got: %d",
				testCase.s, testCase.expectedIndex, index)
		}
	}
}

func Test_deconstructArrayPath_panic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			s := r.(interface{})
			assert.Equal(t, "Voorhees: Array Path | array[not good] is not a valid array denotion", s)
		}
	}()

	deconstructArrayPath("array[not good]")
}
