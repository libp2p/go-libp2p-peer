// This file contains Protobuf and JSON serialization/deserialization methods for peer IDs.
package peer

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"
)

var _ proto.Marshaler = (*ID)(nil)
var _ proto.Unmarshaler = (*ID)(nil)
var _ json.Marshaler = (*ID)(nil)
var _ proto.Unmarshaler = (*ID)(nil)

func (id ID) Marshal() ([]byte, error) {
	return []byte(id), nil
}

func (id ID) MarshalTo(data []byte) (n int, err error) {
	return copy(data, []byte(id)), nil
}

func (id *ID) Unmarshal(data []byte) (err error) {
	*id, err = IDFromBytes(data)
	return err
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(IDB58Encode(id))
}

func (id *ID) UnmarshalJSON(data []byte) (err error) {
	var v string
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	*id, err = IDB58Decode(v)
	return err
}
