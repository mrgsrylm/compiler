package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/database"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/model"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/util"
)

var (
	appJSON = "application/json"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	user := model.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"full_name": user.FullName,
		"email":     user.Email,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := util.GetContentType(ctx)
	_, _ = db, contentType
	User := model.User{}
	password := ""

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := util.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	token := util.GenerateToken(User.ID, User.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
