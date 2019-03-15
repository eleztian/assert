package point

import (
	"testing"

	"github.com/eleztian/assert"
)

func TestAsserts(t *testing.T) {
	p1 := Point{1, 1}
	p2 := Point{2, 1}

	assert.Equal(t, p1, p2)
}

func TestTestTable(t *testing.T) {
	tb := assert.TestTable{
		Point{1, 1}: "1,1",
		Point{2, 2}: "2,2",
		Point{3, 2}: "3,2",
	}

	tb.Test(t, func(cond interface{}) (i interface{}, e error) {
		return cond.(Point).String(), nil
	})
}
