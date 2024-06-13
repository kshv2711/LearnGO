// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func responseSize(url string) {
// 	fmt.Println("Step1: ", url)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Step2: ", url)
// 	defer response.Body.Close()

// 	fmt.Println("Step3: ", url)
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Step4: ", len(body))
// }

// func main() {
// 	go responseSize("https://www.golangprograms.com")
// 	go responseSize("https://coderwall.com")
// 	go responseSize("https://stackoverflow.com")
// 	time.Sleep(10 * time.Second)
// }
package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)

var (
	garPercent = flag.Int("garC", 50, "Collect information about recent garbage collections")
)

func main() {
	debug.SetGCPercent(*garPercent)
	debug.PrintStack()	
	
	var garC debug.GCStats
	debug.ReadGCStats(&garC)
	fmt.Printf("\nLastGC:\t%s", garC.LastGC) // time of last collection
	fmt.Printf("\nNumGC:\t%d", garC.NumGC) // number of garbage collections
	fmt.Printf("\nPauseTotal:\t%s", garC.PauseTotal) // total pause for all collections
	fmt.Printf("\nPause:\t%s", garC.Pause) // pause history, most recent first	
}