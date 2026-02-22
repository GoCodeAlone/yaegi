package unsafe2_test

import (
	"reflect"
	"testing"

	"github.com/GoCodeAlone/yaegi/internal/unsafe2"
)

func TestSwapFieldType(t *testing.T) {
	f := []reflect.StructField{
		{
			Name: "A",
			Type: reflect.TypeFor[int](),
		},
		{
			Name: "B",
			Type: reflect.PointerTo(unsafe2.DummyType),
		},
		{
			Name: "C",
			Type: reflect.TypeFor[int64](),
		},
	}
	typ := reflect.StructOf(f)
	ntyp := reflect.PointerTo(typ)

	unsafe2.SetFieldType(typ, 1, ntyp)

	if typ.Field(1).Type != ntyp {
		t.Fatalf("unexpected field type: want %s; got %s", ntyp, typ.Field(1).Type)
	}
}
