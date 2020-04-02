package emoji

import (
	"fmt"
	"testing"
)

func TestDelEmoji(t *testing.T) {
	fmt.Println(DelEmoji("test -> 😅🤗 hello 😅🤗 <- test"))
}

func TestEncode(t *testing.T) {
	fmt.Println(Encode("test -> 😅🤗 hello 😅🤗 <- test"))
}

func TestDecode(t *testing.T) {
	fmt.Println(Decode(Encode("test -> 😅🤗 hello 😅🤗 <- test")))
}
