<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p>

<h3 align="center">tinyurl</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/sammy1881/tinyurl.svg)](https://github.com/sammy1881/tinyurl/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/sammy1881/tinyurl.svg)](https://github.com/sammy1881/tinyurl/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Practice code for tinyurl service
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Built Using](#built_using)
- [TODO](../TODO.md)
- [Authors](#authors)

## 🧐 About <a name = "about"></a>

Just a practice code for learning golang. 

 ## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

```
go
docker (optional)
```

### Installing

Self-built binary is in their respective OS folder. They can be run directly as:

```
./tinyurl 
```
You can also build a docker image with provided Dockerfile
```
docker build -t tinyurl:latest .
docker run -d -p 8080:8080 tinyurl:latest
```

You can also pass config variables as ENVIRONMENT_VARIABLES

```
Supported Environment Variables
TINYURL_DB          // path to db file
TINYURL_BUCKET      // Name of bucket in DB
TINYURL_HOSTNAME    // hostname of the service
TINYURL_PORT        // Port for the service
TINYURL_ID_LENGTH   // length of random string 
TINYURL_ID_ALPHABET // alphabets in random string
```

End with an example of getting some data out of the system or using it for a little demo.


## 🎈 Usage <a name="usage"></a>

Creating a tinyurl can be done with a POST call
```
$ curl -X POST http://localhost:8080/addurl?link=https://www.google.com

"localhost:8080/By2D7i"
```

You can also view the list of created short URL by going to http://localhost:8080 on your browser. 


## ⛏️ Built Using <a name = "built_using"></a>

- [Gin](https://github.com/gin-gonic/gin) - HTTP Web Framework
- [Bolt](https://github.com/boltdb/bolt) - Embedded K/V database
- [nanoid](github.com/matoous/go-nanoid) - Golang random IDs generator

## ✍️ Authors <a name = "authors"></a>

- [@sammy1881](https://github.com/sammy1881) - Idea & Initial work

See also the list of [contributors](https://github.com/sammy1881/tinyurl/contributors) who participated in this project.
