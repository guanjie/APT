// string functions.go file, for learning purpose

package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

    p("Contains:      ", s.Contains("test", "est"))
    // string.Count(str1, str2) can directly count the elements.
    p("Count:         ", s.Count("ericericeric", "eric"))
    p("HasPrefx:      ", s.HasPrefix("This is", "Th"))
    p("HasSuffix:     ", s.HasSuffix("End this with", "with"))
    // strs := []string{"1st", "2nd", "3rd"}
    strs := []string{"1ststring", "2ndstring"}
    // Append will append and drop directly. Attach the value back.
    strs = append(strs, "asdfasdf")
    p(strs)
    p("Join:          ", s.Join(strs, "-"))
    p("Index:         ", s.Count("Strrrinrrg", "rrr"))
    p("Repeat:        ", s.Repeat("this is great\n", 5))
    p("Split:         ", s.Split("thisisit cool", " "))
    p("ToLower:       ", s.ToLower("THISISCOOL"))
    p("TOUpper:       ", s.ToUpper("this is cool man"))
    p()

    p("Len:           ", len("guanjiehe"))
    // TRICK: "string"[2] will return the char int value
    p("Char:          ", string("geanjiehe"[1]))
    newstr := "asdf" + "guanjie"
    p(newstr)

}
