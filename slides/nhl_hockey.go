package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

func NHLHockey(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "NHL", GetHeader(models.Icehockey_nhl), CreateH2HTable("NHL Hockey", tableData)
}
