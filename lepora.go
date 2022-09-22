package lepora

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Augani/lepora/database"
	"github.com/Augani/lepora/network"
	"github.com/Augani/lepora/util"
	"gorm.io/gorm"
)

type arguments map[string]interface{}

type LeporaMode string

var (
	Managed LeporaMode = "managed"
	Self	LeporaMode = "self"
	Matrix	LeporaMode = "matrix"
	Local	LeporaMode = "local"
)

type LeporaOptions struct {
	//the method of logging, local, managed, self-managed or matrix which is a combination of local and managed logging
	Method LeporaMode
	//the path to the log file
	Path string
	//the name of the log file
	Name string
	//the maximum size of the log file
	MaxSize int
	//the maximum number of log files
	MaxFiles int
	//the maximum number of days to keep the log files
	MaxDays int
	//debug mode on or off
	Debug bool
	//the level of logging
	Level string
	//should logs be persisted
	Persist *bool
	//AutoClean log files
	AutoClean *bool
}

type logger struct {
	db *gorm.DB
	options LeporaOptions
}

type Logger interface {
	Log(args ...interface{})
}

func init() {
	fmt.Println("Lepora initializing.....")
}

func Setup(options LeporaOptions) (Logger, error) {
	var lpDb *gorm.DB
	var err error
	if util.BoolValue(options.Persist) {
		//get db connection by initializing db
		//get postgres db options from Environment
		dbOptions := database.DatabaseOptions{
			Host:         os.Getenv("LEPORA_DB_HOST"),
			Port:         util.StringToInt(os.Getenv("LEPORA_DB_PORT")),
			UserName:     os.Getenv("LEPORA_DB_USER"),
			Password:     os.Getenv("LEPORA_DB_PASSWORD"),
			DatabaseName: os.Getenv("LEPORA_DB_NAME"),
			SSLMode:      os.Getenv("LEPORA_DB_SSL_MODE"),
		}
		if lpDb, err = database.InitDatabase(dbOptions); err != nil {
			return nil, err
		}
	}

	//create a temporary logs folder incase a path is not provided
	if options.Path == "" {
		//get the current working directory
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		//check if folder exists if not then create it
		if _, err := os.Stat(filepath.Join(wd, "logs")); os.IsNotExist(err) {
			if err := os.Mkdir(filepath.Join(wd, "logs"), 0755); err != nil {
				return nil, err
			}
		}
		options.Path = filepath.Join(wd, "logs")
	}

	//use name of application incase name is not provided in options
	if options.Name == "" {
		//get the name of the application
		_, filename, _, ok := runtime.Caller(1)
		if !ok {
			return nil, fmt.Errorf("could not get the name of the application")
		}
		options.Name = filepath.Base(filename)
	}

	//set default max number of days for logs to 3 days if its not provided
	if options.MaxDays == 0 {
		options.MaxDays = 3
	}

	//set default max size of files if max size is not provided
	if options.MaxSize == 0 {
		options.MaxSize = 1024
	}
	newLogger := &logger{
		db: lpDb,
		options: options,
	}

	//set Autoclean to true if autoclean is not provided
	if util.BoolValue(options.AutoClean) {
		//check for old files and delete them
		//get all files in the logs folder
		ticker := time.NewTicker(1 * time.Minute)
		go func() {
			for range ticker.C {
				files, err := ioutil.ReadDir(options.Path)
				if err != nil {
					fmt.Println(err)
				}
				for _, f := range files {
					//check if file is older than 3 days
					if time.Since(f.ModTime()).Hours() > float64(options.MaxDays * 24) {
						//delete file
						if err := os.Remove(filepath.Join(options.Path, f.Name())); err != nil {
							fmt.Println(err)
						}
					}
				}
			}
		}()
	}
	newLogger.options = options
	
	return newLogger, nil
}

