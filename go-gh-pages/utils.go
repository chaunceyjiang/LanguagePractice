package go_gh_pages

import (
	"github.com/labstack/gommon/log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Git struct {
	Cwd    string //  Repository directory
	Cmd    string //exe Git executable (full path if not already on path).
	Err    error
	Output string
}
type U struct {
	Name  string
	Email string
}
type Options struct {
	Remote   string
	Depth    int
	Branch   string
	Dest     string
	Dotfiles bool
	Add      bool
	Repo     string
	Tag      string
	Message  string
	User     U
	Push     bool
	Silent   bool
	Cmd      string
}

func (g *Git) exec(args ...string) {

	b, err := exec.Command(g.Cmd, args...).Output()
	g.Err = err
	if g.Err != nil {
		log.Error(g.Err)
	}
	g.Output = string(b)
}
func (g *Git) Init() {
	g.exec("init")
}
func (g Git) Add(files []string) {
	args := []string{"add"}
	args = append(args, files...)
	g.exec(args...)
}

func (g Git) Clear() {
	g.exec("clean", "-f", "-d")
}

func (g Git) Reset(remote, branch string) {
	g.exec("reset", "--hard", remote+"/"+branch)
}

func (g *Git) Commit(message string) {
	g.exec("diff-index", "--quiet", "HEAD")
	g.exec("commit", "-m", message)
}
func (g *Git) Rm(files []string) {
	args := []string{"rm", "--ignore-unmatch", "-r", "-f"}
	args = append(args, files...)
	g.exec(args...)
}

func (g *Git) Fetch(remote string) {
	g.exec("fetch",remote)
}
func (g Git) Checkout(remote, branch string) {
	treeish := remote + "/" + branch
	g.exec("ls-remote", "--exit-code", ".", treeish)
	if ee, ok := g.Err.(*exec.ExitError); ok {
		if ee.ExitCode() == 2 {
			g.exec("checkout", "--orphan", branch)
		} else {
			g.exec("checkout", branch)
			g.Clear()
			g.Reset(remote, branch)
		}
	} else {
		g.exec("checkout", "--orphan", branch)
	}
}
func (g *Git) Push(remote, branch string) {
	g.exec("push", "--tags", remote, branch)

}

func (g *Git) getRemoteUrl(remote string) string {
	g.exec("config", "--get", "remote."+remote+".url")
	return strings.Split(g.Output, "\n")[0]
}

func (g *Git) Clone(repo, dir, branch string, options Options) *Git {
	gg := NewGit()
	gg.Cwd = dir
	gg.Cmd = options.Cmd
	if gg.Cmd == "" {
		gg.Cmd = "git"
	}

	if exists(dir) {
		return gg
	} else {
		if err := os.MkdirAll(dir, 0777); err != nil {
			log.Fatal(err)
		}
		args := []string{
			"clone",
			repo,
			dir,
			"--branch",
			branch,
			"--single-branch",
			"--origin",
			options.Remote,
			"--depth",
			strconv.Itoa(options.Depth),
		}
		err := exec.Command(options.Cmd, args...).Start()
		if err != nil {
			log.Fatal(err)
		}
		return gg
	}
}

func exists(dirname string) bool {
	_, err := os.Stat(dirname)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func NewGit() *Git {

	g := &Git{
		Cwd: ".",
		Cmd: "git",
	}
	return g
}
