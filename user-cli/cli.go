package main

import (
	"context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	mClient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"log"
	"os"
	pb "shippy/user-service/proto/user"
)

func main() {
	cmd.Init()
	client := pb.NewUserServiceClient("go.micro.srv.user", mClient.DefaultClient)

	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "you full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		))

	// Start as service
	service.Init(
		micro.Action(func(c *cli.Context) {
			//name := c.String("name")
			//email := c.String("email")
			//password := c.String("password")
			//company := c.String("company")
			// Call our user service
			r, err := client.Create(context.TODO(), &pb.User{
				Name:     "xiaohan",
				Email:    "95028555@qq.com",
				Password: "123",
				Company:  "haha",
			})
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %s", r.User.Uid)
			getAll, err := client.GetAll(context.Background(), &pb.Request{})
			if err != nil {
				log.Fatalf("Could not list users: %v", err)
			}
			for _, v := range getAll.Users {
				log.Println(v)
			}
			os.Exit(0)
		}),
	)
	// Run the server
	if err := service.Run(); err != nil {
		log.Println("log--->", err)
	}
}
