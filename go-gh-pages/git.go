package go_gh_pages

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func getCacheDir() string {
	return "/tmp/.cache"

}

var _defaultOpts = Options{
	Dest:     "dist",
	Add:      false,
	Cmd:      "git",
	Depth:    1,
	Dotfiles: false,
	Branch:   "bmsV5.1",
	Remote:   "origin",
	Src:      "**/*",
	Push:     true,
	Message:  "Updates",
	Silent:   false,
}

func getRepo(opt *Options) string {
	if opt.Repo != "" {
		return opt.Repo
	}
	g := NewGit()
	return g.getRemoteUrl(opt.Remote)
}

func Publish(basePath string, opt *Options, callback func(err error)) {
	opts := opt
	if opts == nil {
		opts = &_defaultOpts
	}
	curRepo := NewGit()

	// FIXME 下次再优化这里
	curRepo.Pull(opts.Remote, opts.Branch)
	_ = exec.Command("npm", "run", "build").Run()



	if !exists(basePath) {
		callback(errors.New(`'The "basePath" option must be an existing directory'`))
		return
	}
	var repoUrl string
	var u *U
	if opts.User.Name == "" && opts.User.Email == "" {
		u = getUser()
	}
	repoUrl = getRepo(opts)
	projectName := getUrlProjectName(repoUrl)
	newrepoUrl := filepath.Join(getCacheDir(), projectName)
	g := Git{}.Clone(repoUrl, newrepoUrl, opts.Branch, opts)
	clonerepourl := g.getRemoteUrl(opts.Remote)
	if clonerepourl != repoUrl {
		log.Fatal(errors.New("Remote url mismatch .Got [" + clonerepourl + "]" + "but expected " + repoUrl + "in " + g.Cwd))
	}
	log.Info("Cleaning")
	msg := g.LastCommit()
	log.Info("Fetch " + opts.Remote)
	g.Fetch(opts.Remote)
	g.Checkout(opts.Remote, "releaseV5.1")
	g.Clear()
	if !opts.Add {
		log.Info("Remoing files")
		g.Rm([]string{"."})
	}
	if u != nil {
		log.Debug(u.Name)
		log.Debug(u.Email)
	}
	log.Info("Copying files")
	if err := copyDir(filepath.Join(basePath, opts.Dest), newrepoUrl); err != nil {
		log.Error(err)
	}
	g.Add([]string{"."})
	g.SetUser(u)
	g.Commit("release for commit:" + msg)

	g.Push(opts.Remote, "releaseV5.1")
	tokenurl := `https://oapi.dingtalk.com/robot/send?access_token=13d97d8f8d16f930f3305080594c5d6b7f73fea2989ed2b474164382e260caa45`
	client := &http.Client{}
	body := Msg{"text", T{"release: "+msg}}

	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(body)
	req, err := http.NewRequest("POST", tokenurl, buf)

	if err != nil {
		log.Info(err)
	}

	req.Header.Set("Content-Type", "application/json")

	res, e := client.Do(req)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()
	_, _ = io.Copy(os.Stdout, res.Body)
}
