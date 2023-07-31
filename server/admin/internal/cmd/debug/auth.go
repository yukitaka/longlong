package debug

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/server/admin/internal/interface/server/jwt"
	"github.com/yukitaka/longlong/server/core/pkg/cli"
	"github.com/yukitaka/longlong/server/core/pkg/util"
	"strconv"
)

type Options struct {
	CmdParent string
	cli.IOStream
}

func NewCmdDebug(parent string, streams cli.IOStream) *cobra.Command {
	o := newDebugOptions(parent, streams)

	cmd := &cobra.Command{
		Use:   "debug",
		Short: "Debug Longlong",
		RunE: func(cmd *cobra.Command, args []string) error {
			util.CheckErr(o.Run(args))
			return nil
		},
	}

	return cmd
}

func newDebugOptions(parent string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		IOStream:  streams,
	}
}

func (o *Options) Run(args []string) error {
	secret, err := util.GetEnvironmentValue("JWT_SECRET")
	if err != nil {
		panic(err)
	}

	individualId, _ := strconv.Atoi(args[0])
	organizationId, _ := strconv.Atoi(args[1])
	token, err := jwt.CreateToken(individualId, organizationId, secret)
	if err != nil {
		return err
	}
	fmt.Println(token)
	return nil
}
