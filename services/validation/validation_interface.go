package validation

type Validation interface {
	Validate(i interface{}) error
}
