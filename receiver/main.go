package main

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func main() {
	sVals := map[int]S{1: {"A"}}

	// You can only call Read using a value
	sVals[1].Read()

	// This will not compile:
	//  sVals[1].Write("test")

	sPtrs := map[int]*S{1: {"A"}}

	// You can call both Read and Write using a pointer
	sPtrs[1].Read()
	sPtrs[1].Write("test")

}
