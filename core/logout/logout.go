package logout

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NingYuanLin/cau_go/core/status"
	"github.com/NingYuanLin/cau_go/utils"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Run() error {
	// 先检查是否已经登录，如果没有登录，也就不需要登出了
	loginInfo, err := status.GetLoginInfo()
	if err != nil {
		return err
	}
	if loginInfo.Login == 0 {
		fmt.Println("当前未登录，不需要登出")
		return nil
	}
	// 登录了，就获取登录名
	ac := loginInfo.Ac
	sessionId, err := getLocalSessionId(ac)
	if err != nil {
		return err
	}
	if sessionId == -1 {
		return errors.New("未在登录设备中找到本机")
	}
	err = forceLogout(ac, sessionId)
	if err != nil {
		return err
	}
	fmt.Println("登出成功")
	return nil
}

func getLocalSessionId(ac string) (sessionId int, err error) {
	// 获取当前用户的sessionId
	// 获取在线设备列表
	sessionId = -1

	liveDevices, err := getLiveDevices(ac)
	if err != nil {
		return
	}

	hostIps, err := utils.GetLocalhostIps()
	if err != nil {
		return
	}
	if len(hostIps) > 1 {
		// TODO:会出现这种情况吗？
		fmt.Println("检测到本机可能有多个网卡")
	}
	for _, liveDevice := range liveDevices {
		loginip := liveDevice.Loginip
		for _, hostIp := range hostIps {
			if loginip == hostIp {
				sessionId = liveDevice.Sessionid
				return
			}
		}
	}
	return
}

func getLiveDevices(ac string) ([]LiveDevicesResponseData, error) {
	var liveDevicesResponseData []LiveDevicesResponseData
	milliTimeStamp := time.Now().UnixMilli()
	milliTimeStampStr := strconv.FormatInt(milliTimeStamp, 10)
	url := fmt.Sprintf("http://10.3.38.7:801/eportal/?c=ServiceInterface&a=loadOnlineDevice&callback=jQuery111307422641681895024_%s&account=%s&_=%s", milliTimeStampStr, ac, milliTimeStampStr)
	response, err := http.Get(url)
	if err != nil {
		return liveDevicesResponseData, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return liveDevicesResponseData, err
	}
	body, err = utils.GbkToUtf8(body)
	if err != nil {
		return liveDevicesResponseData, err
	}
	bodyString := string(body)
	bodyString = strings.TrimSpace(bodyString)
	startIndex := strings.Index(bodyString, "(") + 1
	// jQuery111309758939784360352_1633746542401({"result ... }) => {"result ... }
	bodyString = bodyString[startIndex : len(bodyString)-1]
	// {"result":"ok","msg":"查询在线设备成功！","data":[{"sessionid":1619,"logintime":"2022-09-28 22:05:00","loginip":"10.6.109.242","loginmac":"D0817AC0EB64","devicetype":"PC"}]}
	//fmt.Println(bodyString)
	var liveDevicesResponse LiveDevicesResponse
	err = json.Unmarshal([]byte(bodyString), &liveDevicesResponse)
	if err != nil {
		return liveDevicesResponseData, err
	}
	if liveDevicesResponse.Result != "ok" {
		return liveDevicesResponseData, errors.New(liveDevicesResponse.Msg)
	}
	liveDevicesResponseData = liveDevicesResponse.Data
	return liveDevicesResponseData, nil
}

type LiveDevicesResponseData struct {
	Sessionid  int    `json:"sessionid"`
	Logintime  string `json:"logintime"`
	Loginip    string `json:"loginip"`
	Loginmac   string `json:"loginmac"`
	Devicetype string `json:"devicetype"`
}

type LiveDevicesResponse struct {
	Result string                    `json:"result"`
	Msg    string                    `json:"msg"`
	Data   []LiveDevicesResponseData `json:"data"`
}

func forceLogout(ac string, sessionId int) error {
	var logoutResponse LogoutResponse
	milliTimeStamp := time.Now().UnixMilli()
	milliTimeStampStr := strconv.FormatInt(milliTimeStamp, 10)
	//fmt.Println("milliTimeStampStr", milliTimeStampStr)
	url := fmt.Sprintf("http://10.3.38.7:801/eportal/?c=ServiceInterface&a=offlineUserDevice&callback=jQuery111308603583689084444_%s&account=%s&sessionid=%d&_=%s", milliTimeStampStr, ac, sessionId, milliTimeStampStr)

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	body, err = utils.GbkToUtf8(body)
	if err != nil {
		return err
	}
	// jQuery111308603583689084444_1664378291912({"result":"ok","msg":"设备强制离线成功！"})
	bodyString := string(body)
	bodyString = strings.TrimSpace(bodyString)
	startIndex := strings.Index(bodyString, "(") + 1
	bodyString = bodyString[startIndex : len(bodyString)-1]
	err = json.Unmarshal([]byte(bodyString), &logoutResponse)
	if err != nil {
		return err
	}
	if logoutResponse.Result == "ok" {
		return nil
	} else {
		return errors.New(logoutResponse.Msg)
	}

}

type LogoutResponse struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}
