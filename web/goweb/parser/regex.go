// syntax: http://code.google.com/p/re2/wiki/Syntax

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
)

func main() {
    // get html then regexp
    resp, err := http.Get("http://www.google.com")
    if err != nil {
        fmt.Println("http get error")
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("http read error")
        return
    }

    src := string(body)

    re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
    src = re.ReplaceAllStringFunc(src, strings.ToLower)

    re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
    src = re.ReplaceAllString(src, "")

    re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
    src = re.ReplaceAllString(src, "")

    fmt.Println(strings.TrimSpace(src))

    // find regexp
    fmt.Println("=====")
    a := "I am learning Go language"

    reg, _ := regexp.Compile("[a-z]{2,4}")
    one := reg.Find([]byte(a))
    fmt.Println("Find", string(one))

    all := reg.FindAll([]byte(a), -1)
    fmt.Println("findall", all)

    index := reg.FindIndex([]byte(a))
    fmt.Println("FindIndex", index)

    allindex := reg.FindAllIndex([]byte(a), -1)
    fmt.Println("FindAllIndex", allindex)

    // Expand
    text := []byte(`
        call hello alice
        hello bob
        call hello eve
    `)
    pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
    res := []byte{}
    for _, s := range pat.FindAllSubmatchIndex(text, -1) {
        res = pat.Expand(res, []byte("$cmd('$arg')\n"), text, s)
    }
    fmt.Println(string(res))
}
