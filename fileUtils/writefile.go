package fileUtils

import (
	"os"
	"fmt"
	"log"
)

const TEMPPATH string = "/var/spool/sms/temp/"
const OUTPATH string = "/var/spool/sms/outgoing/"

func isError(err error, pre string) error {
	if err != nil {
		log.Printf("%v: %v", pre, err)
	}
	return err
}

func CreateFile(path string) {
	// detect if file exists
	var _, err = os.Stat(TEMPPATH + path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(TEMPPATH + path)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}

	fmt.Println("==> done creating file", TEMPPATH +  path)
}

func WriteFile(path string, text string) {
	CreateFile(path)
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(TEMPPATH + path, os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("==> done writing to file")
	MoveFile(path)
}

func MoveFile(path string) {
	// moves file
	err :=  os.Rename(TEMPPATH+path, OUTPATH+path)
	// error handling
	if err != nil {
		fmt.Println(err)
		return
	}
}