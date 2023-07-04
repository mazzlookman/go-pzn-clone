package helper

import (
	"context"
	"fmt"
	"gopkg.in/vansante/go-ffprobe.v2"
	"log"
	"os"
	"strings"
	"time"
)

func GetLessonContentVideoDuration(path string) string {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, path)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}

	s := data.Format.Duration().Round(data.Format.Duration()).String()
	var duration string
	before, _, found := strings.Cut(s, ".")

	if found {
		split := strings.Split(before, "m")
		if len(split[1]) == 1 {
			second := fmt.Sprintf("0%s", split[1])
			split[1] = second
		}
		join := strings.Join(split, ":")
		duration = join
	}

	return duration
}

func DeleteLessonContent(path string) string {
	err := os.Remove(path)
	if err != nil {
		return err.Error()
	}

	return path + " was deleted"
}
