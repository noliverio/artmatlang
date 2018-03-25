package artmatlang

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
