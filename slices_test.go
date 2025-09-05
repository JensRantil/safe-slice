package slices_test

import (
	"reflect"
	"testing"

	slices "github.com/JensRantil/immutable-slices"
)

func TestBlogPostSlice(t *testing.T) {
	startingSlice := []int{1, 2, 3}
	headSlice := startingSlice[0:2]
	tailSlice := startingSlice[1:3]

	headSlice = append(headSlice, 4)

	// Assert startingSlice == [1 2 4]
	if expected := []int{1, 2, 4}; !reflect.DeepEqual(startingSlice, expected) {
		t.Errorf("startingSlice = %v, want %v", startingSlice, expected)
	}

	// Assert headSlice == [1 2 4]
	if expected := []int{1, 2, 4}; !reflect.DeepEqual(headSlice, expected) {
		t.Errorf("headSlice = %v, want %v", headSlice, expected)
	}

	// Assert tailSlice == [2 4]
	if expected := []int{2, 4}; !reflect.DeepEqual(tailSlice, expected) {
		t.Errorf("tailSlice = %v, want %v", tailSlice, expected)
	}
}

// Test that SafeSlice.AppendAll does not modify the original slice.
func TestSafeSliceAppendAll(t *testing.T) {
	// GIVEN:

	startingSlice := slices.NewSafeSlice([]int{1, 2, 3})
	headSlice := startingSlice.SubSlice(0, 2)
	tailSlice := startingSlice.SubSlice(1, 3)

	// Assert tailSlice == [2 3]
	if expected := []int{2, 3}; !reflect.DeepEqual(tailSlice.GetSlice(), expected) {
		t.Errorf("tailSlice = %v, want %v", tailSlice.GetSlice(), expected)
	}

	// WHEN:
	headSlice = headSlice.AppendAll(4)

	// THEN:

	// Assert startingSlice == [1 2 3]
	if expected := []int{1, 2, 3}; !reflect.DeepEqual(startingSlice.GetSlice(), expected) {
		t.Errorf("startingSlice = %v, want %v", startingSlice.GetSlice(), expected)
	}

	// Assert headSlice == [1 2 4]
	if expected := []int{1, 2, 4}; !reflect.DeepEqual(headSlice.GetSlice(), expected) {
		t.Errorf("headSlice = %v, want %v", headSlice.GetSlice(), expected)
	}

	// Assert tailSlice == [2 3]
	if expected := []int{2, 3}; !reflect.DeepEqual(tailSlice.GetSlice(), expected) {
		t.Errorf("tailSlice = %v, want %v", tailSlice.GetSlice(), expected)
	}
}

// Test that SafeSlice.Set does not modify the original slice.
func TestSafeSliceSet(t *testing.T) {
	// GIVEN:
	startingSlice := slices.NewSafeSlice([]int{1, 2, 3})
	subSlice := startingSlice.SubSlice(0, 2)
	subSlice.Set(0, 4)

	// THEN:

	// Assert startingSlice == [1 2 3]
	if expected := []int{1, 2, 3}; !reflect.DeepEqual(startingSlice.GetSlice(), expected) {
		t.Errorf("startingSlice = %v, want %v", startingSlice.GetSlice(), expected)
	}

	// Assert subSlice == [4 2]
	if expected := []int{4, 2}; !reflect.DeepEqual(subSlice.GetSlice(), expected) {
		t.Errorf("subSlice = %v, want %v", subSlice.GetSlice(), expected)
	}
}
