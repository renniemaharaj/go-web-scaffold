package elements

type Style struct {
	Selection string
	Styles    map[string]string
}

func MakeStyle(selection string, styles map[string]string) *Style {
	return &Style{
		Selection: selection,
		Styles:    styles,
	}
}

func (s *Style) BuildMarkup() string {
	markup := "\t\t<style>\n"

	markup += "\t\t\t" + s.Selection + " {\n"
	for key, value := range s.Styles {
		markup += "\t\t\t\t" + key + ": " + value + ";\n"
	}
	markup += "\t\t\t}\n"

	markup += "\t\t</style>\n"

	return markup
}
