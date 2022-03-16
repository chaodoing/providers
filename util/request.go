package util

import (
	`io`
	`io/ioutil`
	`net/http`
	`strings`
)

// request 带数据的请求
func request(method, url, query string, headers ...map[string]string) (body string, err error) {
	var data io.Reader = strings.NewReader(query)
	request, err := http.NewRequest(strings.ToUpper(method), url, data)
	if err != nil {
		return "", err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			request.Header.Add(key, value)
		}
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	var Body []byte
	Body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return string(Body), nil
}

// GET 发送Get请求
func GET(query string, headers ...map[string]string) (body string, err error) {
	var request *http.Request
	request, err = http.NewRequest("GET", query, nil)
	if err != nil {
		return "", err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			request.Header.Add(key, value)
		}
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	var Body []byte
	Body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	return string(Body), nil
}

// POST 发送POST请求
func POST(query string, data string, headers ...map[string]string) (body string, err error) {
	return request("post", query, data, headers...)
}

// PUT 发送PUT请求
func PUT(query string, data string, headers ...map[string]string) (body string, err error) {
	return request("put", query, data, headers...)
}

// DELETE 发送DELETE请求
func DELETE(query string, data string, headers ...map[string]string) (body string, err error) {
	return request("delete", query, data, headers...)
}

