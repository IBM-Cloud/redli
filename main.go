package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug      = kingpin.Flag("debug", "Enable debug mode.").Bool()
	longprompt = kingpin.Flag("long", "Enable long prompt with host/port").Bool()
	redisurl   = kingpin.Arg("url", "URL to connect To.").Required().URL()
)

func main() {
	kingpin.Parse()
	c, err := redis.DialURL((*redisurl).String())
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	rediscommands := Commands{}

	// redisCommandString, err := ioutil.ReadFile("./commands.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	json.Unmarshal([]byte(redisCommandsJSON), &rediscommands)

	reply, err := redis.String(c.Do("INFO"))
	if err != nil {
		log.Fatal(err)
	}

	info := redisParseInfo(reply)

	fmt.Printf("Connected to %s\n", info["redis_version"])
	rl, err := readline.New(getPrompt())

	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		// ToDo: parse the line

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
			fmt.Printf("TBD: no value\n")
		case []interface{}:
			fmt.Printf("TBD: array\n")
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

func PrintAsJSON(toprint interface{}) {
	jsonstr, _ := json.MarshalIndent(toprint, "", " ")
	fmt.Println(string(jsonstr))
}

type Commands map[string]Command

type Command struct {
	Summary    string     `json:"summary"`
	Complexity string     `json:"complexity"`
	Arguments  []Argument `json:"arguments"`
	Since      string     `json:"since"`
	Group      string     `json:"group"`
}

type Argument struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Enum     string `json:"enum,omitempty"`
	Optional bool   `json:"optional"`
}
