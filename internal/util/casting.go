package util

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func CastStruct[T any](i any) *T {
	v := new(T)
	b, _ := json.Marshal(i)
	_ = json.Unmarshal(b, v)

	return v
}

func CastSlice[T any](i any) T {
	v := new(T)
	b, _ := json.Marshal(i)
	_ = json.Unmarshal(b, v)

	return *v
}

func CastStructToMap(s any) (m map[string]interface{}) {
	b, _ := json.Marshal(s)
	_ = json.Unmarshal(b, &m)

	return m
}

// func CastInterfaceToProtoAny(i interface{}) *pAny.Any {
// 	anyValue := &pAny.Any{}
// 	bytes, _ := json.Marshal(i)
// 	bytesValue := &wrappers.BytesValue{
// 		Value: bytes,
// 	}
// 	anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})

// 	return anyValue
// }

// CastToAnyMap casts a map[string]interface{} to map[string]*anypb.Any.
func CastToAnyMap(data map[string]interface{}) (map[string]*anypb.Any, error) {
	anyMap := make(map[string]*anypb.Any)
	for key, value := range data {
		anyValue, err := CastInterfaceToAny(value)
		if err != nil {
			return nil, err
		}
		anyMap[key] = anyValue
	}
	return anyMap, nil
}

// CastInterfaceToAny casts an interface{} to *anypb.Any.
func CastInterfaceToAny(value interface{}) (*anypb.Any, error) {
	switch v := value.(type) {
	case string:
		return anypb.New(wrapperspb.String(v))
	case int:
		return anypb.New(wrapperspb.Int32(int32(v)))
	case int32:
		return anypb.New(wrapperspb.Int32(v))
	case int64:
		return anypb.New(wrapperspb.Int64(v))
	case float32:
		return anypb.New(wrapperspb.Float(v))
	case float64:
		return anypb.New(wrapperspb.Double(v))
	case bool:
		return anypb.New(wrapperspb.Bool(v))
	case proto.Message:
		return anypb.New(v)
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}
}
