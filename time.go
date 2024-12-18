package cliversion

import (
	"errors"
	"time"
)

// Time wraps time.Time to change the method of marshalling and unmarshalling Time.
//
//nolint:recvcheck // Marshal and Unmarshal methods.
type Time struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in the RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is reported.
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339Nano)+len(`""`))
	b = append(b, '"')
	b = t.AppendFormat(b, time.RFC3339Nano)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time must be a quoted string in the RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	// TODO(https://go.dev/issue/47353): Properly unescape a JSON string.
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	var err error
	*t, err = Parse(time.RFC3339Nano, string(data))
	return err
}

// Parse parses a formatted string and returns the time value it represents.
// See the documentation for the constant called Layout to see how to
// represent the format. The second argument must be parseable using
// the format string (layout) provided as the first argument.
//
// The example for Time.Format demonstrates the working of the layout string
// in detail and is a good reference.
//
// When parsing (only), the input may contain a fractional second
// field immediately after the seconds field, even if the layout does not
// signify its presence. In that case either a comma or a decimal point
// followed by a maximal series of digits is parsed as a fractional second.
// Fractional seconds are truncated to nanosecond precision.
//
// Elements omitted from the layout are assumed to be zero or, when
// zero is impossible, one, so parsing "3:04pm" returns the time
// corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
// 0, this time is before the zero Time).
// Years must be in the range 0000..9999. The day of the week is checked
// for syntax but it is otherwise ignored.
//
// For layouts specifying the two-digit year 06, a value NN >= 69 will be treated
// as 19NN and a value NN < 69 will be treated as 20NN.
//
// The remainder of this comment describes the handling of time zones.
//
// In the absence of a time zone indicator, Parse returns a time in UTC.
//
// When parsing a time with a zone offset like -0700, if the offset corresponds
// to a time zone used by the current location (Local), then Parse uses that
// location and zone in the returned time. Otherwise it records the time as
// being in a fabricated location with time fixed at the given zone offset.
//
// When parsing a time with a zone abbreviation like MST, if the zone abbreviation
// has a defined offset in the current location, then that offset is used.
// The zone abbreviation "UTC" is recognized as UTC regardless of location.
// If the zone abbreviation is unknown, Parse records the time as being
// in a fabricated location with the given zone abbreviation and a zero offset.
// This choice means that such a time can be parsed and reformatted with the
// same layout losslessly, but the exact instant used in the representation will
// differ by the actual zone offset. To avoid such problems, prefer time layouts
// that use a numeric zone offset, or use ParseInLocation.
func Parse(layout, value string) (Time, error) {
	ts, err := time.Parse(layout, value)
	return Time{Time: ts}, err
}
