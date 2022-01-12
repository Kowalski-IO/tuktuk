package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"

	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"./database"
	"./models"
)

func main() {
	db := database.Init()

	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	r.PUT("/*filename", func(c *gin.Context) {
		filename := strings.TrimPrefix(c.Param("filename"), "/")
		data, err := c.GetRawData()

		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Errorf("Unable to read raw data")

			c.JSON(500, gin.H{"message": "An internal error as occurred. Please try again later."})
			return
		}

		id := xid.New().String()
		fp := sha256.Sum256(data)

		metadata := models.File{
			ID:               id,
			AdministrationID: uuid.New(),
			Filename:         filename,
			Type:             http.DetectContentType(data),
			Size:             len(data),
			Fingerprint:      hex.EncodeToString(fp[:]),
			SelfDestruct:     false,
			Status:           models.Active,
			CreatedAt:        time.Now(),
			DestroyOn:        time.Now().Add(time.Hour * time.Duration(24)),
		}

		err = ioutil.WriteFile("uploads/"+id, data, os.FileMode(0755))
		if err != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Errorf("Unable to save file")

			c.JSON(500, gin.H{"message": "An internal error as occurred. Please try again later."})
			return
		}

		res := db.Create(&metadata)

		if res.Error != nil {
			log.WithFields(log.Fields{
				"error": err.Error(),
			}).Errorf("Unable to save file metadata")

			// Delete file logic here
		}

		c.JSON(200, gin.H{
			"metadata": metadata,
		})
	})

	err := r.Run()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Errorf("Unable to start TukTuk Server!")
	}
}
