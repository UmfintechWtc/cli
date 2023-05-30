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

### Usage with pod
```python
[root@tianciwang:cli]# ./cli exec_pod
Error: requires at least 1 arg(s), only received 0
Usage:
  cli exec_pod [flags]

Flags:
      -- string            指明需要执行的CLI Command
  -c, --container string   Container名称
  -h, --help               help for exec_pod
  -m, --mode string        当前执行环境类型 (default "host")
  -n, --namespace string   NameSpace名称 (default "default")
  -s, --service string     Pod名称

requires at least 1 arg(s), only received 0
```

### Usage with ssh
```python
[root@tianciwang:cli]# ./cli exec_ssh
Error: requires at least 1 arg(s), only received 0
Usage:
  cli exec_ssh [flags]

Flags:
      -- string           指明需要执行的CLI Command,多个命令以分号分割并引用双引号
  -a, --address string    目标主机IP (default "127.0.0.1")
  -h, --help              help for exec_ssh
  -P, --password string   目标主机用户密码
  -p, --port string       目标主机端口 (default "22")
  -u, --user string       目标主机用户 (default "root")

requires at least 1 arg(s), only received 0
```

