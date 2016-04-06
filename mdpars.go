package main
import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "os"
    "bufio"
)

func init() {
    readConfig();
}

var (
    default_dir = "/home/nikonor/Yandex.Disk/#iA/"
    mdfiles []string
    err error
)

    type mdstring struct {
        Level int 
        Text string
        C int
    }


func readConfig() {
    fmt.Println("call readConfig");
}

func getFiles () ([]string, error) {
    var mdfs []string
    files, err := ioutil.ReadDir(default_dir)

    for _, file := range files {
        n := file.Name()
        if !file.IsDir() && strings.HasSuffix(n, ".md") {
            mdfs = append(mdfs,n)
        }
    }    
    return mdfs,err
}

func parseFile (fname string) (string,error) {
    var (
        ret string
        filedata []mdstring
        err error
        file *os.File
    )
    ret = ret + "# "+fname
    file, err = os.Open(default_dir+fname)
    defer file.Close()
    filedata,err = readMDFile(file)
    fmt.Println(filedata);

    return ret,err
}

func readMDFile(file *os.File) ([]mdstring,error) {
    var (
        ret []mdstring
        err error
        level = 1
    )

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)   
    for scanner.Scan() {
        // fmt.Println(scanner.Text());
        c := 22
        ret = append(ret,mdstring{Level: level, Text: scanner.Text(), C: c})
    }    

    return ret,err
}

func main() {
    fmt.Println("Start")

    // получаем список файлов для обработки
    mdfiles,err = getFiles()
    if err != nil {
        log.Fatal(err)
    }
    // обработка файлов
    // TK - пока один
    var o string
    o,err = parseFile("ToDo.md")
    fmt.Println("--- begin for ToDo.md ---\n",o,"\n--- end for ToDo.md ---");

    fmt.Println("Finish")
}