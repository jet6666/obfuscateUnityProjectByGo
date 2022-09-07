package main

import (
	"fmt"
	"io/ioutil"
	"m2utils"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	unityFileDir := "../Stages/"
	panelsDir := "../Prefabs/"
	//1. remember all *.unity/*.prefab (YAML format) eg:(  m_Script: {fileID: 11500000, guid: e9d0b5f3bbe925a408bd595c79d0bf63, type: 3})
	//
	//replace all cs files !!!!
	fmt.Println("  ========== ")
	//version := ""

	fmt.Println("done ,next step Open unity and wait for compile ")

	guidMap := map[string]string{}
	m2utils.GetMapFromFile("map.json", &guidMap)
	//json.Unmarshal([]byte( string(content)), &info)
	//fmt.Println(info)
	for oldFileGuid, newClassPath := range guidMap {
		fmt.Println(oldFileGuid, newClassPath)

		//newClassPath = strings.ReplaceAll(newClassPath, "test2", "Scripts")
		newFileGuid := m2utils.GetFileGUID(newClassPath)
		if oldFileGuid == "" {
			panic("cant get guid " + oldFileGuid)
		}
		fmt.Println(oldFileGuid, newClassPath, newFileGuid)
		guidMap[oldFileGuid] = newFileGuid

	}

	//replace guid related cs file from  *.unity  Stages
	filepath.Walk(unityFileDir, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			if m2utils.GetExtName(info.Name()) != "unity" {
				return nil
			}
			contentByte, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("eeee" + err.Error())
			}
			content := string(contentByte)

			for oldUid, newUid := range guidMap {
				content = strings.ReplaceAll(content, "guid: "+oldUid, "guid: "+newUid)
			}

			//save
			f, err := os.Create(path)
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

	//replace all guid from *.prefab, Prefabs
	filepath.Walk(panelsDir, func(path string, info2 os.FileInfo, err error) error {
		fmt.Println(path)
		if info2 == nil {
			return nil
		}
		if !info2.IsDir() {
			if m2utils.GetExtName(info2.Name()) != "prefab" {
				return nil
			}
			contentByte, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("fffff" + err.Error())
			}
			content := string(contentByte)

			for oldUid, newUid := range guidMap {
				content = strings.ReplaceAll(content, "guid: "+oldUid, "guid: "+newUid)
			}

			//save
			f, err := os.Create(path)
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

	//replace all guid from *.prefab, Prefabs
	/*panelsDir = "../Resources/panles/"
	filepath.Walk(panelsDir, func(path string, info2 os.FileInfo, err error) error {
		fmt.Println(path)
		if info2 == nil {
			return nil
		}
		if !info2.IsDir() {
			if m2utils.GetExtName(info2.Name()) != "prefab" {
				return nil
			}
			contentByte, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("fffff" + err.Error())
			}
			content := string(contentByte)

			for oldUid, newUid := range guidMap {
				content = strings.ReplaceAll(content, "guid: "+oldUid, "guid: "+newUid)
			}

			//save
			f, err := os.Create(path)
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
	})*/

}
