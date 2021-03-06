package routers

import (
    //"net/http"
    //"time"

    "github.com/spark8899/deploy-agent/global"

    "github.com/gin-gonic/gin"
    "github.com/spark8899/deploy-agent/internal/middleware"
    "github.com/spark8899/deploy-agent/internal/routers/api"
)

func NewRouter() *gin.Engine {
    r := gin.New()
    if global.ServerSetting.RunMode == "debug" {
        r.Use(gin.Logger())
        r.Use(gin.Recovery())
    } else {
        r.Use(middleware.AccessLog())
        r.Use(middleware.Recovery())
    }

    r.Use(middleware.Tracing())
    r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))

    upload := api.NewUpload()
    r.GET("/debug/vars", api.Expvar)
    r.POST("/upload/file", upload.UploadFile)
    r.POST("/command", api.PostExecCommand)
    //r.StaticFS("/static", http.Dir(global.AppSetting.DeployPath))

    return r
}
