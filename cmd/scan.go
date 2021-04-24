package cmd

import (
    "github.com/gluttony/lib"
    "github.com/gluttony/log"
    "github.com/gluttony/runner"
    "github.com/panjf2000/ants"
    "sync"

    "github.com/spf13/cobra"
    "os"
)

func init() {
    scanCmd.SetHelpFunc(rootHelp)
    rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
    Use:   "scan",
    RunE: runScan,
}

var urls []string
var pocs []string
var paseroption = lib.Parseoptions{}

func runScan(cmd *cobra.Command, _ []string) error {
    checkEmpty()
    paseroption.Urls = parseurl(options.Url)
    paseroption.Pocs = parsepoc(options)
    scan()
    return nil
}

func scan()  {
    var wg sync.WaitGroup
    p, _ := ants.NewPoolWithFunc(options.Thread, func(i interface{}) {
        runner.Run(i, options)
        wg.Done()
    }, ants.WithPreAlloc(true))
    defer p.Release()

    for _,url := range paseroption.Urls {
        for _,poc := range  paseroption.Pocs{
            wg.Add(1)
            job := lib.Job{Url: url, Poc: poc}
            _ = p.Invoke(job)
        }
    }

    wg.Wait()
}

func check()  {
    if len(urls) ==0 || len(pocs)==0 {
        log.Log.Warn("no url or poc")
        log.Log.Warn("to see usege")
        os.Exit(1)
    }
}

func checkEmpty()  {
    if options.Url =="" || options.Poc=="" {
        log.Log.Warn("no url or poc")
        log.Log.Warn("to see usege")
        os.Exit(1)
    }
}

