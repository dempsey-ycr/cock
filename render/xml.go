package render

import (
	"encoding/xml"
	"net/http"
)

type XML struct {
	Data interface{}
}

var xmlContentType = []string{"application/xml; charset=utf-8"}

// Render (XML) encoder the given interface object and writes data with custom ContentType.
func (r XML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	return xml.NewEncoder(w).Encode(r.Data)
}

// WriteContentType (XML) writes XML ContentType for response.
func (r XML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, xmlContentType)
}