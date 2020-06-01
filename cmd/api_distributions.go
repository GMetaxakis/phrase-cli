package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	helpers "github.com/phrase/phrase-cli/helpers"
	api "github.com/phrase/phrase-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	initDistributionCreate()
	initDistributionDelete()
	initDistributionShow()
	initDistributionUpdate()
	initDistributionsList()

	rootCmd.AddCommand(DistributionsApiCmd)
}

var DistributionsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Distributions"),
	Short: "Distributions API",
}

func initDistributionCreate() {
	params := viper.New()
	var DistributionCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionCreate", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Create a distribution",
		Long:  `Create a new distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			distributionCreateParameters := api.DistributionCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &distributionCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", distributionCreateParameters)
			}
			data, api_response, err := client.DistributionsApi.DistributionCreate(auth, accountId, distributionCreateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	DistributionsApiCmd.AddCommand(DistributionCreate)

	AddFlag(DistributionCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(DistributionCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(DistributionCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(DistributionCreate.Flags())
}
func initDistributionDelete() {
	params := viper.New()
	var DistributionDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionDelete", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Delete a distribution",
		Long:  `Delete an existing distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.DistributionsApi.DistributionDelete(auth, accountId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	DistributionsApiCmd.AddCommand(DistributionDelete)

	AddFlag(DistributionDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(DistributionDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(DistributionDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(DistributionDelete.Flags())
}
func initDistributionShow() {
	params := viper.New()
	var DistributionShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionShow", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Get a single distribution",
		Long:  `Get details on a single distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.DistributionsApi.DistributionShow(auth, accountId, id, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	DistributionsApiCmd.AddCommand(DistributionShow)

	AddFlag(DistributionShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(DistributionShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(DistributionShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(DistributionShow.Flags())
}
func initDistributionUpdate() {
	params := viper.New()
	var DistributionUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionUpdate", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "Update a distribution",
		Long:  `Update an existing distribution.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			distributionUpdateParameters := api.DistributionUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &distributionUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", distributionUpdateParameters)
			}
			data, api_response, err := client.DistributionsApi.DistributionUpdate(auth, accountId, id, distributionUpdateParameters, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	DistributionsApiCmd.AddCommand(DistributionUpdate)

	AddFlag(DistributionUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(DistributionUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(DistributionUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(DistributionUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(DistributionUpdate.Flags())
}
func initDistributionsList() {
	params := viper.New()
	var DistributionsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("DistributionsList", strings.TrimSuffix("DistributionsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("DistributionsApi", "Api"), "s"))),
		Short: "List distributions",
		Long:  `List all distributions for the given account.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.DistributionsListOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			data, api_response, err := client.DistributionsApi.DistributionsList(auth, accountId, &localVarOptionals)

			if api_response.StatusCode == 200 {
				jsonBuf, jsonErr := json.MarshalIndent(data, "", " ")
				if jsonErr != nil {
					fmt.Printf("%v\n", data)
					HandleError(err)
				}

				fmt.Printf("%s\n", string(jsonBuf))
			}
			if err != nil {
				HandleError(err)
			}

			if Config.Debug {
				fmt.Printf("%+v\n", api_response) // &{Response:0xc00011ccf0 NextPage:2 FirstPage:1 LastPage:4 Rate:{Limit:1000 Remaining:998 Reset:2020-04-25 00:35:00 +0200 CEST}}
			}
		},
	}

	DistributionsApiCmd.AddCommand(DistributionsList)

	AddFlag(DistributionsList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(DistributionsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(DistributionsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(DistributionsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(DistributionsList.Flags())
}
