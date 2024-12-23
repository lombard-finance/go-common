package config

import (
	"fmt"
	"reflect"

	"github.com/jpillora/longestcommon"
	"github.com/pkg/errors"
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

func ProtoEnumFromString(val string, typeMapping map[string]int32) (int32, error) {
	if v, ok := typeMapping[val]; ok {
		return v, nil
	}
	key := fmt.Sprintf("%s%s", commonPrefix(typeMapping), val)
	if v, ok := typeMapping[key]; ok {
		return v, nil
	}
	return 0, errors.Errorf("can't map proto enum from (%s)", val)
}
