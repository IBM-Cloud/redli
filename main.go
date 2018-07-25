package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/chzyer/readline"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug         = kingpin.Flag("debug", "Enable debug mode.").Bool()
	longprompt    = kingpin.Flag("long", "Enable long prompt with host/port").Bool()
	redisurl      = kingpin.Arg("url", "URL to connect To.").Required().URL()
	rediscertfile = kingpin.Flag("certfile", "Self-signed certificate file for validation").Envar("REDIS_CERTFILE").File()
	rediscertb64  = kingpin.Flag("certb64", "Self-signed certificate string as base64 for validation").Envar("REDIS_CERTB64").String()
)

var (
	rediscommands = Commands{}
)

func main() {
	kingpin.Parse()

	cert := []byte{}

	if *rediscertfile != nil {
		mycert, err := ioutil.ReadAll(*rediscertfile)
		if err != nil {
			log.Fatal(err)
		}
		cert = mycert
	} else if rediscertb64 != nil {
		mycert, err := base64.StdEncoding.DecodeString((*rediscertb64))
		if err != nil {
			log.Fatal("What", err)
		}
		cert = mycert
	}

	config := &tls.Config{RootCAs: x509.NewCertPool(),
		ClientAuth: tls.RequireAndVerifyClientCert}
	ok := config.RootCAs.AppendCertsFromPEM(cert)
	if !ok {
		log.Fatal("Couldn't load cert data")
	}

	c, err := redis.DialURL((*redisurl).String(), redis.DialTLSConfig(config))
	if err != nil {
		log.Fatal("Dial", err)
	}
	defer c.Close()

	json.Unmarshal([]byte(redisCommandsJSON), &rediscommands)

	completer := readline.NewPrefixCompleter()

	keys := reflect.ValueOf(rediscommands).MapKeys()
	commands := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		commands[i] = keys[i].String()
	}

	sort.Strings(commands)

	//fmt.Println(commands)

	for _, v := range reflect.ValueOf(rediscommands).MapKeys() {
		children := completer.GetChildren()
		children = append(children, readline.PcItem(v.String()))
		completer.SetChildren(children)
	}

	reply, err := redis.String(c.Do("INFO"))
	if err != nil {
		log.Fatal(err)
	}

	info := redisParseInfo(reply)

	fmt.Printf("Connected to %s\n", info["redis_version"])
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       getPrompt(),
		AutoComplete: completer,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		if len(line) == 0 {
			continue // Ignore no input
		}

		parts := strings.Split(line, " ")

		if len(parts) == 0 {
			continue // Ignore no input
		}

		if parts[0] == "help" {
			fmt.Printf("Help coming soon\n")
			continue
		}

		if parts[0] == "exit" {
			break
		}

		var args = make([]interface{}, len(parts[1:]))
		for i, d := range parts[1:] {
			args[i] = d
		}

		result, err := c.Do(parts[0], args...)

		switch v := result.(type) {
		case redis.Error:
			fmt.Printf("%s\n", v.Error())
		case int64:
			fmt.Printf("%d\n", v)
		case string:
			fmt.Printf("%s\n", v)
		case []byte:
			fmt.Printf("%s\n", string(v))
		case nil:
			fmt.Printf("nil\n")
		case []interface{}:
			for i, j := range v {
				fmt.Printf("%d) %s\n", i+1, j)
			}
		}
	}
}

func redisParseInfo(reply string) map[string]string {
	lines := strings.Split(reply, "\r\n")
	values := map[string]string{}
	for _, line := range lines {
		if len(line) > 0 && line[0] != '#' {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				values[parts[0]] = parts[1]
			}
		}
	}
	return values
}

func getPrompt() string {
	if *longprompt {
		return fmt.Sprintf("%s:%s> ", (*redisurl).Hostname(), (*redisurl).Port())
	}

	return "> "
}

func printAsJSON(toprint interface{}) {
	jsonstr, _ := json.MarshalIndent(toprint, "", " ")
	fmt.Println(string(jsonstr))
}

//Commands is a holder for Redis Command structures
type Commands map[string]Command

//Command is a holder for Redis Command data includint arguments
type Command struct {
	Summary    string     `json:"summary"`
	Complexity string     `json:"complexity"`
	Arguments  []Argument `json:"arguments"`
	Since      string     `json:"since"`
	Group      string     `json:"group"`
}

//Argument is a holder for Redis Command Argument data
type Argument struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Enum     string `json:"enum,omitempty"`
	Optional bool   `json:"optional"`
}
