package backend

const DataNotFoundError = BackendError("backend: data not found")

type BackendError string

func (e BackendError) Error() string { return string(e) }
