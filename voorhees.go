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
type Voorhees struct {
	JSON map[string]interface{}
}

// NewVoorhees creates a new Voorhees instance, accepting the initial map[string]interface{} for manipulation
// This preliminary value is deep copied, so any subsequence modificiations do not effect the original struct.
func NewVoorhees(x map[string]interface{}) *Voorhees {
	json := deepCopy(x)

	return &Voorhees{json}
}

// Add creates an addition property of the provided value at path.
// Add will also create any nodes that are not present in the path during traversal.
func (v *Voorhees) Add(path string, val interface{}) (map[string]interface{}, error) {
	toAdd := finalPropertyOfPath(path)
	level := &v.JSON

	if toAdd != path {
		path = trimLastPropertyFromPath(path)
		var err error
		level, err = v.navigateToPath(path, "add")
		if err != nil {
			return nil, err
		}
	}

	intermediatry := *level
	intermediatry[toAdd] = val
	level = &intermediatry

	return v.JSON, nil
}

// Change replaces the property denoted at the end of the provided JSON path with the value provided.
func (v *Voorhees) Change(path string, val interface{}) (map[string]interface{}, error) {
	toChange := finalPropertyOfPath(path)
	level := &v.JSON

	if toChange != path {
		path = trimLastPropertyFromPath(path)
		var err error
		level, err = v.navigateToPath(path, "change")
		if err != nil {
			return nil, err
		}
	}

	intermediatry := *level

	if _, exists := intermediatry[toChange]; !exists {
		return nil, fmt.Errorf("[Voorhees]: Unable to change %s because it doesn't exist at path %s",
			toChange, path)
	}

	intermediatry[toChange] = val
	level = &intermediatry

	return v.JSON, nil
}

// Delete removes the property denoted at the end of the provided JSON path.
func (v *Voorhees) Delete(path string) (map[string]interface{}, error) {
	toDelete := finalPropertyOfPath(path)
	level := &v.JSON

	if toDelete != path {
		path = trimLastPropertyFromPath(path)
		var err error
		level, err = v.navigateToPath(path, "delete")
		if err != nil {
			return nil, err
		}
	}

	delete(*level, toDelete)

	return v.JSON, nil
}

func deepCopy(in map[string]interface{}) map[string]interface{} {
	bytes, _ := json.Marshal(in)

	var out interface{}
	json.Unmarshal(bytes, &out)

	return out.(map[string]interface{})
}

func (v *Voorhees) navigateToPath(path, parentOp string) (level *map[string]interface{}, err error) {
	propThatPanicked := "" // keep updating this field so if we panic, we can propagate a meaningful err message

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case *runtime.TypeAssertionError:
				err = fmt.Errorf("[Voorhees]: Unable to navigate to %s. Node: %s was not a map[string]interface{}",
					path, propThatPanicked)
			default:
				err = errors.New("[Voorhees]: Unknown panic")
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
				return nil, fmt.Errorf("[Voorhees]: Unable to navigate to %s. Failed to find node: %s", path, prop)
			}
			continue
		}

		l := *level
		_, exists := l[prop]

		if !exists {
			if parentOp == "add" { // during an add, we create any nonexistant nodes in the path
				l[prop] = make(map[string]interface{})
			} else {
				return nil, fmt.Errorf("[Voorhees]: Unable to navigate to %s. Failed to find node: %s", path, prop)
			}
		}

		intermediary := l[prop].(map[string]interface{}) // potential panic
		level = &intermediary
	}

	return level, nil
}

func navigateIntoArray(level *map[string]interface{}, prop, op string) (*map[string]interface{}, error) {
	var arrayIndex int
	prop, arrayIndex, err := deconstructArrayPath(prop)

	if err != nil {
		return nil, err
	}

	l := *level
	val, exists := l[prop]

	if !exists {
		if op == "add" { // during an add, we create any nonexistant nodes in the path
			val = make([]interface{}, arrayIndex+1)
			intermediary := val.([]interface{})
			intermediary[arrayIndex] = map[string]interface{}{}
			val = intermediary
			l[prop] = val // we must add our newly created array and item at arrayIndex to the source
		} else {
			return nil, fmt.Errorf("[Voorhees]: Unable to find array: %s", prop)
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

func deconstructArrayPath(s string) (string, int, error) {
	openingBraceIndex := strings.Index(s, "[")
	re := regexp.MustCompile("[0-9]+")
	result := re.FindAllString(s, 1)

	if len(result) == 0 {
		return "", 0, fmt.Errorf("[Voorhees]: Array Path | %s is not a valid array denotion", s)
	}

	arrayIndex, err := strconv.Atoi(result[0])

	if err != nil {
		return "", 0, fmt.Errorf("[Voorhees]: Array Path | %s is not a valid array denotion", s)
	}

	return s[0:openingBraceIndex], arrayIndex, nil
}
