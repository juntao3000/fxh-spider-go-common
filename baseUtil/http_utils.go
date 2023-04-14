package baseUtil

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/juntao3000/fxh-spider-go-common/baseCommon"
	"golang.org/x/net/proxy"
	"io"
	"net/http"
	"net/url"
)

func GetHttpClientByProxy() (*http.Client, error) {
	if baseCommon.BaseConfig.Socks5ProxyPort > 0 && len(baseCommon.BaseConfig.Socks5ProxyHost) > 0 {
		addr := fmt.Sprintf("%s:%d", baseCommon.BaseConfig.Socks5ProxyHost, baseCommon.BaseConfig.Socks5ProxyPort)

		dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
		if err != nil {
			return nil, fmt.Errorf("can not connect to socks5 proxy:%v", err)
		}

		return &http.Client{Transport: &http.Transport{
			Dial:            dialer.Dial,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}, nil
	}

	if baseCommon.BaseConfig.HttpProxyPort > 0 && len(baseCommon.BaseConfig.HttpProxyHost) > 0 {
		addr := fmt.Sprintf("http://%s:%d", baseCommon.BaseConfig.HttpProxyHost, baseCommon.BaseConfig.HttpProxyPort)

		proxyUrl, err := url.Parse(addr)
		if err != nil {
			return nil, fmt.Errorf("parse http proxy url failed:%v", err)
		}

		return &http.Client{Transport: &http.Transport{
			Proxy:           http.ProxyURL(proxyUrl),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}, nil
	}

	return &http.Client{}, nil
}

func HttpGetByClient(httpClient *http.Client, url string, result any) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Close = true
	req.Header.Set("Connection", "Close")
	req.Header.Set("Accept-Encoding", "identity")
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http request failed,status code:%d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read http response body failed:%v", err)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("unmarshal http response body json failed:%v", err)
	}

	return nil
}

func HttpGetByProxy(url string, result any) error {
	httpClient, err := GetHttpClientByProxy()
	if err != nil {
		return err
	}

	return HttpGetByClient(httpClient, url, result)
}

func HttpGetBySocks5Proxy(url string, result any, proxyServer string) error {
	dialer, err := proxy.SOCKS5("tcp", proxyServer, nil, proxy.Direct)
	if err != nil {
		return fmt.Errorf("can not connect to socks5 proxy:%v", err)
	}

	httpClient := &http.Client{Transport: &http.Transport{
		Dial:            dialer.Dial,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	return HttpGetByClient(httpClient, url, result)
}

func HttpGetBySsrProxy(url string, result any) error {
	addr := fmt.Sprintf("%s:%d", baseCommon.BaseConfig.SsrSocks5ProxyHost, baseCommon.BaseConfig.SsrSocks5ProxyPort)

	dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	if err != nil {
		return fmt.Errorf("can not connect to ssr socks5 proxy:%v", err)
	}

	httpClient := &http.Client{Transport: &http.Transport{
		Dial:            dialer.Dial,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	return HttpGetByClient(httpClient, url, result)
}

func HttpGet(url string, result any) error {
	return HttpGetByClient(&http.Client{}, url, result)
}
