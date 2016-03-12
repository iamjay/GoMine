package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"unicode/utf8"
)

type Magic [16]byte

var MagicValue Magic = Magic{
	0x00, 0xff, 0xff, 0x00, 0xfe, 0xfe, 0xfe, 0xfe,
	0xfd, 0xfd, 0xfd, 0xfd, 0x12, 0x34, 0x56, 0x78,
}

func MarshalPacket(p interface{}) (data []byte, err error) {
	s := reflect.ValueOf(p)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	if s.Kind() != reflect.Struct {
		return nil, fmt.Errorf("invalid packet type: %s", s.Kind())
	}

	var b bytes.Buffer
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Kind() {
		case reflect.Int8:
			binary.Write(&b, binary.LittleEndian, int8(f.Int()))
		case reflect.Int16:
			binary.Write(&b, binary.LittleEndian, int16(f.Int()))
		case reflect.Int32:
			binary.Write(&b, binary.LittleEndian, int32(f.Int()))
		case reflect.Int64:
			binary.Write(&b, binary.LittleEndian, int64(f.Int()))
		case reflect.String:
			str := f.String()
			binary.Write(&b, binary.LittleEndian, int16(utf8.RuneCountInString(str)))
			b.WriteString(str)
		case reflect.Array:
			e := f.Type().Elem()
			if e.Kind() != reflect.Uint8 {
				return nil, fmt.Errorf("invalid array type: %s", e.Kind())
			}
			if f.Type() == reflect.TypeOf(MagicValue) {
				binary.Write(&b, binary.LittleEndian, MagicValue)
			} else {
				binary.Write(&b, binary.LittleEndian, f.Interface().([]byte))
			}
		}
	}

	return b.Bytes(), nil
}

func UnmarshalPacket(data []byte, p interface{}) error {
	s := reflect.ValueOf(p)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("invalid packet type: %s", s.Kind())
	}

	b := bytes.NewBuffer(data)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Kind() {
		case reflect.Int8:
			var i8 int8
			binary.Read(b, binary.LittleEndian, &i8)
			f.SetInt(int64(i8))
		case reflect.Int16:
			var i16 int16
			binary.Read(b, binary.LittleEndian, &i16)
			f.SetInt(int64(i16))
		case reflect.Int32:
			var i32 int32
			binary.Read(b, binary.LittleEndian, &i32)
			f.SetInt(int64(i32))
		case reflect.Int64:
			var i64 int64
			binary.Read(b, binary.LittleEndian, &i64)
			f.SetInt(i64)
		case reflect.String:
			var len int16
			binary.Read(b, binary.LittleEndian, &len)
			buf := make([]byte, len)
			binary.Read(b, binary.LittleEndian, buf)
			f.SetString(string(buf))
		case reflect.Array:
			e := f.Type().Elem()
			if e.Kind() != reflect.Uint8 {
				return fmt.Errorf("invalid array type: %s", e.Kind())
			}
			binary.Read(b, binary.LittleEndian, f.Addr().Interface())
		}
	}

	return nil
}
