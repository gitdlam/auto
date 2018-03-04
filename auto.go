package main

//-ldflags="-H windowsgui -linkmode external"
import (
	"fmt"
	"strings"

	"time"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "Automate Keystrokes",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true, Text: "Please type or modify existing text on the left."},
				},
			},
			PushButton{
				Text: "Run",
				OnClicked: func() {

					robotgo.ShowAlert("message", "Please click on the target window within the next 5 seconds.")
					time.Sleep(5 * time.Second)
					processInput(inTE.Text())

					//outTE.SetText(strings.ToUpper(inTE.Text()))
					outTE.SetText("Done.")
				},
			},
		},
	}.Run()
}

func processInput(input string) {
	title := robotgo.GetTitle()
	fmt.Println(title)
	lines := strings.Split(input, "\n")
	robotgo.ActiveName(title)
	for _, l := range lines {
		robotgo.TypeString(l)
	}

}
