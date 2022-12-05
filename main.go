// 打开图片
package main

import (
	"fmt"
	cv2 "gocv.io/x/gocv"
	"main/util"
	"sort"
)

var Port = "8080"

func main() {
	err := util.OpenBrowser("http://localhost:" + Port)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	StartHttp()

}

func Crop(file string, config *StepConfig) bool {
	img := cv2.IMRead(file, cv2.IMReadColor)
	if img.Empty() {
		return false
	}
	gray := cv2.NewMat()
	cv2.CvtColor(img, &gray, cv2.ColorBGRToGray)
	binary := cv2.NewMat()
	cv2.Threshold(gray, &binary, config.Threshold.Thresh, config.Threshold.Maxvalue, config.Threshold.ThresholdType)
	contours := cv2.FindContours(binary, cv2.RetrievalTree, cv2.ChainApproxSimple)
	contoursArea := ArrSortArea{}
	for i := 0; i < contours.Size(); i++ {
		temp_area := cv2.ContourArea(contours.At(i))
		contoursArea = append(contoursArea, PointsVectorArea{contours.At(i), temp_area})
	}
	sort.Sort(contoursArea)
	contour_max = contoursArea[config.ContourIndex].PointsVector
	min_rect := cv2.MinAreaRect(contour_max)
	rect := DealRect(min_rect.BoundingRect, img.Size())
	crop := img.Region(rect)
	cv2.IMWrite(ChangeFileName(file, "crop", true), crop)
	return true
}
