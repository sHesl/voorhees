# voorhees [![Build Status](https://travis-ci.org/sHesl/voorhees.svg?branch=master)](https://travis-ci.org/sHesl/voorhees)
Voorhees is a package for slicing and dicing JSON (i.e map[string]interface{}). It relies on type assertions 
and some pointer trickery for manipulations, providing Add, Change and Delete functionality to any property on 
any node inside the underlying map.


## Installation
`go get github.com/sHesl/voorhees`


## Functionality:

### Add
Add will insert a new property into the input map at the requested path. If the requested path contains
nodes that do not currently exist, they will be added as they are encountered. Add can navigate into or create
arrays.
```
myMap := map[string]interface{}{}

withNewProperty := NewVoorhees(myMap).Add("newProperty", "added") 
// {"newProperty":"added"}

withNewNodeAndProperty := NewVoorhees(myMap).Add("newLayer.newProperty", "added") 
// {"newLayer":{"newProperty":"added"}}

withNewNodeAndNewArrayAndNewProperty := NewVoorhees(myMap).Add("newLayer.newArray[0].newProperty", "added") 
// {"newLayer":{"newArray":[{"newProperty": "added"}]}}
```

### Change
Change will update an existing property at the requested path. If the requested property does not exists, an
error will occur. Change can navigate into arrays and update the node at the specifed index.
```
myMap := map[string]interface{}{
  "changeMe": 123,
}

withUpdatedProperty, _ := NewVoorhees(myMap).Change("changeMe", "changed") 
// {"changeMe":"changed"}

nestedMap :=  map[string]interface{}{
  "layer1": map[string]interface{}{
    "changeMe": 123,
  },
}

withUpdatedPropertyInsideNode, _ := NewVoorhees(nestedMap).Change("layer1.changeMe", "changed") 
// {"layer1":{"changeMe":"changed"}}

withArray :=  map[string]interface{}{
  "layer1": map[string]interface{}{
    "array1": [map[string]interface{}{
      "changeMe": 123,
    }],
  },
}

withUpdatedPropertyInsideArray, _ := NewVoorhees(withArray).Change("layer1.array1[0].changeMe", "changed")
// {"layer1":{"array1":[{"changeMe": "added"}]}}
```

### Delete
Delete will delete the existing property at the requested path. If the requested property does not exists, an
error will occur. Delete can navigate into arrays and delete the node at the specifed index.
```
myMap := map[string]interface{}{
  "delete": 123,
}

withDeletedProperty, _ := NewVoorhees(myMap).Delete("changeMe") 
// {}

nestedMap :=  map[string]interface{}{
  "layer1": map[string]interface{}{
    "deleteMe": 123,
  },
}

withDeletePropertyInsideNode, _ := NewVoorhees(nestedMap).Delete("layer1.deleteMe") 
// {"layer1":{}}

withArray :=  map[string]interface{}{
  "layer1": map[string]interface{}{
    "array1": [map[string]interface{}{
      "changeMe": 123,
    }],
  },
}

withDeletedPropertyInsideArray, _ := NewVoorhees(withArray).Delete("layer1.array1[0].changeMe")
// {"layer1":{"array1":[{}]}}
```

### PanickerVoorhees
Ideomatic Go always follow the practice of packages returning errors to back the calling code, and never
panicking from inside a package without recovery. Voorhees was writing to specifically speed up
unit testing in the scenario where there were thousands of test cases that needed to be covered, all involving
minor changes to a source map. Returning multiple values from a func (like a value and an error), prevents
assigning these test cases in a single line, and as such, results in a massive bloating of test code.

To counter this specific scenario, there is a PanickerVoorhees implementation that panics as opposed to 
returning errors. Obviously, this should never be used in production code, but it is available for use
inside tests.

```
testCases := []map[string]interface{}{
  NewPanickerVoorhees(origin).Add("newProperty"),
  NewPanickerVoorhees(origin).Add("newProperty2"),
  NewPanickerVoorhees(origin).Add("newProperty3"),
  NewPanickerVoorhees(origin).Change("changeProperty1"),
  NewPanickerVoorhees(origin).Change("changeProperty2"),
  NewPanickerVoorhees(origin).Delete("deleteProperty1"),
}
```

