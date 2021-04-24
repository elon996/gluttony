package utils

import (
	"bytes"
	"github.com/gluttony/lib"
	"strings"
	"text/template"
)

func Preparerequests(poc *lib.Poc)  {
	Prepareset(poc)
}

func Prepareset(poc *lib.Poc)  {
	for n,request :=range poc.Requests {
		var req lib.Request
		setvalue:=make(map[string]string)

		for _,v :=range request.Set {
			for key, value := range v {
				setvalue[key] = value
			}
		}

		var p []string
		for _, path := range request.Path {
			p =  append(p, Setvalue(path, setvalue))
		}
		request.Path = p
		request.Body = Setvalue(request.Body, setvalue)
		for num,header :=range request.Headers {
			for key,value :=range header {
				request.Headers[num][key] = Setvalue(value, setvalue)
			}
		}

		req = request
		poc.Requests[n] = req
	}
}

func Setvalue(s string ,data map[string]string) string {
	temp, err := template.New("").Parse(s)
	t := template.Must(temp, err)

	buf := &bytes.Buffer{}
	err = t.Execute(buf, data)
	if err != nil {
		return s
	}
	return buf.String()
}

func Prepareurl(url string, poc *lib.Poc, op lib.Options) {
	for n, request := range poc.Requests {
		var req lib.Request


		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		if strings.HasSuffix(url, "/") {
			url = url[:len(url)-1]
		}
		request.Url = url

		if len(request.Proxy)==0 {
			request.Proxy = append(request.Proxy, op.Proxy)
		}
		if request.Timeout==0 {
			request.Timeout = op.Timeout
		}

		req = request
		poc.Requests[n] = req

	}
}

