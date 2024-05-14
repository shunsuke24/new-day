package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/new-day/infra"
	"github.com/new-day/presenter"
)

func (cont *controller) Send(c *gin.Context) {
	// Redisに接続する
	redis := infra.NewRedis()
	// Send()処理終了後にRedisとの接続を切断する
	defer redis.CloseRedis()

	// requestInformationをUserInformationで初期化する
	requestInformation := presenter.UserInformation{}

	// 構造体をBINDする
	err := c.Bind(&requestInformation)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "BadRequest"})
	}

	// Redisで使用するキーの作成
	key := requestInformation.ID + ":" + requestInformation.Name

	// 作成した構造体requestInformationをJSONに変換する
	payload, err := json.Marshal(requestInformation)
	if err != nil {
		fmt.Println("JSON Marshal Error : ", err)
		return
	}

	// key, payloadを引数にRedisに追加する
	if err := redis.Set(key, payload); err != nil {
		fmt.Println("Failed to store data in Redis. ", err)
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Successfully added to redis. "})
}
