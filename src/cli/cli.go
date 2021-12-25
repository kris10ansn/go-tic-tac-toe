package cli

import (
	"fmt"
	"os"
)

func InputByte() (byte, error) {
	var result byte
	_, err := fmt.Scanf("%d\n", &result)

	return result, err
}

func PromptByteInput(message string) (byte, error) {
	fmt.Print(message)
	return InputByte()
}

func GetCommandLineArgument(index int, defaultValue string) string {
	if len(os.Args) >= index {
		return os.Args[index]
	} else {
		return defaultValue
	}
}
