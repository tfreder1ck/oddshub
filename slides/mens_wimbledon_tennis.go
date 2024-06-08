package slides

import (
	"oddshub/models"
	"strings"

	"github.com/rivo/tview"
)

func MensWimbledonTennis(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Location|Teams|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Wimbledon", GetHeader(models.Tennis_atp_wimbledon), CreateH2HTable(string(models.Tennis_atp_wimbledon), builder.String())
}
