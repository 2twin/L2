package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("customShell> ")

		lineBytes, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
		}

		line := string(lineBytes)

		if line == "" {
			continue
		}

		if line == `\q` {
			break
		}

		pipes := strings.Split(line, "|")

		for _, pipe := range pipes {
			pipe = strings.TrimSpace(pipe)
			err = Exec(pipe)
			if err != nil {
				log.Println(err)
			}
		}

	}
}

func Exec(line string) error {
	split := strings.Split(line, " ")

	cmd := split[0]
	args := split[1:]

	switch cmd {
	case "cd":
		err := ChangeDir(args)
		if err != nil {
			return err
		}

	case "pwd":
		err := PrintWorkDir()
		if err != nil {
			log.Println(err)
		}

	case "echo":
		Echo(args)

	case "ps":
		err := PrintProcs()
		if err != nil {
			log.Println(err)
		}

	case "kill":
		KillProc(args)
	}

	return nil
}

func ChangeDir(args []string) error {
	if len(args) > 1 {
		return errors.New("cd: too many arguments")
	}

	if len(args) == 0 {
		return errors.New("cd: not enough arguments")
	}

	path := args[0]

	return os.Chdir(path)
}

func PrintWorkDir() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return nil
}

func Echo(args []string) {
	str := strings.Join(args, " ")
	fmt.Println(str)
}

func PrintProcs() error {
	procs, err := ps.Processes()
	if err != nil {
		return err
	}

	fmt.Println("PID\tPPID\tExecutable")

	for _, p := range procs {
		fmt.Printf("%d\t%d\t%s\n", p.Pid(), p.PPid(), p.Executable())
	}

	return nil
}

func KillProc(args []string) error {
	if len(args) > 1 {
		return errors.New("kill: too many arguments")
	}

	if len(args) == 0 {
		return errors.New("kill: not enough arguments")
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	
	return proc.Kill()
}
