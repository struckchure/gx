package adapters

type IEchoAdapter interface {
	Get()
	Post()
	Patch()
	Put()
	Delete()
	Custom(method string)
}
