package scalar

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// Date 日付型
type Date time.Time

var _ graphql.Marshaler = Date{}
var _ graphql.Unmarshaler = (*Date)(nil)

// dateFormat 日付フォーマット
const dateFormat = "2006-01-02"

// MarshalGQL implements Scalarable.
func (d Date) MarshalGQL(w io.Writer) {
	w.Write([]byte(time.Time(d).Format(dateFormat)))
}

// UnmarshalGQL implements Scalarable.
func (d *Date) UnmarshalGQL(v interface{}) error {
	dateStr, ok := v.(string)
	if !ok {
		return fmt.Errorf("Date must be a date formated string")
	}

	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}

	*d = Date(date)

	return nil
}
