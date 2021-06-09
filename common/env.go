package common

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "cbm-ocs-listener/utl/str"
)

// process environment
var (
	fversion                bool = false
	fconfig                 string
	frunpath                string
	fdebug                  string
	foracledbuser           string
	foracledbpassword       string
	foracleservicename      string
	fqueueserverurl         string
	fqueuename              string
	fqueueuser              string
	fqueuepassword          string
	fcbmserverurl           string
	fsystemresponderurl     string
	fsend                   string
	ffilesend               string
	fparse                  string
	ffileparse              string
)

// get invocation flags
func FlagsInit() {
	flag.BoolVar(&fversion, "v", false, "Version check")
	flag.StringVar(&fconfig, "config", "config.json", "Config json file if not in $RUNPATH/config.json")
	flag.StringVar(&frunpath, "runpath", ".", "Run path")
	flag.StringVar(&fdebug, "debug", "0", "Debug level")
	flag.StringVar(&foracledbuser, "oracledbuser", "", "Oracle DB user")
	flag.StringVar(&foracledbpassword, "oracledbpassword", "", "Oracle DB password")
	flag.StringVar(&foracleservicename, "oracleservicename", "", "Oracle service name")
	flag.StringVar(&fqueueserverurl, "queueserverurl", "", "EMS queue server URL address")
	flag.StringVar(&fqueuename, "queuename", "", "EMS queue name")
	flag.StringVar(&fqueueuser, "queueuser", "", "EMS queue user")
	flag.StringVar(&fqueuepassword, "queuepassword", "", "EMS queue password")
	flag.StringVar(&fcbmserverurl, "cbmserverurl", "", "Target CBM server URL address for action triggering")
	flag.StringVar(&fsystemresponderurl, "csystemrespoinderurl", "", "Self URL address for health check")
	flag.StringVar(&fsend, "send", "", "Send the content")
	flag.StringVar(&ffilesend, "filesend", "", "Send the content of the file")
	flag.StringVar(&fparse, "parse", "", "Parse json provided")
	flag.StringVar(&fparse, "fileparse", "", "Parse json provided in the file")
}

// load env variables if they are set otherwise use default values or config file
func EnvInit(version, build, level string) {
	flag.Parse()

	// shortcut, anyway we would print it out
	setVersion(version, build, level)
	if fversion {
		fmt.Printf("The version is: %s\n", GetVersion())
		os.Exit(0)
	}
		
	config := os.Getenv("CONFIG")
	if Empty(config) {
		config = fconfig
	}
	
	InitConfig(config)
	
	//override loaded config values with env variables or command line switches
	AppConfig.RunPath = Nvl(Nvl(os.Getenv("RUNPATH"), frunpath), AppConfig.RunPath)
	AppConfig.Debug = Nvl(Nvl(os.Getenv("DEBUG"), fdebug), AppConfig.Debug)
	AppConfig.OracleDBUser = Nvl(Nvl(os.Getenv("ORACLEDBUSER"), foracledbuser), AppConfig.OracleDBUser)
	AppConfig.OracleDBPassword = Nvl(Nvl(os.Getenv("ORACLEDBPASSWORD"), foracledbpassword), AppConfig.OracleDBPassword)
	AppConfig.OracleServiceName = Nvl(Nvl(os.Getenv("ORACLESERVICENAME"), foracleservicename), AppConfig.OracleServiceName)
	AppConfig.QueueServerUrl = Nvl(Nvl(os.Getenv("QUEUESERVERURL"), fqueueserverurl), AppConfig.QueueServerUrl)
	AppConfig.QueueName = Nvl(Nvl(os.Getenv("QUEUENAME"), fqueuename), AppConfig.QueueName)
	AppConfig.QueueUser = Nvl(Nvl(os.Getenv("QUEUEUSER"), fqueueuser), AppConfig.QueueUser)
	AppConfig.QueuePassword = Nvl(Nvl(os.Getenv("QUEUEPASSWORD"), fqueuepassword), AppConfig.QueuePassword)
	AppConfig.CbmServerUrl = Nvl(Nvl(os.Getenv("CBMSERVERURL"), fcbmserverurl), AppConfig.CbmServerUrl)
	AppConfig.SystemResponderUrl = Nvl(Nvl(os.Getenv("SYSTEMRESPONDERURL"), fsystemresponderurl), AppConfig.SystemResponderUrl)
	AppConfig.Send = Nvl(Nvl(os.Getenv("SEND"), fsend), AppConfig.Send)
	AppConfig.FileSend = Nvl(Nvl(os.Getenv("FILESEND"), ffilesend), AppConfig.FileSend)
	AppConfig.Parse = Nvl(Nvl(os.Getenv("PARSE"), fparse), AppConfig.Parse)
	AppConfig.FileParse = Nvl(Nvl(os.Getenv("FILEPARSE"), ffileparse), AppConfig.FileParse)

	if AppConfig.FileSend != "" {
		b, err := ioutil.ReadFile(AppConfig.FileSend)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		} else {
			AppConfig.Send = string(b)
		}
    }

	if AppConfig.FileParse != "" {
		b, err := ioutil.ReadFile(AppConfig.FileParse)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		} else {
			AppConfig.Parse = string(b)
		}
	}
	
	EnvLog()
}

// show in log working environment
func EnvLog() {
	log.Printf("CBM OCS listener execution environment")
	log.Printf("%s: %s", "RunPath               ", AppConfig.RunPath)
	log.Printf("%s: %s", "Debug                 ", AppConfig.Debug)
	log.Printf("%s: %s", "OracleDBUser          ", AppConfig.OracleDBUser)
	log.Printf("%s: %s", "OracleDBPassword      ", AppConfig.OracleDBPassword)
	log.Printf("%s: %s", "OracleServiceName     ", AppConfig.OracleServiceName)
	log.Printf("%s: %s", "QueueServerUrl        ", AppConfig.QueueServerUrl)
	log.Printf("%s: %s", "QueueUser             ", AppConfig.QueueUser)
	log.Printf("%s: %s", "QueuePassword         ", AppConfig.QueuePassword)
	log.Printf("%s: %s", "CbmServerUrl          ", AppConfig.CbmServerUrl)
	log.Printf("%s: %s", "SystemResponderUrl    ", AppConfig.SystemResponderUrl)
}
