package main

import (
	"fmt"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"golang.org/x/net/context"
)

func main() {
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}

	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.21", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}

	r, err := cli.ContainerCreate(context.Background(),
		&container.Config{
                        Image: "ubuntu", Hostname: "",
			User:         "",
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
			Cmd: []string{"/bin/bash"}, 
			OpenStdin:    false,
			StdinOnce:    true,
			WorkingDir:   ""}, nil, nil, "MyName6")
	if err != nil {
		panic(err)
	}
	if r.ID != "container_id" {
		fmt.Printf("expected `container_id`, got %s \n", r.ID)
		fmt.Println(err)
		fmt.Println(r)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

	err = cli.ContainerStart(context.Background(), r.ID)
	if err != nil {
		panic(err)
	}
	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Println(c.ID)
		fmt.Println(c.Names)
		// fmt.Println(c)
	}
}
