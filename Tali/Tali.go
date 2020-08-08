package Tali

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	models "wordplay/Models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Tali1(filename string) string {
	cmd := exec.Command("tail", "-1", filename)
	// create a pipe for the output of the script
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return "err"
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			in := []byte(scanner.Text())
			var raw map[string]interface{}
			if err := json.Unmarshal(in, &raw); err != nil {
				panic(err)
			}
			raw["count"] = 1
			x := raw["level"]
			out, err := json.Marshal(raw)
			if err != nil {
				panic(err)
			}

			if x == "info" {
				println(string(out))
				Logger := models.Logger{Title: "Something is wrong new", Error: scanner.Text(), Filename: filename}
				models.DB.Create(&Logger)
			}
			// println(string(out))
		}
	}()
	// return scanner.Text()
	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		return "err"
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		return "jj"
	}
	Logger := models.Logger{Title: "Something is wrong", Error: scanner.Text()}
	models.DB.Create(&Logger)

	// birdJson := scanner.Text()
	// var result map[string]interface{}
	// json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	// birds := scanner.Text()
	return scanner.Text()
	// return string(birds
}
