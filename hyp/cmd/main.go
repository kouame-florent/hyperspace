package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/google/uuid"
	hyp "github.com/kouame-florent/hyperspace/hyp/pkg"
)

func main() {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	quarkusBD := quarkusBundle(ctx, cli)
	mariadbBD := mariadbBundle(ctx, cli)

	bundles := []*hyp.Bundle{quarkusBD, mariadbBD}

	for _, b := range bundles {
		_, err := b.Deploy(ctx, cli)
		if err != nil {
			panic(err)
		}
	}

}

func quarkusBundle(ctx context.Context, cli *client.Client) *hyp.Bundle {

	env := []string{
		"PUID=1000",
		"PGID=1000",
		"TZ=Africa/Abidjan",
		"PASSWORD=code",
		"DEFAULT_WORKSPACE=/config/workspace",
	}

	Volumes := []string{
		"/config",
	}

	Ports := []string{
		"8443/tcp",
		"8080/tcp",
	}

	opt := hyp.NewContainerSpec("127.0.0.1:5000/code-java-env:v0.0.8", env, Volumes, []string{}, Ports)
	fmt.Printf("---> IMAGE NAME: %s\n", opt.Name)
	bundle := hyp.NewBundle("quarkus", "quarkus development enrironment", []string{uuid.NewString()}, []hyp.ContainerSpec{*opt})

	return bundle

}

func mariadbBundle(ctx context.Context, cli *client.Client) *hyp.Bundle {

	bundleNet := []string{uuid.New().String()}

	mEnv := []string{
		"MARIADB_ROOT_PASSWORD=test",
	}

	mVolumes := []string{
		"/var/lib/mysql",
	}

	mPorts := []string{
		"3306/tcp",
	}

	mOpt := hyp.NewContainerSpec("127.0.0.1:5000/mariadb:10.9.4", mEnv, mVolumes, bundleNet, mPorts)
	fmt.Printf("---> IMAGE NAME: %s\n", mOpt.Name)
	mBundle := hyp.NewBundle("mariadb", "mariadb server", []string{uuid.NewString()}, []hyp.ContainerSpec{*mOpt})

	adminerDefaultServer := "ADMINER_DEFAULT_SERVER=" + mOpt.Name

	env := []string{
		adminerDefaultServer,
	}

	Volumes := []string{}

	Ports := []string{
		"8080/tcp",
	}

	opt := hyp.NewContainerSpec("127.0.0.1:5000/adminer:4.8.1-standalone", env, Volumes, bundleNet, Ports)
	fmt.Printf("---> IMAGE NAME: %s\n", opt.Name)
	aBundle := hyp.NewBundle("adminer", "adminer server", []string{uuid.NewString()}, []hyp.ContainerSpec{*opt})

	bund := mBundle.Add(aBundle)

	return bund

}

func adminerBundle(ctx context.Context, cli *client.Client) *hyp.Bundle {
	env := []string{
		"ADMINER_DEFAULT_SERVER=mysql",
	}

	Volumes := []string{}

	Ports := []string{
		"8080/tcp",
	}

	opt := hyp.NewContainerSpec("127.0.0.1:5000/adminer:4.8.1-standalone", env, Volumes, []string{}, Ports)
	fmt.Printf("---> IMAGE NAME: %s\n", opt.Name)
	bundle := hyp.NewBundle("adminer", "adminer server", []string{uuid.NewString()}, []hyp.ContainerSpec{*opt})

	return bundle

}
