package main

type API interface {
	Say() string
}

type hiApi struct {
}

func (h *hiApi) Say() string {
	return "say hi"
}

type HelloAPI struct {
}

func (a *HelloAPI) Say() string {
	return "say hello"
}

func APIFactory(i int) API {
	if 1 == i {
		return &hiApi{}
	} else {
		return &HelloAPI{}
	}
}
func main() {
	println(APIFactory(1).Say())
	println(APIFactory(2).Say())
}
