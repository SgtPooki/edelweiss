package iplddata

import (
	"github.com/ipld/edelweiss/def"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
)

type Bool bool

func (Bool) Def() def.Type {
	return def.Bool{}
}

func (v *Bool) Parse(n datamodel.Node) error {
	if n.Kind() != ipld.Kind_Bool {
		return ErrNA
	} else {
		*(*bool)(v), _ = n.AsBool()
		return nil
	}
}

// datamodel.Node implementation

func (Bool) Kind() datamodel.Kind {
	return datamodel.Kind_Bool
}

func (Bool) LookupByString(string) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) LookupByNode(key datamodel.Node) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) LookupByIndex(idx int64) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) LookupBySegment(seg datamodel.PathSegment) (datamodel.Node, error) {
	return nil, ErrNA
}

func (Bool) MapIterator() datamodel.MapIterator {
	return nil
}

func (Bool) ListIterator() datamodel.ListIterator {
	return nil
}

func (Bool) Length() int64 {
	return -1
}

func (Bool) IsAbsent() bool {
	return false
}

func (Bool) IsNull() bool {
	return false
}

func (v Bool) AsBool() (bool, error) {
	return bool(v), nil
}

func (Bool) AsInt() (int64, error) {
	return 0, ErrNA
}

func (Bool) AsFloat() (float64, error) {
	return 0, ErrNA
}

func (Bool) AsString() (string, error) {
	return "", ErrNA
}

func (Bool) AsBytes() ([]byte, error) {
	return nil, ErrNA
}

func (Bool) AsLink() (datamodel.Link, error) {
	return nil, ErrNA
}

func (Bool) Prototype() datamodel.NodePrototype {
	return nil // not needed
}

func TryParseBool(n datamodel.Node) (Bool, error) {
	var b Bool
	return b, b.Parse(n)
}
