# Go example projects

ljc-deploy is wrapper on top of Helm and Public Cloud Provider that contains a collection of Go programs and libraries to deploy your application in pipeline

## Clone the project

```
$ git clone https://github.com/ljcheng999/ljc-deploy.git
$ cd ljc-deploy
```

## Examples:

```
$ cd ljc-deploy
$ make build
$ ./bin/ljc-deploy -help

//output
Available Commands:
  gitlab      A subcommand where the ljc-deploy is to deploy the version control platform
  help        Help about any command

Flags:
  -h, --help       help for ljc-deploy
      --log-json   Display json output format in the console. (default: false)
      --log-text   Display text output format in the console. (default: false)
  -v, --version    version for ljc-deploy
```

```
$ cd ljc-deploy
$ make runh

//Output
Available Commands:
  gitlab      A subcommand where the ljc-deploy is to deploy the version control platform
  help        Help about any command

Flags:
  -h, --help       help for ljc-deploy
      --log-json   Display json output format in the console. (default: false)
      --log-text   Display text output format in the console. (default: false)
  -v, --version    version for ljc-deploy
```

Gitlab subcommand:

```
$ cd ljc-deploy
$ make labh

//Output
Usage:
  ljc-deploy gitlab [flags]

Flags:
      --chart string                Helm chart name. (default: '')
      --chart-registry-url string   Helm chart registry url.
      --chart-version string        Helm chart version. (default: '')
      --cloud-provider string       Public Cloud Provider (default "aws")
  -c, --cluster string              Deployment cluster name. (default: '')
      --env string                  Gitlab deployment environment. (default "develop")
  -f, --file string                 Helm deployment values file. (default: '')
  -h, --help                        help for gitlab
      --kubeconfig string           kubeconfig file (default is $HOME/kubeconfig)
  -n, --namespace string            Deployment namespace. (default "default")
      --password string             Helm chart password.
  -p, --project string              Helm deployment project name.
      --region string               AWS region (default "us-east-1")
  -r, --release string              Helm release name. (default: '')
      --role-arn string             AWS role arn to be used
      --username string             Helm chart username.

Global Flags:
      --log-json   Display json output format in the console. (default: false)
      --log-text   Display text output format in the console. (default: false)
```

The `ljc-deploy` command covers:

- The basic of an executable command
- Logging in text/json format ([log/slog](https://pkg.go.dev/log/slog)) - the `log/slog` package supports structured logging. It features a flexible backend in the form of a `Handler` format. This guide can help you write your own handler.

The `gitlab` subcommand covers many flags, run `--help` flag with `ljc-deploy gitlab --help` for more information
