/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func Test_runRootCmd(t *testing.T) {
	type args struct {
		c   *cobra.Command
		in1 []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runRootCmd(tt.args.c, tt.args.in1)
		})
	}
}
