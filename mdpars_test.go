package main

import (
    "testing"
    // "fmt"
)

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
