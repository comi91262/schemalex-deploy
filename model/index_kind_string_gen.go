// Code generated by "stringer -type=IndexKind -output=index_kind_string_gen.go"; DO NOT EDIT.

package model

import "strconv"

const _IndexKind_name = "IndexKindInvalidIndexKindPrimaryKeyIndexKindNormalIndexKindUniqueIndexKindFullTextIndexKindSpatialIndexKindForeignKey"

var _IndexKind_index = [...]uint8{0, 16, 35, 50, 65, 82, 98, 117}

func (i IndexKind) String() string {
	if i < 0 || i >= IndexKind(len(_IndexKind_index)-1) {
		return "IndexKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IndexKind_name[_IndexKind_index[i]:_IndexKind_index[i+1]]
}
