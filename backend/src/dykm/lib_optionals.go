package main;

import (
	"fmt"
)

// Represents an optional type which can contain an error
type Result[T any] struct {
	ok bool;
	ok_value T;
	err_value string;
}
func Ok[T any](value T) Result[T] {
	return Result[T]{ok: true, ok_value: value};
}
func Err[T any](value string) Result[T] {
	return Result[T]{ok: false, err_value: value};
}
func (result Result[T]) IsOk() bool {
	return result.ok;
}
func (result Result[T]) IsErr() bool {
	return !result.ok;
}

func (result Result[T]) UnwrapOk() T {
	if result.IsOk() {
		return result.ok_value;
	}    
	
	formatted_err := fmt.Sprintf("Attempted to unwrap on an error!\n%s", result.err_value)

	panic(formatted_err);
}
func (result Result[T]) UnwrapErr() string {
	if result.IsErr() {
		return result.err_value;
	}

	panic("Attempted to get an error from an Ok result");
}