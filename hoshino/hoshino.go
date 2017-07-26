package hoshino

import (
	"github.com/hudson6666/nanami/config"
)

const subVersion  = "0"

func HandleCommand(cmd string, cmdArgs string) (res string)  {
	switch cmd {
	case "version":
		res = "Nanami Hoshino/Milestone 2 ver" + config.Version + "." + subVersion + "." + config.Build + "-α"
	case "バージョン":
		res = "ななみ·ほしの/マイルストーン２ ver" + config.Version + "." + subVersion + "." + config.Build + "-α"
	case "版本":
		res = "七海星野/二代 ver" + config.Version + "." + subVersion + "." + config.Build + "-α"
	default:
	}
	return
}