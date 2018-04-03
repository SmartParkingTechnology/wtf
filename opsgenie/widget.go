package opsgenie

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "OpsGenie",
			RefreshedAt: time.Now(),
			RefreshInt:  21600,
		},
	}

	widget.addView()
	go wtf.Refresh(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := Fetch()

	widget.View.SetTitle(" ⏰ On Call ")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(onCallResponse *OnCallResponse) string {
	str := "\n"

	for _, data := range onCallResponse.OnCallData {
		str = str + fmt.Sprintf(" [green]%s[white]\n", data.Parent.Name)
		str = str + fmt.Sprintf(" %s\n", strings.Join(data.Recipients, ", "))
		str = str + "\n"
	}

	return str
}