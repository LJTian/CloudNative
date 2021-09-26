package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func Initpporf() {
	http := http.NewServeMux()
	http.HandleFunc("/debug/pprof/", pprof.Index)
	http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	http.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

// 没找到类似于gin的USE,需要自己实现,就先用URL解析完成吧
func monitorFunc(w http.ResponseWriter, r *http.Request) {

	io.WriteString(os.Stdout, fmt.Sprintf("Url is [%s] \t addr is [%s]\n", r.URL, r.RemoteAddr))
	// 目前只支持一层
	URL := r.URL.String()
	URLS := strings.Split(URL, "/")

	switch URLS[1] {
	case "headleCopy":
		HeadleCopy(w, r)
	case "GetEnv":
		GetEnv(w, r)
	case "healthz":
		Healthz(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func HeadleCopy(w http.ResponseWriter, r *http.Request) {

	w.Header()
	w.WriteHeader(http.StatusOK)
}

func GetEnv(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.Header().Set("GOPATH", os.Getenv("GOPATH"))
	w.WriteHeader(http.StatusOK)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Index(w http.ResponseWriter, r *http.Request) {
	monitorFunc(w, r)
}

func init() {
	//	注册 pporf
	Initpporf()
}

func main() {
	// 监控函数
	http.HandleFunc("/", Index)
	// 绑定监听端口
	//http.HandleFunc("/headleCopy", HeadleCopy)
	//http.HandleFunc("/GetEnv", GetEnv)
	//http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8001", nil)
}
