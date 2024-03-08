package cmd

import (
	"fmt"
	"os"

	"github.com/andrewvota/scanln/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "scanln",
	Short:   "A typing game played over SSH",
	Version: "0.0.1",
	Args:    cobra.MaximumNArgs(1),
	Run: func(_ *cobra.Command, _ []string) {
		m := tui.New()

		var opts []tea.ProgramOption
		opts = append(opts, tea.WithAltScreen())

		p := tea.NewProgram(m, opts...)
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
