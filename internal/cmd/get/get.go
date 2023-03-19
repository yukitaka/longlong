package get

import (
	"fmt"
	"github.com/yukitaka/longlong/internal/domain/usecase"
	"github.com/yukitaka/longlong/internal/interface/repository"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yukitaka/longlong/internal/cli"
	"github.com/yukitaka/longlong/internal/util"

	"gopkg.in/yaml.v3"
)

type Options struct {
	CmdParent string
	Output    string
	cli.IOStream
}

func NewGetOptions(parent, output string, streams cli.IOStream) *Options {
	return &Options{
		CmdParent: parent,
		Output:    output,
		IOStream:  streams,
	}
}

func NewCmdGet(parent string, streams cli.IOStream) *cobra.Command {
	o := NewGetOptions(parent, "yaml", streams)

	cmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "Display one or many resources",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Run(cmd, args))
		},
	}

	organizationCmd := &cobra.Command{
		Use:     "organization",
		Aliases: []string{"organ"},
		Short:   "Display one or many organizations",
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Organization(cmd, args))
		},
	}
	organizationCmd.PersistentFlags().StringP("output", "o", "yaml", "output format")
	cmd.AddCommand(organizationCmd)

	return cmd
}

func (o *Options) Run(cmd *cobra.Command, args []string) error {
	fmt.Printf("Args is %v.", args)
	return nil
}

func (o *Options) Organization(cmd *cobra.Command, args []string) error {
	rep := repository.NewOrganizationsRepository()
	itr := usecase.NewOrganizationFinder(rep)

	var err error
	if output, err := cmd.PersistentFlags().GetString("output"); err == nil {
		if len(args) > 0 {
			if id, err := strconv.ParseInt(args[0], 10, 64); err == nil {
				if organization, err := itr.Find(id); err == nil {
					o.print(output, organization)
				}
			}
		} else {
			if organizations, err := itr.List(); err == nil {
				o.print(output, organizations)
			}
		}
	}

	return err
}

func (o *Options) print(output string, data interface{}) {
	if output == "yaml" {
		if organizationsYaml, err := yaml.Marshal(&data); err == nil {
			fmt.Println(string(organizationsYaml))
		}
	} else {
		fmt.Printf("%v\n", data)
	}
}
