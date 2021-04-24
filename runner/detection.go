package runner

import (
	"errors"
	"github.com/dop251/goja"
	"github.com/gluttony/lib"
	"strconv"
)

func Detections(req lib.Request, resp lib.Response) (result int, err error) {
	vm := goja.New()
	res := vm.NewObject()
	res.Set("body", resp.Body)
	res.Set("code", resp.StatusCode)
	res.Set("length", resp.Length)
	res.Set("header", resp.Headers)
	/*
	res.Set("search", func(c goja.FunctionCall) goja.Value {
		return vm.ToValue(false)
	})*/

	vm.Set("res", res)

	for _, detection := range req.Detections {
		rst, err := vm.RunString(detection)
		if err != nil {
			return 0 ,errors.New("detection RunString error")
		}
		if rst.String()=="true" {
			return 1 ,nil
		}
		if rst.String()=="false" {
			return 0 ,nil
		}
		r , err := strconv.Atoi(rst.String())
		if err != nil {
			return 0 ,errors.New("detection strconv error")
		}
		if r >0 {
			return r ,nil
		}
	}

	return 0 ,nil
}
