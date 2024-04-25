package helper

import (
	"reflect"
	"strings"
	"testing"
)

func TestSliceDiff(t *testing.T) {
	comp := func(a, b int) bool { return a == b }

	// Test case 1: Testing when there are no elements to add or remove.
	source1 := []int{1, 2, 3}
	target1 := []int{1, 2, 3}
	add1, remove1 := SliceDiff(source1, target1, comp)
	if len(add1) != 0 {
		t.Errorf("Expected empty add slice, got %v", add1)
	}
	if len(remove1) != 0 {
		t.Errorf("Expected empty remove slice, got %v", remove1)
	}

	source2 := []int{1, 2, 3}
	target2 := []int{4, 5, 6}
	add2, remove2 := SliceDiff(source2, target2, comp)
	expectedAdd2 := []int{1, 2, 3}
	expectedRemove2 := []int{4, 5, 6}
	if !reflect.DeepEqual(add2, expectedAdd2) {
		t.Errorf("Expected add slice %v, got %v", expectedAdd2, add2)
	}
	if !reflect.DeepEqual(remove2, expectedRemove2) {
		t.Errorf("Expected empty remove slice, got %v", remove2)
	}

	// Test case 3: Testing when there are elements to remove.
	source3 := []int{1, 2, 3}
	target3 := []int{2, 3}
	add3, remove3 := SliceDiff(source3, target3, nil)
	if len(add3) != 1 {
		t.Errorf("Expected 1add slice, got %v", add3)
	}
	if len(remove3) != 0 {
		t.Errorf("Expected empty remove slice, got %v", add3)
	}

	source4 := []string{"apple", "banana", "orange"}
	target4 := []string{"banana", "kiwi"}
	add4, remove4 := SliceDiff(source4, target4, func(a, b string) bool {
		return strings.HasPrefix(a, b)
	})
	expectedAdd4 := []string{"apple", "orange"}
	if !reflect.DeepEqual(add4, expectedAdd4) {
		t.Errorf("Expected add slice %v, got %v", expectedAdd4, add4)
	}
	expectedRemove4 := []string{"kiwi"}
	if !reflect.DeepEqual(remove4, expectedRemove4) {
		t.Errorf("Expected remove slice %v, got %v", expectedRemove4, remove4)
	}
}
