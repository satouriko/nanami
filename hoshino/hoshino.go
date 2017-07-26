package hoshino

func HandleCommand(cmd string, cmdArgs string) (res string)  {
	switch cmd {
	case "version":
		res = "Nanami Hoshino/Milestone 2 ver1.0.2-α"
	case "バージョン":
		res = "ななみ·ほしの/マイルストーン２ ver1.0.2-α"
	case "版本":
		res = "七海星野/二代 ver1.0.2-α"
	default:
	}
	return
}