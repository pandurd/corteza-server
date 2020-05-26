package commands

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cortezaproject/corteza-server/federation/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func SyncRecords(ctx context.Context, moduleList []*config.ConfigStructureMapped) {

}

func SyncData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync-data",
		Short: "sd",

		Run: func(cmd *cobra.Command, args []string) {
			// finish := make(chan bool)
			// get servers
			// go through mapped: and fetch data

			// ctx, cancel := context.WithCancel(context.Background())

			var lastSynced int64 = 0
			spew.Dump(lastSynced)

			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			go func() {
				<-c
				log.Println("Killed, finishing up")
				os.Exit(1)
			}()

			dat, err := ioutil.ReadFile("/home/wrk/Projects/corteza/corteza-server/federation/config/yaml/federation.yaml")

			if err != nil {
				panic(err)
			}

			parser := config.Parser{
				Config: dat,
			}

			settings, err := parser.Settings()

			if err != nil {
				panic(err)
			}

			spew.Dump("SET", settings.Settings.Sync.Minutes(), fmt.Sprintf("%0.f", settings.Settings.Sync.Minutes()))

			federatedServerList, err := parser.Source()
			federatedServer := federatedServerList.FindSource("48241cd7-bd92-4df9-9521-a722d687bf53")

			if federatedServer == nil {
				panic(errors.New("could not find a federated server config"))
			}

			federatedModuleHandleList, err := parser.Structure()

			if err != nil {
				panic(err)
			}

			ticker := time.NewTicker(settings.Settings.Sync)
			defer ticker.Stop()

			federatedModuleList := federatedModuleHandleList.FindModules(federatedServer.ID)

			for {
				for _, module := range federatedModuleList {
					// fetch via rest api
					// use timestamp cutoff
					query := "perPage=1000"

					if lastSynced > 0 {
						query = fmt.Sprintf("%s&lastSynced=%d", query, lastSynced)
					}

					url := &url.URL{
						Scheme:   "http",
						Host:     federatedServer.URI,
						Path:     fmt.Sprintf("federation/namespace/%d/module/%d/record/", module.Source.Namespace, module.Source.ID),
						RawQuery: query,
					}
					spew.Dump(url.String())
					response, err := http.Get(url.String())

					if err != nil {
						fmt.Println("could not get data, waiting", err)
						break
					}

					responseData, err := ioutil.ReadAll(response.Body)

					if err != nil {
						log.Fatal(err)
						ticker.Stop()
						return
					}

					log.Printf("Fetched response successfuly", string(responseData))
					lastSynced = time.Now().Unix()
				}

				select {
				case t := <-ticker.C:
					fmt.Println("Current time: ", t)
				}
			}
		},
	}

	return cmd
}
