package forms

type errors map[string][]string

// Add adds a validation error to a field.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error for a field, or an empty string if there are no errors.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
