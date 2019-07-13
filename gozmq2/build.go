package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	BUILD_DIR    = "app"
	PROJECT_NAME = "cbms"
)

var (
	parDir    string
	base      string
	bmsBinary string
	bmsTopo   string
	build_dir string
	err       error
	helps     []string
)

type Version struct {
	Version string `json:"version"`
	Base    string `json:"base"`
	Topo    string `json:"topo"`
}

func main() {

	base, err = filepath.Abs(filepath.Dir(os.Args[0]))
	//base = "/home/chauncey/BMS/bmsStats"
	os.Chdir(base)
	if err != nil {
		log.Fatal(err)
	}
	parDir, _ = filepath.Split(base)
	checkrepo()
	build_dir = filepath.Join(base, BUILD_DIR)
	err = os.RemoveAll(build_dir)
	if err != nil {
		log.Fatal(err)
	}

	mkdir("install")
	mkdir(PROJECT_NAME, "bin")
	mkdir(PROJECT_NAME, "etc")
	mkdir(PROJECT_NAME, "var", "log")
	mkdir(PROJECT_NAME, "var", "run")
	mkdir(PROJECT_NAME, "var", "update")

	fmt.Println(color.GreenString("building"), "install script...")
	err = copyDir(filepath.Join(base, "build"), filepath.Join(build_dir, "install"))
	if err != nil {
		log.Error(err)
	}

	fmt.Println(color.GreenString("building"), "licenses file...")
	if err = copyFile(filepath.Join(base, "LICENSES"), filepath.Join(build_dir, "install", "LICENSES")); err != nil {
		log.Error(err)
	}
	if err = copyFile(filepath.Join(base, "LICENSE-CH.txt"), filepath.Join(build_dir, "install", "LICENSE-CH.txt")); err != nil {
		log.Error(err)
	}

	helps = []string{"pktminer", "datacollection", "datatransform", "web"}
	for _, d := range helps {
		fmt.Println(color.GreenString("building"), fmt.Sprintf("%s project...", d))
		if err = copyDir(filepath.Join(base, "src", d), filepath.Join(build_dir, PROJECT_NAME, d)); err != nil {
			log.Error(err)
		}
	}

	if err = copyDir(filepath.Join(bmsTopo, "static"), filepath.Join(build_dir, PROJECT_NAME, "web", "static")); err != nil {
		log.Error(err)
	}

	files, _ := filepath.Glob(filepath.Join(bmsTopo, "*.html"))
	for _, f := range files {
		_, filename := filepath.Split(f)
		if err = copyFile(f, filepath.Join(build_dir, PROJECT_NAME, "web", "templates", filename)); err != nil {
			log.Error(err)
		}
	}

	helps = []string{"utility", "utils", "datastore", "topo", "distcall", "jobber", "variable", "dsapi"}
	for _, d := range helps {
		if err = copyDir(filepath.Join(base, "src", "web", d), filepath.Join(build_dir, PROJECT_NAME, "datatransform", d)); err != nil {
			log.Error(err)
		}
	}
	if err = copyFile(filepath.Join(base, "src", "web", "settings.py"), filepath.Join(build_dir, PROJECT_NAME, "datatransform", "settings.py")); err != nil {
		log.Error(err)
	}
	if err = copyDir(filepath.Join(base, "bin"), filepath.Join(build_dir, PROJECT_NAME, "bin")); err != nil {
		log.Error(err)
	}

	fmt.Println(color.GreenString("building"), "rpm package...")
	dirs, _ := filepath.Glob(filepath.Join(bmsBinary, "install", "centos7", "*"))
	for _, d := range dirs {
		_, filename := filepath.Split(d)
		if err = copyDir(d, filepath.Join(build_dir, PROJECT_NAME, filename)); err != nil {
			log.Error(err)
		}
	}
	fmt.Println(color.GreenString("building"), "config file...")

	dirs, _ = filepath.Glob(filepath.Join(base, "config", "*"))
	for _, d := range dirs {
		_, filename := filepath.Split(d)
		if err = copyDir(d, filepath.Join(build_dir, PROJECT_NAME, "etc", filename)); err != nil {
			log.Error(err)
		}

	}
	helps = []string{"bin", "var"}
	for _, d := range helps {
		if err = copyDir(filepath.Join(base, d), filepath.Join(build_dir, PROJECT_NAME, d)); err != nil {
			log.Error(err)
		}
	}

	buildGenVersion()

	fmt.Println(color.GreenString("compressing"), "web project...")
	helps = []string{
		"web",
		"rpm",
		"etc",
		"pypy"}
	var opts []string
	for _, o := range helps {
		opts = append(opts, filepath.Join(build_dir, PROJECT_NAME, o))
	}
	compressTarGz("web.tar.gz", opts)

	fmt.Println(color.GreenString("compressing"), "dataprovider project...")
	helps = []string{
		"datacollection",
		"rpm",
		"jzmq",
		"etc",
		"pktminer",
		"pypy"}
	opts = []string{}
	for _, o := range helps {
		opts = append(opts, filepath.Join(build_dir, PROJECT_NAME, o))
	}
	compressTarGz("dataprovider.tar.gz", opts)

	fmt.Println(color.GreenString("compressing"), "mongoDB project...")

	opts = []string{
		filepath.Join(build_dir, PROJECT_NAME, "mongodb"),
		filepath.Join(build_dir, PROJECT_NAME, "etc", "service", "mongodb.service"),
		filepath.Join(build_dir, PROJECT_NAME, "etc", "mongodb"),
		filepath.Join(build_dir, PROJECT_NAME, "rpm", "dialog-1.2-4.20130523.el7.x86_64.rpm"),
	}
	compressTarGz("mongodb.tar.gz", opts)

	fmt.Println(color.GreenString("compressing"), "transform project...")
	helps = []string{
		"datatransform",
		"rpm",
		"etc",
		"pypy"}
	opts = []string{}
	for _, o := range helps {
		opts = append(opts, filepath.Join(build_dir, PROJECT_NAME, o))
	}
	compressTarGz("datatransform.tar.gz", opts)
}
func compressTarGz(filename string, opts []string) {
	opts = append(opts,
		filepath.Join(build_dir, "install"),
		filepath.Join(build_dir, PROJECT_NAME, "bin"),
		filepath.Join(build_dir, PROJECT_NAME, "var"),
	)
	if err = TarGz(filepath.Join(build_dir, filename), filepath.Join(build_dir), opts); err != nil {
		log.Error(err)
	}
}
func buildGenVersion() {
	versionFile, err := os.Open(filepath.Join(base, "var", "version"))
	defer versionFile.Close()
	if err != nil {
		log.Error(err)
	}

	jsonDate, _ := ioutil.ReadAll(versionFile)
	var v Version
	if err = json.Unmarshal(jsonDate, &v); err != nil {
		log.Error(err)
	}

	v.Base = gitRevList(base)
	v.Topo = gitRevList(bmsTopo)
	jsonDate, _ = json.Marshal(v)
	varVersionFile, err := os.OpenFile(filepath.Join(build_dir, PROJECT_NAME, "var", "version"), os.O_RDWR, 0666)
	if err != nil {
		log.Error(err)
	}
	varVersionFile.Write(jsonDate)

	defer varVersionFile.Close()
}

