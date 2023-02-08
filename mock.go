package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"xmirror.cn/iast/goat/model"
)

const (
	BaseUrl    = "http://192.168.172.30:8080"
	LoginUrl   = BaseUrl + "/iast/api-v1/system/user/login"
	GetNodeUrl = BaseUrl + "/iast/api-v1/app/application/node/select"
	GetVulUrl  = BaseUrl + "/iast/api-v1/node/common/vul/list"
	DelVulUrl  = BaseUrl + "/iast/api-v1/vul/result/delete"
	UserName   = "goat" //"xmirror"
	PassWord   = "Xmirror!@#123"
	XmNodeID   = "FGEQY3PWK3816XJD"
	SleepTime  = 30 * time.Second
)

var (
	token  string
	nodeId string
	Ip     = flag.String("ip", "", "IP of the target host")
	Port   = flag.Int("p", 0, "Port of the target site")
)

func clientDo(request *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	resBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return resBytes, nil
}

func getToken() error {
	params := map[string]interface{}{
		"username":   UserName,
		"password":   PassWord,
		"forceLogin": 1,
		"rememberMe": 0,
		"openLogin":  1,
	}
	paramBytes, _ := json.Marshal(params)

	jsonParam := strings.NewReader(string(paramBytes))
	request, err := http.NewRequest(http.MethodPost, LoginUrl, jsonParam)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	resBytes, err := clientDo(request)
	if err != nil {
		return err
	}
	var logInfo model.LoginInfo
	err = json.Unmarshal(resBytes, &logInfo)
	if err != nil {
		return err
	}
	if logInfo.Code != 0 {
		return errors.New(logInfo.Message)
	}
	token = logInfo.Data.Token
	return nil
}

func getNodeId() error {
	req, err := http.NewRequest(http.MethodGet, GetNodeUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	params := req.URL.Query()
	params.Add("status", "")
	params.Add("applicationId", "")
	req.URL.RawQuery = params.Encode()
	resBytes, err := clientDo(req)
	if err != nil {
		return err
	}
	var nodeList model.NodeInfo
	err = json.Unmarshal(resBytes, &nodeList)
	if err != nil {
		return err
	}
	if nodeList.Code != 0 {
		return errors.New(nodeList.Message)
	}
	for _, node := range nodeList.Data {
		if node.XmNodeID == XmNodeID {
			nodeId = node.NodeID
			return nil
		}
	}
	return errors.New("failed to find the node id")
}

func getNodeVul() ([]model.Records, error) {
	var vulRecords []model.Records
	req, err := http.NewRequest(http.MethodGet, GetVulUrl, nil)
	if err != nil {
		return vulRecords, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	params := req.URL.Query()
	params.Add("nodeId", nodeId)
	params.Add("type", "1")
	params.Add("pageSize", "100")
	params.Add("detectEngineIdList", "6")
	req.URL.RawQuery = params.Encode()
	resBytes, err := clientDo(req)
	if err != nil {
		return vulRecords, err
	}
	var vulInfo model.VulInfo
	err = json.Unmarshal(resBytes, &vulInfo)
	if err != nil {
		return vulRecords, err
	}
	if vulInfo.Code != 0 {
		return vulRecords, errors.New(vulInfo.Message)
	}
	vulRecords = vulInfo.Data.Records
	return vulRecords, nil
}

func delNodeVul(vulIDs map[string]bool) error {
	if len(vulIDs) <= 0 {
		return nil
	}
	request, err := http.NewRequest(http.MethodPost, DelVulUrl, nil)
	if err != nil {
		return err
	}
	var vulResultIdList = request.URL.Query()
	for k, _ := range vulIDs {
		vulResultIdList.Add("vulResultIdList", k)
	}
	request.URL.RawQuery = vulResultIdList.Encode()
	request.Header.Add("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	resBytes, err := clientDo(request)
	if err != nil {
		return err
	}
	var delVul model.DelVul
	err = json.Unmarshal(resBytes, &delVul)
	if err != nil {
		return err
	}
	if delVul.Code != 0 {
		return errors.New(delVul.Message)
	}
	return nil
}

func getConfigData() (configVulData []model.ConfigVulData, err error) {

	file, err := os.OpenFile("mock.json", os.O_RDONLY, 0666)
	if err != nil {
		return configVulData, err
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return configVulData, err
	}
	var configVul model.ConfigVul
	err = json.Unmarshal(fileBytes, &configVul)
	if err != nil {
		return configVulData, err
	}
	configVulData = configVul.Data
	return configVulData, nil
}

func prepare() error {
	err := getToken()
	if err != nil {
		return err
	}
	err = getNodeId()
	if err != nil {
		return err
	}
	vulRecords, err := getNodeVul()
	if err != nil {
		return err
	}
	var vulIDs = make(map[string]bool)
	for _, record := range vulRecords {
		vulIDs[record.VulResultID] = true
	}
	return delNodeVul(vulIDs)
}

func sendRequest() error {
	args := []string{"run", ".", "-c=1", "-n=2"}
	if len(*Ip) > 0 {
		if parsed := net.ParseIP(*Ip); parsed != nil {
			args = append(args, "-ip="+parsed.String())
		}
	}
	args = append(args, "-p="+strconv.Itoa(*Port))
	cmd := exec.Command("go", args...)
	cmd.Dir = "./cmd/test/"
	return cmd.Run()
}

func printAgentVersion() string {
	cmd := exec.Command("xmirror-config", "-v")
	cmd.Dir = "."
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed to obtain the Agent version  Procedure %s \n", err)
	}
	return out.String()
}

func verificationVul() {
	vulRecords, err := getNodeVul()
	if err != nil {
		fmt.Printf("get vul info failure : %s \n", err.Error())
	}
	fmt.Printf("successfully acquired %d vulnerabilities \n", len(vulRecords))
	configVulData, err := getConfigData()
	if err != nil {
		fmt.Printf("get config failure : %s \n", err.Error())
	}
	fmt.Println(printAgentVersion())
	fmt.Printf("==================================假阴性==========================================\n")
	for _, vul := range configVulData {
		var through bool
		for _, record := range vulRecords {
			if strings.Contains(record.VulURL, vul.VulURL) && record.VulName == vul.VulName {
				through = true
			}
		}
		if !through {
			fmt.Printf("验证假阴性: unsafe 靶点: %s 未检测到漏洞 %s \n", vul.VulURL, vul.VulName)
		}
	}
	fmt.Printf("\n==================================假阳性==========================================\n")
	for _, record := range vulRecords {
		if strings.Contains(record.VulURL, "/safe/") {
			fmt.Printf("验证假阳性: safe 靶点: %s 检测到漏洞 %s \n", record.VulURL, record.VulName)
		}
	}
}

func main() {
	flag.Parse()
	// 1.等待靶机启动上线
	fmt.Println("Wait for the target to go online")
	time.Sleep(SleepTime)
	err := prepare()
	if err != nil {
		fmt.Printf("failed to verify the preliminary preparation, err:%s \n", err)
		return
	}
	fmt.Println("Successfully deleted the original vulnerability list")
	err = sendRequest()
	if err != nil {
		fmt.Printf("target aircraft request failed %s \n", err)
		return
	}
	fmt.Printf("After sending the request, wait for the vulnerability to enter the repository\n")
	//2. 等待漏洞入库
	time.Sleep(SleepTime)
	verificationVul()
}
