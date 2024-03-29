/*
 * photo viewer
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetPhotos - 
func GetPhotos(c *gin.Context) {
	imageDir := os.Getenv("IMAGEDIR")
	photos := loadFiles(imageDir)
	res := Photos{
		Photos: photos,
	}
	c.JSON(http.StatusOK, res)
}

func loadFiles(dir string) (photos []PhotoProperties) {
	baseUrl := os.Getenv("BASEURL")
	id := 1
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".jpg") {
			pp := PhotoProperties{
				Id: id,
				Title: file.Name(),
				Url: baseUrl + "/" + file.Name(),
			}
			photos = append(photos, pp)
			id += 1
		}
	}
	return
}
