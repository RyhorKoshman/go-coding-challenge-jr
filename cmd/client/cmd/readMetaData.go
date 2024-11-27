/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"challenge/pkg/proto"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
)

// readMetaDataCmd represents the readMetaData command
var readMetaDataCmd = &cobra.Command{
	Use:   "readMetaData -keyName",
	Short: "Prints a value with the key equal to -keyName",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()
		c := proto.NewChallengeServiceClient(conn)
		md := metadata.Pairs("i-am-random-key", "lol?")
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		data, _ := c.ReadMetadata(ctx, &proto.Placeholder{Data: "i-am-random-key"})
		fmt.Println(data.Data)
	},
}

func init() {
	rootCmd.AddCommand(readMetaDataCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readMetaDataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readMetaDataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
