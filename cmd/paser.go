package cmd

import (
    "github.com/gluttony/lib"
    "github.com/gluttony/log"
    "github.com/gluttony/utils"
)

func parseurl(s string) []string {
    var urls []string
    if utils.IsFile(s) {
        urls , err := utils.Readfile(s)
        if err == nil{
            return urls
        }
    }
    urls = append(urls, s)
    log.Log.Info("Url Loaded: %v", len(urls))
    return urls
}

func parsepoc(op  lib.Options) []lib.Poc {
    var pocs []string
    var Pocs []lib.Poc
    if utils.IsFolder(op.Poc) {
        pocs = utils.Readfolder(op.Poc)
    } else {
        pocs = append(pocs, op.Poc)
    }
    log.Log.Info("Poc Loaded: %v", len(pocs))
    for _, v := range pocs {
        yam := utils.Parseyaml(v)
        yam = utils.Prepareyaml(yam, op)
        Pocs = append(Pocs, yam)
    }
    return Pocs
}


























