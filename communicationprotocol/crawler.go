package main

import (
	"compress/gzip"
	"crypto/md5"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

/**
 * Created by @CaomaoBoy on 2020/2/20.
 *  email:<115882934@qq.com>
 */



func ParseWeb(url string) []string{
	fmt.Println(url)
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	//增加header选项
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.5 Safari/605.1.15")
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Host", "www.toutiao.com")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate, br")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	if err != nil{
		fmt.Println("读取网页失败！")
		return []string{}
	}
	body, err := gzip.NewReader(response.Body)
	if err != nil{
		panic(err)
	}
	b,err := ioutil.ReadAll(body)
	if err != nil{
		return []string{}
	}
	fmt.Println(string(b))
	file,err := os.Create("data.txt")
	file.Write(b)
	file.Close()
	//newfile,_ := os.Open("1.txt")
	//doc,err := html.Parse(newfile)
	//tmplist := []string{}
	//for _,link := range visit(nil,doc){
	//	if !strings.Contains(link,"javascript") {
	//		link = strings.Replace(link, "”", "\"", -1)
	//		if len(link) >= 4 && link[:4] != "http" {
	//			link = "http://bbs.tianya.cn" + link
	//		}
	//	}
	//	fmt.Println(link)
	//	tmplist = append(tmplist,link)
	//}

	return nil
}

func readjs(){
	filePath := "/Users/qiao/go/src/qqsort/communicationprotocol/js/signature.js"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	if err!=nil {
		panic(err)
	}

	data := "123234354354"
	//encodeInp是JS函数的函数名
	value, err := vm.Call("get_signature", nil, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(value.String())

}
func getHoney() (string, string) {
	t := time.Now().Unix()
	i := fmt.Sprintf("%X", t)
	h := md5.New()
	h.Write([]byte(strconv.Itoa(int(t))))
	e := fmt.Sprintf("%X", h.Sum(nil))
	if len(i) != 8 {
		return "479BB4B7254C150", "7E0AC8874BB0985"
	}
	var n, l string
	for a, r, s, o := 0, 0, e[:5], e[len(e)-5:]; 5 > a; a++ {
		n += string(s[a]) + string(i[a])
		l += string(i[r+3]) + string(o[r])
		r++
	}
	return fmt.Sprintf("A1%s%s", n, i[len(i)-3:]), fmt.Sprintf("%s%sE1", i[:3], l)
}
func main(){
	//readjs()
//https://www.toutiao.com/api/pc/feed/?min_behot_time=0&category=__all__&utm_source=toutiao&widen=1&tadrequire=true&as=A155CE143E81B9E&cp=5E4EB12B197E7E1&_signature=
	_signature := "sWGJIgAAAABOnnbduiixRrFhiTAAO89"
	as,cp := "A1F5AE342EF6860","5E4E06B826C04E1"
	ParseWeb(fmt.Sprintf("https://www.toutiao.com/api/pc/feed/?min_behot_time=0&category=__all__&utm_source=toutiao&widen=1&tadrequire=true&as=%s&cp=%s&_signature=%s",as,cp,_signature))
	//ParseWeb(fmt.Sprintf("www.toutiao.com/api/pc/feed/?min_behot_time=0&category=__all__&utm_source=toutiao&widen=1&tadrequire=true&as=%s&cp=%s&_signature=yVk1GAAgEBDfITNjGl224slZdAAAJcgmgDM8eFN.LW7kwjXO2fQsk4MoUFdUrb8tgSBCGQmxrMrbz.MjzBycT5TOcVJu1ZVRc4nC5G750qdwWV2XLnaY0wUI22XtlfKFWy6",as,cp))
}