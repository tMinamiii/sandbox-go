package main

import (
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

// MapStruct represents Map
type MapStruct struct {
	Str     string  `map:"str"`
	StrPtr  *string `map:"strPtr"`
	Bool    bool    `map:"bool"`
	BoolPtr *bool   `map:"boolPtr"`
	Int     int     `map:"int"`
	IntPtr  *int    `map:"intPtr"`
}

func main() {
	src := MapStruct{
		Str:    "string-value",
		StrPtr: &[]string{"string-ptr-value"}[0],
	}

	dest := map[string]string{}
	Encode(dest, &src)
	log.Println(dest)
}

func Encode(target map[string]string, src interface{}) error {
	v := reflect.ValueOf(src)
	e := v.Elem()
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 埋め込まれた構造体は再帰処理
		if f.Anonymous {
			if err := Encode(target, src); err != nil {
				return err
			}
			if err := Encode(target, e.Field(i).Addr().Interface()); err != nil {
				return err
			}
			continue
		}
		key := f.Tag.Get("map")
		// タグがなければフィールド姪をそのまま使う
		if key == "" {
			key = f.Name
		}

		// 子供が構造体だったら再帰処理(名前は引き継ぐ)
		if f.Type.Kind() == reflect.Struct {
			if err := Encode(target, e.Field(i).Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		// フィールドの型を取得
		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			isP = true
			// ポインターのポインターは無視
			if k == reflect.Ptr {
				continue
			}
		}

		switch k {
		case reflect.String:
			if isP {
				// nil ならデータは読み込まない
				if e.Field(i).Pointer() != 0 {
					target[key] = *(*string)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				target[key] = e.Field(i).String()
			}
		case reflect.Bool:
			var b bool
			if isP {
				if e.Field(i).Pointer() != 0 {
					b = *(*bool)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				b = e.Field(i).Bool()
			}
			target[key] = strconv.FormatBool(b)
		case reflect.Int:
			var n int64
			if isP {
				if e.Field(i).Pointer() != 0 {
					n = int64(*(*int)(unsafe.Pointer(e.Field(i).Pointer())))
				}
			} else {
				n = e.Field(i).Int()
			}
			target[key] = strconv.FormatInt(n, 10)
		}
	}
	return nil
}

func Decode(target interface{}, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	return decode(e, src)
}

func decode(e reflect.Value, src map[string]string) error {
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 埋め込まれた構造体は再帰処理
		if f.Anonymous {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// 子供が構造体だったら再帰処理
		if f.Type.Kind() == reflect.Struct {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// タグがなければフィール名をそのまま使う
		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		// 元データになければスキップ
		sv, ok := src[key]
		if !ok {
			continue
		}

		// フィールドの型を取得
		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			// ポインターのポインターは無視
			if k == reflect.Ptr {
				continue
			}
			isP = true
		}
		switch k {
		case reflect.String:
			if isP {
				e.Field(i).Set(reflect.ValueOf(&sv))
			} else {
				e.Field(i).SetString(sv)
			}
		case reflect.Bool:
			b, err := strconv.ParseBool(sv)
			if err == nil {
				if isP {
					e.Field(i).Set(reflect.ValueOf(&b))
				} else {
					e.Field(i).SetBool(b)
				}
			}
		case reflect.Int:
			n64, err := strconv.ParseInt(sv, 10, 64)
			if err == nil {
				if isP {
					n := int(n64)
					e.Field(i).Set(reflect.ValueOf(&n))
				} else {
					e.Field(i).SetInt(n64)
				}
			}
		}
	}
	return nil
}
