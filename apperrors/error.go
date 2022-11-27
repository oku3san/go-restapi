package apperrors

type MyAppError struct {
  ErrCode
  Message string
}

func (myErr *MyAppError) Error() string {
  return myErr.Message
}
