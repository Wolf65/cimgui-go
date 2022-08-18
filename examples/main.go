package main

import (
	"fmt"

	"github.com/AllenDang/cimgui-go"
)

var (
	showDemoWindow bool
	value1         int32
	value2         int32
	values         [2]*int32 = [2]*int32{&value1, &value2}
	content        string    = "Let me try"
)

func callback(data cimgui.ImGuiInputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

func loop() {
	if showDemoWindow {
		cimgui.ShowDemoWindow(&showDemoWindow)
	}

	cimgui.SetNextWindowSize(cimgui.NewImVec2(300, 300), cimgui.ImGuiCond_Once)
	cimgui.Begin("Window 1", nil, 0)
	cimgui.Checkbox("Show demo window", &showDemoWindow)
	cimgui.InputTextWithHint("Name", "write your name here", &content, 0, callback)
	cimgui.Text(content)
	cimgui.DragInt("Drag int", &value1, 1, 0, 100, "%d", 0)
	cimgui.DragInt2("Drag int2", &values, 1, 0, 100, "%d", 0)
	cimgui.End()
}

func main() {
	window := cimgui.CreateGlfwWindow("Hello from cimgui-go", 1200, 900)

	window.Run(loop)
}
