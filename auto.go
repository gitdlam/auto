package main

//-ldflags="-H windowsgui -linkmode external"
import (
	"bufio"
	"io/ioutil"
	"log"
	"path/filepath"
	//"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)

	var inTE, outTE *walk.TextEdit

	type MyMainWindow struct {
		*walk.MainWindow
	}

	mw := new(MyMainWindow)

	text := ""
	defaultText := `Hello [space] world [enter] 1 [tab] 2 [tab] 3 [control-s]`
	b, err := ioutil.ReadFile(exPath + "\\auto.txt")
	if err == nil {
		text = string(b)
	} else {
		text = defaultText
	}
	err = MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Automate Keystrokes",
		MinSize:  Size{450, 400},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, Text: text, Font: Font{Family: "Arial", PointSize: 11}},
					TextEdit{AssignTo: &outTE, ReadOnly: true, Font: Font{Family: "Arial", PointSize: 11}, Text: "Please replace existing text on the left if needed.\r\n\r\nAvailable special tags:\r\n[space][tab][enter][escape][control-s]\r\n[f1] to [f12]\r\n[shift-f1] to [shift-f12]\r\n[alt-f1] to [alt-f12]"},
				},
			},
			PushButton{
				Text: "Run",
				Font: Font{Family: "Arial", PointSize: 16, Bold: true},
				OnClicked: func() {
					walk.MsgBox(mw, "Info", "Please click on the target window within 5 seconds after clicking OK.", walk.MsgBoxIconInformation)

					time.Sleep(5 * time.Second)
					processInput(inTE.Text())

					//outTE.SetText(strings.ToUpper(inTE.Text()))
					outTE.SetText("Done.")
				},
			},
			Label{Text: "open source software: github.com/gitdlam/auto", Font: Font{Family: "Arial", PointSize: 8}},
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
	input = strings.Replace(input, "[", " [", -1)
	input = strings.Replace(input, "]", "] ", -1)
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
			case "[shift-f1]":
				robotgo.KeyTap("f1", "shift")
			case "[shift-f2]":
				robotgo.KeyTap("f2", "shift")
			case "[shift-f3]":
				robotgo.KeyTap("f3", "shift")
			case "[shift-f4]":
				robotgo.KeyTap("f4", "shift")
			case "[shift-f5]":
				robotgo.KeyTap("f5", "shift")
			case "[shift-f6]":
				robotgo.KeyTap("f6", "shift")
			case "[shift-f7]":
				robotgo.KeyTap("f7", "shift")
			case "[shift-f8]":
				robotgo.KeyTap("f8", "shift")
			case "[shift-f9]":
				robotgo.KeyTap("f9", "shift")
			case "[shift-f10]":
				robotgo.KeyTap("f10", "shift")
			case "[shift-f11]":
				robotgo.KeyTap("f11", "shift")
			case "[shift-f12]":
				robotgo.KeyTap("f12", "shift")
			case "[escape]":
				robotgo.KeyTap("escape")

			case "[alt-f1]":
				robotgo.KeyTap("f1", "alt")
			case "[alt-f2]":
				robotgo.KeyTap("f2", "alt")
			case "[alt-f3]":
				robotgo.KeyTap("f3", "alt")
			case "[alt-f4]":
				robotgo.KeyTap("f4", "alt")
			case "[alt-f5]":
				robotgo.KeyTap("f5", "alt")
			case "[alt-f6]":
				robotgo.KeyTap("f6", "alt")
			case "[alt-f7]":
				robotgo.KeyTap("f7", "alt")
			case "[alt-f8]":
				robotgo.KeyTap("f8", "alt")
			case "[alt-f9]":
				robotgo.KeyTap("f9", "alt")
			case "[alt-f10]":
				robotgo.KeyTap("f10", "alt")
			case "[alt-f11]":
				robotgo.KeyTap("f11", "alt")
			case "[alt-f12]":
				robotgo.KeyTap("f12", "alt")

			case "[f1]":
				robotgo.KeyTap("f1")
			case "[f2]":
				robotgo.KeyTap("f2")
			case "[f3]":
				robotgo.KeyTap("f3")
			case "[f4]":
				robotgo.KeyTap("f4")
			case "[f5]":
				robotgo.KeyTap("f5")
			case "[f6]":
				robotgo.KeyTap("f6")
			case "[f7]":
				robotgo.KeyTap("f7")
			case "[f8]":
				robotgo.KeyTap("f8")
			case "[f9]":
				robotgo.KeyTap("f9")
			case "[f10]":
				robotgo.KeyTap("f10")
			case "[f11]":
				robotgo.KeyTap("f11")
			case "[f12]":
				robotgo.KeyTap("f12")

			}

		} else {
			robotgo.TypeStrDelay(s, 1500)
		}

	}

}
