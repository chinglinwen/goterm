# GoTerm

A ssh client terminal written in go

with replace for expect tool in minds.

## Problems

I used to use expect for machine jumps
It got tricky issue that expect can't properly handle 
terminal resize or other issues.

Therefore this tool, also for easier hosts management.

## Features

* Friendly with tmux(terminal resize)
* Replace expect tool
* Support password and key based auth
* Share credential for many hosts
* Config file for hosts and credentials
* One line ssh login(with password, if you like)

## usage

```bash
Usage: goterm <name>
       goterm -ip <127.0.0.1> -user <user> [ -pass <pass> ]
       goterm -ip <127.0.0.1> -user <user> [ -pass <pass> ] [ -p 2222 ]
```

## Config

config example

```
$ cat ~/.goterm/config.yaml 
creds:
- name: vm
  user: root
  pass: password
  keypath: 
hosts:
- name: 11
  host: 10.47.120.11
  cred: vm
  port: 22
```

## Acknowledgements

Thanks to https://mritd.me/2018/11/09/go-interactive-shell/
