<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">Lepora</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> The best logger fo your project written in Go, with support for postgres and mongodb persistence.</p>
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)

## 🧐 About <a name = "about"></a>

This is a logger that has support for local file writes and database persistence. It has additional dashboards for viewing logs in a web browser.


## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

### Prerequisites

A system running go. You can check if you have go on your system by running

```
go version
```

### Installing

Run

```
go get github.com/augani/lepora@v0.1.0
```

## 🎈 Usage <a name="usage"></a>

```go
package main

import (
  "github.com/augani/lepora"
)

func main(){
  lep, err := lepora.Setup(lepora.LeporaOptions{
    // Options
    // Local means a local file will be created for logging
		Method: lepora.Local,
    //Name is your app name that will be used to name the log file
		Name: "AppTest",
    // Max size is the maximum size of the log file in bytes
		MaxSize: 1024,
    // Max files is the maximum number of log files you want to have
		MaxFiles: 5,
    // max days is the maximum number of days you want to keep logs for
		MaxDays: 7,
    // Debug is a boolean that determines if you want to log debug messages
		Debug: true,
	})

  if err != nil {
    panic(err)
  }

  lep.Log("key", "value", "key2", "value2")

  // Logs will be in key value pairs of any number
}
```

## ⛏️ Built Using <a name = "built_using"></a>

- [Go](https://golang.org/) - Go


## ✍️ Authors <a name = "authors"></a>

- [@augani](https://github.com/augani) - Idea & Initial work

<!-- And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo.

## 🔧 Running the tests <a name = "tests"></a>

Explain how to run the automated tests for this system.

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## 🎈 Usage <a name="usage"></a>

Add notes about how to use the system.

## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system.

## ⛏️ Built Using <a name = "built_using"></a>

- [MongoDB](https://www.mongodb.com/) - Database
- [Express](https://expressjs.com/) - Server Framework
- [VueJs](https://vuejs.org/) - Web Framework
- [NodeJs](https://nodejs.org/en/) - Server Environment

## ✍️ Authors <a name = "authors"></a>

- [@kylelobo](https://github.com/kylelobo) - Idea & Initial work

See also the list of [contributors](https://github.com/kylelobo/The-Documentation-Compendium/contributors) who participated in this project.

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References -->
