package routing

type ControllerInterface interface {
	Get()
	Post()
	Put()
	Delete()
}

type Controller struct {
}

func (controller *Controller) Get() {
	panic("Method GET not allowed")
}

func (controller *Controller) Post() {
	panic("Method POST not allowed")
}

func (controller *Controller) Put() {
	panic("Method PUT allowed")
}

func (controller *Controller) Delete() {
	panic("Method DELETE allowed")
}
