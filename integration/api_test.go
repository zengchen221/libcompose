package integration

import (
	. "gopkg.in/check.v1"

	"golang.org/x/net/context"

	"github.com/zengchen221/libcompose/docker"
	"github.com/zengchen221/libcompose/docker/ctx"
	"github.com/zengchen221/libcompose/project"
	"github.com/zengchen221/libcompose/project/options"
)

func init() {
	Suite(&APISuite{})
}

type APISuite struct{}

func (s *APISuite) TestVolumeWithoutComposeFile(c *C) {
	service := `
service:
  image: busybox
  command: echo Hello world!
  volumes:
    - /etc/selinux:/etc/selinux`

	project, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeBytes: [][]byte{[]byte(service)},
			ProjectName:  "test-volume-without-compose-file",
		},
	}, nil)

	c.Assert(err, IsNil)

	err = project.Up(context.Background(), options.Up{})
	c.Assert(err, IsNil)
}
