package illmenu

// Crash out on non-recoverable errors
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
