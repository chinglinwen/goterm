package main

import (
	"flag"
	"fmt"
	"goterm/config"
	"goterm/ssh"
	"os"
	"os/user"
	"path/filepath"

	"k8s.io/klog/v2"
)

func helpfunc() {
	flag.PrintDefaults()
	fmt.Print(`
Usage: goterm <name>
       goterm -ip <127.0.0.1> -user <user> [ -pass <pass> ]
       goterm -ip <127.0.0.1> -user <user> [ -pass <pass> ] [ -p 2222 ]
`)
}
func main() {
	ip := flag.String("ip", "", "host to connect")
	port := flag.String("port", "22", "port to connect")
	user := flag.String("user", "root", "user to auth")
	pass := flag.String("pass", "", "pass to auth, if empty, fall to key based auth")
	keyPath := flag.String("keypath", defaultKeyPath(), "private key auth")
	cmds := flag.String("initcmds", "", "init cmds after login")
	flag.Usage = helpfunc
	klog.InitFlags(flag.CommandLine)

	flag.Parse()

	// if ip provided, use it without config
	if len(*ip) != 0 {
		if len(*user) == 0 {
			exit("user not provided")
		}
		if len(*pass) == 0 {
			klog.V(2).Infof("keypath: %v", *keyPath)
		}
		// no pass provided, will use keybased
		startssh(*ip, *port, *user, *pass, *keyPath, *cmds)
		return
	}
	// if ip not provided, get it from config
	args := flag.Args()
	if len(args) == 0 {
		exit("no name or ip to connect")
	}

	klog.V(2).Infof("keypath: %v", *keyPath)

	klog.V(2).Info("args: ", args)
	name := args[0]
	klog.V(2).Infof("connecting to %v ...", name)

	c, err := config.ParseConfig()
	if err != nil {
		exiterr("parse config ", err)
	}
	chost, cport, ccred, ccmds := c.GetHost(name)
	if len(chost) == 0 {
		exit("there's no config for " + name)
	}
	cuser, cpass, ckeypath := c.GetCred(ccred)
	startssh(chost, cport, cuser, cpass, ckeypath, ccmds)
}

func startssh(ip, port, user, pass, keypath, cmds string) {
	klog.Infof("connecting to ip: %v...", ip)
	t := ssh.New(ip, user, pass,
		ssh.SetPort(port),
		ssh.SetKeyPath(keypath),
		ssh.SetInitCmds(cmds))
	err := t.Start()
	if err != nil {
		exiterr("start term", err)
	}
	t.Close()
}

func exit(msg string) {
	klog.Errorf(msg)
	os.Exit(1)
}

func exiterr(msg string, err error) {
	klog.Errorf("%v, err: %v", msg, err)
	os.Exit(1)
}

func defaultKeyPath() string {
	return filepath.Join(homedir(), ".ssh/id_rsa")
}
func homedir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}
