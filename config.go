package main

import (
	cv2 "gocv.io/x/gocv"
	"os"
	"path/filepath"
	"strings"
)

type PointsVectorArea struct {
	PointsVector cv2.PointVector
	Area         float64
}
type ArrSortArea []PointsVectorArea

func (a ArrSortArea) Len() int           { return len(a) }
func (a ArrSortArea) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ArrSortArea) Less(i, j int) bool { return a[i].Area > a[j].Area }

type AnsF struct {
	Path  string `json:"path"`
	Statu bool   `json:"statu"`
}

type Ans struct {
	F     []*AnsF `json:"f"`
	Statu bool    `json:"statu"`
}

type Threshold struct {
	Thresh        float32           `json:"thresh"`
	Maxvalue      float32           `json:"maxvalue"`
	ThresholdType cv2.ThresholdType `json:"threshold_type"`
}

type StepConfig struct {
	Demo         string `json:"demo"`
	Port         string
	Threshold    Threshold `json:"threshold"`
	ContourIndex int       `json:"contour_index"`
	Dir          string    `json:"dir"`
}

func NewStepConfig() *StepConfig {
	return &StepConfig{Demo: "demo.png"}
}

func ChangeFileName(sPath string, name string, addDir bool) string {

	if !addDir {
		return filepath.Join(filepath.Dir(sPath), strings.TrimSuffix(filepath.Base(sPath), filepath.Ext(sPath))+name+filepath.Ext(sPath))
	} else {
		dir := filepath.Join(filepath.Dir(sPath), name)
		os.Mkdir(dir, os.ModePerm)
		return filepath.Join(dir, filepath.Base(sPath))
	}

}
