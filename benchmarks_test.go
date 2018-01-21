package voorhees

import (
	"testing"
)

func BenchmarkAddOneLayer(b *testing.B) {
	x := map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Add("newProperty", 123)
	}
}

func BenchmarkAddTwoLayers(b *testing.B) {
	x := map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Add("newLayer.newProperty", 123)
	}
}

func BenchmarkAddFiveLayers(b *testing.B) {
	x := map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Add("1.2.3.4.5.newProperty", 123)
	}
}

func BenchmarkChangeOneLayer(b *testing.B) {
	x := map[string]interface{}{
		"changeMe": 1,
	}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Change("changeMe", 123)
	}
}

func BenchmarkChangeTwoLayers(b *testing.B) {
	x := map[string]interface{}{
		"layer1": map[string]interface{}{
			"changeMe": 1,
		},
	}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Change("layer1.changeMe", 123)
	}
}

func BenchmarkChangeFiveLayers(b *testing.B) {
	x := map[string]interface{}{
		"1": map[string]interface{}{
			"2": map[string]interface{}{
				"3": map[string]interface{}{
					"4": map[string]interface{}{
						"5": map[string]interface{}{
							"changeMe": 1,
						},
					},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Change("1.2.3.4.5.changeMe", 123)
	}
}

func BenchmarkDeleteOneLayer(b *testing.B) {
	x := map[string]interface{}{
		"deleteMe": 1,
	}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Delete("deleteMe")
	}
}

func BenchmarkDeleteTwoLayers(b *testing.B) {
	x := map[string]interface{}{
		"layer1": map[string]interface{}{
			"deleteMe": 1,
		},
	}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Delete("layer1.deleteMe")
	}
}

func BenchmarkDeleteFiveLayers(b *testing.B) {
	x := map[string]interface{}{
		"1": map[string]interface{}{
			"2": map[string]interface{}{
				"3": map[string]interface{}{
					"4": map[string]interface{}{
						"5": map[string]interface{}{
							"deleteMe": 1,
						},
					},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		v := NewVoorhees(x)
		v.Delete("1.2.3.4.5.deleteMe")
	}
}
