package main

//-ldflags="-H windowsgui -linkmode external"
import (
	"bufio"
	"log"
	//"fmt"
	"strings"

	"time"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	type MyMainWindow struct {
		*walk.MainWindow
	}

	mw := new(MyMainWindow)

	err := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Automate Keystrokes",
		MinSize:  Size{400, 300},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, Text: `Hello [space] world [enter] 1 [tab] 2 [tab] 3 [control-s]`},
					TextEdit{AssignTo: &outTE, ReadOnly: true, Text: `Please modify existing text on the left.

Important to have space separating the [...] tags.`},
				},
			},
			PushButton{
				Text: "Run",
				OnClicked: func() {
					walk.MsgBox(mw, "Info", "Please click on the target window within the next 5 seconds.", walk.MsgBoxIconInformation)

					time.Sleep(5 * time.Second)
					processInput(inTE.Text())

					//outTE.SetText(strings.ToUpper(inTE.Text()))
					outTE.SetText("Done.")
				},
			},
		},
	}.Create()

	if err != nil {
		log.Println(err.Error())
	}

	mw.Run()
}

func processInput(input string) {
	//title := robotgo.GetTitle()
	//a := robotgo.GetActive()

	//robotgo.SetHandle(robotgo.GetHandle())

	//robotgo.ActivePID(int32(robotgo.GetPID()))
	//fmt.Println(title)
	//lines := strings.Split(input, "\n")
	//robotgo.ActiveName(title)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		if s[0] == byte('[') {
			switch s {
			case "[space]":
				robotgo.KeyTap("space")
			case "[tab]":
				robotgo.KeyTap("tab")
			case "[enter]":
				robotgo.KeyTap("enter")
			case "[control-s]":
				robotgo.KeyTap("s", "control")

			}

		} else {
			robotgo.TypeStrDelay(s, 1500)
		}

	}

}
