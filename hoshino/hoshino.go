package hoshino

func HandleCommand(cmd string, cmdArgs string) (res string)  {
	switch cmd {
	case "version":
		res = "Nanami Hoshino/Milestone 2 ver1.0.1-α"
	case "バージョン":
		res = "ななみ·ほしの/マイルストーン２ ver1.0.1-α"
	case "版本":
		res = "七海星野/二代 ver1.0.1-α"
	default:
		res = HandleText(cmd + " " + cmdArgs)
	}
	return
}

func HandleText(text string) (res string) {
	res = "嗯……"
	return
}
