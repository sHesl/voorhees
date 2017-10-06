package voorhees

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// Voorhees allows json path denoted manipulations of a map[string]interface{}
//
// Examples:
// NewVoorhees(myMap).Add("firstLayer.nestedObj.anArray[1].newField", "new val")
// NewVoorhees(myMap).Change("firstLayer.nestedObj.aProperty", "new val")
// NewVoorhees(myMap).Delete("firstLayer.deleteMe")
//
// Voorhees will panic if an invalid json path is provided as the initial parameter
// i.e if a property that is requested does not exist, so it goes without saying that this should
// only be used in testing.
type Voorhees struct {
	JSON map[string]interface{}
}

// NewVoorhees creates a new Voorhees, taking in a map[string]interface{}
// and deep copying (so as not to manipulate the provided JSON).
// It is so important that the JSON provided in this factory method is valid, otherwise all is for naught.
func NewVoorhees(validJSON map[string]interface{}) *Voorhees {
	return &Voorhees{deepCopy(validJSON)}
}

// Add creates an addition property at the location denoted at the end of the provided JSON path, with the value provided.
// Also creates any nodes that are not present in the path during traversal.
func (v *Voorhees) Add(path string, val interface{}) map[string]interface{} {
	toAdd := finalPropertyOfPath(path)
	level := &v.JSON

	if toAdd != path {
		path = trimLastPropertyFromPath(path)
		var err error
		level, err = v.navigateToPath(path, "add")
		if err != nil {
			panic(fmt.Sprintf("Voorhees: Add | %s", err.Error()))
		}
	}

	intermediatry := *level
	intermediatry[toAdd] = val
	level = &intermediatry

	return v.JSON
}

// AddThen performs an add operation, identical to .Add(), but instead of returning the result,
// it returns the JSON back to allow more manipulations to be performed.
func (v *Voorhees) AddThen(path string, val interface{}) *Voorhees {
	v.JSON = v.Add(path, val)

	return v
}

// Change replaces the property denoted at the end of the provided JSON path with the value provided.
// Panics if it is unable to traverse along the path.
func (v *Voorhees) Change(path string, val interface{}) map[string]interface{} {
	toChange := finalPropertyOfPath(path)
	level := &v.JSON

	if toChange != path {
		path = trimLastPropertyFromPath(path)
		var err error
		level, err = v.navigateToPath(path, "change")
		if err != nil {
			panic(fmt.Sprintf("Voorhees: Change | %s", err.Error()))
		}
	}

	intermediatry := *level

	if _, exists := intermediatry[toChange]; !exists {
		panic(fmt.Sprintf("Voorhees: Change | Unable to change %s because it doesn't exist at path %s",
			toChange, path))
	}

	intermediatry[toChange] = val
	level = &intermediatry

	return v.JSON
}

// ChangeThen performs a change operation, identical to .Change(), but instead of returning the result,
// it returns the JSON back to allow more manipulations to be performed.
func (v *Voorhees) ChangeThen(path string, val interface{}) *Voorhees {
	v.JSON = v.Change(path, val)

	return v
}

// Delete removes the property denoted at the end of the provided JSON path.
// Panics if it is unable to traverse along the path.
func (v *Voorhees) Delete(path string) map[string]interface{} {
	toDelete := finalPropertyOfPath(path)
	level := &v.JSON

	if toDelete != path {
		path = trimLastPropertyFromPath(path)
		var err error
		level, err = v.navigateToPath(path, "delete")
		if err != nil {
			panic(fmt.Sprintf("Voorhees: Delete | %s", err.Error()))
		}
	}

	delete(*level, toDelete)

	return v.JSON
}

// DeleteThen performs a delete operation, identical to .Delete(), but instead of returning the result,
// it returns the JSON back to allow more manipulations to be performed.
func (v *Voorhees) DeleteThen(path string) *Voorhees {
	v.JSON = v.Delete(path)

	return v
}

func deepCopy(in map[string]interface{}) map[string]interface{} {
	bytes, err := json.Marshal(in)
	if err != nil {
		panic("Unable to marshal the provided 'valid JSON' into a JSON object.")
	}

	var out interface{}
	err = json.Unmarshal(bytes, &out)
	if err != nil {
		panic("Unable to unmarshal the provided 'valid JSON' into an interface{}.")
	}

	return out.(map[string]interface{})
}

func (v *Voorhees) navigateToPath(path, parentOp string) (level *map[string]interface{}, err error) {
	propThatPanicked := "" // keep updating this field so if we panic, we can propagate a meaningful err message

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case *runtime.TypeAssertionError:
				err = fmt.Errorf("Unable to navigate to %s. Node: %s was not a map[string]interface{}", path, propThatPanicked)
			default:
				err = errors.New("Unknown panic")
			}
			level = nil
		}
	}()

	if path == "" || path == "." {
		return &v.JSON, nil
	}

	nestedProperties := strings.Split(path, ".")

	level = &v.JSON
	for _, prop := range nestedProperties {
		propThatPanicked = prop

		if denotesArray(prop) {
			level, err = navigateIntoArray(level, prop, parentOp)
			if err != nil {
				return nil, fmt.Errorf("Unable to navigate to %s. Failed to find node: %s", path, prop)
			}
			continue
		}

		l := *level
		_, exists := l[prop]

		if !exists {
			if parentOp == "add" { // during an add, we create any nonexistant nodes in the path
				l[prop] = make(map[string]interface{})
			} else {
				return nil, fmt.Errorf("Unable to navigate to %s. Failed to find node: %s", path, prop)
			}
		}

		intermediary := l[prop].(map[string]interface{}) // potential panic
		level = &intermediary
	}

	return level, nil
}

func navigateIntoArray(level *map[string]interface{}, prop, parentOp string) (*map[string]interface{}, error) {
	var arrayIndex int
	prop, arrayIndex = deconstructArrayPath(prop)

	l := *level
	val, exists := l[prop]

	if !exists {
		if parentOp == "add" { // during an add, we create any nonexistant nodes in the path
			val = make([]interface{}, arrayIndex+1)
			intermediary := val.([]interface{})
			intermediary[arrayIndex] = map[string]interface{}{}
			val = intermediary
			l[prop] = val // we must add our newly created array and item at arrayIndex to the source
		} else {
			return nil, fmt.Errorf("Unable to find array: %s", prop)
		}
	}

	a := val.([]interface{})
	l = a[arrayIndex].(map[string]interface{})
	return &l, nil
}

func finalPropertyOfPath(path string) string {
	split := strings.Split(path, ".")
	return split[len(split)-1]
}

func trimLastPropertyFromPath(path string) string {
	split := strings.Split(path, ".")

	if len(split) == 1 {
		return path
	}

	return strings.Join(split[:len(split)-1], ".")
}

func denotesArray(path string) bool {
	return strings.Contains(path, "[") && strings.Contains(path, "]")
}

func deconstructArrayPath(s string) (string, int) {
	openingBraceIndex := strings.Index(s, "[")
	re := regexp.MustCompile("[0-9]+")
	result := re.FindAllString(s, 1)

	if len(result) == 0 {
		panic(fmt.Sprintf("Voorhees: Array Path | %s is not a valid array denotion", s))
	}

	arrayIndex, err := strconv.Atoi(result[0])

	if err != nil {
		panic(fmt.Sprintf("Voorhees: Array Path | %s is not a valid array denotion", s))
	}

	return s[0:openingBraceIndex], arrayIndex
}
