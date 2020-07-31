package main

import (
  "unsafe"
  "encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	js.Global().Set("goAdd", addWrapper())
	js.Global().Set("goNewThing", newThingWrapper())

	<-make(chan bool)
}

func add(a, b int) int {
	sum := a + b

	return sum
}

func addWrapper() js.Func {
	wrapper := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 2 {
			return "ERROR::need more args"
		}

		fmt.Printf("this is: %v\n", this)
		fmt.Printf("args is: %v\n", args)

		a := args[0].Int()
		b := args[1].Int()
		result := add(a, b)

		return result
	})

	return wrapper
}

// Thing :: a thing, a thing, know what i mean?
type Thing struct {
	Name string `json:"name"`
	Code int    `json:"code"`
}

func (t Thing) String() string {
	str := fmt.Sprintf("%s %d", t.Name, t.Code)

	return str
}

// ToValueOf :: 
func (t Thing) ToValueOf() map[string]interface{} {
  value := map[string]interface{}{
    "name": t.Name,
    "code": t.Code,
  }

  return value
}

func newThing(name string, code int) Thing {
	t := Thing{
		Name: name,
		Code: code,
	}

	return t
}

func newThingWrapper() js.Func {
	wrapper := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Printf("this is: %v\n", this)
		fmt.Printf("args is: %v\n", args)

		name := args[0].String()
		code := args[1].Int()
		fmt.Printf("args is %s and %d\n", name, code)
		result := newThing(name, code)

    fmt.Println("returning newThingWrapper")
    vo := result.ToValueOf()
    fmt.Printf("result is %v\n", vo)
    
    thingBytes, err := json.Marshal(result)
    if err != nil {
      return err.Error()
    }
    stringified := string(thingBytes)
    fmt.Printf("stringified: %s\n", stringified)

    fmt.Printf("size of struct:                     %d\n", unsafe.Sizeof(result))
    fmt.Printf("size of map[string]interface{}:     %d\n", unsafe.Sizeof(vo))
    fmt.Printf("size of json string:                %d\n", unsafe.Sizeof(stringified))
  
		return js.ValueOf(vo)
	})

	return wrapper
}
