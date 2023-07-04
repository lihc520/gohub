package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lihc520/gohub/app/http/controllers/api/v1"
	"github.com/lihc520/gohub/pkg/captcha"
	"github.com/lihc520/gohub/pkg/logger"
	"net/http"
)

// VerifyCodecontroller 用户控制器
type VerifycodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifycodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
