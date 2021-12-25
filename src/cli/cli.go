package cli

import "fmt"

func InputByte() (byte, error) {
	var result byte
	_, err := fmt.Scanf("%d\n", &result)

	return result, err
}

func PromptByteInput(message string) (byte, error) {
	fmt.Print(message)
	return InputByte()
}
