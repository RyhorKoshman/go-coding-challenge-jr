package server

import (
	"challenge/pkg/proto"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (*Server) MakeShortLink(ctx context.Context, in *proto.Link) (*proto.Link, error) {
	client := &http.Client{}
	raw := `{"long_url": `
	raw = raw + in.Data + "}"
	data := strings.NewReader(raw)
	req, err := http.NewRequest("POST", "https://api-ssl.bitly.com/v4/shorten", data)
	if err != nil {
		panic(fmt.Errorf("NewRequest erorr: %v", err))
	}
	token := viper.GetString("BITLY_OAUTH_TOKEN")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer {%s}", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	///TODO: incorrect message decode
	fmt.Printf("%s\n", bodyText)
	return nil, nil
}
