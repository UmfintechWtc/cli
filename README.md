### go version
```python
go version go1.20.1 linux/amd64
```

### go build
```python
make exec cli binary file
```

### Description
> Pod executes shell commands

### Usage 
```python
[root@wangtianci:cli_0]# ./cli exec
Error: requires at least 1 arg(s), only received 0
Usage:
  cli exec [flags]

Flags:
      -- string            指明需要执行的CLI Command
  -c, --container string   Container名称
  -h, --help               help for exec
  -n, --namespace string   NameSpace名称 (default "default")
  -s, --service string     Pod名称

requires at least 1 arg(s), only received 0
```