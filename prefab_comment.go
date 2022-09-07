package main

import (
	"bufio"
	"fmt"
	"io"
	"m2utils"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("  ========== ")
	prefabDir := "../"
	//comment post conent
	timeUnix := strconv.Itoa(int(time.Now().Unix()))
	comment1 := "#prefab Jd Media Games   "
	comment2 := "#prefab  Jd Media Games Inc."
	comment3 := "#prefab  Das Dad is a technology startup based in Sao Paulo, Brazil. Our first product was launched in"

	//add comments, and rename all prefab files
	filepath.Walk(prefabDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if m2utils.GetExtName(info.Name()) != "prefab" {
				return nil
			}
			fmt.Println(path, info.Name())
			source, err := os.Open(path)
			if err != nil {
				fmt.Println(" !read file error  " + path)
			} else {
				defer source.Close()

				buf := bufio.NewReader(source)
				commentAdded1 := false
				commentAdded2 := false
				commentAdded3 := false
				lineNo := 0
				contents := make([]string, 0)
				for {
					a, _, c := buf.ReadLine()
					if c == io.EOF {
						break
					}
					contents = append(contents, string(a))

					if commentAdded1 == false && strings.Contains(string(a), comment1) {
						commentAdded1 = true
					}
					if commentAdded2 == false && strings.Contains(string(a), comment2) {
						commentAdded2 = true
					}
					if commentAdded3 == false && strings.Contains(string(a), comment3) {
						commentAdded3 = true
					}
					lineNo++
				}

				if commentAdded1 == false {
					contents[2] += "\n" + comment1 + timeUnix
				}

				if commentAdded2 == false {
					m := int(math.Floor(float64(lineNo / 2)))
					contents[ m] += "\n" + comment2 + timeUnix
				}
				if commentAdded3 == false {
					contents[lineNo-1] += "\n" + comment3 + timeUnix
				}

				if commentAdded1 == false || commentAdded2 == false || commentAdded3 == false {

					fileContent := strings.Join(contents, "\n")
					//fmt.Println(fileContent)
					fw, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
					w := bufio.NewWriter(fw)
					w.WriteString(fileContent)
					if err != nil {
						fmt.Println("write content errr : " + path + " , " + err.Error())

					}
					w.Flush()
				}
			}
			source.Close()
		}
		return nil
	})
}
