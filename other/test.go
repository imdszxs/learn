package main

import (
    "fmt"
)

func dealStr(strings []string) [][]string{
    if 1 == len(strings) {
        return [][]string{strings}
    }

    var res [][]string

    for k,v := range strings{
        temp := make([]string, len(strings) - 1)
        copy(temp[0:], strings[0:k])
        copy(temp[k:], strings[k+1:])
        fmt.Println(temp)

        sub := dealStr(temp)

        for _,subV := range sub{
            res = append(res, append(subV, v))
        }
    }

    return res
}


func main() {
    str := "abc"

    // 处理字符串为切片
    var aaa []string
    for _,v := range []rune(str){
        aaa = append(aaa, string(v))
    }

    // 计算无重复字符串的全排列
    arr := dealStr(aaa)

    // 输出全排列的总数
    fmt.Printf("全排列的总数%v\n", len(arr))

    // 输出结果
    fmt.Println(arr)
}