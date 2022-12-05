package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cv2 "gocv.io/x/gocv"
	"image"
	"image/color"
	"main/util"
	"path"
	"path/filepath"
	"sort"
)

var img image.Image
var mat, mat_gray, mat_binary, mat_org cv2.Mat
var min_rect cv2.RotatedRect
var contour_max cv2.PointVector
var saveConfig StepConfig
var ans Ans = Ans{}

func init() {
	mat_gray = cv2.NewMat()
	mat_binary = cv2.NewMat()
}

func picture(c *gin.Context) {
	if img == nil {
		config := NewStepConfig()
		img = util.OpenImg(config.Demo)
	}
	util.WriteImage(c.Writer, &img)
}

func load(c *gin.Context) {
	var config StepConfig
	c.ShouldBind(&config)
	fmt.Printf("%v\n", config)
	mat = cv2.IMRead(config.Demo, cv2.IMReadColor)
	if mat.Empty() {
		c.JSON(200, "加载文件错误")
		return
	}
	mat_org = mat.Clone()
	img, _ = mat.ToImage()
	util.WriteImage(c.Writer, &img)
}

func gray(c *gin.Context) {
	var config StepConfig
	c.ShouldBind(&config)
	cv2.CvtColor(mat, &mat_gray, cv2.ColorBGRToGray)
	img, _ = mat_gray.ToImage()
	util.WriteImage(c.Writer, &img)
}

func binary(c *gin.Context) {
	var config StepConfig
	c.ShouldBind(&config)
	cv2.Threshold(mat_gray, &mat_binary, config.Threshold.Thresh, config.Threshold.Maxvalue, config.Threshold.ThresholdType)
	img, _ = mat_binary.ToImage()
	util.WriteImage(c.Writer, &img)
}

func ContourArea(c *gin.Context) {
	var config StepConfig
	c.ShouldBind(&config)
	contours := cv2.FindContours(mat_binary, cv2.RetrievalTree, cv2.ChainApproxSimple)
	contoursArea := ArrSortArea{}
	for i := 0; i < contours.Size(); i++ {
		temp_area := cv2.ContourArea(contours.At(i))
		contoursArea = append(contoursArea, PointsVectorArea{contours.At(i), temp_area})
	}
	sort.Sort(contoursArea)
	contour_max = contoursArea[config.ContourIndex].PointsVector
	points := cv2.NewPointsVector()
	points.Append(contour_max)
	mat_contour := mat.Clone()
	cv2.DrawContours(&mat_contour, points, -1, color.RGBA{0, 255, 0, 0}, 1) // red
	img, _ = mat_contour.ToImage()
	util.WriteImage(c.Writer, &img)
}

func MinAreaRect(c *gin.Context) {
	var config StepConfig
	c.ShouldBind(&config)
	min_rect = cv2.MinAreaRect(contour_max)
	mat_rect := mat.Clone()
	cv2.Rectangle(&mat_rect, min_rect.BoundingRect, color.RGBA{255, 0, 0, 0}, 2)
	img, _ = mat_rect.ToImage()
	util.WriteImage(c.Writer, &img)
}
func DealRect(rect image.Rectangle, max []int) image.Rectangle {
	if rect.Min.X < 0 {
		rect.Min.X = 0
	}
	if rect.Min.Y < 0 {
		rect.Min.Y = 0
	}
	if rect.Max.X > max[1] {
		rect.Max.X = max[1]
	}
	if rect.Max.Y > max[0] {
		rect.Max.Y = max[0]
	}
	return rect
}
func crop(c *gin.Context) {
	var config StepConfig
	c.ShouldBind(&config)
	rect := DealRect(min_rect.BoundingRect, mat_org.Size())
	crop := mat_org.Region(rect)
	img, _ = crop.ToImage()
	util.WriteImage(c.Writer, &img)
}

func SaveParam(c *gin.Context) {
	c.ShouldBind(&saveConfig)
	c.JSON(200, "ok")
}

func getstatu(c *gin.Context) {
	c.JSON(200, ans)
}

func batch(c *gin.Context) {
	var config StepConfig
	c.ShouldBind(&config)
	ans = Ans{}

	BatchCrop(&config)
	ans.Statu = true
	c.String(200, "ok")
}

func BatchCrop(config *StepConfig) {
	png, _ := filepath.Glob(path.Join(config.Dir, "*.png"))
	jpg, _ := filepath.Glob(path.Join(config.Dir, "*.jpg"))
	bmp, _ := filepath.Glob(path.Join(config.Dir, "*.bmp"))
	png = append(png, jpg...)
	png = append(png, bmp...)
	for i := 0; i < len(png); i++ {
		ans.F = append(ans.F, &AnsF{png[i], false})
	}
	for i := 0; i < len(ans.F); i++ {

		ans.F[i].Statu = Crop(ans.F[i].Path, config)
	}

}

func StartHttp() {
	r := gin.Default()
	r.StaticFile("/", "./index.html")
	r.POST("/picture", picture)
	r.POST("/load", load)
	r.POST("/gray", gray)
	r.POST("/binary", binary)
	r.POST("/ContourArea", ContourArea)
	r.POST("/MinAreaRect", MinAreaRect)
	r.POST("/SaveParam", SaveParam)
	r.POST("/crop", crop)
	r.GET("/getstatu", getstatu)
	r.POST("/batch", batch)
	r.Run(":" + Port)
}
