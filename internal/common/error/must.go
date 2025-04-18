package errs

func MustNoErr(err error, msg string) {
	if err != nil {
		panic(NewPanic(500, msg, err))
	}
}
