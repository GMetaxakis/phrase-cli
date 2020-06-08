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
	initStyleguideCreate()
	initStyleguideDelete()
	initStyleguideShow()
	initStyleguideUpdate()
	initStyleguidesList()

	rootCmd.AddCommand(StyleGuidesApiCmd)
}

var StyleGuidesApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("StyleGuides"),
	Short: "StyleGuides API",
}

func initStyleguideCreate() {
	params := viper.New()
	var StyleguideCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideCreate", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Create a style guide",
		Long:  `Create a new style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			styleguideCreateParameters := api.StyleguideCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &styleguideCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", styleguideCreateParameters)
			}
			data, api_response, err := client.StyleGuidesApi.StyleguideCreate(auth, projectId, styleguideCreateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	StyleGuidesApiCmd.AddCommand(StyleguideCreate)

	AddFlag(StyleguideCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(StyleguideCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(StyleguideCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(StyleguideCreate.Flags())
}
func initStyleguideDelete() {
	params := viper.New()
	var StyleguideDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideDelete", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Delete a style guide",
		Long:  `Delete an existing style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.StyleGuidesApi.StyleguideDelete(auth, projectId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	StyleGuidesApiCmd.AddCommand(StyleguideDelete)

	AddFlag(StyleguideDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(StyleguideDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(StyleguideDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(StyleguideDelete.Flags())
}
func initStyleguideShow() {
	params := viper.New()
	var StyleguideShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideShow", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Get a single style guide",
		Long:  `Get details on a single style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.StyleGuidesApi.StyleguideShow(auth, projectId, id, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	StyleGuidesApiCmd.AddCommand(StyleguideShow)

	AddFlag(StyleguideShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(StyleguideShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(StyleguideShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(StyleguideShow.Flags())
}
func initStyleguideUpdate() {
	params := viper.New()
	var StyleguideUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguideUpdate", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "Update a style guide",
		Long:  `Update an existing style guide.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguideUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			styleguideUpdateParameters := api.StyleguideUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &styleguideUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", styleguideUpdateParameters)
			}
			data, api_response, err := client.StyleGuidesApi.StyleguideUpdate(auth, projectId, id, styleguideUpdateParameters, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	StyleGuidesApiCmd.AddCommand(StyleguideUpdate)

	AddFlag(StyleguideUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(StyleguideUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(StyleguideUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(StyleguideUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(StyleguideUpdate.Flags())
}
func initStyleguidesList() {
	params := viper.New()
	var StyleguidesList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("StyleguidesList", strings.TrimSuffix("StyleGuidesApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("StyleGuidesApi", "Api"), "s"))),
		Short: "List style guides",
		Long:  `List all styleguides for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.StyleguidesListOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			data, api_response, err := client.StyleGuidesApi.StyleguidesList(auth, projectId, &localVarOptionals)

			if api_response.StatusCode >= 200 && api_response.StatusCode < 300 {
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

	StyleGuidesApiCmd.AddCommand(StyleguidesList)

	AddFlag(StyleguidesList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(StyleguidesList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(StyleguidesList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(StyleguidesList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(StyleguidesList.Flags())
}
