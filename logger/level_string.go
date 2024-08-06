// Code generated by "stringer -type Level -output level_string.go -linecomment"; DO NOT EDIT.

package logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LevelSilent-1]
	_ = x[LevelDebug-2]
	_ = x[LevelInfo-3]
	_ = x[LevelWarn-4]
	_ = x[LevelError-5]
}

const _Level_name = "silentdebuginfowarnerror"

var _Level_index = [...]uint8{0, 6, 11, 15, 19, 24}

func (i Level) String() string {
	i -= 1
	if i < 0 || i >= Level(len(_Level_index)-1) {
		return "Level(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}
