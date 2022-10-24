package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
	s "github.com/sammy1881/tinyurl/storage"
)

func getDBWriter() s.StorageService {
	dbw, err := s.NewStrorageService()
	if err != nil {
		log.Panicf("Could not open connection to Database")
		return nil
	}
	return dbw
}

// Home - Home http request
func Home(c *gin.Context) {
	dbw := getDBWriter()
	count, _ := dbw.Count()
	response, err := dbw.GetAllRecords()
	if err != nil {
		c.AbortWithError(500, err)
	}
	response += fmt.Sprintf("The total links in the DB is: %d", count)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
}

// addLink - Add a link to the linkList and generate a shorter link
// Example: localhost:8080/addurl?link=https://google.com
func AddURL(c *gin.Context) {
	dbw := getDBWriter()
	var err error
	key := c.Query("link")
	log.Printf("Adding Link: %s \n", key)
	if key == "" {
		err = fmt.Errorf("could not read the link to be added")
		c.AbortWithError(500, err)
	}
	if utf8.RuneCountInString(key) > 2000 {
		err = fmt.Errorf("input link too long")
		c.AbortWithError(500, err)
	}
	if !validLink(key) {
		err = fmt.Errorf("could not create shortlink need absolute path link. Ex: /addurl?link=https://github.com/")
		c.AbortWithError(500, err)
	}

	parsedUrl, err := url.Parse(key)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "https"
	}

	id, err := gonanoid.Generate(c.MustGet("IdAlphabet").(string), c.MustGet("IdLength").(int))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = dbw.Put(id, parsedUrl.String())

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	result := fmt.Sprintf(c.MustGet("ShortenerHostname").(string) + ":" + c.MustGet("Port").(string) + "/" + id)

	c.JSON(http.StatusOK, result)

}

// validLink - check that the link we're creating a shortlink for is a absolute URL path
func validLink(link string) bool {
	r, err := regexp.Compile("^(http|https)://")
	if err != nil {
		return false
	}
	link = strings.TrimSpace(link)
	log.Printf("Checking for valid link: %s", link)
	// Check if string matches the regex
	return r.MatchString(link)
}

// getURL - Find link that matches the shortened link in the db
func GetURL(c *gin.Context) {
	dbw := getDBWriter()
	shorturl := c.Param("shorturl")
	log.Printf("Getting Link: %s \n", shorturl)
	destination, err := dbw.Get(shorturl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	if destination == "" {
		c.AbortWithError(http.StatusNotFound, nil)
	}
	log.Println(destination)
	c.Redirect(302, destination)
}
