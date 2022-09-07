package main

import (
	"fmt"
	"io/ioutil"
	"m2utils"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

/**
rename files and waiting unity to generate meta data
 */
func main() {
	//hash  sand
	version := "_C"

	//find all cs file
	cSharpFileDir := "../Scripts/"
	prefabDir := "../Resources/panles/"

	prefabPathMap := make(map[string]string)
	prefabNameMap := make(map[string]string)
	//加入comment内容,并把prefab 文件名改成md5
	//都在同一个目录下，不存在重名的问题
	filepath.Walk(prefabDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if m2utils.GetExtName(info.Name()) != "prefab" {
				return nil
			}
			fileName := info.Name()[0 : len(info.Name())-7  ]
			newName :=  m2utils.Md5sum(strings.ToLower(fileName + version))

			newPath := strings.ReplaceAll(path, fileName+".prefab", newName+".prefab")
			prefabPathMap[m2utils.GetFileGUID(path)] = newPath
			prefabNameMap[fileName] = newName
			fmt.Println(path, info.Name())

			//
			//rename
			if len(fileName) == 32 {

			} else {
				//new prefab name
				//newPath := strings.ReplaceAll(path, fileName, newName)
				fmt.Println("oldpath " + path + " , new path " + newPath)
				err1 := os.Rename(path, newPath)
				if err1 != nil {
					//panic(err1)
					fmt.Println("rename fail =====================================" + info.Name() + err1.Error())
				} else {
					fmt.Println("rename success" + info.Name())
				}
			}
		}
		return nil
	})

	//保存进去等untty生成
	if len(prefabPathMap) > 0 {
		m2utils.SaveMapFile(prefabPathMap, "panelPrefab.json")
	}

	//将名字按长度排序
	oldNameKeys := make([]string ,len(prefabNameMap))
	for oldPrefabName, _ := range prefabNameMap {
		oldNameKeys = append(oldNameKeys ,oldPrefabName)
	}
	sort.Sort(SortByNameLen(oldNameKeys))

	//change all cs files which contains  prefab
	filepath.Walk(cSharpFileDir, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			if m2utils.GetExtName(info.Name()) != "cs" {
				return nil
			}
			contentByte, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("eeee" + err.Error())
			}
			content := string(contentByte)

			//for oldPrefabName, newPrefabName := range prefabNameMap {
			//	//content = strings.ReplaceAll(content, "\"panles/"+oldPrefabName, "\"panles/"+newPrefabName)
			//	re := regexp.MustCompile(`(?i)"panles\/` + oldPrefabName)
			//	content = re.ReplaceAllString(content, "\"panles/"+newPrefabName)
			//}
			for _,oldPrefabName := range oldNameKeys {
				re := regexp.MustCompile(`(?i)"panles\/` + oldPrefabName)
				content = re.ReplaceAllString(content, "\"panles/"+prefabNameMap[oldPrefabName])
			}

			//save
			f, err := os.Create(path)
			if err != nil {
				fmt.Println("save -> open failed :  " + err.Error())
			} else {
				defer f.Close()
				_, err2 := f.WriteString(content)
				if err2 != nil {
					fmt.Println(" write file failed :  " + err2.Error())
				} else {
					fmt.Println("write back " + path)
				}
			}

		}
		return nil
	})
	fmt.Println("open unity and wati to compile complete[new *.prefab.meta] ,and then start step2 ")

}



type SortByNameLen []string

func (a SortByNameLen) Len() int {
	return len(a)
}

func (a SortByNameLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortByNameLen) Less(i, j int) bool {
	//return a[i].Age > a[j].Age //降序
	return len(a[i]) > len(a[j])
}