package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/spf13/viper"
)

type getbody struct { // 用于将string格式转成json格式，取出token
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type msgbody struct {
	Projectname 	string `json:"projectname"`
	Builddetails	string `json:"builddetails"`
	Buildbranch	    string `json:"buildbranch"`
	Buildusers		string `json:"buildusers"`
	Buildtime		string `json:"buildtime"`
	Buildresult		string `json:"buildresult"`
	Webbookurl       string `json:"webbookurl"`
	Webbooksecrets    string `json:"webbooksecrets"`
}

//项目名称
//构建详情
//构建分支
//构建用户
//构建时间
//构建结果
//webbook地址
//webbook密钥

var viperConfig *viper.Viper
var (
	projectname 	string
	builddetails	string
	buildbranch	    string
	buildusers		string
	buildtime		string
	buildresult		string
	webbookurl       string
	webbooksecrets    string
)

const (
	envwebbookurl = "https://open.feishu.cn/open-apis/bot/v2/hook/60948abd-4d51-4355-aa3d-069d10427810"
	//m = "项目名称:12345\n构建详情:678910\n构建分支:110192********9283\n 构建用户:0001.00\n 构建时间:ss\n 构建结果:推送为测试,请知晓\n"
)



func SendDingMsg(msg msgbody) {
	//fmt.Println(msg)
	webHook := webbookurl

content := `{
  "msg_type": "post",
  "content": {
    "post": {
      "zh-CN": {
        "title": "Workflows Notification",
        "content": [
          [
            {
              "tag": "text",
              "text": "项目名称:` + msg.Projectname + `\n"
            },
            {
              "tag": "text",
              "text": "构建详情:"
            },
            {
              "tag": "a",
              "text": "点击查看\n",
              "href": "` + msg.Builddetails + `"
            },
            {
              "tag": "text",
              "text": "构建分支:` + msg.Buildbranch + `\n"
            },
            {
              "tag": "text",
              "text": "构建用户:` + msg.Buildusers + `\n"
            },
            {
              "tag": "text",
              "text": "构建结果:` + msg.Buildresult + `\n"
            },
            {
              "tag": "text",
              "text": "构建时间:` + msg.Buildtime + `\n"
            }
          ]
        ]
      }
    }
  }
}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")

	//发送请求
	//timeStr:=time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("时间：%s-------开始推送消息-------\n",time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println("开始推送消息-------")
	resp, err := client.Do(req)

	//关闭请求
	defer resp.Body.Close()
	//var Resqerr string
	if err != nil {
		// handle error
		log.Fatal(err)
		fmt.Println("resperror")
		//Resqerr = "resperror"
	}
	body, err := ioutil.ReadAll(resp.Body)

	var bodyJson getbody
	bodystr := string(body)
	errJson := json.Unmarshal([]byte(bodystr), &bodyJson) // 将string 格式转成json格式
	if errJson != nil {
		failOnError(err, "Failed get body ")
	}
	fmt.Println(bodyJson.Errcode )
	//fmt.Println("resp Body is:" + bodystr)
	if bodyJson.Errcode == 0 {
		fmt.Printf("时间：%s-------消息推送成功-------\n",time.Now().Format("2006-01-02 15:04:05"))

		//fmt.Println("消息推送成功-------")
	} else {
		fmt.Printf("时间：%s-------消息推送失败-------\n",time.Now().Format("2006-01-02 15:04:05"))
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var  msgbodyvar msgbody
	msgbodyvar.Projectname = projectname
	msgbodyvar.Webbookurl = webbookurl
	msgbodyvar.Buildbranch = buildbranch
	msgbodyvar.Buildusers = buildusers
	msgbodyvar.Buildtime = buildtime
	msgbodyvar.Buildresult = buildresult
	//a := fmt.Sprintf("%s", projectname)
	SendDingMsg(msgbodyvar)

}

func init() {
	flag.StringVar(&projectname, "name", "Unknown", "-name 指定项目名称")
	flag.StringVar(&webbookurl, "url", envwebbookurl, "-url 指定webbookurl")
	flag.StringVar(&buildbranch, "branch", "master", "-branch 指定branch")
	flag.StringVar(&buildusers, "user", "Unknown", "-user 指定发起者")
	flag.StringVar(&buildtime, "time", "Unknown", "-time 指定时间")
	flag.StringVar(&buildresult, "result", "Unknown", "-result 指定结果")
	flag.Parse()
}
