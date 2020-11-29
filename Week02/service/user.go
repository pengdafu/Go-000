package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strconv"
	"week02/dao"
)

type Resp struct {
	code int
	msg  interface{}
	err  error
}

func FindUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(200, Resp{
			code: 400,
			err:  err,
		})
		return
	}
	u, err := dao.FindUserById(id)
	r := Resp{
		msg:  u,
		err:  err,
		code: 200,
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("%+v", err)
		r.code = 404
	} else if err != nil {
		log.Printf("%+v", err)
		r.code = 500
	}
	ctx.JSON(200, r)
}
