package test

type Person struct {
	Id    int64
	Name  string
	Birth string
}

type People interface {
	Show()
}

type Student struct {
}

func (s *Student) Show() {

}
