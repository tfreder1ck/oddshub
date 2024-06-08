package slides

import (
	"oddshub/models"
  "strings"
	"github.com/rivo/tview"
)

func EPLSoccer(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "EPL Soccer", GetHeader(models.Soccer_epl), CreateH2HTable(string(models.Soccer_epl), builder.String())
}
