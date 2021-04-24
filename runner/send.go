package runner

import (
	"errors"
	"github.com/gluttony/lib"
	"github.com/gluttony/utils"
	"github.com/levigross/grequests"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Send(req lib.Request) (res grequests.Response, err error) {
	requestsoption := grequests.RequestOptions{}

	//method
	method := strings.ToLower(strings.TrimSpace(req.Method))

	//header
	headers := utils.ParseHeaders(req.Headers)
	requestsoption.Headers = headers
	if method == "post" {
		if len(requestsoption.Headers["Content-Type"]) ==0 {
			requestsoption.Headers["Content-Type"] = "application/x-www-form-urlencoded"
		}
	}

	//auth
	if len(req.Authuser)>0 {
		requestsoption.Auth = []string{req.Authuser , req.Authpassword}
	}

	requestsoption.InsecureSkipVerify = false

	//timeout
	t := time.Duration(req.Timeout) * time.Second
	requestsoption.RequestTimeout = t

	//file
	if len(req.File)>0 {
		fd, err := grequests.FileUploadFromDisk(req.File)
		if err != nil {
			return res, err
		}
		requestsoption.Files = fd
	}

	//proxy
	requestsoption.Proxies = map[string]*url.URL{}
	if len(req.Proxy) >0 {
		for _,proxy := range req.Proxy {
			if len(proxy)==0 {
				break
			}

			if !strings.HasPrefix(proxy, "http://") && !strings.HasPrefix(proxy, "https://") {
				return res, errors.New("proxy format error")
			}

			proxyURL, err := url.Parse(proxy) // Proxy URL
			if err != nil {
				return res, errors.New("proxy error")
			}
			requestsoption.Proxies[proxyURL.Scheme] = proxyURL
		}
	}

	//data
	if req.Body != "" {
		requestsoption.RequestBody = strings.NewReader(req.Body)
	}
	

	//redirect
	requestsoption.RedirectLimit = 20
	if req.Redirect !="" {
		redirect ,err:= strconv.Atoi("req.Redirect")
		if err != nil {
			return res, err
		}
		requestsoption.RedirectLimit = redirect
	}

	//send
	var resp *grequests.Response
	switch method {
	case "get":
		resp, err = grequests.Get(req.Url,
			&requestsoption,
		)
		break
	case "post":
		resp, err = grequests.Post(req.Url,
			&requestsoption,
		)
		break
	case "head":
		resp, err = grequests.Head(req.Url,
			&requestsoption,
		)
		break
	case "options":
		resp, err = grequests.Options(req.Url,
			&requestsoption,
		)
		break
	case "patch":
		resp, err = grequests.Patch(req.Url,
			&requestsoption,
		)
		break
	case "put":
		resp, err = grequests.Put(req.Url,
			&requestsoption,
		)
		break
	case "delete":
		resp, err = grequests.Delete(req.Url,
			&requestsoption,
		)
		break
	}



	if err != nil {
		return res, err
	}

	res = *resp
	return res, nil
}


