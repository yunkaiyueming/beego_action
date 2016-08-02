package controllers

import (
	"fmt"
	"os"
	"runtime"
)

type PerformanceController struct {
	BaseController
}

func (this *PerformanceController) Construct() {
	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/classic.html"
	this.sidebarFile = "include/sidebar/classic_sidebar.html"
}

func (this *PerformanceController) MyRender(viewFile string) {
	this.Construct()
	this.BaseController.MyRender(viewFile)
}

func (this *PerformanceController) Index() {
	fmt.Println(os.Hostname())

	if name, ok := os.Hostname(); ok == nil {
		this.Data["hostName"] = name
	}

	this.Data["numCpu"] = runtime.NumCPU()

	this.Data["cpuProfile"] = runtime.CPUProfile()
	fmt.Println(this.Data["cpuProfile"])

	this.Data["envs"] = os.Environ()
	this.MyRender("perfm/view_index.html")
}

func (this *PerformanceController) GetGoEnv() {
	this.Data["goVersion"] = runtime.Version()
	this.Data["goRoot"] = runtime.GOROOT()
	this.Data["goArch"] = runtime.GOARCH
	this.Data["goMaxProcs"] = runtime.GOMAXPROCS(-1)
	this.Data["goPath"] = os.Getenv("GoPath")
	fmt.Println("--get go env success--")

	this.MyRender("perfm/view_goenv.html")
}
