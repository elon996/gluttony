package cmd

import (
    "fmt"
    "github.com/gluttony/log"
    "github.com/spf13/cobra"
    "github.com/gluttony/color"
    "github.com/gluttony/lib"
    "os"

)

var rootCmd = &cobra.Command{
    Use:   "gluttony",
}

var options = lib.Options{}

func init() {
    cobra.OnInitialize(initConfig)
    color.Redf("Gluttony %v by %v\n", version, author)

    rootCmd.PersistentFlags().StringVarP(&options.Url, "url", "u", "", "url")
    rootCmd.PersistentFlags().StringVarP(&options.Poc, "poc", "p", "", "poc yaml")
    //rootCmd.PersistentFlags().StringVarP(&options.Config, "config","c","", "config file")
    rootCmd.PersistentFlags().BoolVarP(&options.Debug, "debug", "d", false, "debug")
    rootCmd.PersistentFlags().StringVarP(&options.Proxy, "proxy", "", "", "proxy")
    rootCmd.PersistentFlags().IntVarP(&options.Thread, "thread", "t", 15, "thread")
    rootCmd.PersistentFlags().IntVarP(&options.Pocthread, "pocthread", "", 15, "thread")
    rootCmd.PersistentFlags().IntVarP(&options.Timeout, "timeout", "", 60, "Timeout")
    //rootCmd.PersistentFlags().StringVarP(&options.Output, "output", "", "", "output folder")
    rootCmd.SetHelpFunc(rootHelp)

}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}



func initConfig()  {
    //Debug init
    if options.Debug {
        log.Log.SetLevel(log.DebugLevel)
    } else {
        log.Log.SetLevel(log.InfoLevel)
    }
}



































