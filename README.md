Assert
=====

this is a simple assertion library much in the style of gocheck, but without
the suite stuff.

for example:
```go

import (
	"testing"

	. "github.com/hrkipp/assert"
)

func TestAssertsFail(t *testing.T) {
    Assert(t, 1, Equals, 1)
    Assert(t, 1, Not(Equals), 2)
    var nilString *string
    Assert(t, nilString, IsNil)
    Assert(t, []string{"foo", "bar"}, Contains, "foo"}
    Assert(t, []string{"foo", "bar"}, HasSameElementAs , []string{"bar", "foo"}}
}
```
