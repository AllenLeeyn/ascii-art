package generator

import (
	"ascii-art/pkg/fileMgr"

	"fmt"
	"strings"
)

var style = make(map[rune][]string)

// GenArt generate ascii art string by combining [][8]artStr into one string.
func GenArt(txt, styleNm string) string {
	getStyle(styleNm)
	for _, rn := range txt {
		if rn < 32 || rn > 127 {
			fmt.Println("Character not an ASCII char:", string(rn))
			return ""
		}
	}
	txtLns := strings.Split(txt, "\\n")
	return genArtStrs(txtLns)
}

// getStyle read <styleName>.txt and store the ascii art runes in a map[rune][]string.
func getStyle(styleNm string) {
	rawStyle := strings.Split(fileMgr.ReadFile("./assets/"+styleNm+".txt"), "\n")
	for i := 1; i < len(rawStyle); i = i + 9 {
		curChar := rune(32 + i/9)
		style[curChar] = rawStyle[i : i+8]
	}
}

// genArtStr generate [8]string of ascii art for each txtLn.
func genArtStrs(txtLns []string) string {
	art := ""
	for i, txtLn := range txtLns {
		artStrs := [8]string{}
		if txtLn == "" {
			if i > 0 {
				art += "\n"
			}
			continue
		}
		for _, rn := range txtLn {
			for i := 0; i < 8; i++ {
				artStrs[i] += style[rn][i]
			}
		}
		for i := range 8 {
			art += artStrs[i] + "\n"
		}
	}
	return art
}
