package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TakeInputFromUser(inputPrompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputPrompt)
	rawInput, _ := reader.ReadString('\n')

	// Trim the input, otherwise it'll contain the input will contain a `\n` character
	trimmedInput := strings.TrimSpace(rawInput)
	return trimmedInput
}
