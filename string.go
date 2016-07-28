package handyPointers


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
	*s = src.(String)
	return nil
}
func (s String) Value() (driver.Value, error) {
	return string(s).Value()
}

func (s *String) MarshalJSON() ([]byte, error) {
	return s.StringPtr().MarshalJSON()
}

