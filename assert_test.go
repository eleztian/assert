package assert

import (
	"errors"
	"testing"
)

func TestLineNumbers(t *testing.T) {
	Equal(t, "foo", "foo", "msg!")
	//Equal(t, "foo", "bar", "this should blow up")
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, "foo", "bar", "msg!")
	//NotEqual(t, "foo", "foo", "this should blow up")
}

func TestTestTable_Test(t *testing.T) {
	tb := TestTable{
		"a":   1,
		"bb":  2,
		"":    0,
		"   ": 3,
		"zxc": 3,
	}
	tb.Test(t, func(cond interface{}) (i interface{}, e error) {
		str, ok := cond.(string)
		if !ok {
			return nil, errors.New("not a string")
		}
		return len(str), nil
	})
}
