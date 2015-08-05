package options

import (
	"fmt"
	"reflect"

	"github.com/guregu/null"
)

type (
	Option map[string]string

	Encoder interface {
		EncodeOption() Option
	}

	encoding struct {
		Options []Option
		names   map[int]string
	}
)

func Encode(s interface{}) []Option {
	en := encoding{
		Options: []Option{},
		names:   map[int]string{},
	}

	v := reflect.ValueOf(s)

	switch v.Kind() {
	case reflect.Slice:
		el := v.Type().Elem()
		if el.Kind() == reflect.Struct {
			en.analyze(el)
			for i := 0; i < v.Len(); i++ {
				e := v.Index(i)
				o := en.encode(e)
				en.Options = append(en.Options, o)
			}
		}
	case reflect.Struct:
		en.analyze(v.Type())
		o := en.encode(v)
		en.Options = append(en.Options, o)
	}

	return en.Options
}

func (e encoding) analyze(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("option")
		if tag != "" {
			e.names[i] = tag
		}
	}
}

func (en encoding) encode(e reflect.Value) Option {
	if encoder, ok := e.Interface().(Encoder); ok {
		return encoder.EncodeOption()
	}

	o := Option{}
	for i, name := range en.names {
		val := e.Field(i).Interface()
		if s, ok := val.(null.String); ok {
			val = s.String
		}
		if i, ok := val.(null.Int); ok {
			val = i.Int64
		}
		o[name] = fmt.Sprintf("%v", val)
	}
	return o
}
