/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/ljcheng999/ljc-app-deploy/cmd"
	_ "github.com/ljcheng999/ljc-app-deploy/cmd/gitlab"
)

func main() {
	cmd.Execute()
}
