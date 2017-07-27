package haruka

import (
	"strconv"
	"strings"
	db "github.com/hudson6666/nanami/database"
	"github.com/hudson6666/nanami/config"
	"fmt"
)

const subVersion = "2"

func HandleCommand(cmd string, cmdArgs string, from int, chat int64) (res string) {
	switch cmd {
	case "version":
		res = "Nanami Haruka/Milestone 1 ver" + config.Version + "." + subVersion + "." + config.Build
	case "バージョン":
		res = "ななみ·はるか/マイルストーン１ ver" + config.Version + "." + subVersion + "." + config.Build
	case "版本":
		res = "七海春歌/初代 ver" + config.Version + "." + subVersion + "." + config.Build
	case "memo":
		res = handleMemo(cmdArgs, chat)
	default:
		if ret, msg := HandleText(cmd+" "+cmdArgs, from); ret {
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
		db.Set("person", strconv.Itoa(from), person)
		return
	}
	res = false
	return
}

func meetPerson(text string, from int) (res bool, msg string, person Person) {
	db.Get("person", strconv.Itoa(from), &person)
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

func handleMemo(cmd string, chat int64) (res string) {
	cmd = strings.Replace(cmd, "@nanami_nanabot", "", -1)
	ls := strings.Fields(cmd)
	if len(ls) == 0 {
		ls = append(ls, "")
	}
	tagFlag := false
	archFlag := false
	switch ls[0] {
	case "add":
		if len(ls) == 1 {
			res = "没有要 memo 的内容呀"
			return
		} else {
			var t string
			for k, v := range ls {
				if k >= 1 {
					t += v
					t += " "
				}
			}
			id := db.SetIncr("memo-incr")
			db.Push("memo", strconv.FormatInt(chat, 10), strconv.Itoa(id))
			var md MemoDetail = MemoDetail{Content:t}
			db.Set("memo-detail", strconv.Itoa(id), md)
		}
	case "tag":
		if len(ls) <= 2 {
			res = "没有要 tag 的内容呀"
			return
		} else {
			i, _ := strconv.Atoi(ls[1])
			t := ls[2]
			var memos Memo
			db.List("memo", strconv.FormatInt(chat, 10), &memos)
			i = len(memos) - 1 - i
			if i < 0 || i >= len(memos) {
				res = "没有这个索引诶"
				return
			}
			db.Push("tags", memos[i], t)
			db.Push("tag", t + "@" + strconv.FormatInt(chat, 10), memos[i])
		}
	case "rmtag":
		if len(ls) <= 2 {
			res = "没有要移除的 tag 呀"
			return
		} else {
			i, _ := strconv.Atoi(ls[1])
			t := ls[2]
			var memos Memo
			db.List("memo", strconv.FormatInt(chat, 10), &memos)
			i = len(memos) - 1 - i
			if i < 0 || i >= len(memos) {
				res = "没有这个索引诶"
				return
			}
			db.Remove("tags", memos[i], t)
			db.Remove("tag", t + "@" + strconv.FormatInt(chat, 10), memos[i])
		}
	case "append":
		if len(ls) <= 2 {
			res = "没有要追加的内容呀"
			return
		} else {
			i, _ := strconv.Atoi(ls[1])
			var memos Memo
			db.List("memo", strconv.FormatInt(chat, 10), &memos)
			i = len(memos) - 1 - i
			if i < 0 || i >= len(memos) {
				res = "没有这个索引诶"
				return
			}
			var t string
			for k, v := range ls {
				if k >= 2 {
					t += v
					t += " "
				}
			}
			var md MemoDetail
			db.Get("memo-detail", memos[i], &md)
			md.Content += t
			db.Set("memo-detail", memos[i], md)
		}
	case "arch":
		if len(ls) <= 1 {
			archFlag = true
		} else {
			i, _ := strconv.Atoi(ls[1])
			var memos Memo
			db.List("memo", strconv.FormatInt(chat, 10), &memos)
			i = len(memos) - 1 - i
			if i < 0 || i >= len(memos) {
				res = "没有这个索引诶"
				return
			}
			db.Push("memo-arch", strconv.FormatInt(chat, 10), memos[i])
			db.Remove("memo", strconv.FormatInt(chat, 10), memos[i])
			var tgs Tags
			db.List("tags", memos[i], &tgs)
			for _, t := range tgs {
				db.Remove("tag", t + "@" + strconv.FormatInt(chat, 10), memos[i])
				db.Push("tag-arch", t + "@" + strconv.FormatInt(chat, 10), memos[i])
			}
		}
	case "":
		// No tag
	default:
		tagFlag = true
	}
	var memos Memo
	if archFlag {
		db.List("memo-arch", strconv.FormatInt(chat, 10), &memos)
	} else if tagFlag {
		db.List("tag", ls[0] + "@" + strconv.FormatInt(chat, 10), &memos)
	} else {
		db.List("memo", strconv.FormatInt(chat, 10), &memos)
	}
	for r := len(memos) - 1; r >= 0; r-- {
		k, v := len(memos) - 1 - r, memos[r];
		var md MemoDetail
		db.Get("memo-detail", v, &md)
		if tagFlag {
			res += fmt.Sprintf("(%v): %v", k, md.Content)
		} else {
			res += fmt.Sprintf("%v): %v", k, md.Content)
		}
		if c := db.Count("tags", v); c != 0 {
			res += " ["
			var tgs Tags
			db.List("tags", v, &tgs)
			for i, t := range tgs {
				res += t
				if i != len(tgs) - 1 {
					res += "|"
				} else {
					res += "]"
				}
			}
		}
		res += fmt.Sprintf("\n")
	}
	return
}
