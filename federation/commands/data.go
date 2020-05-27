package commands

import (
	"context"
	"encoding/json"
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
	"github.com/cortezaproject/corteza-server/federation/rest"
	"github.com/cortezaproject/corteza-server/federation/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

const (
	FederationRecordStatusCreated FederationRecordStatus = iota
	FederationRecordStatusUpdated
	FederationRecordStatusDeleted
	FederationRecordStatusUnknown
)

type (
	FederationRecordStatus int

	FederationMappingContext struct {
		Status            FederationRecordStatus
		SourceRecord      *types.FederatedRecord
		DestinationRecord *types.FederatedRecord
		Mapping           interface{}
	}

	federationMapper interface {
		GetStatus(record *types.FederatedRecord) error
		Transform(context FederationMappingContext) error
	}

	FMapper struct{}
)

func (fm *FMapper) GetStatus(record *types.FederatedRecord) (FederationRecordStatus, error) {
	createdAt := !record.CreatedAt.IsZero()
	updatedAt := record.UpdatedAt != nil && !record.UpdatedAt.IsZero()
	deletedAt := record.DeletedAt != nil && !record.DeletedAt.IsZero()

	switch {
	case createdAt && !updatedAt && !deletedAt:
		return FederationRecordStatusCreated, nil
	case createdAt && updatedAt && !deletedAt:
		return FederationRecordStatusUpdated, nil
	case deletedAt:
		return FederationRecordStatusDeleted, nil
	}

	return FederationRecordStatusUnknown, errors.New("this should not happen")
}

func (fm *FMapper) Transform(context FederationMappingContext) (*types.FederatedRecord, error) {
	// use status of the record
	// fetch record from db
	// merge with mapping data
	// return federatedrecord
	return nil, nil
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
			ctx := context.TODO()

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

			federationRestService := rest.Record{}.New()

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

			fmapper := &FMapper{}

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

					// unmarshal to FederatedRecord
					responseRecord := &types.FederatedRecord{}

					err = json.Unmarshal(responseData, &responseRecord)

					recordSet, err := federationRestService.DecodeFilterPayload(ctx, responseData)

					log.Printf("Fetched response successfuly")

					if len(recordSet) == 0 {
						log.Printf("No results to crunch, skipping")
						break
					}

					for _, frec := range recordSet {
						recordStatus, _ := fmapper.GetStatus(frec)
						mappingContext := &FederationMappingContext{
							Status:            recordStatus,
							SourceRecord:      frec,
							DestinationRecord: frec,
							Mapping:           federatedModuleList,
						}

						spew.Dump(fmapper.Transform(*mappingContext))
					}

					// spew.Dump(recordSet)

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
