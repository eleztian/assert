# assert
NO LONGER MAINTAINED - Just use Go's testing package.

[![Build Status][1]][2]
[![codecov][3]][4]
[![goreportcard for eleztian/assert][5]][6]
[![godoc for eleztian/assert][7]][8]
[![MIT Licence][9]][10]

Fork Form github.com/bmizerany/assert


## Assertions for Go tests

## Install

    $ go get github.com/eleztian/assert

## Use

**point.go**

```go
package point

type Point struct {
    x, y int
}
func (p Point) String() string {
    return fmt.Sprintf("%d,%d", p.X, p.Y)
}
```

**point_test.go**

```go
package point

import (
    "testing"
    "github.com/bmizerany/assert"
)

func TestAsserts(t *testing.T) {
    p1 := Point{1, 1}
    p2 := Point{2, 1}

    assert.Equal(t, p1, p2)
}
func TestTestTable(t *testing.T) {
    tb := assert.TestTable{
        Point{1,1}:"1,1",
        Point{2,2}:"2,2",
        Point{3,2}:"3,2",
    }

    tb.Test(t, func(cond interface{}) (i interface{}, e error) {
        return cond.(Point).String(), nil
    })
}
```
 

**output**
```bash
$ go test
 --- FAIL: TestAsserts (0.00 seconds)
 assert.go:15: /Users/flavio.barbosa/dev/stewie/src/point_test.go:12
     assert.go:24: ! X: 1 != 2
 FAIL
```


[1]: https://travis-ci.com/eleztian/assert.svg?branch=master
[2]: https://travis-ci.com/eleztian/assert
[3]: https://codecov.io/gh/eleztian/assert/branch/master/graph/badge.svg
[4]: https://codecov.io/gh/eleztian/assert
[5]: https://goreportcard.com/badge/github.com/eleztian/assert
[6]: https://goreportcard.com/report/github.com/eleztian/assert
[7]: https://godoc.org/github.com/eleztian/assert?status.svg
[8]: https://godoc.org/github.com/eleztian/assert
[9]: https://badges.frapsoft.com/os/mit/mit.svg?v=103
[10]: https://opensource.org/licenses/mit-license.php
