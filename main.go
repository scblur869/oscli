package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"src/local/oscli/admin"
	"src/local/oscli/cat"
	"src/local/oscli/configure"
	"src/local/oscli/structs"

	color "github.com/logrusorgru/aurora"
)

func printUsage() {
	fmt.Println(color.Bold(color.Cyan("usage:")))
	fmt.Println()
	fmt.Println(color.Bold(color.BrightWhite("admin \n\t ls-role-mapping | ls-tenants | ls-security ")))
	fmt.Println(color.Bold(color.BrightWhite("configure \n\t { allows for basic auth and elastic host url }")))
	fmt.Println(color.Bold(color.BrightWhite("role-mapping \n\t new -name=role_name -user=elasticsearch_user -backend-role=kibana_user -host={nil or host match}")))
	fmt.Println(color.Bold(color.BrightWhite("indice \n\t delete -name=indexName")))
	fmt.Println(color.Bold(color.BrightWhite("template \n\t view -name=templateName")))
	fmt.Println(color.Bold(color.BrightWhite("role \n\t new -name=test_role -clusterperms=indices_monitor,cluster_ops -index=\"movies-*\" -indexperm=read -tenant=sales,marketing -tenantperms=kibana_all_read")))
	fmt.Println(color.Bold(color.BrightWhite("cat \n\t cluster-info | stats | nodes | templates | indices | shards | health | disk | recovery | master | count | field_data | alias")))
	fmt.Println(color.Bold(color.Yellow("help: { list this message :-) }")))

}

