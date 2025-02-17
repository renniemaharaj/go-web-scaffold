package elements

// MetaTag represents a meta tag in the HTML head.
type Meta struct {
	Attribute string
	Values    []string
	Content   string
}

// Returns a meta tag <meta attribute="value" content="content">. Omit content if necessary
func MakeMeta(attribute string, values []string, content string) *Meta {
	return &Meta{
		Attribute: attribute,
		Values:    values,
		Content:   content,
	}
}

func (m *Meta) BuildMarkup() string {
	markup := "\t\t<meta "

	markup += m.Attribute + "=\""
	for i, value := range m.Values {
		markup += value
		if i < len(m.Values)-1 {
			markup += " "
		}
	}
	markup += "\""

	if m.Content != "" {
		markup += " content=\"" + m.Content + "\""
	}

	markup += ">\n"

	return markup
}
