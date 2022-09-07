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

func getNewClassName(name string, version string) string {
	return "BR1" + version + "_" + m2utils.Md5sum(name+version)
}

type TargetFile struct {
	Name    string
	Path    string
	NewName string
	Guid    string
	FullName string
}


//type
type SortByNameLen2 []TargetFile

func (a SortByNameLen2) Len() int {
	return len(a)
}

func (a SortByNameLen2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortByNameLen2) Less(i, j int) bool {
	//sort by file name's length
	return len(a[i].Name) > len(a[j].Name)
}


/*
1. 修改/scripts/stages/cs文件名(xxxx.cs) 为yyyy.cs
   查找xxxx.cs.meta 中的:   guid: abcdef 保存备用
   修改stages/下所有cs文件内public class xxxx 为yyyy
   修改其他cs文件内的关联 public(private) xxxx ...为 yyy
   保存yyy.cs，删除xxxx.cs ,删除xxxx.cs.meta
   开启unity
   读取yyy.cs.meta ，替换所有*.unity *.prefab中 guid: abcedf 保存
*/
func main() {
	//1. read and store all *.unity/*.prefab (YAML format) eg:(  m_Script: {fileID: 11500000, guid: e9d0b5f3bbe925a408bd595c79d0bf63, type: 3})

	//hash sand
	version := "C"
	//find all c# class files in this directory
	cSharpFileDir := "../scripts/Stages/"
	//find all related c# files
	//cSharpAllFileDir := "../scripts/"


	/////////////////class  files
	targetFiles := make([]TargetFile ,0)
	filepath.Walk(cSharpFileDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if m2utils.GetExtName(info.Name()) != "cs" {
				return nil
			}
			//fmt.Println(path, info.Name())
			className := info.Name()[:len(info.Name())-3]

			//ignore some files
			if className == "XXXXX" {
				return nil
			}

			//old guid -> md5(new file name +version)
			newClassName := getNewClassName(className, version) //"A_" + className //+ m2utils.Md5sum(path+version)
			oldFileGuid := m2utils.GetFileGUID(path)
			if oldFileGuid == "" {
				panic("cant get guid " + path)
			}
			targetFiles = append(targetFiles, TargetFile{
				Name:    className,
				Path:    path,
				NewName: newClassName,
				Guid:    oldFileGuid,
				FullName:info.Name(),
			})
		}
		return nil
	})


	//sort on file name's length
	sort.Sort(SortByNameLen2(targetFiles))

	//for _,v:= range targetFiles {
	//	fmt.Println(v)
	//}
	//return

	guidMap := make(map[string]string)
	//find all classname and rename ,then replace related name
	for _,v:= range targetFiles {
		className :=v.Name
		path :=v.Path
		newClassName:=v.NewName
		oldFileGuid:=v.Guid
		fullName := v.FullName

		//find all related class
		cSharpDir := cSharpFileDir

		//you can custom replace files
		//if   className == "SettingPanel" {
		//	cSharpDir = cSharpAllFileDir
		//}
		filepath.Walk(cSharpDir, func(path2 string, info2 os.FileInfo, err error) error {
			if !info2.IsDir() {
				if m2utils.GetExtName(info2.Name()) != "cs" {
					return nil
				}
				//find class name
				contentByte, err := ioutil.ReadFile(path2)
				if err != nil {
					fmt.Println("eeee" + err.Error())
				}
				content := string(contentByte)

				if info2.Name() == fullName {
					//m2 := regexp.MustCompile(`(public|private)(\s+class\s+)(` + className + `)\s+`)
					m2 := regexp.MustCompile(`(public|private)(\s+class\s+)(` + className + `)([\s|:]+)`)
					Str := "${1}${2}" + newClassName + "${4}"
					content = m2.ReplaceAllString(content, Str)
				}

				m2 := regexp.MustCompile(`([\s|.|\(\,\[]+|<)(` + className + `)([\s|.|>|\[\)]+)`)
				//m := regexp.MustCompile("^(.*?)Geeks(.*)$")
				//Str := "${1}GEEKS$2"
				Str := "${1}" + newClassName + "${3}"
				content = m2.ReplaceAllString(content, Str)

				//
				//content = strings.ReplaceAll(content, "Menu."+newClassName, "Menu."+className)
				//content = strings.ReplaceAll(content, "Cache."+newClassName, "Cache."+className)

				f, err := os.Create(path2)
				if err != nil {
					fmt.Println("err2 " + err.Error())
				} else {
					defer f.Close()
					_, err2 := f.WriteString(content)
					if err2 != nil {
						fmt.Println("err3 " + err2.Error())
					}
				}
			}
			return nil
		})

		//rename files
		newPath := strings.ReplaceAll(path, className+".cs", newClassName+".cs")
		fmt.Println("old path " + path + " , new path " + newPath)
		err1 := os.Rename(path, newPath)
		if err1 != nil {
			//panic(err1)
			fmt.Println("rename fail ====================================="+ fullName + err1.Error())
		} else {
			fmt.Println("rename success -> " + fullName)
		}
		guidMap[oldFileGuid] = newPath

	}




	//path2 := "../scripts/Stages/Index/" + getNewClassName("IndexMain", version) + ".cs"
	//contentByte, err := ioutil.ReadFile(path2)
	//if err != nil {
	//	fmt.Println("eeee" + err.Error())
	//}
	//content := string(contentByte)
	//
	//list := []string{"EquipPanel"} //, "AnniversaryPanel"}
	//for _, v := range list {
	//	newClassName := getNewClassName(v, version)
	//	content = strings.ReplaceAll(content, newClassName, v)
	//}
	//// private AnniversaryPanel _anniversaryPanel;
	//
	//f, err := os.Create(path2)
	//if err != nil {
	//	fmt.Println("err2 " + err.Error())
	//} else {
	//	defer f.Close()
	//	_, err2 := f.WriteString(content)
	//	if err2 != nil {
	//		fmt.Println("err3 " + err2.Error())
	//	}
	//}

	//content = strings.ReplaceAll(content,"A_AnniversaryPanel\n","AnniversaryPanel")




	//write to file
	m2utils.SaveMap2File(guidMap)
	fmt.Println("done ,next step Open unity and wait for compile ")
}
