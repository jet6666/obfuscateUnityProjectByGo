package main

import (
	"bufio"
	"fmt"
	"io"
	"m2utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)
/**
add spam code into c#
 */
func main() {
 //hash  sand
	version := "version1"
	//find all cs files in this directory
	cSharpFileDir := "../Scripts/"
	spamTag := "   //  Jd Media Games Inc."

	i := 0
	filepath.Walk(cSharpFileDir, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			if m2utils.GetExtName(info.Name()) != "cs" {
				return nil
			}

			//ignore files
			//if strings.Contains(path, "\\Tools\\") {
			//	return nil
			//}

			fmt.Println(path, info.Name())

			//find content
			source, err := os.Open(path)
			if err != nil {
				fmt.Println(" !open file error : " + path)
			} else {
				defer source.Close()

				buf := bufio.NewReader(source)
				var content = ""
				var replaceAble = false
				for {
					a, _, c := buf.ReadLine()
					if c == io.EOF {
						break
					}

					lineContent := string(a)
					lineContentAppend := ""
					if !strings.Contains(lineContent, spamTag) {

						//find all fields .(public /private ) .............
						reg1 := regexp.MustCompile(`(public|private)(\s*)([a-zA-Z]+)(\s*)([a-zA-Z0-9]+);`)
						if reg1 == nil {
							fmt.Println("regexp errr -------------")
							//content += lineContent + "\n"
						} else {
							result1 := reg1.FindAllStringSubmatch(lineContent, -1)
							//fmt.Println("result1 = ", result1 ,lineContent)
							if len(result1) > 0 {
								len2 := len(result1[0])
								if len(result1) >= 1 && len2 >= 6 {
									//fileName := result1[0][1]
									fmt.Println(" result2 = ", " ="+result1[0][5]+"= ", strings.Join(result1[0], ";"), "===========", lineContent)
									//fieldName := result1[0][5]
									//content += "private bool C" + Md5sum(fieldName) + "; \n " + lineContent + "\n"
									i++
									switch i % 5 {
									case 1:
										lineContentAppend = "private bool ux9" + m2utils.Md5sum(string(i)+version) + "; \n "
									case 2:
										lineContentAppend = "private string opx9" + m2utils.Md5sum(string(i)+version) + "; \n "
									case 3:
										lineContentAppend = "public int Eaa3" + m2utils.Md5sum(string(i)+version) + "; \n "
									case 4:
										lineContentAppend = "private byte di99a4" + m2utils.Md5sum(string(i)+version) + "; \n "
									case 0:
										lineContentAppend = "public bool Retry5" + m2utils.Md5sum(string(i)+version) + "; \n "
									}
									replaceAble = true
								} else {
									//content += lineContent + "\n"
								}
							} else {
								//content += lineContent + "\n"
							}
						}

						//find all methods .(public /private ) .............
						reg2 := regexp.MustCompile(`(public|private)(\s*)([a-zA-Z]+)(\s*)([a-zA-Z0-9]+)([a-zA-Z0-9\s\(]+)\)`)
						if reg2 == nil {
							fmt.Println("regexp errr -------------")
							//content += lineContent + "\n"
						} else {
							result1 := reg2.FindAllStringSubmatch(lineContent, -1)
							//fmt.Println("result1 = ", result1 ,lineContent)
							if len(result1) > 0 {
								len2 := len(result1[0])
								if len(result1) >= 1 && len2 >= 2 {
									//fileName := result1[0][2]
									fmt.Println("="+result1[0][5]+"= ALL methods  = ", len(result1), len(result1[0]), strings.Join(result1[0], " ;"), "===========", lineContent)
									//methodName := result1[0][5]
									//content += "private void MM" + Md5sum(methodName) + "; \n " + lineContent + "\n"
									lineContentAppend = "public  void Klio" + m2utils.Md5sum(string(i)+version) + "() {} ; \n "
									i++
									switch i % 7 {
									case 1:
										lineContentAppend = "private bool mioa" + m2utils.Md5sum(string(i)+version) + "(int a){ return false ;} \n "
									case 2:
										lineContentAppend = "public string yu71a" + m2utils.Md5sum(string(i)+version) + "(){ return \"\" ;}  \n "
									case 3:
										lineContentAppend = "private int nili3" + m2utils.Md5sum(string(i)+version) + "(string b) { return 0 ;} \n "
									case 4:
										lineContentAppend = "public void BO904" + m2utils.Md5sum(string(i)+version) + "(){ } \n "
									case 5:
										lineContentAppend = "private void nopio" + m2utils.Md5sum(string(i)+version) + "(bool e) { } \n "
									case 6:
										lineContentAppend = "public void Bviooas" + m2utils.Md5sum(string(i)+version) + "(){ } \n "
									case 0:
										lineContentAppend = "public bool YTdaasdf" + m2utils.Md5sum(string(i)+version) + "(string x){ return true ;} \n "
									}
									replaceAble = true
								} else {
									//content += lineContent + "\n"
								}
							} else {
								//content += lineContent + "\n"
							}
						}

						content += lineContentAppend + lineContent + "\n"

					} else {
						break
					}
				}

				//write back
				if replaceAble == true {
					fmt.Println(path, "========================replace========")
					content2 := spamTag + "\n" + content
					//fmt.Println(content)
					fw, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
					w := bufio.NewWriter(fw)
					w.WriteString(content2)
					if err != nil {
					}
					w.Flush()
				}

			}
			//for test once
		}
		return nil
	})
}
