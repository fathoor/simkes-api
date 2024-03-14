package exception

func PanicIfError(e error) {
	if e != nil {
		switch e.Error() {
		case "duplicated key not allowed":
			panic(BadRequestError{
				Message: "Data already exists",
			})
		case "record not found":
			panic(NotFoundError{
				Message: "Data not found",
			})
		default:
			panic(e)
		}
	}
}
