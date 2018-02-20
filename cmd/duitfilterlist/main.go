// Duitfilterlist demonstrates a Filterlist.
package main

import (
	"log"

	"github.com/mjl-/duit"
	"github.com/mjl-/filterlist"
)

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

func main() {
	dui, err := duit.NewDUI("duitfilterlist", nil)
	check(err, "new dui")

	var list *duit.List
	list = &duit.List{
		Values: []*duit.ListValue{
			{Text: "apple"},
			{Text: "lemon"},
			{Text: "lime"},
			{Text: "orange"},
			{Text: "kiwi"},
			{Text: "banana"},
			{Text: "mango"},
			{Text: "grapes"},
			{Text: "pear"},
			{Text: "pomegranate"},
			{Text: "strawberry"},
			{Text: "cherry"},
			{Text: "papaya"},
			{Text: "blackberry"},
			{Text: "blueberry"},
		},
		Changed: func(index int) (e duit.Event) {
			var lv duit.ListValue
			if index >= 0 {
				lv = *list.Values[index]
			}
			log.Printf("selection changed, now %q, selected %v\n", lv.Text, lv.Selected)
			return
		},
	}
	fl := filterlist.NewFilterlist(dui, list)
	dui.Top.UI = fl
	dui.Render()
	dui.Focus(fl.Search)

	for {
		select {
		case e := <-dui.Inputs:
			dui.Input(e)

		case err, ok := <-dui.Error:
			if !ok {
				return
			}
			log.Printf("duit: %s\n", err)
		}
	}
}
