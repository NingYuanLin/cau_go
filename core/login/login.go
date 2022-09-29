package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NingYuanLin/cau_go/core/status"
	"github.com/NingYuanLin/cau_go/utils"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
)

func Run() error {
	// 检查格式
	err := checkFormat()
	if err != nil {
		return err
	}
	loginInfo, err := status.GetLoginInfo()

	if err != nil {
		return err
	}
	//fmt.Println(loginInfo)
	//fmt.Println(reflect.TypeOf(loginInfo))

	if loginInfo.Login == 1 {
		fmt.Println("账号已经登录，当前登录的用户为:", loginInfo.Username)
		return nil
	}
	username := viper.GetString("username")
	password := viper.GetString("password")
	realIp, err := status.GetRealIp()
	if err != nil {
		return err
	}
	isSuccessBool, errMsg, err := doLogin(username, password, realIp)
	if err != nil {
		return err
	}
	// 在部分网关下，需要在username后加上"@cau"才能正确登录
	if isSuccessBool == false {
		isSuccessBool, errMsg, err = doLogin(username+"@cau", password, realIp)
	}

	if isSuccessBool == false {
		fmt.Println("登录失败:", errMsg)
	} else {
		fmt.Println("登录成功")
	}

	return nil
}

// LoginResponse TODO:为什么：Name starts with the package name？
type LoginResponse struct {
	Result int    `json:"result"`
	Aolno  int    `json:"aolno"`
	M46    int    `json:"m46"`
	V46IP  string `json:"v46ip"`
	Myv6IP string `json:"myv6ip"`
	Sms    int    `json:"sms"`
	Nid    string `json:"NID"`
	Olmac  string `json:"olmac"`
	Ollm   int    `json:"ollm"`
	Olm1   string `json:"olm1"`
	Olm2   string `json:"olm2"`
	Olm3   int    `json:"olm3"`
	Olmm   int    `json:"olmm"`
	Olm5   int    `json:"olm5"`
	Gid    int    `json:"gid"`
	Oltime int    `json:"oltime"`
	Olflow int64  `json:"olflow"`
	Lip    string `json:"lip"`
	Stime  string `json:"stime"`
	Etime  string `json:"etime"`
	UID    string `json:"uid"`
	Sv     int    `json:"sv"`
	Msga   string `json:"msga"` // 当登录失败时才有这个
}

func doLogin(username string, password string, realIp string) (isSuccessBool bool, errorMsg string, err error) {
	err = nil

	url := fmt.Sprintf("http://%s/drcom/login?callback=dr1003&DDDDD=%s&upass=%s&0MKKey=123456&R1=0&R3=0&R6=0&para=00&v6ip=&v=6817", realIp, username, password)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Referer", "http://10.3.38.7/a79.htm?isReback=1")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Add("Cookie", "PHPSESSID=mve8c7kvtso03dqcknjsu22bv1")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	// dr1003({"result":1,"aolno":18263,"m46":0,"v46ip":"10.6.109.242","myv6ip":"","sms":0,"NID":"","olmac":"d0817ac0eb64","ollm":0,"olm1":"00000000","olm2":"0000","olm3":0,"olmm":2,"olm5":0,"gid":24,"oltime":172800000,"olflow":4294967295,"lip":"","stime":"","etime":"","uid":"S20213081520@cau","sv":0})
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	// gbk to utf8
	body, err = utils.GbkToUtf8(body)
	if err != nil {
		return
	}
	bodyString := string(body)
	bodyString = strings.TrimSpace(bodyString)
	startIndex := strings.Index(bodyString, "(") + 1
	bodyString = bodyString[startIndex : len(bodyString)-1]

	var loginResponse LoginResponse
	err = json.Unmarshal([]byte(bodyString), &loginResponse)
	if err != nil {
		return
	}
	isSuccess := loginResponse.Result
	if isSuccess == 0 {
		isSuccessBool = false
	} else {
		isSuccessBool = true
	}
	errorMsg = loginResponse.Msga // 只有当请求错误即result=0时，才有这个
	return isSuccessBool, errorMsg, nil
}

func checkFormat() error {
	username := viper.Get("username")
	password := viper.Get("password")
	if username == "" || password == "" {
		return errors.New("需要同时指定username和password或使用cau_go config -c创建配置文件")
	}
	return nil
}
