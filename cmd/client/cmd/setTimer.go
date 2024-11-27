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
	"log"
	"strconv"
)

// setTimerCmd represents the setTimer command
var setTimerCmd = &cobra.Command{
	Use:   "setTimer -sec -freq -name",
	Short: "Sets a new Timer with name -name, -sec seconds left and updates every -freq seconds",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		sec, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln(args[0], " is not a number: ", err)
		}
		freq, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln(args[1], " is not a number: ", err)
		}
		timer := &proto.Timer{
			Name:      args[2],
			Seconds:   int64(sec),
			Frequency: int64(freq),
		}
		client := proto.NewChallengeServiceClient(conn)
		stream, err := client.StartTimer(context.Background(), timer)
		for val, err := stream.Recv(); err == nil; val, err = stream.Recv() {
			fmt.Println(val.Seconds)
		}
	},
}

func init() {
	rootCmd.AddCommand(setTimerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setTimerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setTimerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
