// A package for working with string slices
package slices

// Returns a boolean value indicating whether the given string exists in the slice or not
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
