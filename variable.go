package easyscripts

type Variable interface {
	Name() string
	Value() string
}

func NewVariable(name, value string) Variable {
	return &variable{name: name, value: value}
}

type variable struct {
	name  string
	value string
}

func (v *variable) Name() string {
	return v.name
}
func (v *variable) Value() string {
	return v.value
}
