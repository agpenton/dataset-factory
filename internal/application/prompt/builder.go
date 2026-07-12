package prompt

import "fmt"

func BuildInstructionFromAnswer(answer string) string {
	return fmt.Sprintf(
		"Generate the user instruction that would produce the following answer:\n\n%s",
		answer,
	)
}
