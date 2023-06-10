package helper

import "github.com/go-playground/validator/v10"

//func ErrRepository(err error) (domain.User, error) {
//	if err != nil {
//		return domain.User{}, errors.New("Repository Error: " + err.Error())
//	}
//}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ValidationError(err error) map[string][]string {
	var sliceErrors []string
	errors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, fieldError := range errors {
			sliceErrors = append(sliceErrors, fieldError.Error())
		}
	}

	return map[string][]string{
		"errors": sliceErrors,
	}
}
