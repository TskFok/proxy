package controller

import (
	"ProxyHttp/app/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

func Proxy(ctx *gin.Context) {
	direct, exist := ctx.GetQuery("url")

	if !exist {
		return
	}
	proxyUrl, err := url.Parse(global.ProxyUrl)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	req, err := http.NewRequest("GET", direct, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client := http.Client{Transport: transport}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if response.StatusCode == http.StatusOK {
		defer response.Body.Close()

		by, err := io.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		ctx.Data(200, "xml", by)
	}
}
