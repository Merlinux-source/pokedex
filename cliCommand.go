package main

type cliCommand struct {
	name        string
	description string
	callback    func(conf *CommandConf) error
	conf        *CommandConf
}
