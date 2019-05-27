package lib

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/tidwall/gjson"
	"github.com/xyzj/gopsu"
)

const (
	bwhStatusURL = "https://api.64clouds.com/v1/getServiceInfo?veid=%s&api_key=%s"
	bwhAPIKey    = "42Xjj9rD6ZN45lubxhYV0mFah3ZvX7kqRvKVLjrQAINUKOzzQMR7Tlm/M1XrXmtl+ZT00lAmZsUQ48tb3JYbHmw"
	bwhVeid      = "979913"
)

func RemoteIP(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.String(200, strings.Split(c.Request.RemoteAddr, ":")[0])
	case "POST":
		f, ex := os.OpenFile(".ipcache", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		defer f.Close()
		if ex == nil {
			f.WriteString(strings.Split(c.Request.RemoteAddr, ":")[0])
		}
		c.String(200, "success")
	}
}

func IPCache(c *gin.Context) {
	b, _ := ioutil.ReadFile(".ipcache")
	c.String(200, string(b))
}

func Ssh(c *gin.Context) {
	b, _ := ioutil.ReadFile(".ipcache")
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://%s:1983/ssh/host/127.0.0.1", strings.TrimSpace(string(b))))
}

func Kod(c *gin.Context) {
	b, _ := ioutil.ReadFile(".ipcache")
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://%s:1981", strings.TrimSpace(string(b))))
}

func Deluge(c *gin.Context) {
	b, _ := ioutil.ReadFile(".ipcache")
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://%s:1980", strings.TrimSpace(string(b))))
}

func Mldonkey(c *gin.Context) {
	b, _ := ioutil.ReadFile(".ipcache")
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://%s:1982", strings.TrimSpace(string(b))))
}

func VpsInfo(c *gin.Context) {
	// use beego httplib
	// req := httplib.Get(fmt.Sprintf(urlVpsStatus, beego.AppConfig.String("veid"), gopsu.DecodeString(beego.AppConfig.String("apikey"))))
	// req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// req.SetTimeout(5*time.Second, 5*time.Second)
	// rep, ex := req.Response()

	// use golang net
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, ex := client.Get(fmt.Sprintf(bwhStatusURL, bwhVeid, gopsu.DecodeString(bwhAPIKey)))
	if ex == nil {
		if d, ex := ioutil.ReadAll(resp.Body); ex == nil {
			a := gjson.ParseBytes(d)
			c.Set("plan", a.Get("plan").String())
			c.Set("vmtype", a.Get("vm_type").String())
			c.Set("os", a.Get("os").String())
			c.Set("hostname", a.Get("hostname").String())
			c.Set("location", a.Get("node_location").String())
			c.Set("datacenter", a.Get("node_datacenter").String())
			c.Set("plan_monthly_data", a.Get("plan_monthly_data").Float()/1024.0/1024.0/1024.0)
			c.Set("data_counter", fmt.Sprintf("%.03f", a.Get("data_counter").Float()/1024.0/1024.0/1024.0))
			c.Set("ivp6", a.Get("location_ipv6_ready").String())
			c.Set("error", a.Get("error").String())
			c.Set("ipv4", a.Get("ip_addresses").Array()[0].String()+":26937")
		} else {
			println(ex.Error())
		}
	}
	c.HTML(200, "vpsinfo", c.Keys)
}

func SohoCache(c *gin.Context) {
	c.String(200, "180.167.245.233")
}

func SohoKod(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://180.167.245.233:20080")
}
