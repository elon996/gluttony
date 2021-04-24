package cmd

import (
    "github.com/gluttony/color"
    "github.com/spf13/cobra"

)

func rootHelp(cmd *cobra.Command, _ []string) {
    color.Yellow(`usege: 
gluttony scan -u 192.168.1.1 -p cve-2018-10225.yaml
gluttony scan -u 192.168.1.1 -p rce

flags:
-u  --url       a url,can be a flie ,or custom plug-in file
-p  --poc       you need to scan the file
-t  --thread    thread
--pocthread     poc thread
--debug         more info
--proxy         http proxy
--timeout       timeout
`)

}




































