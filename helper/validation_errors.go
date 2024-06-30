package helper

type ValidationErrorsMap map[string]map[string]string

func ValidationErrors(input any) (errors ValidationErrorsMap) {
	errors = make(ValidationErrorsMap, 0)
	if rules, ok := input.(map[string]any); ok {
		for fieldName, list := range rules {
			if listString, ok := list.(map[string]any); ok {
				errors[fieldName] = make(map[string]string, len(listString))
				for rule, errMessage := range listString {
					errors[fieldName][rule] = errMessage.(string)
				}
			}
		}
	}
	return errors
}
