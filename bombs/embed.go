package bombs

import "embed"

//go:embebd 1G.gzip
//go:embebd 10G.gzip
//go:embebd 1T.gzip
var Bombs embed.FS

var BombFileNameList = []string{`1G`, `10G`, `1T`}

func Exists(filename string) bool {
	for _, v := range BombFileNameList {
		if v == filename {
			return true
		}
	}
	return false
}
