## Sample code layout for cli application use go module style

### Usage
```$bash
$ git clone https://github.com/alimy/cli
$ cd cli
$ make

$ ./cli
a sample cli application base on go module style

Usage:
  cli [command]

Available Commands:
  hello       sample sub-command for cli
  help        Help about any command
  version     version of application

Flags:
  -h, --help   help for cli

Use "cli [command] --help" for more information about a command.
$ ./cli version
0.0.0 (APIVersion:v1alphal)
BuildTime:2018-12-27 11:58:13 UTC
BuildGitSHA:6f61a2376753e270b943d5ab63eaeb85357b8486

$ ./cli hello
2018-12-27T20:00:06.390+0800    INFO    hello@v0.0.0-20181227113021-4b91de8a165a/hello.go:6     hello world!
```

### Custom your code base this repository

* change module name and others
* add sub-module in module directory
* import sub-module in main.go to register sub-command
* and so on...