func (logger *logger) Log(args ...interface{}) {
	opt := logger.options
	//get last created file if any from path
	lastFile, err := logger.getLastCreatedFile(opt.Path)
	fmt.Println(lastFile, "hhhhhhhhhh")
	if err != nil {
		fmt.Println(err, "error")
	}
	var filePath string
	var logFile *os.File
	if lastFile == ""{
		//create a new file
		logFile, err = os.Create(filepath.Join(opt.Path, fmt.Sprintf("%s__%s.log", opt.Name, time.Now().Format("2006-01-02_15:04:05"))))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(logFile.Name(), "file name")
		filePath = logFile.Name()
	}
	if filePath == ""{
		filePath = opt.Path + "/" + lastFile
	}
	fmt.Println(filePath, "file path")
	//get key value pairs from args passed
	argArray := parseArguments(args...)
	fmt.Println(argArray)
	//check if file exists at path, else create file at path
	//check if file size is greater than max size, else create new file
	//write to file
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//check file size to make sure it is less than max size or will be equal to max size after writing
	//if file size is greater than max size, create new file
	//write to file
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	//get the caller and line of the function when it was called
	_, fileCall, line, ok := runtime.Caller(1)
	if !ok {
		fileCall = "???"
		line = 0
	}

	var log string
	for _, arg := range argArray {
		for key, value := range arg {
			//if on the last argument don't append with comma
			//get key of argument of map
			if arg[key] == argArray[len(argArray)-1][key] {
				log = fmt.Sprintf("%s %s:%v", log, key, value)
			} else {
				log = fmt.Sprintf("%s %s:%v,", log, key, value)
			}
		}
	}

	logPrefix := fmt.Sprintf("Info logged from %s on line %v, at %s", filepath.Base(fileCall), line, time.Now().Format("2006-01-02 15:04:05"))
	log = fmt.Sprintf("%s||%s", logPrefix, log)

	//check size of file + the size of new string to be written to see if its more than the max size allowed
	//if it is, create new file

	//if Method is managed, send logs to Server using network package
	//if Method is self-managed, send logs to Server using network package
	//if Method is matrix, send logs to Server using network package and also write to local file
	//if Method is local, write to local file
	//if Method is not specified, write to local file
	if opt.Method != Local {
		switch opt.Method {
		case Managed:
			//send logs to server
			//sendLogs(log, network.NetworkOptions{})
			//get network properties from env
			url := os.Getenv("LEPORA_URL")
			token := os.Getenv("LEPORA_TOKEN")
			networkOptions := network.NetworkOptions{
				URL: url,
				AccessToken: token,
			}
			go logger.processLogs(networkOptions, log)
		case Self:
			//send logs to server
			//sendLogs(log, network.NetworkOptions{})
			break
		case Matrix:
			//send logs to server
			//sendLogs(log, network.NetworkOptions{})
			//write to local file
		}
	}

	if stat.Size()+int64(len(log)) > int64(opt.MaxSize) {
		//create new file in given path with today's date and timestamp
		//write to file
		fmt.Println("file size is greater than max size")
		newFile, err := os.Create(filepath.Join(opt.Path, fmt.Sprintf("%s__%s", opt.Name, time.Now().Format("2006-01-02_15:04:05"))))
		if err != nil {
			fmt.Println(err)
		}
		defer newFile.Close()
		//write to file
		// _, err = newFile.WriteString(log)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		file = newFile
		
	}
	fmt.Println("writing to file")
	//if not greater than max size write to file on new line
	_, err = file.WriteString(log + "\n")
	if err != nil {
		fmt.Println(err)
	}
}

func parseArguments(args ...interface{}) []arguments {
	//parse args to an array of key value pairs passed as key, value, key, value in args
	//every odd number is a key and even number is a value
	//return an array of key value pairs
	var argArray []arguments
	for i := 0; i < len(args); i++ {
		argArray = append(argArray, arguments{args[i].(string): args[i+1]})
		i++
	}
	return argArray
}

func (l *logger) processLogs(net network.NetworkOptions, log string) {
	//process logs
	//send logs to server
	//sendLogs(log, network.NetworkOptions{})
	if err := network.PostLogs(log, net); err != nil {
		fmt.Println(err)
	}
}




func (l *logger) getLastCreatedFile(filepath string) (string, error) {
	//get last created file in the directory
	//return the file name
	var lastCreatedFile string
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return "", err
	}
	var mdTime time.Time
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if f.ModTime().After(mdTime) {
			mdTime = f.ModTime()
			lastCreatedFile = f.Name()
		}
	}
	return lastCreatedFile, nil
}
