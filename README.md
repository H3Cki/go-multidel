# go-multidel 
This package provides the most basic functions for deleting multiple slice elements at once

## Usage
```go
import "github.com/H3Cki/go-multidel"
```

## Examples
```go
type thing struct {
    name string
}

func main() {
    // Delete
    items := []int{1, 2, 3, 4, 5}
    multidel.Delete(&items, 0, 2)

    fmt.Println(items) // [2, 4, 5]

    // DeleteFunc
    things := []thing{
        {name: "well-being"},
        {name: "luxury"},
        {name: "poverty"},
        {name: "friends"},
        {name: "family"},
        {name: "cancer"},
        {name: "happiness"},
    }
    multidel.DeleteFunc(&things, func(i thing) bool {
        return i.name == "poverty" || i.name == "cancer"
    })

    fmt.Println(items) // [{well-being} {luxury} {friends} {family} {happiness}]
}
```