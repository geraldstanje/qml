package main

import (
	"fmt"
	"gopkg.in/qml.v1"
	"os"
)

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	controls, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}

  ctrl := Control{}

  context := engine.Context()
  context.SetVar("ctrl", &ctrl)

	window := controls.CreateWindow(nil)
  window.On("heightChanged", func(hight int) {
    fmt.Println("Hight:", hight)
  })

  ctrl.Root = window.Root()

  window.Show()	
	window.Wait()
	return nil
}

type Control struct {
  Root    qml.Object
  //Message string
}

func (ctrl *Control) Loaded(obj qml.Object) {
  fmt.Println("onCompleted() called")

  fmt.Println(obj.Int("count"))

  fmt.Println(ctrl.Root)

  fmt.Println(obj.TypeName())

  //obj.Call("append", "\"name\":\"Person \"+0 , \"age\": 20, \"gender\": \"Male\"")

  //obj.Call("append", "\"cost\": 5.95, \"name\":\"Pizza\", \"name\":\"Pizza\"")

  //fmt.Println(obj.Int("count"))

  //obj.Call("update")

  //for i:=0; i < 500; i++ {
  //  largeModel.Set("append", \""cost": 5.95, "name":"Pizza"\")

    //largeModel.append("test")//name":"Person ", i, "age:", Math.round(Math.random()*100), "gender": Math.random()>0.5 ? "Male" : "Female")
  //}
}
