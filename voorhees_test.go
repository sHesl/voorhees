package voorhees

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
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
		result, err := NewVoorhees(testCase.input).Add(testCase.path, testCase.value)

		assert.NoError(t, err)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected .Add(%s, %s) to correctly add property '%s' to map with value: %s",
				testCase.path, testCase.value, finalPropertyOfPath(testCase.path), testCase.value)
		}

		if reflect.DeepEqual(result, testCase.input) {
			t.Errorf("Expected .Add() to not to modify input map")
		}
	}
}

func TestChange(t *testing.T) {
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
		result, err := NewVoorhees(testCase.input).Change(testCase.path, testCase.value)

		assert.NoError(t, err)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected .Change(%s, %s) to correctly change property '%s' to value: %s",
				testCase.path, testCase.value, finalPropertyOfPath(testCase.path), testCase.value)
		}

		if reflect.DeepEqual(result, testCase.input) {
			t.Errorf("Expected .Change() to not to modify input map")
		}
	}
}

func TestDelete(t *testing.T) {
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
		result, err := NewVoorhees(testCase.input).Delete(testCase.path)

		assert.NoError(t, err)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Expected .Delete(%s) to correctly delete property '%s'",
				testCase.path, finalPropertyOfPath(testCase.path))
		}

		if reflect.DeepEqual(result, testCase.input) {
			t.Errorf("Expected .Delete() to not to modify input map")
		}
	}
}

func TestPropertyInPathIsNotAMap(t *testing.T) {
	expected := "[Voorhees]: Unable to navigate to integer.uhoh. Node: integer was not a map[string]interface{}"

	testCase := map[string]interface{}{
		"integer": 123,
	}

	result, err := NewVoorhees(testCase).Add("integer.uhoh.added", "excellent")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestArrayIsNotPresentInPath(t *testing.T) {
	expected := "[Voorhees]: Unable to navigate to notreal[0].uhoh. Failed to find node: notreal[0]"

	testCase := map[string]interface{}{
		"integer": 123,
	}

	result, err := NewVoorhees(testCase).Change("notreal[0].uhoh.added", "excellent")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestChangeInvalidPath(t *testing.T) {
	expected := "[Voorhees]: Unable to navigate to layer1.uhoh. Failed to find node: uhoh"

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe": "please",
			},
		},
	}

	result, err := NewVoorhees(testCase).Change("layer1.uhoh.layer2", "x")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestChangeNonexistantProperty(t *testing.T) {
	expected := "[Voorhees]: Unable to change changeMe because it doesn't exist at path layer1.layer2"

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe": "please",
			},
		},
	}

	result, err := NewVoorhees(testCase).Change("layer1.layer2.changeMe", "x")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestDeleteInvalidPath(t *testing.T) {
	expected := "[Voorhees]: Unable to navigate to layer1.layer2.uhoh. Failed to find node: uhoh"

	testCase := map[string]interface{}{
		"layer1": map[string]interface{}{
			"layer2": map[string]interface{}{
				"keepMe":   "please",
				"deleteMe": "bye",
			},
		},
	}

	result, err := NewVoorhees(testCase).Delete("layer1.layer2.uhoh.deleteMe")

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestFinalPropertyOfPath(t *testing.T) {
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

		assert.Equal(t, testCase.expected, result,
			"Expected finalPropertyOfPath(\"%s\") to return \"%s\". Got: %s",
			testCase.input, testCase.expected, result)
	}
}

func TestTrimLastPropertyFromPath(t *testing.T) {
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

		assert.Equal(t, testCase.expected, result,
			"Expected trimLastPropertyFromPath(\"%s\") to return \"%s\". Got: %s",
			testCase.input, testCase.expected, result)
	}
}

func TestDeconstructArrayPath(t *testing.T) {
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
		name, index, err := deconstructArrayPath(testCase.s)

		assert.NoError(t, err)

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

func TestDeconstructArrayPathPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			s := r.(interface{})
			assert.Equal(t, "[Voorhees]: Array Path | array[not good] is not a valid array denotion", s)
		}
	}()

	deconstructArrayPath("array[not good]")
}
