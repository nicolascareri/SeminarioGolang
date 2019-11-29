package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/seminario/lib"
    "net/http"
)

var agencia lib.Agencia

func main() {
    router := gin.Default()

    router.GET("/hello/:name", GetHello)
    router.GET("/auto/:auto", SearchAuto)

    agencia = lib.NewAgencia()

    MiAuto := lib.Auto{"ABC123", "rojo"}
    MiAuto1 := lib.Auto{"DEF123", "verde"}
    MiAuto2 := lib.Auto{"GHI123", "azul"}
    MiAuto3 := lib.Auto{"JKL123", "rojo"}


    agencia.AddAuto(MiAuto)
    agencia.AddAuto(MiAuto1)
    agencia.AddAuto(MiAuto2)
    agencia.AddAuto(MiAuto3)


    //    var unmarshalledCar Car
    //    err := json.Unmarshal(marshalledCar, &unmarshalledCar)
    //    if err != nil {
    //    fmt.Println(err)
    //    }
    if err := router.Run(); err != nil {
        fmt.Println(err)
    }
}


func SearchAuto(ctx *gin.Context) {
    plate := ctx.Param("auto")
    Auto := agencia.BuscarAuto(plate)
    fmt.Println("search")
    if Auto != nil {
        ctx.JSON(http.StatusOK, gin.H{
            "Auto": Auto,
        })
    }
}

func GetHello(ctx *gin.Context) {
	name := ctx.Param("name")
    ctx.JSON(http.StatusOK, gin.H{
        "message": "hello: " + name,
    })
}
