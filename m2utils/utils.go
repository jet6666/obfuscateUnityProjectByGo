package m2utils

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func Md5sum(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}

func GetExtName(str string) string {
	list := strings.Split(str, ".")
	len := len(list)
	if len > 1 {
		ext := make([]string, 0)
		for i := 1; i < len; i++ {
			ext = append(ext, list[i])
		}
		return strings.Join(ext, ".")

	} else {
		return ""
	}
}

func GetFileGUID(csPath string) string {
	path := csPath + ".meta"
	source, err := os.Open(path)
	if err != nil {
		fmt.Println(" !读取源文件出错: " + path)
	} else {
		defer source.Close()

		buf := bufio.NewReader(source)
		for {
			a, _, c := buf.ReadLine()
			if c == io.EOF {
				break
			}
			lineContent := string(a)
			reg2 := regexp.MustCompile(`guid:\s*([a-zA-Z0-9]+)`)
			if reg2 == nil {
				fmt.Println("regexp errr -------------")
			} else {
				result1 := reg2.FindAllStringSubmatch(lineContent, -1)
				if len(result1) > 0 {
					len2 := len(result1[0])
					if len2 >= 1 && len(result1[0]) > 1 {
						//fmt.Println("result1 ", strings.Join(result1[0], "="))
						//fmt.Println(result1[0][1])
						return result1[0][1]
					}
				}
			}
		}
	}
	return ""
}

func SaveMap2File(guidMap map[string]string) {
	jsonStr, err := json.Marshal(guidMap)
	if err != nil {
		fmt.Println("errr " + err.Error())
	} else {
		fmt.Println("ok ")
		f, err := os.Create("map.json")
		if err != nil {
			fmt.Println("err2 " + err.Error())
		} else {
			defer f.Close()
			_, err2 := f.WriteString(string(jsonStr))
			if err2 != nil {
				fmt.Println("err3 " + err2.Error())
			}
		}
	}
}

func SaveMapFile(guidMap map[string]string ,fileName string ) {
	jsonStr, err := json.Marshal(guidMap)
	if err != nil {
		fmt.Println("errr " + err.Error())
	} else {
		fmt.Println("ok ")
		f, err := os.Create( fileName)
		if err != nil {
			fmt.Println("err2 " + err.Error())
		} else {
			defer f.Close()
			_, err2 := f.WriteString(string(jsonStr))
			if err2 != nil {
				fmt.Println("err3 " + err2.Error())
			}
		}
	}
}

func GetMapFromFile(filename string, info interface{}) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("eeee" + err.Error())
	}
	//fmt.Println(string(content))

	//info := map[string]string{}
	json.Unmarshal([]byte( string(content)), &info)
	//fmt.Println(info)
	//for k, v := range info {
	//	fmt.Println(k, v)
	//}
}
