package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johskw/meigen_bot/model"
)

func PostNickname(c *gin.Context) {
	var nickname model.Nickname
	err := c.Bind(&nickname)
	if err != nil {
		log.Print(err)
	}
	nickname.CharacterID, _ = strconv.Atoi(c.Param("character_id"))
	_, err = nickname.Create()
	if err != nil {
		log.Print(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/characters/"+c.Param("character_id"))
}