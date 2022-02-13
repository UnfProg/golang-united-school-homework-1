package solution

import "github.com/kyokomi/emoji"

func GetMessage() string {
	str := emoji.Sprintf("Hello :world_map:!")
	return str
}
