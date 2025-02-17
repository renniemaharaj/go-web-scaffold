package elements

// HtmlHead struct represents the entire head section of a document.
type Head struct {
	Title   string
	Metas   []Meta
	Links   []Link
	Styles  []Style
	Scripts []Script
}

func (h *Head) BuildMarkup() string {
	markup := "\t<head>\n"

	markup += "\t\t<title>" + h.Title + "</title>\n"

	for _, meta := range h.Metas {
		markup += meta.BuildMarkup()
	}

	for _, link := range h.Links {
		markup += link.BuildMarkup()
	}

	for _, style := range h.Styles {
		markup += style.BuildMarkup()
	}

	for _, script := range h.Scripts {
		markup += script.BuildMarkup()
	}

	markup += "\t</head>\n"

	return markup
}
