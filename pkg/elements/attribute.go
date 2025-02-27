package elements

//An attribute for name=value matches
type Attribute struct {
	Name  string
	Value string
}

//CreateAttribute will create attributes in map[name]value
func CreateAttribute(name, value string) Attribute {
	return Attribute{
		Name:  name,
		Value: value,
	}
}

func (a *Attribute) BuildMarkup() string {
	return a.Name + "=\"" + a.Value + "\" "
}
