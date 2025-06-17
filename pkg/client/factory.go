package client

import "github.com/spf13/pflag"

type factory struct {
	flags *pflag.FlagSet

	clusterName string
	baseName    string
	namespace   string
}
