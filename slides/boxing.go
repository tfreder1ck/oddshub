package slides

import (
	"oddshub/models"
  "strings"
	"github.com/rivo/tview"
)

// Boxing creates a slide for boxing odds.
func Boxing(games []models.Event, nextSlide func()) (string, string, tview.Primitive) {
	var builder strings.Builder
	builder.WriteString("Commencement Date|Ranking|Players|Bookmaker|Spread|Money|Total\n")

	for _, game := range games {
		builder.WriteString(FormatTeamEvent(game))
	}

	return "Boxing", GetHeader(models.Boxing), CreateH2HTable("Boxing", builder.String())
}
