package shell

import (
	"yunion.io/x/onecloud/pkg/util/azure"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type SecurityGroupListOptions struct {
		Classic bool `help:"List classic secgroups"`
		Limit   int  `help:"page size"`
		Offset  int  `help:"page offset"`
	}
	shellutils.R(&SecurityGroupListOptions{}, "security-group-list", "List security group", func(cli *azure.SRegion, args *SecurityGroupListOptions) error {
		if args.Classic {
			secgrps, err := cli.GetClassicSecurityGroups()
			if err != nil {
				return err
			}
			printList(secgrps, len(secgrps), args.Offset, args.Limit, []string{})
			return nil
		}
		secgrps, err := cli.GetSecurityGroups()
		if err != nil {
			return err
		}
		printList(secgrps, len(secgrps), args.Offset, args.Limit, []string{})
		return nil
	})

	type SecurityGroupOptions struct {
		ID string `help:"ID or name of security group"`
	}
	shellutils.R(&SecurityGroupOptions{}, "security-group-show", "Show details of a security group", func(cli *azure.SRegion, args *SecurityGroupOptions) error {
		if secgrp, err := cli.GetSecurityGroupDetails(args.ID); err != nil {
			return err
		} else {
			printObject(secgrp)
			return nil
		}
	})

	shellutils.R(&SecurityGroupOptions{}, "security-group-rule-list", "List security group rules", func(cli *azure.SRegion, args *SecurityGroupOptions) error {
		if secgroup, err := cli.GetSecurityGroupDetails(args.ID); err != nil {
			return err
		} else if rules, err := secgroup.GetRules(); err != nil {
			return err
		} else {
			printList(rules, len(rules), 0, 30, []string{})
			return nil
		}
	})

	type SecurityGroupCreateOptions struct {
		NAME  string `help:"Security Group name"`
		TagId string `help:"Add a id tag to secgroup"`
	}

	shellutils.R(&SecurityGroupCreateOptions{}, "security-group-create", "Create security group", func(cli *azure.SRegion, args *SecurityGroupCreateOptions) error {
		if secgrp, err := cli.CreateSecurityGroup(args.NAME, args.TagId); err != nil {
			return err
		} else {
			printObject(secgrp)
			return nil
		}
	})
}