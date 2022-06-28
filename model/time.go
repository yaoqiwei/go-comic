package model

import (
	"database/sql/driver"
	"errors"
	"fehu/util/stringify"
	"time"
)

type Time time.Time

const RFC = "2006-01-02 15:04:05"

func (t2 Time) MarshalJSON() ([]byte, error) {

	t := time.Time(t2)
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(RFC)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, RFC)
	b = append(b, '"')
	return b, nil
}

func (t2 *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+RFC+`"`, string(data))
	*t2 = Time(t)
	return err
}

func (t2 Time) Value() (driver.Value, error) {
	return time.Time(t2), nil
}

func (t2 *Time) Scan(src interface{}) error {
	switch s := src.(type) {
	case string:
		t, err := time.Parse(RFC, s)
		*t2 = Time(t)
		return err
	case []byte:
		t, err := time.Parse(RFC, string(s))
		*t2 = Time(t)
		return err
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		*t2 = Time(time.Unix(stringify.ToInt(s), 0))
	case time.Time:
		*t2 = Time(s)
	default:
		return errors.New("mismatched time type")
	}
	return nil
}
