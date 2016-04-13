package main

import (
    "testing"
    "fmt"
    // "strings"
)

func TestFindTags(t *testing.T) {
    cases := []struct{
        in string
        wait []tag
    }{
        {"строка без тэгов", []tag{} },
        {"строка с одним тэгом без даты @na", []tag{{"na",""}} },
        {"строка с двумя тэгами без даты @na @wait", []tag{{"na",""},{"wait",""}}},
        {"строка с двумя тэгами один без даты, другой с датой @na @done(2016-04-11)", []tag{{"na",""},{"done","2016-04-11"}} },
    }

    for _,c := range cases {
        out := FindTags(c.in)
        fmt.Println("\tC.in=",c.in);
        fmt.Println("\t\tin =",c.wait);
        fmt.Println("\t\tout=",out,"\n");

        if len(out) != len(c.wait) {
            t.Errorf("FindTags wrong %d != %d",len(out), len(c.wait))
        }
        for i,o := range(out) {
            if o.Tag != c.wait[i].Tag || o.Date != c.wait[i].Date {
                t.Errorf("FindTags wrong %v != %v",o, c.wait[i])
            }
        }
    }
}

func TestCheckString(t *testing.T) {
    cases := []struct{
        in mdstring
        wait bool
    }{
        {mdstring{Level:6, Text:"- [ ] стол в спалью", Tags:[]tag(nil), Show:false},
            false },
        {mdstring{Level:1, Text:"- [ ] страница про профессии @na", Tags:[]tag{tag{Tag:"na", Date:""}}, Show:false},
            true },
        {mdstring{Level:2, Text:"- [ ] 2016 @start(2016-04-10)", Tags:[]tag{tag{Tag:"start", Date:"2016-04-10"}}, Show:false},
            true },
        {mdstring{Level:2, Text:"- [ ] 2017 ivi.ru @start(2017-04-10)", Tags:[]tag{tag{Tag:"start", Date:"2017-04-10"}}, Show:false},
            false },
        {mdstring{Level:2, Text:"- [ ] что-то сдалать @done(2016-04-10)", Tags:[]tag{tag{Tag:"done", Date:"2016-04-10"}}, Show:false},
            false },
    }
    for _,c := range cases {
        out := CheckString(c.in,today)
        fmt.Printf("\tCheckString\n\t%s\n",c.in.Text)
        fmt.Printf("\t\treturned %v, we waited %v\n",out,c.wait)

        if out != c.wait {
            t.Errorf("CheckString returned %v != %v",out,c.wait)
        }
    }

}

func TestHowManyTabs(t *testing.T) {
    cases := []struct{
        str string
        tab string
        out int
    }{
        {"string","\t",0},
        {"\tstring","\t",1},
        {"\t\tstring","\t",2},
        {"string"," ",0},
        {"  string"," ",2},
        {"    string"," ",4},
    }

    for _,c := range cases {
        out := HowManyTabs(c.str,c.tab)
        if c.out != out {
            t.Errorf("HowManyTabs wrong: %d!=%d",c.out,out)
        }
    }

}
