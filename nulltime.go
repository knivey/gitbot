package gitbot

import (
	"strings"
	"time"
)

// NullTime allows us to parse time OR a null from JSON
type NullTime struct {
	Valid bool
	Time  time.Time
}

// MarshalJSON turns null time into json
func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte(`null`), nil
	}

	return []byte(n.Time.Format(time.RFC3339)), nil
}

// UnmarshalJSON parses time from json
func (n *NullTime) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)

	if str == "null" {
		n.Valid = false
		return nil
	}

	var err error
	if n.Time, err = time.Parse(time.RFC3339, str); err != nil {
		return err
	}

	n.Valid = true
	return nil
}

// String for fmt.Stringer
func (n NullTime) String() string {
	if !n.Valid {
		return ""
	}
	return n.Time.Format(time.RFC3339)
}
