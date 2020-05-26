package commands

import (
	"github.com/spf13/cobra"
)

func SyncStructure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync-structure",
		Short: "SyncStructure",

		Run: func(cmd *cobra.Command, args []string) {

			// dat, err := ioutil.ReadFile("/home/wrk/Projects/corteza/corteza-server/federation/config/yaml/federation.yaml")

			// if err != nil {
			// 	panic(err)
			// }

			// parser := config.Parser{
			// 	Config: dat,
			// }

			// federatedServerList, err := parser.Source()
			// federatedServer := federatedServerList.FindSource("48241cd7-bd92-4df9-9521-a722d687bf53")

			// if federatedServer == nil {
			// 	panic(errors.New("could not find a federated server config"))
			// }

			// federatedModuleHandleList, err := parser.Modules()

			// if err != nil {
			// 	panic(err)
			// }

			// federatedModuleList := federatedModuleHandleList.FindModules(federatedServer.ID)

			// for _, module := range federatedModuleList {
			// 	// fetch via rest api
			// 	response, err := http.Get(fmt.Sprintf("%s/federation/namespace/145125472630472705/module/%d/record/?perPage=10", federatedServer.URI, module.Source.ID))

			// 	if err != nil {
			// 		fmt.Print(err.Error())
			// 		os.Exit(1)
			// 	}

			// 	responseData, err := ioutil.ReadAll(response.Body)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}

			// 	fmt.Println(string(responseData))
			// }

		},
	}

	return cmd
}
