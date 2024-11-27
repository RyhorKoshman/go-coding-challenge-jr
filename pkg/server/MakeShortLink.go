package server

import (
	"bytes"
	"challenge/pkg/proto"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"strings"
)

func (*Server) MakeShortLink(ctx context.Context, in *proto.Link) (*proto.Link, error) {
	client := &http.Client{}

	data := strings.NewReader(fmt.Sprintf(`{"long_url": "%s", "domain": "bit.ly"}`, in.Data))
	fmt.Println("Data to link: ", data)
	urlAPI := "https://api-ssl.bitly.com/v4/shorten"
	req, err := http.NewRequest("POST", urlAPI, data)
	if err != nil {
		return nil, err
	}
	token := viper.GetString("BITLY_OAUTH_TOKEN")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf("Responce status not success: %s\n", resp.Status)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	v := viper.New()
	v.SetConfigType("json")
	v.ReadConfig(bytes.NewBuffer(bodyText))
	s := v.GetString("link")
	return &proto.Link{Data: s}, nil
}
