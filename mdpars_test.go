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
