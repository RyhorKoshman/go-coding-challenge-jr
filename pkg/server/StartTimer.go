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

const url = "https://timercheck.io/"

var client = &http.Client{}

func iterate(freq int64, s grpc.ServerStreamingServer[proto.Timer], name string) {
	req, _ := http.NewRequest("GET", url+name, nil)
	data, _ := client.Do(req)
	v := viper.New()
	v.SetConfigType("json")
	err := v.ReadConfig(data.Body)
	if err != nil {
		fmt.Println("v.ReadConfig err: ", err)
	}
	now := v.GetInt64("seconds_remaining")
	for {

		time.Sleep(time.Duration(freq) * time.Second)
		now -= freq
		if now <= 0 {
			break
		}
		fmt.Println("Left for ", name, " ", now)
		e := s.Send(&proto.Timer{
			Name:      name,
			Frequency: freq,
			Seconds:   now,
		})
		if e != nil {
			fmt.Println("s.Send err: ", e)
			return
		}
	}
}

func setNewTimer(t *proto.Timer, s grpc.ServerStreamingServer[proto.Timer]) error {
	req, _ := http.NewRequest("GET", url+t.GetName()+"/"+strconv.FormatInt(t.Seconds, 10), nil)
	data, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do err ", err)
		return err
	}
	fmt.Print("Set new Timer: ", t)
	defer data.Body.Close()
	iterate(t.Frequency, s, t.Name)
	return nil
}

func connectExistingTimer(t *proto.Timer, s grpc.ServerStreamingServer[proto.Timer]) error {
	iterate(t.Frequency, s, t.Name)
	return nil
}

func (*Server) StartTimer(t *proto.Timer, s grpc.ServerStreamingServer[proto.Timer]) error {
	req, _ := http.NewRequest("GET", url+t.GetName(), nil)
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
