package apperrors

type ErrCode string

const (
  Unknown          ErrCode = "U000"
  InsertDataFailed ErrCode = "S001"
)
