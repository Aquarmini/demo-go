package controller

type ControllerInterface interface {
	Handle() (r string, err error)
}