func gitRevList(gitPath string) string {
	os.Chdir(gitPath)
	cmd1 := exec.Command("git", "rev-list", "HEAD")
	var out bytes.Buffer
	cmd1.Stdout = &out

	if err != nil {
		log.Error(err)
	}
	cmd1.Run()
	head := make([]byte, 40)
	out.Read(head)
	os.Chdir(base)
	return string(head)

}
func mkdir(elem ...string) {
	elems := []string{base, BUILD_DIR}
	elems = append(elems, elem...)
	var err error
	err = os.MkdirAll(filepath.Join(elems...), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func checkrepo() {
	bmsBinary = filepath.Join(parDir, "bmsbinary")
	bmsTopo = filepath.Join(parDir, "bmsTopo")
	fmt.Println()

	fmt.Println("Backend repository  : ", base)
	fmt.Println("Binary repository   : ", bmsBinary)
	fmt.Println("Front-end repository: ", bmsTopo)
	fmt.Printf("Check Repository ")
	exist(base, bmsBinary, bmsTopo)
	color.Green("Success...")
	fmt.Println()
}
func exist(filepaths ...string) {
	for _, p := range filepaths {
		_, err = os.Stat(p)
		if err != nil {
			if os.IsNotExist(err) {
				color.Red("%s not exist!", p)
				os.Exit(-1)
			}
		}
	}

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

func TarGz(dst string, src string, opts []string) (err error) {
	// 创建文件
	fw, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fw.Close()

	// 将 tar 包使用 gzip 压缩，其实添加压缩功能很简单，
	// 只需要在 fw 和 tw 之前加上一层压缩就行了，和 Linux 的管道的感觉类似
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 创建 Tar.Writer 结构
	tw := tar.NewWriter(gw)
	// 如果需要启用 gzip 将上面代码注释，换成下面的

	defer tw.Close()

	// 下面就该开始处理数据了，这里的思路就是递归处理目录及目录下的所有文件和目录
	// 这里可以自己写个递归来处理，不过 Golang 提供了 filepath.Walk 函数，可以很方便的做这个事情
	// 直接将这个函数的处理结果返回就行，需要传给它一个源文件或目录，它就可以自己去处理
	// 我们就只需要去实现我们自己的 打包逻辑即可，不需要再去路径相关的事情
	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		// 因为这个闭包会返回个 error ，所以先要处理一下这个
		if err != nil {
			return err
		}

		// 这里就不需要我们自己再 os.Stat 了，它已经做好了，我们直接使用 fi 即可
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		// 这里需要处理下 hdr 中的 Name，因为默认文件的名字是不带路径的，
		// 打包之后所有文件就会堆在一起，这样就破坏了原本的目录结果
		// 例如： 将原本 hdr.Name 的 syslog 替换程 log/syslog
		// 这个其实也很简单，回调函数的 fileName 字段给我们返回来的就是完整路径的 log/syslog
		// strings.TrimPrefix 将 fileName 的最左侧的 / 去掉，
		// 熟悉 Linux 的都知道为什么要去掉这个
		for _, opt := range opts {
			if strings.Contains(fileName, opt) {
				hdr.Name = strings.TrimPrefix(fileName, build_dir)

				// 写入文件信息
				if err := tw.WriteHeader(hdr); err != nil {
					return err
				}

				// 判断下文件是否是标准文件，如果不是就不处理了，
				// 如： 目录，这里就只记录了文件信息，不会执行下面的 copy
				if !fi.Mode().IsRegular() {
					return nil
				}

				// 打开文件
				fr, err := os.Open(fileName)
				defer fr.Close()
				if err != nil {
					return err
				}
				log.Debug(fileName)
				// copy 文件数据到 tw
				_, err = io.Copy(tw, fr)
				if err != nil {
					return err
				}

				return nil
			}

		}

		return nil
	})
}
