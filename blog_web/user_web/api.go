package user_web

import (
	"log"
	"net/http"

	"blog-go/blog_srv/user_srv/dao"
	form "blog-go/blog_web/user_web/forms"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

//HandleValidatorErr 表单验证错误处理返回
func HandleValidatorErr(c *gin.Context, err error) {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errs, //"errs.Translate(global.Trans)"
	})
}

//登录
func Login(ctx *gin.Context) {
	//表单
	var forms form.LoginForm
	if err := ctx.BindJSON(&forms); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}

	//获取数据库密码
	res, err := dao.MobileToUser(&dao.UserInfo{
		Mobile: forms.Mobile,
	})
	if err != nil {
		log.Fatal("手机号没有注册", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "手机号没有注册",
		})
		return
	}

	//验证密码
	isCheck := dao.CheckPassWord(res.PassWord, forms.Password)
	if !isCheck {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"password": "登录失败",
		})

	} else {
		//TODO 密码验证通过，颁发签名，返回token
	}
}
