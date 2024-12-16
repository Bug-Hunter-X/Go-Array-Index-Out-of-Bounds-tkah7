# Go Array Index Out of Bounds

This repository demonstrates a common error in Go: accessing an array index out of bounds.  The code attempts to iterate past the end of an array resulting in a runtime panic.

## The Bug

The `bug.go` file contains a loop that iterates from 0 to 5 (inclusive), attempting to assign values to a 5-element array. Because arrays in Go are zero-indexed, index 5 is out of bounds, leading to a runtime panic.

## The Solution

The `bugSolution.go` file demonstrates the correct approach.  The loop is modified to iterate up to, but not including, the length of the array, preventing the out-of-bounds access.