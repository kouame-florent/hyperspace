package hyp_test

/*

func runCmd(command string, args []string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Print(string(output))
		return err
	}

	log.Print(string(output))
	return nil

}

func createClient(t *testing.T) (*client.Client, error) {

	client.NewClientWithOpts()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	t.Cleanup(func() {
		cli.Close()
	})

	return cli, nil
}

func cmdNetworkCreate(t *testing.T, name string) {
	err := runCmd("docker", []string{"network", "create", name})
	if err != nil {
		t.Fatal(err)
	}
}

func cmdNetworkRemove(t *testing.T, name string) {
	err := runCmd("docker", []string{"network", "rm", name})
	if err != nil {
		t.Fatal(err)
	}
}

func cmdVolumeCreate(t *testing.T, name string) {
	err := runCmd("docker", []string{"volume", "create", name})
	if err != nil {
		t.Fatal(err)
	}
}

func cmdVolumeRemove(t *testing.T, name string) {
	err := runCmd("docker", []string{"volume", "create", name})
	if err != nil {
		t.Fatal(err)
	}
}

// e.g docker run -d --name test_nginx nginx:1.23.2-alpine
func cmdContainerCreate(t *testing.T, image, name string) {
	err := runCmd("docker", []string{"run", "-d", "--name", name, image})
	if err != nil {
		t.Fatal(err)
	}
}

// e.g docker container stop test_nginx
func cmdContainerStop(t *testing.T, name string) {
	err := runCmd("docker", []string{"container", "stop", name})
	if err != nil {
		t.Fatal(err)
	}
}

// e.g docker container rm test_nginx
func cmdContainerRemove(t *testing.T, name string) {
	err := runCmd("docker", []string{"container", "rm", name})
	if err != nil {
		t.Fatal(err)
	}
}
*/
