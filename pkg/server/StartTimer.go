package server

import (
	"challenge/pkg/proto"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

const timerUrl = "https://timercheck.io/"

var client = &http.Client{}

func iterate(timer *proto.Timer, s grpc.ServerStreamingServer[proto.Timer]) {
	req, _ := http.NewRequest("GET", timerUrl+timer.Name, nil)
	data, _ := client.Do(req)
	v := viper.New()
	v.SetConfigType("json")
	err := v.ReadConfig(data.Body)
	if err != nil {
		fmt.Println("v.ReadConfig err: ", err)
		return
	}
	timer.Seconds = v.GetInt64("seconds_remaining")
	for {

		time.Sleep(time.Duration(timer.Frequency) * time.Second)
		timer.Seconds -= timer.Frequency
		if timer.Seconds <= 0 {
			timer.Seconds = 0
			break
		}
		e := s.Send(&proto.Timer{
			Name:      timer.Name,
			Frequency: timer.Frequency,
			Seconds:   timer.Seconds,
		})
		if e != nil {
			fmt.Println("s.Send err: ", e)
			return
		}
	}
}

func setNewTimer(t *proto.Timer, s grpc.ServerStreamingServer[proto.Timer]) error {
	req, _ := http.NewRequest("GET", timerUrl+t.GetName()+"/"+strconv.FormatInt(t.Seconds, 10), nil)
	data, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do err ", err)
		return err
	}
	fmt.Print("Set new Timer: ", t)
	defer data.Body.Close()
	iterate(t, s)
	return nil
}

func connectExistingTimer(t *proto.Timer, s grpc.ServerStreamingServer[proto.Timer]) error {
	iterate(t, s)
	return nil
}

func (*Server) StartTimer(t *proto.Timer, s grpc.ServerStreamingServer[proto.Timer]) error {
	req, _ := http.NewRequest("GET", timerUrl+t.GetName(), nil)
	data, err := client.Do(req)
	if err != nil {
		return err
	}
	defer data.Body.Close()
	if data.Status != "200 OK" {
		return setNewTimer(t, s)
	} else {
		return connectExistingTimer(t, s)
	}
}
