package util

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func GetP(id string) {
	url := "https://oj.czos.cn/p/" + id
	method := "GET"

	payload := &bytes.Buffer{}

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	cookiesB, _ := os.ReadFile("files/cookie.txt")
	cookieString := string(cookiesB)
	//fmt.Println("\nsavedCookie:\n" + cookieString)
	req.Header.Add("Cookie", cookieString)
	res, _ := client.Do(req)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, _ := io.ReadAll(res.Body)

	_ = os.WriteFile("files/p.html", body, 0644)
}

func Tran(id string) {
	txtBytes, _ := os.ReadFile("files/DFBYHead.html")
	txtString := string(txtBytes)
	headList := strings.Split(txtString, "\n")

	txtBytes, _ = os.ReadFile("files/p.html")
	txtString = string(txtBytes)
	bodyList := strings.Split(txtString, "\n")

	var saveList []string

	//indexFile, _ := os.Create("files/index.html")
	for i := 0; i < len(headList)-2; i++ {
		saveList = append(saveList, headList[i]+"\n")
	}

	flag := false
	for i, str := range bodyList {
		if strings.Contains(str, "<div class=\"col-md-9 problem-view\">") {
			flag = true
		}
		if strings.Contains(bodyList[i+1], "<span>来源<i class='fa fa-ellipsis-v pull-right'") {
			flag = false
			break
		}
		if flag {
			saveList = append(saveList, str+"\n")
		}
	}
	saveList = append(saveList, headList[len(headList)-2]+"\n")
	saveList = append(saveList, headList[len(headList)-1]+"\n")
	indexFile, _ := os.Create("files/" + id + ".html")
	_, _ = indexFile.WriteString(strings.Join(saveList, ""))
	_ = indexFile.Close()
}

func GetId(id string, subCode string) string {

	url := "https://oj.czos.cn/p/" + id
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("Solution[language]", "1")
	_ = writer.WriteField("Solution[source]", subCode)
	_ = writer.Close()

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	cookiesB, _ := os.ReadFile("files/cookie.txt")
	cookieString := string(cookiesB)
	//fmt.Println("\nsavedCookie:\n" + cookieString)
	req.Header.Add("Cookie", cookieString)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := client.Do(req)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, _ := io.ReadAll(res.Body)

	txtString := string(body)
	//fmt.Println(txtString)
	//查找第一个"/solution/result?id="
	index := strings.Index(txtString, "/solution/result?id=")
	//id为7位数字
	getId := txtString[index+len("/solution/result?id=") : index+len("/solution/result?id=")+9]
	return getId
}
func GetResult(id string) string {
	url := "https://oj.czos.cn/solution/detail?id=" + id
	method := "GET"
	payload := &bytes.Buffer{}
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	cookiesB, _ := os.ReadFile("files/cookie.txt")
	cookieString := string(cookiesB)
	//fmt.Println("\nsavedCookie:\n" + cookieString)
	req.Header.Add("Cookie", cookieString)
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	res, _ := client.Do(req)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, _ := io.ReadAll(res.Body)

	txtString := string(body)
	//查找第一个<th>C++</th>
	index := strings.Index(txtString, "<th>C++</th>")
	//找到下一个<th> </th>所包裹的值
	index2 := strings.Index(txtString[index+len("<th>C++</th>"):], "<th>") + index + len("<th>C++</th>")
	index3 := strings.Index(txtString[index2+len("<th>"):], "</th>") + index2 + len("<th>")
	//fmt.Println(index, index2, index3)
	tableRes := txtString[index2+len("<th>") : index3]
	//去除空格和空行等
	tableRes = strings.ReplaceAll(tableRes, " ", "")
	tableRes = strings.ReplaceAll(tableRes, "\n", "")
	tableRes = strings.ReplaceAll(tableRes, "\t", "")
	return tableRes
}
