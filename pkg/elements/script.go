package elements

// Script represents a script tag in the HTML head.
type Script struct {
	Src   string
	Async bool
	Defer bool
}

//This function will return a single script tag.
func MakeScript(src string, async, deferring bool) *Script {
	return &Script{
		Src:   src,
		Async: async,
		Defer: deferring,
	}
}

func (s *Script) BuildMarkup() string {
	markup := "\t\t<script "

	if s.Src != "" {
		markup += "src=\"" + s.Src + "\""
	}

	if s.Async {
		markup += " async"
	}

	if s.Defer {
		markup += " defer"
	}

	markup += "></script>\n"

	return markup
}
