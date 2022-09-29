package errors

// ValidationError ...
type ValidationError struct {
	Err error
}

// Error ...
func (r *ValidationError) Error() string {
	return r.Err.Error()
}

// TimeoutError ...
type TimeoutError struct {
	Err error
}

// Error ...
func (r *TimeoutError) Error() string {
	return r.Err.Error()
}

// URLNotFoundError ...
type URLNotFoundError struct {
	Err error
}

// Error ...
func (r *URLNotFoundError) Error() string {
	return r.Err.Error()
}

// EntityNotFoundError ...
type EntityNotFoundError struct {
	Err error
}

// Error ...
func (r *EntityNotFoundError) Error() string {
	return r.Err.Error()
}

// InternalError ...
type InternalError struct {
	Err error
}

// Error ...
func (r *InternalError) Error() string {
	return r.Err.Error()
}

// Unauthorized ...
type Unauthorized struct {
	Err error
}

// Error ...
func (r *Unauthorized) Error() string {
	return r.Err.Error()
}

type BadRequest struct {
	Err error
}

func (r *BadRequest) Error() string {
	return r.Err.Error()
}
