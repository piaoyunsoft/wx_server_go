package constants

type ErrCode int

const (
	Success ErrCode = 1000 + iota
	InvalidParams
	DataNull
	IdExist
	DBError
	InvalidToken
	LoginFail
)
