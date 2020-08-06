package tail

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"reflect"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Tail() string {
	cmd := exec.Command("tail", "-1", "gin.log")
	// create a pipe for the output of the script
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return "err"
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("\t > %s\n", scanner.Text())
			fmt.Println("var1 = ", reflect.TypeOf(scanner.Text()))
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

	return scanner.Text()
}
