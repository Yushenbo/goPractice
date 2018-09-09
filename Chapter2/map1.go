package main
import "fmt"

type PersonInfo struct {
    ID string
    Name string
    Address string
}

func main() {
    var personDB map[string] PersonInfo
    personDB = make(map[string] PersonInfo)

    // Insert datas into map
    personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203, ..." }
    personDB["1"] = PersonInfo{"1", "Jack", "Room 101, ..."}

    // Query info of key "1234"
    key := "12345"
    person, ok := personDB[key]

    // Ok return a boolean value, if true returned, found this data in map
    if ok {
        fmt.Println("Found person", person.Name, "with ID", person.ID)
    } else {
        fmt.Println("Can not find person with ID", key)
    }
}
