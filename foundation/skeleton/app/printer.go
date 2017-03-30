package app

type PrinterService struct {
	Message string
}

func (printer *PrinterService) Render() string {
	return "<h1>" + printer.Message + "</h1>"
}
