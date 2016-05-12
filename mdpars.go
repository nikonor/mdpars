package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	today = time.Now().Format("2006-01-02")
}

var (
	default_dir = "/Users/nikonor/Яндекс.Диск/#iA/"
	default_tab = " "
	mdfiles     []string
	err         error
	today       string
)

type mdstring struct {
	Level int    // уровень записи, состоит из уровня предка и кол-ва отступов
	Text  string // собственно строка, без отступов и т.д.
	Tags  []tag  // список тэгов
	Show bool
}
type tag struct {
	Tag  string
	Date string
}

func (mds mdstring) String () string {
	s := ""
	for i := 0; i< (mds.Level-1); i++ {
		s += " "
	}
	s += mds.Text
	// for _,t := range mds.Tags {
	// 	s += "\t"+t.Tag
	// 	if t.Date != "" {
	// 		s+= "\n\t"+t.Date+""
	// 	}
	// }
	return s
}

func getFiles() ([]string, error) {
	var mdfs []string
	files, err := ioutil.ReadDir(default_dir)

	for _, file := range files {
		n := file.Name()
		if !file.IsDir() && strings.HasSuffix(n, ".md") {
			mdfs = append(mdfs, n)
		}
	}
	return mdfs, err
}

func CheckString(s mdstring, t string) bool {
	if ElInArray("done",s.Tags) {
		// шаг 1: если есть done, то сразу нет
		return false
	} else {
		// ищем @na
		if ElInArray("na",s.Tags) {
			return true
		}
		// ищем @start
		if StartInArray(today,s.Tags) {
			return true
		}
	}
	return false
}

func ElInArray(el string, tt []tag) bool {
	for _,t := range tt {
		if el == t.Tag {
			return true
		}
	}
	return false
}


func StartInArray(today string, tt []tag) bool {
	for _,t := range tt {
		if t.Tag=="start" && t.Date <= today {
			return true
		}
	}
	return false
}


func parseFile(fname string) (string, error) {
	var (
		ret      []string
		filedata []mdstring
		err      error
		file     *os.File
	)
	// ret = append(ret, ("# " + fname))
	file, err = os.Open(default_dir + fname)
	defer file.Close()
	filedata, err = readMDFile(file)
	for _,s := range filedata {
		ret = append(ret,s.String())
	}
	return strings.Join(ret,"\n"), err
}

func readMDFile(file *os.File) ([]mdstring, error) {
	var (
		data []mdstring
		ret []mdstring
		err error
		l   = 0
	)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		t := scanner.Text()
		h := HowManyTabs(t, "#")
		c := 0
		if h > 0 {
			l = h
		} else {
			c = HowManyTabs(t, default_tab) + 1
		}
		t = strings.TrimSpace(t)
		tt := FindTags(t)
		if len(t) > 0 {
			s := mdstring{Level: (l + c), Text: t, Tags: tt, Show: false}
			if CheckString(s,today) {
				s.Show = true
			}
			data = append(data, s)
		}
	}

	for i,d := range data {
		if d.Show {
			// прогон вверх, отмечает все, что выше имеют Level меньше текущего
			go_up(data, i, d.Level)
			// TK - прогон вниз. отмечаем все записи, Level которых больше.
			go_down(data, i, d.Level)
		}
	}


	for _,d := range data {
		if d.Show {
			ret = append(ret,d)
		}
	}

	return ret, err
}


func go_up(d []mdstring, n, level int) {
	// прогон вверх, отмечает все, что выше имеют Level меньше текущего
	cur_level := d[n].Level
	for i := n; i>=0; i-- {
		if d[i].Level < cur_level {
			d[i].Show = true
			cur_level=d[i].Level
		}
	}
}

func go_down(d []mdstring, n,level int) {
	cur_level := d[n].Level
	for i := n; i < len(d); i++ {
		if d[i].Level < cur_level || strings.HasPrefix(d[i].Text,"#") {
			break
		}
		if d[i].Level >= cur_level {
			cur_level = d[i].Level
			d[i].Show = true
		}
	}
}


func ParseTag(t string) tag {
	ret := tag{}

	t = t[1:]
	tt := strings.Split(t, "(")

	ret.Tag = strings.ToLower(tt[0])
	if len(tt) > 1 {
		ret.Date = tt[1][:len(tt[1])-1]
	}
	return ret
}

func FindTags(s string) []tag {
	var (
		ret []tag
	)

	for _, w := range strings.Fields(s) {
		if strings.HasPrefix(w, "@") {
			ret = append(ret, ParseTag(w))
		}
	}
	return ret
}

func HowManyTabs(s string, tab string) int {
	q := 0
	emMax := 10
	for strings.HasPrefix(s, tab) {
		q++
		s = s[len(tab):]
		if q >= emMax {
			break
		}
	}
	return q
}

func main() {
	fmt.Println("Start")

	// получаем список файлов для обработки
	mdfiles, err = getFiles()
	if err != nil {
		log.Fatal(err)
	}
	// обработка файлов
	// TK - пока один
	var o string
	for _,f := range mdfiles {
		// fmt.Printf("--- begin for %s ---\n",f)
		o, err = parseFile(f)
		fmt.Println(o)
		// fmt.Printf("--- end for %s ---\n",f)
	}

	fmt.Println("Finish")
}
