package expand

import "strings"

// expand 函数将s中的"foo"替换为f("foo")的返回值
func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "foo", f("foo"))
}
