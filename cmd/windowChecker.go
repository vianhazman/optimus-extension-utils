package cmd

import (
	"fmt"
	"github.com/odpf/optimus/models"
	"github.com/spf13/cobra"
	"time"
)

type windowCheckerParameters struct {
	schedule string
	size     string
	offset   string
	truncate string
}

func windowCheckerCmd() *cobra.Command {

	params := &windowCheckerParameters{}

	cmd := &cobra.Command{
		Use:     "window-checker",
		Aliases: []string{"wc"},
		Short:   "Optimus Window Checker",
		RunE: func(cmd *cobra.Command, args []string) error {

			size, _ := time.ParseDuration(params.size)
			offset, _ := time.ParseDuration(params.offset)

			layout := "2006-01-02T15:04"
			schedule, err := time.Parse(layout, params.schedule)

			if err != nil {
				fmt.Println(err)
				return nil
			}

			win := &models.JobSpecTaskWindow{
				Size:       size,
				Offset:     offset,
				TruncateTo: params.truncate,
			}
			windowStart := win.GetStart(schedule)
			windowEnd := win.GetEnd(schedule)
			fmt.Printf("Job will run with DSTART %s and DEND %s\n", windowStart, windowEnd)
			return nil
		},
	}

	cmd.Flags().StringVar(&params.schedule, "schedule", params.schedule, "Schedule in ISO date format")
	cmd.Flags().StringVar(&params.size, "size", params.size, "Optimus window size")
	cmd.Flags().StringVar(&params.offset, "offset", params.offset, "Optimus window offset")
	cmd.Flags().StringVar(&params.truncate, "truncate", params.truncate, "Truncate configuration for windowing")
	cmd.MarkFlagRequired("name")

	return cmd
}
