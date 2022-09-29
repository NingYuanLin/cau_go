package status

import (
	"encoding/json"
	"fmt"
	"github.com/NingYuanLin/cau_go/utils"
	"io"
	"net/http"
	"strings"
)

func Run() error {
	loginInfo, err := GetLoginInfo()
	if err != nil {
		return err
	}
	if loginInfo.Login == 1 {
		fmt.Println("您当前已登录，登录用户为：", loginInfo.Username)
	} else {
		fmt.Println("您当前未登录")
	}
	return nil
}

type responseStatus struct {
	Result int    `json:"result"`
	Time   int    `json:"time"`
	Flow   int    `json:"flow"`
	Fsele  int    `json:"fsele"`
	Fee    int    `json:"fee"`
	M46    int    `json:"m46"`
	V46IP  string `json:"v46ip"`
	Myv6IP string `json:"myv6ip"`
	Oltime int64  `json:"oltime"`
	Olflow int64  `json:"olflow"`
	Lip    string `json:"lip"`
	Stime  string `json:"stime"`
	Etime  string `json:"etime"`
	UID    string `json:"uid"`
	V6Af   int    `json:"v6af"`
	V6Df   int    `json:"v6df"`
	V46M   int    `json:"v46m"`
	V4IP   string `json:"v4ip"`
	V6IP   string `json:"v6ip"`
	Ac     string `json:"AC"`
	Ss5    string `json:"ss5"`
	Ss6    string `json:"ss6"`
	Vid    int    `json:"vid"`
	Ss1    string `json:"ss1"`
	Ss4    string `json:"ss4"`
	Cvid   int    `json:"cvid"`
	Pvid   int    `json:"pvid"`
	Hotel  int    `json:"hotel"`
	Aolno  int    `json:"aolno"`
	Eport  int    `json:"eport"`
	Eclass int    `json:"eclass"`
	Zxopt  int    `json:"zxopt"`
	Nid    string `json:"NID"`
	Olno   int    `json:"olno"`
	Udate  string `json:"udate"`
	Olmac  string `json:"olmac"`
	Ollm   int    `json:"ollm"`
	Olm1   string `json:"olm1"`
	Olm2   string `json:"olm2"`
	Olm3   int    `json:"olm3"`
	Olmm   int    `json:"olmm"`
	Olm5   int    `json:"olm5"`
	Gid    int    `json:"gid"`
	ActM   int    `json:"actM"`
	Actt   int    `json:"actt"`
	Actdf  int    `json:"actdf"`
	Actuf  int    `json:"actuf"`
	Act6Df int    `json:"act6df"`
	Act6Uf int    `json:"act6uf"`
	Allfm  int    `json:"allfm"`
	D1     int    `json:"d1"`
	U1     int    `json:"u1"`
	D2     int    `json:"d2"`
	U2     int    `json:"u2"`
	O1     int    `json:"o1"`
	Nd1    int    `json:"nd1"`
	Nu1    int    `json:"nu1"`
	Nd2    int    `json:"nd2"`
	Nu2    int    `json:"nu2"`
	No1    int    `json:"no1"`
}

type LoginInfo struct {
	Login    int
	Username string
	Ac       string // 用户名
}

var realIp string

func GetRealIp() (string, error) {
	// TODO:检查是否可以正常获取到redirect地址
	if realIp != "" {
		// 单例
		return realIp, nil
	}
	response, err := http.Get("http://10.3.38.7")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	// 10.3.191.8
	realIp = response.Request.URL.Host
	//fmt.Println("realIp:", realIp)
	return realIp, nil
}

func GetCheckStatusInterfaceRet() (string, error) {
	realIp, err := GetRealIp()
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("http://%s/drcom/chkstatus?callback=dr1002&v=10311", realIp)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Referer", "http://10.3.38.7/")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Add("Cookie", "PHPSESSID=mve8c7kvtso03dqcknjsu22bv1")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// gbk to utf8
	body, err = utils.GbkToUtf8(body)
	if err != nil {
		return "", err
	}

	bodyStr := string(body)

	//fmt.Println("bodyStr:", bodyStr)

	//fmt.Println(bodyStr)
	return bodyStr, nil
}

func GetLoginInfo() (LoginInfo, error) {
	var loginInfo LoginInfo

	statusString, err := GetCheckStatusInterfaceRet()
	if err != nil {
		return loginInfo, err
	}
	// dr1002({"result":1,"time":1080,"flow":194456,"fsele":1,"fee":3400,"m46":0,"v46ip":"10.6.5.92","myv6ip":"","oltime":4294967295,"olflow":4294967295,"lip":"10.6.5.92","stime":"2022-09-27 16:12:43","etime":"2022-09-27 16:17:45","uid":"s20213081520","v6af":0,"v6df":0,"v46m":0,"v4ip":"10.6.5.92","v6ip":"::","AC":"s20213081520","ss5":"10.6.5.92","ss6":"10.3.38.7","vid":0,"ss1":"000d48448ce8","ss4":"000000000000","cvid":0,"pvid":0,"hotel":0,"aolno":31956,"eport":-1,"eclass":1,"zxopt":1,"NID":"��Զ��","olno":0,"udate":"","olmac":"d0817ac0eb64","ollm":10,"olm1":"00080010","olm2":"0000","olm3":0,"olmm":1,"olm5":0,"gid":24,"actM":1,"actt":0,"actdf":0,"actuf":0,"act6df":0,"act6uf":0,"allfm":1,"d1":0,"u1":0,"d2":0,"u2":0,"o1":0,"nd1":580272,"nu1":169036,"nd2":0,"nu2":0,"no1":0})
	statusString = strings.TrimSpace(statusString)
	startIndex := strings.Index(statusString, "(") + 1
	statusString = statusString[startIndex : len(statusString)-1]
	//fmt.Println("loginInfo", statusString)
	responseStatus := responseStatus{}
	err = json.Unmarshal([]byte(statusString), &responseStatus)
	if err != nil {
		return loginInfo, err
	}
	loginInfo.Login = responseStatus.Result
	loginInfo.Username = responseStatus.Nid
	loginInfo.Ac = responseStatus.Ac
	return loginInfo, nil
}
