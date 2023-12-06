package maps

import (
	"reflect"
	"testing"
)

func TestNewIntMap(t *testing.T) {
	m := map[int]float64{1: 3, 2: 2, 3: 1}
	expectK := []int{1, 2, 3}
	expectV := []float64{3, 2, 1}
	intMap := NewIntMap(m)
	gotK, gotV := intMap.Items()
	if !reflect.DeepEqual(gotK, expectK) {
		t.Errorf("gotK=%v, expectK=%v", gotK, expectK)
	}
	if !reflect.DeepEqual(gotV, expectV) {
		t.Errorf("gotV=%v, expectV=%v", gotV, expectV)
	}

	i := 0
	for item := range intMap.Items2() {
		gotK, gotV := item.Key, item.Value
		if !reflect.DeepEqual(gotK, expectK[i]) {
			t.Errorf("gotK=%v, expectK=%v", gotK, expectK)
		}
		if !reflect.DeepEqual(gotV, expectV[i]) {
			t.Errorf("gotV=%v, expectV=%v", gotV, expectV)
		}
		i++
	}
}
