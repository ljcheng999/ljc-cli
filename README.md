# ljc-cli projects

ljc-cli is wrapper on top of Helm and Public Cloud Provider that contains a collection of Go programs and libraries to deploy your application in pipeline

## Clone the project

```
$ git clone https://github.com/ljcheng999/ljc-cli.git
$ cd ljc-cli
```

## Examples:

```
$ cd ljc-cli
$ make build
$ ./bin/ljc-cli -help

//output
Usage:
  ljc-cli [flags]
  ljc-cli [command]

Available Commands:
  deploy      A subcommand where the ljc-cli is to deploy the application
  help        Help about any command

Flags:
  -h, --help       help for ljc-cli
      --log-json   Display json output format in the console. (default: false)
      --log-text   Display text output format in the console. (default: false)
  -v, --version    version for ljc-cli
```

```
$ cd ljc-deploy
$ make runh

//Output
Usage:
  ljc-cli [flags]
  ljc-cli [command]

Available Commands:
  deploy      A subcommand where the ljc-cli is to deploy the application
  help        Help about any command

Flags:
  -h, --help       help for ljc-cli
      --log-json   Display json output format in the console. (default: false)
      --log-text   Display text output format in the console. (default: false)
  -v, --version    version for ljc-cli
```

Gitlab subcommand:

```
$ cd ljc-cli
$ make dh

//Output
Usage:
  ljc-cli deploy [flags]

Flags:
      --chart string                Helm chart name. (default: '')
      --chart-registry-url string   Helm chart registry url.
      --chart-version string        Helm chart version. (default: '')
      --cloud-provider string       Public Cloud Provider (default "aws")
  -c, --cluster string              Deployment cluster name. (default: '')
  -f, --file string                 Helm deployment values file. (default: '')
      --git string                  git version control platform (default '')
  -h, --help                        help for deploy
      --kubeconfig string           kubeconfig file (default is $HOME/kubeconfig)
  -n, --namespace string            Deployment namespace. (default "default")
      --password string             Helm chart password.
      --region string               AWS region (default "us-east-1")
  -r, --release string              Helm release name. (default: '')
      --role-arn string             AWS role arn to be used
      --username string             Helm chart username.

Global Flags:
      --log-json   Display json output format in the console. (default: false)
      --log-text   Display text output format in the console. (default: false)
```

The `ljc-cli` command covers:

- The basic of an executable command
- Logging in text/json format ([log/slog](https://pkg.go.dev/log/slog)) - the `log/slog` package supports structured logging. It features a flexible backend in the form of a `Handler` format. This guide can help you write your own handler.

The `deploy` subcommand covers many flags, run `--help` flag with `ljc-cli deploy --help` for more information
