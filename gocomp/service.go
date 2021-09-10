package gocomp

type GoCompResponse struct {
	Name string
	Data map[string][]string
}

type GoCompService struct {
}

func (g *GoCompService) GetResponse() *GoCompResponse {
	return &GoCompResponse{
		Name: "result",
		Data: map[string][]string{
			"data1": {"1", "2", "3"},
			"data2": {"4", "5", "6"},
			"data3": {"7", "8", "9"},
		},
	}
}
