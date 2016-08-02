package controllers

import (
	"bufio"
	_ "flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type FileController struct {
	BaseController
}

func (this *FileController) Construct() {
	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/classic.html"
	this.sidebarFile = "include/sidebar/classic_sidebar.html"
}

func (this *FileController) MyRender(viewFile string) {
	this.Construct()
	this.BaseController.MyRender(viewFile)
}

func (this *FileController) GetFileCon() {
	file := "E:/GO_PATH/src/beego_action/beego_action.exe"

	start := time.Now()
	ReadByOsRead(file)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n", t1.Sub(start))

	ReadByBufio(file)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n", t2.Sub(t1))

	ReadByIoutil(file)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n", t3.Sub(t2))

	this.Data["fileInfo"] = ReadByIoutil("E:/GO_PATH/src/beego_action/conf/app.conf")
	this.Data["t1"] = fmt.Sprintf("Cost time %v\n", t1.Sub(start))
	this.Data["t2"] = fmt.Sprintf("Cost time %v\n", t2.Sub(t1))
	this.Data["t3"] = fmt.Sprintf("Cost time %v\n", t3.Sub(t2))
	this.MyRender("file/index.html")
}

func (this *FileController) ListDir() {
	dirPath := "E:/GO_PATH/src/beego_action"
	rFile, _ := os.Open(dirPath)
	fileInfos, _ := rFile.Readdir(-1)

	this.Data["fileInfos"] = fileInfos
	this.MyRender("file/view_list_dir.html")
}

func ReadByOsRead(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func ReadByBufio(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

func ReadByIoutil(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}
