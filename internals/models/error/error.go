package error

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *BaseError) Error() string {
	return e.Message
}

func NewBaseError(code int, desc string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: desc,
	}
}

var GenericError = BaseError{
	Code:    5000,
	Message: "Generic error",
}

var GenericSuccess = BaseError{
	Code:    1000,
	Message: "Success",
}

var ProductNotEnoughError = BaseError{
	Code:    5100,
	Message: "Product is not enough",
}

var ProductInvalid = BaseError{
	Code:    5101,
	Message: "Payment status invalid",
}

var ProductNotFound = BaseError{
	Code:    5001,
	Message: "Payment status invalid",
}

var PaymentNotFound = BaseError{
	Code:    5001,
	Message: "Payment status invalid",
}
