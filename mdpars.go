package main
import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func init() {
    readConfig();
}

var (
    default_dir = "/home/nikonor/Yandex.Disk/#iA/"
    mdfiles []string
)


func readConfig() {
    fmt.Println("call readConfig");
}

func main() {
    fmt.Println("Start")
    files, err := ioutil.ReadDir(default_dir)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        n := file.Name()
        if !file.IsDir() && strings.HasSuffix(n, ".md") {
            mdfiles = append(mdfiles,n)
        }
    }    

    fmt.Println(mdfiles);

    fmt.Println("Finish")
}