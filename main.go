package main

import (
	"fmt"
	"gitlab.com/jinfagang/colorgo"
	flag "github.com/ogier/pflag"
	"encoding/json"
	"github.com/atotto/clipboard"
)

type ResultData struct {
	Path string `json:"path"`
	Url string `json:"url"`
}

type ResultJson struct {
	Code string `json:"code"`
	ResultBody ResultData `json:"data"`
}


func main() {
	fmt.Print(cg.BoldStart)
	cg.Foreground(cg.Blue, true)
	fmt.Println("gooooup - upload images to cloud for bedding.")
	fmt.Print(cg.BoldEnd)

	fmt.Println(`
  ____  ___   ___   ___   ___  _   _ ____
 / _  |/ _ \ / _ \ / _ \ / _ \| | | |  _ \
( (_| | |_| | |_| | |_| | |_| | |_| | |_| |
 \___ |\___/ \___/ \___/ \___/|____/|  __/
(_____|                             |_|    `)
	fmt.Print(cg.BoldStart)
	cg.Foreground(cg.Red, true)
	fmt.Println("Author JinTian. @wechat: jintianiloveu")
	fmt.Print(cg.BoldEnd)



	//cg.PrintlnYellow("gooooup will upload image to sms: " + cg.BoldStart + "https://sm.ms" + cg.BoldEnd)

	var uploadUrl = "https://sm.ms/api/upload"

	var markdown bool
	var showHistory bool


	flag.BoolVarP(&markdown, "markdown", "m",false, "result in markdown format.")
	flag.BoolVarP(&showHistory, "history", "h", false, "show upload history.")
	flag.Parse()


	posArgs := flag.Args()
	if len(posArgs) == 0 {

		// if not pos args, indicates user have no arg or using flag
		if showHistory {
			// show upload history
			ShowHistory()
		} else {
			cg.PrintlnRed("please provide a file path.")
		}
	} else {
		filePath := posArgs[0]
		if Exists(filePath) {
			// do the upload process
			// many I should check the format of image

			msg, r := UploadFile(uploadUrl, nil, "smfile", filePath, false)
			if r {
				fmt.Print(cg.BoldStart)
				cg.Foreground(cg.Green, true)
				fmt.Println("upload success!")
				fmt.Print(cg.BoldEnd)

				var result map[string]interface{}
				if err := json.Unmarshal([]byte(msg), &result); err != nil {
					cg.PrintlnRed("error in parse json data.")
				}


				var data = result["data"]
				var dataMap map[string]interface{}
				dataMap = data.(map[string]interface{})

				var url string
				url = dataMap["url"].(string)

				// save this url to history
				newRecord := Record{FilePath:filePath, Url:url}
				SaveToHistory(newRecord)

				if markdown {
					// write markdown url to pasteboard
					url = "![picture](" + url + ")"
					clipboard.WriteAll(url)

					fmt.Print(cg.BoldStart)
					cg.Foreground(cg.Yellow, true)
					fmt.Print("url: ")
					cg.Foreground(cg.Green, true)
					fmt.Println(url)
					fmt.Print(cg.BoldEnd)
				} else {
					// write url directly to pasteboard
					clipboard.WriteAll(url)
					fmt.Print("url: ")

					fmt.Print(cg.BoldStart)
					cg.Foreground(cg.Yellow, true)
					cg.Foreground(cg.Green, true)
					fmt.Println(url)
					fmt.Print(cg.BoldEnd)
				}

				fmt.Print(cg.BoldStart)
				cg.Foreground(cg.Yellow, true)
				fmt.Println("Done! just paste it!")
				fmt.Print(cg.BoldEnd)


			} else {
				cg.PrintlnRed(msg)
			}

		} else {
			cg.PrintlnRed(filePath + " are you sure this is a file?")
		}
	}
}
