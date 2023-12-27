package logger

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/smartystreets/goconvey/convey"
	"io"
	"os"
	"testing"
	"time"
)

type logMsg struct {
	Time  time.Time `json:"time"`
	Level string    `json:"level"`
	Msg   string    `json:"msg"`
	Id    int       `json:"id"`
}

func TestSlogLogger(t *testing.T) {
	convey.Convey("Test slogger", t, func() {
		logFile, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			t.Fatal(err)
		}
		defer logFile.Close()

		logOpts := []Option{
			WithLevel(ErrorLevel),
			WithFile(logFile),
		}

		ctx := context.TODO()

		log := NewSlogLogger(logOpts...)
		log.Info(ctx, "Info msg", "id", 333)
		log.Error(ctx, "Error", "id", 777)

		logFile.Seek(0, io.SeekStart)

		scanner := bufio.NewScanner(logFile)

		cnt := 0
		defer func() {
			os.Remove("log.log")
		}()
		for scanner.Scan() {
			cnt++
			line := scanner.Bytes()

			m := logMsg{}
			err := json.Unmarshal(line, &m)
			if err != nil {
				t.Fatal(err)
			}

			convey.So(m.Id, convey.ShouldEqual, 777)
		}
		convey.So(cnt, convey.ShouldEqual, 1)
	})

}
