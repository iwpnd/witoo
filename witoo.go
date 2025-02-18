package witoo

import "fmt"

func SayHiTo(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	return fmt.Sprintf("Hi, %s", name), nil
}
