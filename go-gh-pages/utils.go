package go_gh_pages

import (
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
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
	Src      string //The minimatch pattern or array of patterns used to select which files should be published.
}

func (g *Git) exec(args ...string) {
	if err := os.Chdir(g.Cwd); err != nil {
		log.Warn(err)
	}
	b, err := exec.Command(g.Cmd, args...).CombinedOutput()
	g.Err = err
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
	//g.exec("diff-index", "--quiet", "HEAD")
	g.exec("commit", "-m", message)
}
func (g *Git) Rm(files []string) {
	args := []string{"rm", "--ignore-unmatch", "-r", "-f"}
	args = append(args, files...)
	g.exec(args...)
}

func (g *Git) Fetch(remote string) {
	g.exec("-c", " http.sslVerify=false", "fetch", remote)
}
func (g Git) Checkout(remote, branch string) {

	g.exec("checkout", "--orphan", branch)

	if g.Err != nil {
		log.Warn(g.Err)
	}
}
func (g *Git) Push(remote, branch string) {
	g.DeleteReomoteBranch(remote, branch)
	g.exec("-c", "http.sslVerify=false", "push", "--tags", remote, branch)
	if g.Err != nil {
		log.Warn(g.Err)
	}
}

func (g *Git) getRemoteUrl(remote string) string {
	g.exec("config", "--get", "remote."+remote+".url")
	return strings.Split(g.Output, "\n")[0]
}

func (g Git) Clone(repo, dir, branch string, options *Options) *Git {
	gg := NewGit()
	gg.Cwd = dir
	gg.Cmd = options.Cmd
	if gg.Cmd == "" {
		gg.Cmd = "git"
	}

	if exists(dir) {
		return gg
	} else {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal(err)
		}
		args := []string{
			"-c",
			"http.sslVerify=false",
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
		b, err := exec.Command(options.Cmd, args...).CombinedOutput()
		if err != nil {
			log.Fatal(string(b), err)
		}
		return gg
	}
}
func (g *Git)YesterdayCommit() string {
	g.exec("log","--after=yesterday")
	return g.Output
}
func (g Git) LastCommit() string {
	g.exec("rev-parse", "HEAD")
	return g.Output
}
func (g *Git) Pull(remote, branch string) {
	g.exec("-c", "http.sslVerify=false", "pull", remote, branch+":"+branch)
}
func (g *Git) SetUser(u *U) {
	g.exec("config", "user.email", u.Email)
	g.exec("config", "user.name", u.Name)
}
func (g *Git) DeleteReomoteBranch(remote, branch string) {
	g.exec("-c", "http.sslVerify=false", "push", remote, "--delete", branch)
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

func getUser() *U {
	name, err := exec.Command("git", "config", "user.name").CombinedOutput()
	if err != nil {
		log.Error(err)
	}
	email, err := exec.Command("git", "config", "user.email").CombinedOutput()
	return &U{
		Name:  string(name),
		Email: string(email),
	}
}

func getUrlProjectName(repoUrl string) string {
	rest := strings.Split(repoUrl, "/")
	result := rest[len(rest)-1]
	return strings.Split(result, ".")[0]
}

func copyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	if !srcinfo.IsDir() {
		if err = copyFile(src, dst); err != nil {
			return err
		}
		return nil
	}
	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				log.Error(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				log.Error(err)
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

type T struct {
	Content string `json:"content"`
}
type Msg struct {
	Msgtype string `json:"msgtype"`
	Text    T      `json:"text"`
}
