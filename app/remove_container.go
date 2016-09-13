package app

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func (app *App) RemoveContainer(containerID string) string {
	ctx := context.Background()
	if err := app.Docker.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true}); err != nil {
		return err.Error()
	}
	return "Successfully removed " + containerID
}
