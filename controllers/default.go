package controllers

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
	"github.com/xyzj/mxgo"
)

const (
	urlVpsStatus = "https://api.64clouds.com/v1/getServiceInfo?veid=%s&api_key=%s"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
}

func (c *MainController) Post() {
	scode := c.GetString("scode")
	if len(strings.TrimSpace(scode)) == 0 {
		c.Ctx.WriteString("ok")
		c.StopRun()
	}
	t := time.Now().UTC()
	md5ctx := md5.New()
	md5ctx.Write([]byte(fmt.Sprintf("%04d-%02d-%02d updatemycloudipaddress %02d:00:00", t.Year(), t.Month(), t.Day(), t.Hour())))
	cs := md5ctx.Sum(nil)
	if hex.EncodeToString(cs) != scode {
		c.StopRun()
	}
	// ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	f, ex := os.OpenFile(".ipcache", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer f.Close()
	if ex == nil {
		f.WriteString(strings.Split(c.Ctx.Request.RemoteAddr, ":")[0])
	}
	c.Ctx.WriteString("success")
}

func (c *MainController) WebTools() {
	f, ex := os.Open(".ipcache")
	defer f.Close()
	if ex == nil {
		b := make([]byte, 15)
		if n, ex := f.Read(b); ex == nil {
			c.Data["hostname"] = strings.TrimSpace(string(b[:n]))
			c.TplName = "webtool.tpl"
			c.Render()
		}
	}
}

func (c *MainController) IPInfo() {
	println("in ipinof")
	f, ex := os.Open(".ipcache")
	defer f.Close()
	if ex == nil {
		b := make([]byte, 15)
		if n, ex := f.Read(b); ex == nil {
			c.Ctx.WriteString(strings.TrimSpace(string(b[:n])))
		}
	}
}

func (c *MainController) Ssh() {
	f, ex := os.Open(".ipcache")
	defer f.Close()
	if ex == nil {
		b := make([]byte, 15)
		if n, ex := f.Read(b); ex == nil {
			c.Redirect(fmt.Sprintf("http://%s:1983/ssh/host/127.0.0.1", strings.TrimSpace(string(b[:n]))), 302)
		}
	}
}

func (c *MainController) Kod() {
	f, ex := os.Open(".ipcache")
	defer f.Close()
	if ex == nil {
		b := make([]byte, 15)
		if n, ex := f.Read(b); ex == nil {
			c.Redirect(fmt.Sprintf("http://%s:1981", strings.TrimSpace(string(b[:n]))), 302)
		}
	}
}

func (c *MainController) Deluge() {
	f, ex := os.Open(".ipcache")
	defer f.Close()
	if ex == nil {
		b := make([]byte, 15)
		if n, ex := f.Read(b); ex == nil {
			c.Redirect(fmt.Sprintf("http://%s:1980", strings.TrimSpace(string(b[:n]))), 302)
		}
	}
}

func (c *MainController) Mldonkey() {
	f, ex := os.Open(".ipcache")
	defer f.Close()
	if ex == nil {
		b := make([]byte, 15)
		if n, ex := f.Read(b); ex == nil {
			c.Redirect(fmt.Sprintf("http://%s:1982", strings.TrimSpace(string(b[:n]))), 302)
		}
	}
}

type VpsController struct {
	beego.Controller
}

func (c *VpsController) VpsInfo() {
	c.TplName = "vpsinfo.tpl"
	// use beego httplib
	req := httplib.Get(fmt.Sprintf(urlVpsStatus, beego.AppConfig.String("veid"), mxgo.DecodeString(beego.AppConfig.String("apikey"))))
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(5*time.Second, 5*time.Second)
	rep, ex := req.Response()
	// use golang net
	// tr := &http.Transport{
	// 	TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	// 	DisableCompression: true,
	// }
	// client := &http.Client{Transport: tr}
	// rep, ex := client.Get(fmt.Sprintf(urlVpsStatus, beego.AppConfig.String("veid"), mxgo.DecodeString(beego.AppConfig.String("apikey"))))
	if ex == nil {
		if d, ex := ioutil.ReadAll(rep.Body); ex == nil {
			a := gjson.ParseBytes(d)
			c.Data["plan"] = a.Get("plan").String()
			c.Data["vmtype"] = a.Get("vm_type").String()
			c.Data["os"] = a.Get("os").String()
			c.Data["hostname"] = a.Get("hostname").String()
			c.Data["location"] = a.Get("node_location").String()
			c.Data["datacenter"] = a.Get("node_datacenter").String()
			c.Data["plan_monthly_data"] = a.Get("plan_monthly_data").Float() / 1024.0 / 1024.0 / 1024.0
			c.Data["data_counter"] = fmt.Sprintf("%.03f", a.Get("data_counter").Float()/1024.0/1024.0/1024.0)
			c.Data["ivp6"] = a.Get("location_ipv6_ready").String()
			c.Data["error"] = a.Get("error").String()
			c.Data["ipv4"] = a.Get("ip_addresses").Array()[0].String() + ":26937"
		}
	}
	c.Render()
}
