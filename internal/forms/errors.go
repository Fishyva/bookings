package forms

//Making a type error thats a map that takes a string key to string slice
type errors map[string][]string

// Add adds an error message for a given form field
func (e errors) Add(field,message string) {
	// adding error with its field and its message to error map
	e[field] = append(e[field], message)
}
// Get returns the first error message
func (e errors) Get(field string) string {
	// Checking if the field has errors
	errorString := e[field]
	if len(errorString) == 0 {
		return ""
	}
	// Returning the first error
	return errorString[0]
}