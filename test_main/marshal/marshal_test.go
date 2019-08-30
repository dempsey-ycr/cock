package marshal

import "testing"

func BenchmarkStdJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StdJSONMarshal()
	}
}

func BenchmarkIterJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterJSONMarshal()
	}
}

func BenchmarkPbMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PbMarshal()
	}
}

func BenchmarkStdUnmarshal(b *testing.B) {
	bs := StdJSONMarshal()
	v := &Student{}
	for i := 0; i < b.N; i++ {
		StdUnmarshal(bs, v)
	}
}

func BenchmarkIterUnmarshal(b *testing.B) {
	bs := IterJSONMarshal()
	v := &Student{}
	for i := 0; i < b.N; i++ {
		IterUnmarshal(bs, v)
	}
}

func BenchmarkPbUnmarshal(b *testing.B) {
	bs := PbMarshal()
	v := &PbStudent{}
	for i := 0; i < b.N; i++ {
		IterUnmarshal(bs, v)
	}
}
