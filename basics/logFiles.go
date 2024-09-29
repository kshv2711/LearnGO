// package main

// import (
// 	"fmt"
// 	"log"
// 	"log/syslog"
// 	"os"
// )
// func main(){
// 	programName:= filepath.base(os.Args[0])
// 	syslog, err:=syslog.New(syslog.LOG_INFO|syslog.LoG_LOCAL7,programName)
// 	if err!=nil{
// 		log.Fatal(err)
		
// 	}else{
// 		log.SetOutput(sysLog)
// 	}
// 	log.Println("LOG_INFO + LOG_LOCAL7: Logging in GO!!")
// 	syslog,err=syslog.New(syslog.LOG_MAIL,"some Program!!")
// 	if err!=nil{
// 		log.Fatal(err)
// 	}else{
// 		log.SetOutput(syslog)
// 	}
// 	log.Println("LOG_MAIL: Logging in Go!")
// 	fmt.Println("will you see this ?")
// }
package main

import (
    "fmt"
    "log"
    "log/syslog"
    "os"
    "path/filepath"
)

func main() {
    programName := filepath.Base(os.Args[0])

    // Create a syslog writer for LOG_INFO and LOG_LOCAL7
    sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
    if err != nil {
        log.Fatal(err)
    } else {
        log.SetOutput(sysLog)
    }

    log.Println("LOG_INFO + LOG_LOCAL7: Logging in GO!!")

    // Create a syslog writer for LOG_MAIL
    mailLog, err := syslog.New(syslog.LOG_MAIL, "some Program!!")
    if err != nil {
        log.Fatal(err)
    } else {
        log.SetOutput(mailLog)
    }

    log.Println("LOG_MAIL: Logging in Go!")
    fmt.Println("will you see this?")
}