func main() {

	catCommand := flag.NewFlagSet("cat", flag.ExitOnError)
	adminCommand := flag.NewFlagSet("admin", flag.ExitOnError)
	configureCommand := flag.NewFlagSet("configure", flag.ExitOnError)
	rolemappingCommand := flag.NewFlagSet("role-mapping", flag.ExitOnError)
	idxManagementCommand := flag.NewFlagSet("indice", flag.ExitOnError)
	roleCommand := flag.NewFlagSet("role", flag.ExitOnError)
	idxTemplateCommand := flag.NewFlagSet("template", flag.ExitOnError)
	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)
	rmUserFlag := rolemappingCommand.String("user", "", "name of user to apply role mapping")
	rmBackendRoleFlag := rolemappingCommand.String("backend-role", "", "some backend role")
	rmHostFlag := rolemappingCommand.String("host", "", "host where the role-mapping has access")
	rmNameFlag := rolemappingCommand.String("name", "", "name of the role")
	idxStringFlag := idxManagementCommand.String("name", "", "name or string match of indice(s)")
	idxTemplateFlag := idxTemplateCommand.String("name", "", "name of the index template to view")
	rClusterPermsFlag := roleCommand.String("clusterperms", "", "predefined cluster permissions")
	rIndexPattFlag := roleCommand.String("index", "", "define index pattern access")
	rIndexPermFlag := roleCommand.String("indexperm", "", "define index permission access")
	rTenantFlag := roleCommand.String("tenant", "", "define tenant for access ..ie Sales, Marketing")
	rTenantPermFlag := roleCommand.String("tenantPerm", "", "define tenant permission.. ie KIBANA_READ_ALL")
	rNameFlag := roleCommand.String("name", "", "name of the role you are creating")
	flag.Parse()

	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)
	}
	switch os.Args[1] {
	//validate
	case "configure":
		configureCommand.Parse(os.Args[1:])
		if configureCommand.Parsed() {
			configure.Setup()
		}
	case "admin":
		if len(os.Args) > 2 {
			adminCommand.Parse(os.Args[2:])
			switch adminCommand.Arg(0) {
			case "ls-role-mapping":
				admin.ListRoleMapping()
			case "ls-tenants":
				admin.ListTenants()
			case "ls-security":
				admin.ListSecurityConfig()
			default:
				fmt.Println(color.Bold(color.BrightWhite("admin \n\t ls-role-mapping | ls-tenants | ls-security ")))
				os.Exit(1)
			}
		}
	// ./oscli role-mapping new -user=someuser -host={can be left bank}  -backend-role=kibana_user
	case "role-mapping":
		if len(os.Args) > 2 {
			rolemappingCommand.Parse(os.Args[2:])
			switch rolemappingCommand.Arg(0) {
			case "new":
				rolemappingCommand.Parse(os.Args[3:])
				var actionPayload structs.RoleMapping
				name := *rmNameFlag
				br := strings.Split(string(*rmBackendRoleFlag), ",")
				us := strings.Split(string(*rmUserFlag), ",")
				ho := strings.Split(string(*rmHostFlag), ",")
				actionPayload.User = us
				actionPayload.Host = ho
				actionPayload.Backend_roles = br

				admin.AddRoleMapping(actionPayload, name)

			default:
				fmt.Println(color.Bold(color.BrightWhite("role-mapping \n\t new -name=role_name -user=elasticsearch_user -backend-role=kibana_user -host={nil or host match}")))
				os.Exit(1)
			}
		}
	case "role":
		if len(os.Args) > 2 {
			roleCommand.Parse(os.Args[2:])
			switch roleCommand.Arg(0) {
			case "new":
				roleCommand.Parse(os.Args[3:])
				emptySlice := make([]string, 0)
				var role structs.Role
				var idx structs.Index_Permissions
				var tenant structs.Tenant_Permissions
				cp := strings.Split(string(*rClusterPermsFlag), ",")
				ip := strings.Split(string(*rIndexPattFlag), ",")
				aa := strings.Split(string(*rIndexPermFlag), ",")
				tp := strings.Split(string(*rTenantFlag), ",")
				ta := strings.Split(string(*rTenantPermFlag), ",")
				role.ClusterPermissions = cp
				idx.IndexPatterns = ip
				idx.AllowedActions = aa
				idx.DLS = ""
				idx.FLS = emptySlice
				idx.MaskedFields = emptySlice
				role.IndexPermissions = append(role.IndexPermissions, idx)
				tenant.TenantPatterns = tp
				tenant.AllowedActions = ta
				role.TenantPermissions = append(role.TenantPermissions, tenant)
				admin.NewRole(role, *rNameFlag)

			default:
				fmt.Println(color.Bold(color.BrightWhite("role new -name=test_role -clusterperms=indices_monitor,cluster_ops -index=\"movies-*\" -indexperm=read -tenant=sales,marketing -tenantperms=kibana_all_read")))
				os.Exit(1)
			}
		}
	case "indice":
		if len(os.Args) > 2 {
			idxManagementCommand.Parse(os.Args[2:])
			switch idxManagementCommand.Arg(0) {
			case "delete":
				idxManagementCommand.Parse(os.Args[3:])
				if *idxStringFlag != "" {
					admin.DeleteIndexRequest(*idxStringFlag)
				} else {
					fmt.Println("delete action must include a name or string match")
					os.Exit(1)
				}

			default:
				fmt.Println(color.Bold(color.BrightWhite("indice \n\t delete -name=stringmatch")))
				os.Exit(1)
			}
		}
	case "template":
		if len(os.Args) > 2 {
			idxTemplateCommand.Parse(os.Args[2:])
			switch idxTemplateCommand.Arg(0) {
			case "view":
				idxTemplateCommand.Parse(os.Args[3:])
				if *idxTemplateFlag != "" {
					admin.GetTemplateDetails(*idxTemplateFlag)
				} else {
					fmt.Println("must include a name or string match with -name=")
					os.Exit(1)
				}

			default:
				fmt.Println(color.Bold(color.BrightWhite("template \n\t view -name=stringmatch")))
				os.Exit(1)
			}
		}
	case "cat":
		if len(os.Args) > 2 {
			catCommand.Parse(os.Args[2:])
			switch catCommand.Arg(0) {
			case "nodes":
				cat.GetNodesInfo()
			case "stats":
				cat.GetClusterCurrentStats()
			case "health":
				cat.GetHealthInfo()
			case "indices":
				cat.GetIndexList()
			case "shards":
				cat.GetShardsInfo()
			case "disk":
				cat.GetDiskAllocation()
			case "recovery":
				cat.GetRecoverStatus()
			case "master":
				cat.GetCurrentMaster()
			case "count":
				cat.GetDocumentCount()
			case "field_data":
				cat.GetFieldDataUsage()
			case "aliases":
				cat.GetAlias()
			case "templates":
				cat.GetIndexTemplates()
			case "cluster-info":
				cat.GetClusterDefaults()
			default:
				fmt.Println(color.Bold(color.BrightWhite("cat  \n\t cluster-info | stats | nodes | templates | indices | shards | health | disk | recovery | master | count | field_data | alias")))
				os.Exit(1)
			}
		} else {
			fmt.Println(color.Bold(color.BrightWhite("cat  \n\t cluster-info | stats | nodes | templates | indices | shards | health | disk | recovery | master | count | field_data | alias")))
			os.Exit(1)
		}
	case "help":
		printUsage()
		helpCommand.PrintDefaults()
	default:
		os.Exit(1)
	}
}
