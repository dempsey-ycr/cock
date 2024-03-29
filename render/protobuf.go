package render

import (
	"net/http"

	"github.com/golang/protobuf/proto"
)

// ProtoBuf contains the given interface object.
type ProtoBuf struct {
	Data interface{}
}

var protobufContentType = []string{"application/x-protobuf"}

// Render (ProtoBuf) marshals the given interface object and writes data with custom ContentType.
func (r ProtoBuf) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := proto.Marshal(r.Data.(proto.Message))
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
