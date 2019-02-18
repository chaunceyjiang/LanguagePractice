package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func Pipeline(cmds ...*exec.Cmd) ([]byte, []byte, error) {
	if len(cmds) < 1 {
		return nil, nil, nil
	}
	var output bytes.Buffer
	var stderr bytes.Buffer
	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		var err error
		if cmds[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return nil, nil, err
		}
		cmd.Stderr = &stderr
	}
	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr
	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return output.Bytes(), stderr.Bytes(), err
		}
	}
	return output.Bytes(), stderr.Bytes(), nil
}
func main() {
	// Collect directories from the command-line
	var dirs []string
	if len(os.Args) > 1 {
		dirs = os.Args[1:]
	} else {
		dirs = []string{"."}
	}

	// Run the command on each directory
	for _, dir := range dirs {
		// find $DIR -type f # Find all files
		ls := exec.Command("find", dir, "-type", "f")

		// | grep -v '/[._]' # Ignore hidden/temporary files
		visible := exec.Command("egrep", "-v", `/[._]`)

		// | sort -t. -k2 # Sort by file extension
		sort := exec.Command("sort", "-t.", "-k2")

		// Run the pipeline
		output, stderr, err := Pipeline(ls, visible, sort)
		if err != nil {
			log.Printf("dir %q: %s", dir, err.Error())
		}

		// Print the stdout, if any
		if len(output) > 0 {
			log.Printf("%q:\n%s", dir, output)
		}

		// Print the stderr, if any
		if len(stderr) > 0 {
			log.Printf("%q: (stderr)\n%s", dir, stderr)
		}
	}
}