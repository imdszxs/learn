package main
import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

// 获取输入的参数
func getInput() string {
    var str string
    fmt.Print("请输入字母字符串：")
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
    str = input.Text()
    fmt.Println("您输入的字符串是：", str)

    if !isLetter(str){
        fmt.Println("您输入的字符串不是纯字母组成")
        getInput()
    }
    return str
}

func getType() string {
    var dealType string
    fmt.Print("请输入处理方式（1.不包含重复字符串全排列 2.包含重复字符串的全排列 3.不计较顺序的排列）：")
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
    dealType = input.Text()

    switch dealType {
    case "1":
        fmt.Println("您选择的处理方式是不包含重复字符串全排列")
        break
    case "2":
        fmt.Println("您选择的处理方式是包含重复字符串的全排列")
        break
    case "3":
        fmt.Println("您选择的处理方式是不计较顺序的排列")
        break
    default:
        fmt.Println("不存在的处理方式！")
        getType()
    }
    return dealType
}

// 字符串是否是字母
func isLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func dealStr(strings []string) [][]string{
    if 1 == len(strings) {
        return [][]string{strings}
    }

    var res [][]string

    for k,v := range strings{
        temp := make([]string, len(strings) - 1)
        copy(temp[0:], strings[0:k])
        copy(temp[k:], strings[k+1:])

        sub := dealStr(temp)

        for _,subV := range sub{
            res = append(res, append(subV, v))
        }
    }

    return res
}


func main() {
    str := getInput()

    dealType := getType()

    switch dealType {
        case "1":
            fmt.Println("您选择的处理方式是不包含重复字符串全排列")
            break
        case "2":
            fmt.Println("您选择的处理方式是包含重复字符串的全排列")
            break
        case "3":
            fmt.Println("您选择的处理方式是不计较顺序的排列")
            break
    }

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