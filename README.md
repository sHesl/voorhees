# voorhees [![Build Status](https://travis-ci.org/sHesl/voorhees.svg?branch=master)](https://travis-ci.org/sHesl/voorhees)
Voorhees is a package for slicing and dicing JSON (i.e map[string]interface{}), designed to make working with these structs less painful. It relies on type assertions and some pointer trickery for manipulation, no reflection and no marshalling. 

By default, it panics instead of returning errors, partially because the first use-case for this I had was in unit testing, but also because it is fitting with the name! There is a 'safe' version of Voorhees that abides by the typical go error handling paradigm and should never panic.

## Examples:
```
// where myMap is a map[string]interface{}

withNewProperty := NewVoorhees(myMap).Add("firstLayer.nestedObj.anArray[1].newField", "new val")
withChangedProperty := NewVoorhees(myMap).Change("firstLayer.nestedObj.aProperty", "new val")
withDeletedProperty := NewVoorhees(myMap).Delete("firstLayer.deleteMe")

withTwoNewProperties := NewVoorhees(myMap).AddThen("firstLayer.fieldOne", 1).Add("firstLayer.fieldTwo", 2)
```