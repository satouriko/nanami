package haruka

func HandleCommand(cmd string, cmdArgs string) (res string)  {
	switch cmd {
	case "version":
		res = "Nanami Haruka/Milestone 1 ver1.0.1"
	case "バージョン":
		res = "ななみ·はるか/マイルストーン１ ver1.0.1"
	case "版本":
		res = "七海春歌/初代 ver1.0.1"
	default:
		res = "没有这个命令にゃー，你可以猜猜都有什么命令(○'ω'○)丿"
	}
	return
}
