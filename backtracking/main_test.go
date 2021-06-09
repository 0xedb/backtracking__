package backtracking

import (
	"log"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	input := []int{1, 2, 3}

	actual := backtrack(input)

	expected := [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
	}

	expected1 := [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2},
	}

	if !reflect.DeepEqual(actual, expected1) && !reflect.DeepEqual(actual, expected) {
		log.Println(expected)
		log.Println(actual)
		log.Println(expected1)
		t.Fail()
	}
}
