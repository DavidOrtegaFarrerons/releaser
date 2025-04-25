package ui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"release-handler/config"
	"release-handler/internal/helper"
	"release-handler/internal/models"
	"release-handler/internal/release"
	"release-handler/internal/scm/azure"
	_ "release-handler/internal/scm/azure"
	"runtime"
)

func CreateReleaseTable() {
	mergedTickets := release.MergeTickets(viper.GetString(config.JiraDefaultRelease))
	ReleaseTable(mergedTickets)
}

// ReleaseTable renders a table with pull requests
func ReleaseTable(tableTickets []models.TableTicket) {
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false)

	headers := []string{"Created By", "Ticket Name", "Jira", "Azure", "isApproved", "Action"}
	for col, title := range headers {
		table.SetCell(0, col, tview.NewTableCell(title).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter).
			SetSelectable(false))
	}

	for row, tableTicket := range tableTickets {
		if tableTicket.PullRequest != nil {
			table.SetCell(row+1, 0, tview.NewTableCell(tableTicket.PullRequest.CreatedBy.DisplayName).SetAlign(tview.AlignLeft))
			table.SetCell(row+1, 1, tview.NewTableCell(tableTicket.Ticket.Key).SetAlign(tview.AlignLeft))
			table.SetCell(row+1, 2, tview.NewTableCell(tableTicket.Ticket.Fields.Status.Name).SetAlign(tview.AlignLeft))
			table.SetCell(row+1, 3, tview.NewTableCell(tableTicket.PullRequest.Status).SetAlign(tview.AlignLeft).SetTextColor(textColorByStatus(tableTicket.PullRequest.Status)))
			table.SetCell(row+1, 4, tview.NewTableCell(azure.GetFinalReviewStatus(tableTicket.PullRequest.Reviewers)).SetAlign(tview.AlignLeft).SetTextColor(textColorByStatus(tableTicket.PullRequest.Status)))
			table.SetCell(row+1, 5, tview.NewTableCell("[Go to PR]").
				SetTextColor(tcell.ColorBlue).
				SetSelectable(true).
				SetAlign(tview.AlignCenter))
		} else {
			table.SetCell(row+1, 0, tview.NewTableCell(tableTicket.Ticket.Fields.Assignee.DisplayName).SetAlign(tview.AlignLeft))
			table.SetCell(row+1, 1, tview.NewTableCell(tableTicket.Ticket.Key).SetAlign(tview.AlignLeft))
			table.SetCell(row+1, 2, tview.NewTableCell(tableTicket.Ticket.Fields.Status.Name).SetAlign(tview.AlignLeft))
			table.SetCell(row+1, 3, tview.NewTableCell("No PR Found!").SetAlign(tview.AlignLeft).SetTextColor(tcell.ColorRed))
			table.SetCell(row+1, 4, tview.NewTableCell("N/A").SetAlign(tview.AlignLeft))
			table.SetCell(row+1, 5, tview.NewTableCell("[Go to JIRA]").
				SetTextColor(tcell.ColorBlue).
				SetSelectable(true).
				SetAlign(tview.AlignCenter))
		}
	}

	table.SetSelectedFunc(func(row, column int) {
		if row > 0 && row <= len(tableTickets) {
			selectedPR := tableTickets[row-1]

			url := ""
			if selectedPR.PullRequest != nil {
				url = helper.GenerateUrl("azure", selectedPR)
			} else {
				url = helper.GenerateUrl("jira", selectedPR)
			}

			err := openInBrowser(url)

			if err != nil {
				fmt.Println("There has been an error generating the url")
				os.Exit(1)
			}
		}
	})

	app := tview.NewApplication().SetRoot(table, true)
	if err := app.Run(); err != nil {
		fmt.Println("Error running application:", err)
	}
}

func openInBrowser(url string) error {
	if url == "" {
		return fmt.Errorf("empty URL, cannot open in browser")
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	return cmd.Start()
}

func textColorByStatus(status string) tcell.Color {
	switch status {
	case "active":
		return tcell.ColorGreen
	default:
		return tcell.ColorWhite
	}
}
