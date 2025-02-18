package witoo

import "fmt"

func SayHiTo(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("Hi, %s", name)
}
