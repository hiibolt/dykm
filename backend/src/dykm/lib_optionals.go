package main;

// Represents an optional type which can contain an error
type Result[T any] struct {
	is_ok bool;
	ok_value T;
	err_value string;
}
func Ok[T any](value T) Result[T] {
	return Result[T]{is_ok: true, ok_value: value};
}
func Err[T any](value string) Result[T] {
	return Result[T]{is_ok: false, err_value: value};
}
func is_ok[T any](result Result[T]) bool {
	return result.is_ok;
}
func is_err[T any](result Result[T]) bool {
	return !result.is_ok;
}
func (result Result[T]) unwrap_ok() T {
	if is_ok(result) {
		return result.ok_value;
	}

	panic("Attempted to unwrap an error");
}
func (result Result[T]) unwrap_err() string {
	if is_err(result) {
		return result.err_value;
	}

	panic("Attempted to get an error from an Ok result");
}