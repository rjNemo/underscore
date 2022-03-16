package underscore

// Result represent the outcome of an operation where failure is possible
type Result[T any] interface {
	isResult() //to seal the Result interface
	ToValue() (*T, error)
	IsSuccess() bool
}

// Ok is the Result that represents success.
type Ok[T any] struct {
	Value *T
}

func (Ok[T]) isResult() {}

func (o Ok[T]) ToValue() (*T, error) {
	return o.Value, nil
}

func (o Ok[T]) IsSuccess() bool {
	return true
}

// Err is the Result that represents failure. It implements the error interface
type Err[T any] struct{ Err error }

func (e Err[T]) ToValue() (*T, error) {
	return nil, e.Err
}

func (e Err[T]) IsSuccess() bool {
	return false
}

func (Err[T]) isResult() {}

func (e Err[T]) Error() string {
	return e.Err.Error()
}

func ToResult[T any](value *T, err error) Result[T] {
	if err != nil {
		return Err[T]{
			Err: err,
		}
	}
	return Ok[T]{Value: value}
}
