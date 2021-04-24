package runner

import (
    "github.com/gluttony/lib"
    "github.com/gluttony/log"
    "github.com/gluttony/utils"
    "github.com/panjf2000/ants"
    "strings"
    "sync"
    "time"
)

func Run(i interface{},op lib.Options)  {
    job := i.(lib.Job)

    utils.Prepareurl(job.Url,&job.Poc,op)
    utils.Preparerequests(&job.Poc)

    if job.Poc.Single {
        Singlerun(job)
    } else {
        Multiplerun(job, op.Pocthread)
    }

}

func Multiplerun(job lib.Job ,thread int)  {
    var wg sync.WaitGroup
    p, _ := ants.NewPoolWithFunc(thread, func(i interface{}) {
        request := i.(lib.Request)
        timeUnix :=time.Now().UnixNano()
        resp ,err := Send(request)
        if err !=nil {
            log.Log.Error("[requests] %s , cve:%s ,error:%s",request.Url, job.Poc.Cve.Id, err)
            wg.Done()
            return
        }
        res := Parseresponse(resp,timeUnix)
        log.Log.Info("[requests] %s , cve:%s ,info:send ,method:%s ,starttime:%s ,endtime:%s",request.Url ,job.Poc.Cve.Id,request.Method,res.Starttime,res.Endtime)

        result ,err := Detections(request ,res)
        if err !=nil {
            log.Log.Error("[detection] %s , cve:%s ,error:%s",request.Url, job.Poc.Cve.Id, err)
            wg.Done()
            return
        }
        if result == 0 {
            log.Log.Debug("[detection] %s , cve:%s ,debug:mismatch",request.Url ,job.Poc.Cve.Id)
            wg.Done()
            return
        }
        log.Log.Info("[detection] %s , cve:%s ,info:match",request.Url ,job.Poc.Cve.Id)
        wg.Done()
    }, ants.WithPreAlloc(true))
    defer p.Release()

    for _, request := range job.Poc.Requests {
        for _, pa := range request.Path{
            wg.Add(1)
            if !strings.HasPrefix(pa, "/") {
                pa = "/" + pa
            }
            request.Url = request.Url + pa
            _ = p.Invoke(request)
        }
    }

    wg.Wait()
}

func Singlerun(job lib.Job)  {
    step := 0
    num := len(job.Poc.Requests)
    for _, request := range job.Poc.Requests {
        step = step + 1
        timeUnix :=time.Now().UnixNano()
        resp ,err := Send(request)
        if err !=nil {
            log.Log.Error("[requests] %s , step:%d , cve:%s ,error:response",request.Url, step ,job.Poc.Cve.Id)
            break
        }

        res := Parseresponse(resp, timeUnix)
        result ,err:= Detections(request, res)
        if err !=nil {
            log.Log.Error("[detection] %s , cve:%s ,error:%s",request.Url, job.Poc.Cve.Id, err)
            break
        }
        if result == 0 {
            log.Log.Debug("[requests] %s , step:%d , cve:%s ,debug:condition mismatch",request.Url, step ,job.Poc.Cve.Id)
            break
        }

        if step == num {
            log.Log.Info("[detection] %s , cve:%s ,info:match",request.Url ,job.Poc.Cve.Id)
        }
    }
}
