package builtins

func Match(cmd string) (func([]string) error, bool) {
	mp := map[string]func([]string) error{
		"echo": Echo,
		"exit": Exit,
		"type": Type,
	}
	f, ok := mp[cmd]
	return f, ok
}