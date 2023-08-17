package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(text string) string {
	fmt.Printf("%s: ", text)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}
