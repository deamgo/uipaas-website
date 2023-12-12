package types

// ResponseWrapper wraps the response with its original payload,
// and sets the Status field to codes.OK if everything is OK, but when
// the response is invalid, ErrorReason could be filled to show the error
// details and in such a case, Status is not codes.OK but a specific error
// code to show the kind.
type ResponseWrapper[T any] struct {
	// Value carries the original data.
	Value T `json:"value,omitempty"`
	// ErrorMsg is the error details, it's exclusive with Payload.
	ErrorMsg string `json:"error_msg,omitempty"`
	// Warning attaches a warning message to the response.
	WarningMsg string `json:"warning_msg,omitempty"`
}

// ListResponseWrapper is a wrapper for list response.
type ListResponseWrapper[T any] struct {
	// List carries the list of the data.
	List []T `json:"list"`
	// Total carries the total number of the data.
	Total int64 `json:"total"`
	// ErrorMsg is the error details, it's exclusive with Payload.
	ErrorMsg string `json:"error_msg,omitempty"`
}

// NewEmptyResponse creates an empty ResponseWrapper object.
func NewEmptyResponse() *ResponseWrapper[any] {
	return &ResponseWrapper[any]{}
}

// NewValidResponse creates a ResponseWrapper object with the given value.
func NewValidResponse[T any](value T) *ResponseWrapper[T] {
	return &ResponseWrapper[T]{
		Value: value,
	}
}

// NewValidListResponse creates a ResponseWrapper object with the given list and total.
func NewValidListResponse[T any](list []T, total int64) *ListResponseWrapper[T] {
	// If the list is nil, we will return empty list instead of null.
	if list == nil {
		list = make([]T, 0)
	}
	return &ListResponseWrapper[T]{
		List:  list,
		Total: total,
	}
}

// NewErrorResponse creates a ResponseWrapper object and sets the fields
// according to the given parameters.
func NewErrorResponse(details string) *ResponseWrapper[any] {
	return &ResponseWrapper[any]{
		ErrorMsg: details,
	}
}

// NewWarningResponse creates a ResponseWrapper object and sets the fields
// according to the given parameters.
func NewWarningResponse(details string) *ResponseWrapper[any] {
	return &ResponseWrapper[any]{
		WarningMsg: details,
	}
}
