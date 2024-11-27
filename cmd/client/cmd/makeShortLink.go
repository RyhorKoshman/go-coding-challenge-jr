/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"challenge/pkg/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"

	"github.com/spf13/cobra"
)

// makeShortLinkCmd represents the makeShortLink command
var makeShortLinkCmd = &cobra.Command{
	Use:   "makeShortLink -link",
	Short: "Prints such -link in the sort format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()
		c := proto.NewChallengeServiceClient(conn)
		md := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{}))
		res, err := c.MakeShortLink(md, &proto.Link{Data: args[0]})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.Data)
	},
}

func init() {
	rootCmd.AddCommand(makeShortLinkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeShortLinkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeShortLinkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
