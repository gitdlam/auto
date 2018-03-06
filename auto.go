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

type MyMainWindow struct {
	*walk.MainWindow
}

var singleKeys map[string]bool
var mw *MyMainWindow

func main() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)

	var inTE, outTE *walk.TextEdit

	singleKeys = make(map[string]bool)

	for _, v := range strings.Split("a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,left,right,up,down,enter,tab,space,escape,f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12", ",") {
		singleKeys[v] = true
	}

	mw = new(MyMainWindow)

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
					TextEdit{AssignTo: &outTE, ReadOnly: true, Font: Font{Family: "Arial", PointSize: 11}, Text: "Please replace existing text on the left if needed.\r\n\r\nAvailable special tags:\r\n[space][tab][enter][escape][left][right][up][down]\r\n[f1] to [f12]\r\n[shift-*][alt-*][control-*] where * is an alphabet or any of the above"},
				},
			},
			PushButton{
				Text: "Run",
				Font: Font{Family: "Arial", PointSize: 16, Bold: true},
				OnClicked: func() {
					problem := processInput(inTE.Text(), true)

					if problem == "" {
						walk.MsgBox(mw, "Info", "Please click on the target window within 5 seconds after clicking OK.", walk.MsgBoxIconInformation)

						time.Sleep(5 * time.Second)

						processInput(inTE.Text(), false)
						outTE.SetText("Done.")
					} else {
						walk.MsgBox(mw, "Info", problem, walk.MsgBoxIconExclamation)
						outTE.SetText(problem)
					}

				},
			},

			PushButton{
				Text: "About",

				OnClicked: func() {
					walk.MsgBox(mw, "Info", "open source software:\r\ngithub.com/gitdlam/auto", walk.MsgBoxIconInformation)

				},
			},
		},
	}.Create()

	if err != nil {
		log.Println(err.Error())
	}

	mw.Run()
}

func processInput(input string, checkOnly bool) string {
	//title := robotgo.GetTitle()
	//a := robotgo.GetActive()
	//robotgo.SetHandle(robotgo.GetHandle())
	//robotgo.ActivePID(int32(robotgo.GetPID()))
	//robotgo.ActiveName(title)
	badTags := ""
	input = strings.Replace(input, "[", " [", -1)
	input = strings.Replace(input, "]", "] ", -1)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		if s[0] == byte('[') {
			matched := false
			key := strings.Trim(s, "[]")

			if singleKeys[key] {
				matched = true
				if !checkOnly {
					robotgo.KeyTap(key)
				}

			}

			if !matched && len(key) > 8 && key[:8] == "control-" && singleKeys[key[8:]] {
				matched = true
				if !checkOnly {
					robotgo.KeyTap(key[8:], "control")
				}

			}

			if !matched && len(key) > 6 && key[:6] == "shift-" && singleKeys[key[6:]] {
				matched = true
				if !checkOnly {
					robotgo.KeyTap(key[6:], "shift")
				}
			}

			if !matched && len(key) > 4 && key[:4] == "alt-" && singleKeys[key[4:]] {
				matched = true
				if !checkOnly {
					robotgo.KeyTap(key[4:], "alt")
				}

			}

			if !matched {
				badTags = badTags + s

			}

		} else {
			if !checkOnly {
				robotgo.TypeStrDelay(s, 1500)
			}
		}

	}

	if badTags == "" {
		return ""
	} else {
		return "Program aborted due to unimplemented tag(s): " + badTags
	}

}
