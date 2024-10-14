package utils

import (
	"fmt"
	"reflect"

	"github.com/jpillora/longestcommon"
	log "github.com/sirupsen/logrus"
)

func commonPrefix(mapping interface{}) string {
	var keys []string
	switch m := mapping.(type) {
	case map[int32]string:
		for _, v := range m {
			keys = append(keys, v)
		}
	case map[string]int32:
		for v := range m {
			keys = append(keys, v)
		}
	case []string:
		keys = m
	default:
		log.Panicf("not supported type (%s) for key extraction", reflect.TypeOf(mapping).String())
	}
	return longestcommon.Prefix(keys)
}

func trimCommonPrefix(value string, mapping interface{}) string {
	prefix := commonPrefix(mapping)
	return value[len(prefix):]
}

func ProtoEnumFromString(val string, typeMapping map[string]int32) (int32, error) {
	if v, ok := typeMapping[val]; ok {
		return v, nil
	}
	key := fmt.Sprintf("%s%s", commonPrefix(typeMapping), val)
	if v, ok := typeMapping[key]; ok {
		return v, nil
	}
	return 0, fmt.Errorf("can't map proto enum from (%s)", val)
}

type isXoEnumLikeReadonly interface {
	String() string
	MarshalText() ([]byte, error)
}

type isXoEnumLike interface {
	String() string
	MarshalText() ([]byte, error)
	UnmarshalText(buf []byte) error
}

type isProtoEnumLikeReadonly interface {
	String() string
	EnumDescriptor() ([]byte, []int)
}

func ToXoEnum(dst isXoEnumLike, src isProtoEnumLikeReadonly, typeMapping map[string]int32) error {
	value := trimCommonPrefix(src.String(), typeMapping)
	if err := dst.UnmarshalText([]byte(value)); err != nil {
		return fmt.Errorf("unable to map from proto enum (%s)", src.String())
	}
	return nil
}

func ToProtoEnum(src isXoEnumLikeReadonly, typeMapping map[string]int32) (int32, error) {
	key := fmt.Sprintf("%s%s", commonPrefix(typeMapping), src.String())
	value, ok := typeMapping[key]
	if !ok {
		return 0, fmt.Errorf("unable to map to proto enum (%s)", key)
	}
	return value, nil
}
