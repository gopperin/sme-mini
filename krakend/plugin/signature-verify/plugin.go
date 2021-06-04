package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HandlerRegisterer is the symbol the plugin loader will try to load. It must implement the Registerer interface
var HandlerRegisterer = registerer("signature-verify")

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(ctx context.Context, extra map[string]interface{}, handler http.Handler) (http.Handler, error) {
	// check the passed configuration and initialize the plugin
	name, ok := extra["name"].([]interface{})
	if !ok {
		return nil, errors.New("wrong config")
	}
	if name[1] != string(r) {
		return nil, fmt.Errorf("unknown register %s", name)
	}

	fmt.Println("register", extra["app"])

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("signature-verify called")

		fmt.Println("handler", extra["app"])

		// 把request的内容读取出来
		var _bodyBytes []byte
		if nil == req.Body {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(406)
			w.Write([]byte(`{"code":406,"msg":"验签错误"}`))
			return
		}

		_bodyBytes, _ = ioutil.ReadAll(req.Body)

		_contentType := req.Header.Get("Content-Type")
		if "" == _contentType {
			_contentType = "application/json"
		}
		_props, err := GenPropsByBody(_contentType, _bodyBytes)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(406)
			w.Write([]byte(`{"code":406,"msg":"验签错误"}`))
			return
		}

		_appID := req.Header.Get("X-APPID")
		_salt := "f90f4ec04b10"
		if "" != _appID {
			_salt = extra["app"].(map[string]interface{})[_appID].(string)
		}
		// 处理各个appid对应的salt

		fmt.Println("salt:", _salt[:4])

		_sign := WechatSign(_props, _salt, "signature")
		fmt.Println("====== api signed : ", _sign, _props["signature"])
		if _props["signature"] != _sign {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(406)
			w.Write([]byte(`{"code":406,"msg":"验签错误"}`))
			return
		}
		// 把刚刚读出来的再写进去
		req.Body = ioutil.NopCloser(bytes.NewBuffer(_bodyBytes))

		handler.ServeHTTP(w, req)

	}), nil
}

func init() {
	fmt.Println("signature-verify handler loaded!!!")
}

func main() {}
