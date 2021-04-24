package runner

import (
	"github.com/gluttony/lib"
	"github.com/levigross/grequests"
	"net/http"
	"time"
)

func Parseresponse(resp grequests.Response, starttime int64) (res lib.Response) {
	timeUnix :=time.Now().UnixNano()

	res.Endtime = time.Unix(0,timeUnix).Format("15:04:05.0000")
	res.Starttime = time.Unix(0,starttime).Format("15:04:05.0000")
	res.Body = resp.String()
	res.Headers = Prepareheaders(resp.Header)
	res.Status = resp.RawResponse.Status
	res.Length = resp.RawResponse.ContentLength
	res.StatusCode = resp.StatusCode
	return res
}

func Prepareheaders(h http.Header) string {
	s :=""
	for k, v := range h {
		s= s+ k +":"
		for _, str := range v {
			s = s+ str+" "
		}
	}
	return s
}
