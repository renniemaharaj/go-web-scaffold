package elements

import "fmt"

// Element struct represents a document body element
type Element struct {
	Tag        string
	Attributes []Attribute
	Children   []Element
}

// This function appends one element as a child to another
func (parent *Element) AppendChild(element *Element) {
	parent.Children = append(parent.Children, *element)
}

// This function appends attributes to elements.
func (element *Element) AppendAttribute(attribute *Attribute) {
	element.Attributes = append(element.Attributes, *attribute)
}

// This will create and element of type specificed, and take a map for attributes.
func CreateElementByAttributes(tag string, attributes *[]Attribute) *Element {
	var element Element = Element{}
	element.Tag = tag
	element.Attributes = *attributes
	return &element
}

func (element *Element) BuildMarkup(depth uint8) string {
	// Indentation
	indent := ""

	// To place as child of the element
	innerHTML := ""

	for i := uint8(0); i < depth; i++ {
		indent += "\t"
	}
	markup := fmt.Sprintf("%v<", indent) + element.Tag

	for _, attribute := range element.Attributes {

		// If the attribute is innerHTML, set the innerHTML variable and continue
		if attribute.Name == "innerHTML" {
			innerHTML = attribute.Value
			continue
		}
		markup += " " + attribute.BuildMarkup()
	}

	markup += ">\n"

	// If there is innerHTML, add it to the markup
	if innerHTML != "" {
		markup += fmt.Sprintf("%v\t%v\n", indent, innerHTML)
	}

	// Build the markup for each other child
	for _, child := range element.Children {
		markup += child.BuildMarkup(depth + 1)
	}

	markup += fmt.Sprintf("%v</", indent) + element.Tag + ">\n"

	return markup
}
