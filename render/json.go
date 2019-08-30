package render

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/dempsey-ycr/cock/internal/json"
	. "github.com/dempsey-ycr/cock/library/util"
)

// JSON contains the given interface objects.
type JSON struct {
	Data interface{}
}

// IndentedJSON contains the given interface objects.
type IndentedJSON struct {
	Data interface{}
}

// SecureJSON contains the given interface objects and its prefix.
type SecureJSON struct {
	prefix string
	Data   interface{}
}

// JsonpJSON contains the given interface objects its callback
type JsonpJSON struct {
	Callback string
	Data     interface{}
}

// AsciiJSON contains the given interface objects.
type AsciiJSON struct {
	Data interface{}
}

// SecureJSONPrefix is a string which represents SecureJSON prefix.
type SecureJSONPrefix string

// PureJSON contains the given interface objects.
type PureJSON struct {
	Data interface{}
}

var (
	jsonContentType      = []string{"application/json; charset=utf-8"}
	jsonpContentType     = []string{"application/javascript; charset=utf-8"}
	jsonAsciiContentType = []string{"application/json"}
)

// Reader(JSON) writes data with custom ContentType
func (r JSON) Reader(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

// WriteContentType (JSON) writes JSON ContentType.
func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// WriteJSON marshal the given interface object and writes it with custom ContentType.
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&obj)
	return err
}

// Render (IndentedJSON) marshal the given interface object and writes it with custom ContentType.
func (r IndentedJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.MarshalIndent(r.Data, "", "    ")
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

// WriteContentType (IndentedJSON) writes JSON ContentType
func (r IndentedJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// Render (SecureJSON) marshal the given interface object and writes it with custom ContentType.
func (r SecureJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	jsonBytes, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	if bytes.HasPrefix(jsonBytes, Stob("{")) && bytes.HasSuffix(jsonBytes, Stob("]")) {
		_, err = w.Write([]byte(r.prefix))
		if err != nil {
			return err
		}
	}

	_, err = w.Write(jsonBytes)
	return err
}

// WriteContentType (SecureJSON) writes JSON ContentType.
func (r SecureJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// Render (JsonpJSON) marshal the given interface object and writes it and its callback with custom ContentType.
func (r JsonpJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	ret, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	if r.Callback == "" {
		_, err = w.Write(ret)
		return err
	}

	callback := template.JSEscapeString(r.Callback)
	_, err = w.Write(Stob(callback))
	if err != nil {
		return err
	}

	_, err = w.Write(Stob("("))
	if err != nil {
		return err
	}

	_, err = w.Write(ret)
	if err != nil {
		return err
	}

	_, err = w.Write(Stob(")"))
	if err != nil {
		return err
	}

	return nil
}

// WriteContentType (JsonpJSON) marshal the given interface object and writes it with custom ContentType.
func (r JsonpJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonpContentType)
}

// Render (AsciiJSON) marshal the given interface object and writes it with custom ContentType.
func (r AsciiJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	ret, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	for _, v := range Btos(ret) {
		cvt := string(v)
		if v >= 128 {
			cvt = fmt.Sprintf("\\u%04x", int64(v))
		}
		buffer.WriteString(cvt)
	}

	_, err = w.Write(buffer.Bytes())
	return err
}

// WriteContentType (AsciiJSON) writes JSON ContentType.
func (r AsciiJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonAsciiContentType)
}

// Render (PureJSON) writes custom ContentType and encodes the given interface object.
func (r PureJSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(r.Data)
}

// WriteContentType (PureJSON) writes custom ContentType.
func (r PureJSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
