package marshal

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
)

var (
	iterJson = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Student struct {
	Age     int     `json:"age"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Salt    float64 `json:"salt"`
	Sub     *Sub    `json:"sub"`
}

type Sub struct {
	Options []byte `json:"options"`
}

var s = &Student{
	Age:     25,
	Name:    "dempsey",
	Address: "黑沙土 china",
	Salt:    28700.85,
	Sub: &Sub{
		Options: []byte("hello, world..."),
	},
}

var ps = &PbStudent{
	Age:     25,
	Name:    "dempsey",
	Address: "黑沙土 china",
	Salt:    28700.85,
	Sub: &PbSub{
		Options: []byte("hello, world..."),
	},
}

func StdJSONMarshal() []byte {
	bs, _ := json.Marshal(s)
	return bs
}

func IterJSONMarshal() []byte {
	bs, _ := iterJson.Marshal(s)
	return bs
}

func PbMarshal() []byte {
	bs, _ := proto.Marshal(ps)
	return bs
}

// ------------------------------------------------------unmarshal------------------------//
func StdUnmarshal(bs []byte, v interface{}) {
	_ = json.Unmarshal(bs, v)
}

func IterUnmarshal(bs []byte, v interface{}) {
	_ = iterJson.Unmarshal(bs, v)
}

func PbUnmarshal(bs []byte, v proto.Message) {
	_ = proto.Unmarshal(bs, v)
}
