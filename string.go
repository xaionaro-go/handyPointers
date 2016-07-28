package handyPointers


import (
	"database/sql/driver"
)

type String string;

func (s *String) String() string {
	if (s == nil) {
		return ""
	}
	return string(*s)
}

func (s *String) StringPtr() *string {
	if (s == nil) {
		return nil
	}
	str := string(*s)
	return &str
}

func (s *String) Scan(src interface{}) error {
	*s = String(src.([]byte))
	return nil
}
func (s String) Value() (driver.Value, error) {
	return []byte(s), nil
}

func (s *String) MarshalJSON() ([]byte, error) {
	if (s == nil) {
		return []byte("null"), nil
	}

	return []byte("\""+*s+"\""), nil
}

