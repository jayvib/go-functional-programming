package domain

type File struct {
	Id int
	Name string `json:"name"`
	ErrorMsg string `json:"error"`
	Contents LogFile `json:"logFile"`
	Bytes []byte `json:"byte"`
}

