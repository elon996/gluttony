package utils

import (
	"bufio"
	"github.com/gluttony/lib"
	"github.com/gluttony/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func IsFile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil{
		return false
	}
	return true
}

func Readfile(file string) ([]string , error) {
	var urls []string
	urls = append(urls, file)
	f, _ := os.Open(file)
	s := bufio.NewScanner(f)
	for s.Scan() {
		urls = append(urls, s.Text())
	}
	err := s.Err()
	if err != nil {
		log.Log.Error("%v",err)
		return urls,err
	}
	if err = f.Close(); err != nil {
		log.Log.Error("%v",err)
		return urls,err
	}
	return urls,nil
}


func IsFolder(foldername string) bool {
	s, err := os.Stat(foldername)
	if err != nil {
		return false
	}
	return s.IsDir()
}


func Readfolder(folder string) []string{
	var s []string
	files, _ := ioutil.ReadDir(folder)
	for _,file := range files{
		if file.IsDir(){
			d:=Readfolder(folder + "/" + file.Name())
			s= append(s ,d...)
		}else{
			s= append(s ,folder + "/" + file.Name())
		}
	}
	return s
}

func Parseyaml(file string) lib.Poc {
	var y = lib.Poc{}
	yam, err := ioutil.ReadFile(file)
	if err != nil {
		log.Log.Error("Error parsing Signature:  %v - %v", err, file)
	}
	err = yaml.Unmarshal(yam, &y)
	if err != nil {
		log.Log.Error("Error parsing Signature:  %v - %v", err, file)
	}
	return y
}

func Prepareyaml(p lib.Poc , op  lib.Options) lib.Poc {
	var reqs []lib.Request
	for _,r := range p.Requests {
		r.Proxy = append(r.Proxy ,op.Proxy)
		r.Timeout = op.Timeout
		reqs = append(reqs, r)
	}
	p.Requests = reqs
	return p
}

func ParseHeaders(rawheaders []map[string]string) map[string]string {
	headers := make(map[string]string)
	for _, header := range rawheaders {
		for key, value := range header {
			headers[key] = value
		}
	}
	return headers
}


/*
func Splitspace(s string) string {
	a := strings.Split(s, " ")[0]
	return a
}
*/
















