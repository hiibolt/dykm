package main;

// Represents an optional type which can contain an error
type Result[T any] struct {
	is_ok bool;
	value T;
}
func Ok[T any](value T) Result[T] {
	return Result[T]{is_ok: true, value: value};
}
func Err[T any](value T) Result[T] {
	return Result[T]{is_ok: false, value: value};
}
func is_ok[T any](result Result[T]) bool {
	return result.is_ok;
}
func is_err[T any](result Result[T]) bool {
	return !result.is_ok;
}
func unwrap_ok[T any](result Result[T]) T {
	if is_ok(result) {
		return result.value;
	}

	panic("Attempted to unwrap an error");
}
func unwrap_err[T any](result Result[T]) T {
	if is_err(result) {
		return result.value;
	}

	panic("Attempted to get an error from an Ok result");
}