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
    today = "2016-04-05"
}

var (
    default_dir = "/home/nikonor/Yandex.Disk/#iA/"
    default_tab = " "
    mdfiles []string
    err error
    today string
)

type mdstring struct {
    Level int  // уровень записи, состоит из уровня предка и кол-ва отступов
    Text string // собственно строка, без отступов и т.д.
    Tags []tag // список тэгов
}
type tag struct {
    Tag string
    Date string
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
    for _,s := range filedata {
        fmt.Println(s.Text,", Level=",s.Level,s.Tags);
    }

    return ret,err
}

func readMDFile(file *os.File) ([]mdstring,error) {
    var (
        ret []mdstring
        err error
        l = 0 
    )

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)   
    for scanner.Scan() {
        t := scanner.Text()
        c := HowManyTabs(t,default_tab)
        h := HowManyTabs(t,"#")
        if h > 0 {
            l = h
        }
        t = strings.TrimSpace(t)
        tt := FindTags(t)
        if len(t) > 0 {
            ret = append(ret,mdstring{Level: (l + c), Text:t, Tags: tt })
        }
    }    

    return ret,err
}

func ParseTag (t string) tag {
    ret := tag{}

    t = t[1:]
    tt := strings.Split(t,"(")

    ret.Tag = tt[0]
    if len(tt)>1 {
        ret.Date = tt[1][:len(tt[1])-1]
    } 
    return ret
}

func FindTags (s string) []tag {
    var (
        ret []tag
    )

    for _,w := range(strings.Fields(s)) {
        if strings.HasPrefix(w,"@") {
            ret = append(ret,ParseTag(w))
        }
    }
    return ret
}

func HowManyTabs(s string,tab string) int{
    q := 0
    emMax := 10
    for strings.HasPrefix(s,tab) {
        q++
        s = s[len(tab):]
        if q >= emMax {
            break;
        }
    }
    return q
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