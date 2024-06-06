package slides

import (
	"oddshub/models"

	"github.com/rivo/tview"
)

// FootballSlide creates a slide for football odds.
func NFLfootball(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	tableData := "Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total" + "\n"
	for _, game := range games {
		tableData += FormatTeamEvent(game) 
	}
	return "NFL", GetHeader(models.Americanfootball_nfl), CreateH2HTable("NFL Football", tableData)
}
