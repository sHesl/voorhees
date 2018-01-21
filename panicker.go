package voorhees

// PanickerVoorhees allows json path denoted manipulations of a map[string]interface{}
type PanickerVoorhees struct {
	v *Voorhees
}

// NewPanickerVoorhees creates a new PanickerVoorhees instance, that behaves identical to a Voorhees instance,
// except it panics instead of returning errors.
// This is not ideomatic Go, and should never be used in production! The intended use for this struct is
// during unit tests that require a significant number of test cases using json manipulation.
func NewPanickerVoorhees(x map[string]interface{}) *PanickerVoorhees {
	json := deepCopy(x)

	return &PanickerVoorhees{&Voorhees{json}}
}

// Add creates an addition property of the provided value at path
// behaving exactly as NewVoorhees.Add(), expect NewPanickerVoorhees().Add()
// panics upon encountering an error.
func (pv *PanickerVoorhees) Add(path string, val interface{}) map[string]interface{} {
	result, err := pv.v.Add(path, val)

	if err != nil {
		panic(err)
	}

	return result
}

// Change replaces the property denoted at the end of the provided JSON path with the value provided,
// behaving exactly as NewVoorhees.Change(), expect NewPanickerVoorhees().Change()
// panics upon encountering an error.
func (pv *PanickerVoorhees) Change(path string, val interface{}) map[string]interface{} {
	result, err := pv.v.Change(path, val)

	if err != nil {
		panic(err)
	}

	return result
}

// Delete removes the property denoted at the end of the provided JSON path,
// behaving exactly as NewVoorhees.Delete(), expect NewPanickerVoorhees().Delete()
// panics upon encountering an error.
func (pv *PanickerVoorhees) Delete(path string) map[string]interface{} {
	result, err := pv.v.Delete(path)

	if err != nil {
		panic(err)
	}

	return result
}
