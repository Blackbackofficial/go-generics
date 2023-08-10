package generics

// AppendStringOrBytes appends string or bytes to a given byte slice and returns the result
func AppendStringOrBytes[S string | []byte](src []byte, elem S) []byte {
	return append(src, elem...)
}
