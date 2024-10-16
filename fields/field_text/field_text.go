package field_text

type Type struct {
	ReadOnly    bool
	Placeholder string
}

type Options func(*Type)

func WithPlaceholder(placeholder string) Options {
	return func(t *Type) {
		t.Placeholder = placeholder
	}
}

func WithReadOnly(readOnly bool) Options {
	return func(t *Type) {
		t.ReadOnly = readOnly
	}
}

func Make(options ...Options) *Type {
	t := &Type{}
	for _, option := range options {
		option(t)
	}
	return t
}
