// Duitfiltergridlist demonstrates a Filtergridlist.
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
	dui, err := duit.NewDUI("duitfiltergridlist", nil)
	check(err, "new dui")

	var gl *duit.Gridlist
	gl = &duit.Gridlist{
		Padding: duit.SpaceXY(4, 2),
		Striped: true,
		Header: &duit.Gridrow{
			Values: []string{"fruit", "color"},
		},
		Rows: []*duit.Gridrow{
			{Values: []string{"apple", "green"}},
			{Values: []string{"lemon", "yellow"}},
			{Values: []string{"lime", "green"}},
			{Values: []string{"orange", "orange"}},
			{Values: []string{"kiwi", "brown/green"}},
			{Values: []string{"banana", "yellow"}},
			{Values: []string{"mango", "green/red"}},
			{Values: []string{"grapes", "blue/white"}},
			{Values: []string{"pear", "green"}},
			{Values: []string{"pomegranate", "purpple"}},
			{Values: []string{"strawberry", "red"}},
			{Values: []string{"cherry", "red"}},
			{Values: []string{"papaya", "orange"}},
			{Values: []string{"blackberry", "blue"}},
			{Values: []string{"blueberry", "blue"}},
		},
		Changed: func(index int) (e duit.Event) {
			var row duit.Gridrow
			if index >= 0 {
				row = *gl.Rows[index]
			}
			log.Printf("selection changed, now %v, selected %v\n", row, row.Selected)
			return
		},
	}
	fl := filterlist.NewFiltergridlist(dui, gl)
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
