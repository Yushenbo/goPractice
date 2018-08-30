//calc.go
package main

import "os" // retrive args from command line
import "fmt"
import "simplemath" // mathmatic library
import "strconv" // str conversion, maybe int/float numbers transmission

var Usage = func() {
    fmt.Println("USAGE: calc command [arguments] ... ")
    fmt.Println("\nThe commands are: \n\tadd\taddition of 2 values.\n\tsqrt\tSquare rot of a non-negative value.")
}

func main(){
    args := os.Args
    if args == nil || len(args) < 2 {
        fmt.Println("Args wrong")
        Usage()
        return
    }

    switch args[0] {
        case "add":
            if len(args) != 3 {
                fmt.Println("args number USAGE: calc add <interger1><integer2>")
                return
            }

            v1, err1 := strconv.Atoi(args[1])
            v2, err2 := strconv.Atoi(args[2])
            if err1 != nil || err2 != nil {
                fmt.Println("USAGE: calc add <integer><integer>")
                return
            }
            ret := simplematch.Add(v1, v2)
            fmt.Println("Result: ", ret)
            break;
        case "sqrt":
            if len(args) != 2 {
                fmt.Println("USAGE: calc sqrt <integer>")
                return
            }
            v, err := strconv.Atoi(args[1])
            if err != nil {
                fmt.Println("USAGE: calc sqrt <integer>")
                return
            }
            ret := simplemath.Sqrt(v)
            fmt.Println("Result: ", ret)
            break;
        default:
            fmt.Println("go to defaut condition")
            Usage()
    }
}
