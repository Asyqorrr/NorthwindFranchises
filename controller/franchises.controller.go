package controller

import (
	db "b30northwindapi/db/sqlc"
	"b30northwindapi/models"
	"b30northwindapi/services"
	"database/sql"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type FranchisesController struct {
	storedb services.Store
}

func NewFranchisesController(store services.Store) *FranchisesController{
	return &FranchisesController{
		storedb: store,
	}
}

func (fr *FranchisesController) GetAllFranchises (c *gin.Context){
	franchises, err := fr.storedb.GetAllFranchises(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}
	c.JSON(http.StatusOK, franchises)
}

func (fr *FranchisesController) GetAllFranchisesById (c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	franchise, err := models.Nullable(fr.storedb.GetFranchisesById(c, int32(id)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	if franchise == nil {
		c.JSON(http.StatusNotFound, models.NewError(models.ErrFranchisesNotFound))
	}

	c.JSON(http.StatusOK, franchise)
}

func (fr *FranchisesController) CreateFranchises (c *gin.Context){
	var payload *models.CreateFranchisesReq
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	// Set timestamps to now
	timestamps := time.Now()
	ts := pgtype.Timestamp{
		Time:   timestamps,
		Valid:  true, // Indicates that the timestamp is valid (not NULL)
	}
	
	args := &db.CreateFranchisesParams{
		FrchName: payload.FrchName,
		FrchDesc: payload.FrchDesc,
		FrchPriceBuyout: payload.FrchPriceBuyout,
		FrchPriceYearly: payload.FrchPriceYearly,
		FrchModified: ts,
		FrchCateID: payload.FrchCateID,
	}

	franchises, err := fr.storedb.CreateFranchises(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
	}

	c.JSON(http.StatusCreated, franchises)
}

type UploadImagesDto struct {
	FrimFilename *string `form:"frim_filename"`
	FrimDefault  *string `form:"frim_default"`
	FrimFrchID   *int32  `form:"frim_frch_id"`
	Filename	 *SingleFileUpload
}

type SingleFileUpload struct {
	Filename *multipart.FileHeader `form:"filename" binding:"required"`
}

type MultipleFileUpload struct{
	Filename []*multipart.FileHeader `form: "filename" binding:"required"`
}

func (fr *FranchisesController) UploadImages(c *gin.Context){
	var payload *UploadImagesDto

	id, _ := strconv.Atoi(c.Param("id"))
	idInt32 := int32(id)

	err := c.ShouldBind(&payload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.NewValidationError(err))
		return
	}

	fileUpload, err := c.FormFile("filename")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	extension := filepath.Ext(fileUpload.Filename)

	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(fileUpload, "./public/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	args := &db.UploadImagesParams{
		FrimFilename:  &newFileName,
		FrimDefault:   payload.FrimDefault,
		FrimFrchID: &idInt32,
	}

	Images, err := fr.storedb.UploadImages(c, *args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, Images)
}

func (fr *FranchisesController) DeleteFranchises(c *gin.Context){
	id, _ := strconv.Atoi((c.Param("id")))

	_, err := fr.storedb.GetFranchisesById(c, int32(id))
	if err != nil {
		if err == sql.ErrNoRows{
			c.JSON(http.StatusNotFound, models.ErrDataNotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	err = fr.storedb.DeleteFranchises(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError(err))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"status": "success", "message": "data has been deleted"})
}


