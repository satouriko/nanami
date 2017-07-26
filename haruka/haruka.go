package haruka

import (
	"strconv"
	"strings"
	db "github.com/hudson6666/nanami/database"
)

func HandleCommand(cmd string, cmdArgs string, from int) (res string)  {
	switch cmd {
	case "version":
		res = "Nanami Haruka/Milestone 1 ver1.1.2"
	case "バージョン":
		res = "ななみ·はるか/マイルストーン１ ver1.1.2"
	case "版本":
		res = "七海春歌/初代 ver1.1.2"
	default:
		if ret, msg := HandleText(cmd + " " + cmdArgs, from); ret {
			res = msg
		} else {
			res = "没有这个命令にゃー，你可以猜猜都有什么命令(○'ω'○)丿"
		}
	}
	return
}

func HandleText(text string, from int) (res bool, msg string) {
	res, msg, person := meetPerson(text, from)
	if res {
		return
	}
	if strings.Contains(text, "ななみ") || strings.Contains(text, "Nanami") || strings.Contains(text, "nanami") || strings.Contains(text, "七海") {
		person.Familiarity++
		msg = "你在叫我嘛，可我不知道你在说什么哦"
		res = true
		return
	}
	res = false
	return
}

func meetPerson(text string, from int) (res bool, msg string, person Person) {
	db.ScanStruct(db.Get("person", strconv.Itoa(from)), &person)
	switch person.Status {
	case "name":
		person.Name = text
		person.Status = ""
		person.Familiarity = 1
		msg = "我是ななみ，请多关照=w="
		res = true
	default:
		if person.Name == "" {
			msg = "初次见面，请问你叫什么名字QwQ～"
			person.Status = "name"
			res = true
		} else {
			res = false
		}
	}
	if res {
		db.Set("person", strconv.Itoa(from), person)
	}
	return
}