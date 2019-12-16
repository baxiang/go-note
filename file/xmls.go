package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)
func main() {
	file,err:= excelize.OpenFile("sim-qu.xlsx")
	if err!= nil {
		fmt.Println(err.Error())
	}
	sheetMap := file.GetSheetMap()
	rows := file.GetRows(sheetMap[1])
	for _,r :=range rows{
		fmt.Println(r[1])
	}

}
