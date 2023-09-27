package main

import (
	"testing"
)

func TestFindMostValuablePath(t *testing.T) {
	// Test case 1: A simple 2x2 triangle
	arr1 := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	expected1 := 237
	result1 := FindMostValuablePath(arr1)
	if result1 != expected1 {
		t.Errorf("Expected %d, but got %d for test case 1", expected1, result1)
	}

}

func TestMax(t *testing.T) {
	// Test case for the max function
	result := max(5, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d for max(5, 3)", expected, result)
	}
}
