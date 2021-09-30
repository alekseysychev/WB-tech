package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
)

const (
	delim      = "|"
	fork       = "&"
	colorGreen = "\033[32m"
	colorWhite = "\033[37m"
	r          = true
)

var (
	writer io.Writer = os.Stdout
	reader io.Reader = os.Stdin
	isUdp  bool
)

func init() {
	flag.BoolVar(&isUdp, "u", false, "Use UDP instead of the default option of TCP.")
}

func main() {
	mainminic()
}

func mainminic() {
	// родительский процесс
	for {

		reader := bufio.NewReader(os.Stdin)
		if !r {
			hat()
		}
		line, _ := reader.ReadString('\n')

		if line == "exit\n" {
			break
		}
		if line == "\n" {
			continue
		}
		if r {

			subdata := strings.Split(line[:len(line)-1], ":")
			netc(subdata[0], subdata[1])
			return
		} else {

			mimicshell(line[:len(line)-1])
		}

	}

}

func mimicshell(line string) {
	subdata := strings.Split(line, fork)
	for i := range subdata {
		ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if err != 0 {
			os.Exit(2)
		}
		if ret == 0 && len(subdata) == 1 {
			// потомок, умирает если & не имеются
			os.Exit(0)
		}

		if ret == 0 || len(subdata) == 1 {
			// потомок, если & имеются
			data := strings.Split(subdata[i], delim)
			rez := " "
			for j := range data {
				var err error

				rez, err = shell(data[j] + rez)
				if err != nil {
					println(err)
				}
			}
			if len(rez) > 1 {
				fmt.Println(rez[1:])
			}
			if ret == 0 {
				os.Exit(0)
			}
		}
	}
}

func shell(command string) (rez string, err error) {

	args := strings.Fields(command)
	if len(args) == 0 {
		return "", nil
	}
	switch args[0] {
	case "pwd":
		rez, err = os.Getwd()
	case "cd":
		err = os.Chdir(args[1])
	case "echo":
		rez = args[len(args)-1]
	case "ps":
		var A, o bool
		psargc := make([]string, 0)
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "-A":
				A = true
			case "-o":
				o = true
				i++
				buf := strings.Split(args[i], ",")
				psargc = append(psargc, buf...)
			}
		}
		var procs []ps.Process
		if !o {
			rez = "PID\tCMD\n"
		}
		procs, err = ps.Processes()
		for _, p := range procs {
			if os.Getppid() == p.Pid() || A {
				if !o {
					rez += fmt.Sprintf("%d\t%s\n", p.Pid(), p.Executable())
				} else {
					for j := range psargc {
						switch psargc[j] {
						case "pid":
							rez += strconv.Itoa(p.Pid())
						case "CMD":
							rez += p.Executable()
						default:
						}
					}
				}
			}
		}
	case "kill":
		for i := range args {
			if pid, _ := strconv.Atoi(args[i]); i > 0 {
				syscall.Kill(pid, syscall.SIGINT)
			} else {
				println(err)
			}
		}
	case "exec":
		rez, err = exec(args[1])
	}
	return " " + rez, err
}

func exec(path string) (string, error) {
	rez := ""
	data, err := ioutil.ReadFile(path)
	if err != nil {
		mimicshell(path)
	} else {
		comm := strings.Split(string(data), "\n")

		for i := 0; i < len(comm); i++ {
			if comm[i] == "" {
				continue
			}
			mimicshell(comm[i])
		}
	}
	return rez, nil
}

func hat() {
	wd, _ := os.Getwd()
	args := strings.Split(wd, "/")
	fmt.Print(string(colorGreen), "[minicshell@minicshell ", string(colorWhite), args[len(args)-1], string(colorGreen), "]$ ", string(colorWhite))
}
