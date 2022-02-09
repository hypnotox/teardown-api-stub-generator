package src

type Writer interface {
	Write(api Api) (string, error)
}
