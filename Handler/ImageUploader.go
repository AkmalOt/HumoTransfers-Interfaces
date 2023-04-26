package Handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
)

func (h *Handler) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("icon")
	if err != nil {

		log.Printf("%s - error in FormFile - icon?", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if file.Size > 10000000 {
		log.Printf("file size is %s, need <50 Kb", file.Size)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file size is bigger than 50 Kb"})
		return
	}

	//Logger.Println(file.Filename)
	randomise := rand.Intn(111111)
	randomiserString := strconv.Itoa(randomise)

	// Путь указывается тут, я не добавил create папки, ибо при указании директории сохранения Иконок, отпадает необходисомть в этом
	Icon := "./icons/" + randomiserString + file.Filename

	fileExt := filepath.Ext(Icon)
	log.Println(Icon, fileExt)

	if fileExt != ".png" {
		log.Println("file is not png")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file is not png"})
		return
	}

	err = ctx.SaveUploadedFile(file, Icon)
	if err != nil {
		log.Println(err, "error in context.SaveUploadedFile")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, Icon)
}
