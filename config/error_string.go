// Code generated by "stringer -type _e -output error_string.go -linecomment"; DO NOT EDIT.

package config

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[_ErrTokenFileReadFailed-1]
	_ = x[ErrNoToken-2]
	_ = x[ErrBadToken-3]
	_ = x[_ErrTokenFileNotAFile-4]
}

const __e_name = "config: token can't be readconfig: no tokenconfig: bad tokenconfig: token file not a file"

var __e_index = [...]uint8{0, 27, 43, 60, 89}

func (i _e) String() string {
	i -= 1
	if i < 0 || i >= _e(len(__e_index)-1) {
		return "_e(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return __e_name[__e_index[i]:__e_index[i+1]]
}
