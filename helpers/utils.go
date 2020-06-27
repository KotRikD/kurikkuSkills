package helpers

import (
	"math"
	"reflect"
)

// GetWeightedValue lol
func GetWeightedValue(vals []float64, decay float64) int {
	weightDecay := 1.0
	result := 0.00
	for i := 0; i < len(vals); i++ {
		result += weightDecay * vals[i]
		weightDecay *= decay
	}
	return int(math.Round(result * (1.0 - decay)))
}

// ExistsInArray will check value exists in array or not
func ExistsInArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
