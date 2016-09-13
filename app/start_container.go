package app

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stringid"
	"golang.org/x/net/context"
	"os"
)

func (app *App) StartContainer(station string, minutes string, channel string) string {
	imageName := os.Getenv("IMAGE_NAME")
	if imageName == "" {
		imageName = "atsnngs/radiko-recorder-s3"
	}
	config := &container.Config{
		Image:        imageName,
		AttachStdin:  false,
		AttachStdout: false,
		AttachStderr: false,
		StdinOnce:    false,
		Env: []string{"SLACK_CHANNEL=" + channel,
			"STATION=" + station,
			"DURATION_MINUTES=" + minutes,
			"RADIKO_LOGIN=" + os.Getenv("RADIKO_LOGIN"),
			"RADIKO_PASSWORD=" + os.Getenv("RADIKO_PASSWORD"),
			"S3_BUCKET=" + os.Getenv("S3_BUCKET"),
			"AWS_ACCESS_KEY_ID=" + os.Getenv("AWS_ACCESS_KEY_ID"),
			"AWS_SECRET_ACCESS_KEY=" + os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"SLACK_WEBHOOK_URL=" + os.Getenv("SLACK_WEBHOOK_URL")}}
	hostConfig := &container.HostConfig{
		AutoRemove: true}
	ctx := context.Background()
	res, err := app.Docker.ContainerCreate(ctx, config, hostConfig, nil, "")
	if err != nil {
		return err.Error()
	}
	if err := app.Docker.ContainerStart(ctx, res.ID, types.ContainerStartOptions{}); err != nil {
		return err.Error()
	}
	return "Started container " + stringid.TruncateID(res.ID)
}
