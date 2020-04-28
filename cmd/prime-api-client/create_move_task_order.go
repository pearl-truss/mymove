package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	openapi "github.com/go-openapi/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	movetaskorderclient "github.com/transcom/mymove/pkg/gen/supportclient/move_task_order"
)

// initCreateMTOFlags initializes flags.
func initCreateMTOFlags(flag *pflag.FlagSet) {
	flag.String(FilenameFlag, "", "Path to the file with the payment request JSON payload")

	flag.SortFlags = false
}

// checkCreateMTOConfig checks the args.
func checkCreateMTOConfig(v *viper.Viper, args []string) error {
	err := CheckRootConfig(v)
	if err != nil {
		return err
	}

	if v.GetString(FilenameFlag) == "" && (len(args) < 1 || len(args) > 0 && !containsDash(args)) {
		return errors.New("create-payment-request expects a file to be passed in")
	}

	return nil
}

// createMTO sends a CreateMoveTaskOrder request to the support endpoint
func createMTO(cmd *cobra.Command, args []string) error {
	v := viper.New()

	// Create the logger - remove the prefix and any datetime data
	logger := log.New(os.Stdout, "", log.LstdFlags)

	errParseFlags := ParseFlags(cmd, v, args)
	if errParseFlags != nil {
		return errParseFlags
	}

	// Check the config before talking to the CAC
	err := checkCreateMTOConfig(v, args)
	if err != nil {
		logger.Fatal(err)
	}

	// Decode json from file that was passed in
	filename := v.GetString(FilenameFlag)
	var createMTOParams movetaskorderclient.CreateMoveTaskOrderParams
	err = decodeJSONFileToPayload(filename, containsDash(args), &createMTOParams)
	if err != nil {
		logger.Fatal(err)
	}
	createMTOParams.SetTimeout(time.Second * 30)

	// Create the client and open the cacStore
	gateway, cacStore, errCreateClient := CreateSupportClient(v)
	if errCreateClient != nil {
		return errCreateClient
	}
	// Defer closing the store until after the API call has completed
	if cacStore != nil {
		defer cacStore.Close()
	}

	// Make the API Call
	resp, err := gateway.MoveTaskOrder.CreateMoveTaskOrder(&createMTOParams)
	if err != nil {
		// If you see an error like "unknown error (status 422)", it means
		// we hit a completely unhandled error that we should handle.
		// We should be enabling said error in the endpoint in swagger.
		// 422 for example is an Unprocessable Entity and is returned by the swagger
		// validation before it even hits the handler.
		if _, ok := err.(*openapi.APIError); ok {
			apiErr := err.(*openapi.APIError).Response.(openapi.ClientResponse)
			logger.Fatal(fmt.Sprintf("%s: %s", err, apiErr.Message()))
		}
		// If it is a handled error, we should be able to pull out the payload here
		data, _ := json.Marshal(err)
		fmt.Printf("%s", data)
		return nil
	}

	// Get the successful response payload and convert to json for output
	payload := resp.GetPayload()
	if payload != nil {
		payload, errJSONMarshall := json.Marshal(payload)
		if errJSONMarshall != nil {
			logger.Fatal(errJSONMarshall)
		}
		fmt.Println(string(payload))
	} else {
		logger.Fatal(resp.Error())
	}

	return nil

}
