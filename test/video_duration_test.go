package test

import (
	"go-pzn-clone/helper"
	"testing"
)

func TestGetVideoDuration(t *testing.T) {
	duration := helper.GetLessonContentVideoDuration("../resources/contents/Belajar_Go-Lang_-_1_Pengenalan_Go-Lang.mov")
	t.Log(duration)
}
