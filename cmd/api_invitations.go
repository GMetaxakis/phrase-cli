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
	initInvitationCreate()
	initInvitationDelete()
	initInvitationResend()
	initInvitationShow()
	initInvitationUpdate()
	initInvitationsList()

	rootCmd.AddCommand(InvitationsApiCmd)
}

var InvitationsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Invitations"),
	Short: "Invitations API",
}

func initInvitationCreate() {
	params := viper.New()
	var InvitationCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationCreate", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Create a new invitation",
		Long:  `Invite a person to an account. Developers and translators need &lt;code&gt;project_ids&lt;/code&gt; and &lt;code&gt;locale_ids&lt;/code&gt; assigned to access them. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))

			invitationCreateParameters := api.InvitationCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &invitationCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", invitationCreateParameters)
			}
			data, api_response, err := client.InvitationsApi.InvitationCreate(auth, accountId, invitationCreateParameters, &localVarOptionals)

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

	InvitationsApiCmd.AddCommand(InvitationCreate)

	AddFlag(InvitationCreate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(InvitationCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(InvitationCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(InvitationCreate.Flags())
}
func initInvitationDelete() {
	params := viper.New()
	var InvitationDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationDelete", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Delete an invitation",
		Long:  `Delete an existing invitation (must not be accepted yet). Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.InvitationsApi.InvitationDelete(auth, accountId, id, &localVarOptionals)

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

	InvitationsApiCmd.AddCommand(InvitationDelete)

	AddFlag(InvitationDelete, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(InvitationDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(InvitationDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(InvitationDelete.Flags())
}
func initInvitationResend() {
	params := viper.New()
	var InvitationResend = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationResend", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Resend an invitation",
		Long:  `Resend the invitation email (must not be accepted yet). Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationResendOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.InvitationsApi.InvitationResend(auth, accountId, id, &localVarOptionals)

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

	InvitationsApiCmd.AddCommand(InvitationResend)

	AddFlag(InvitationResend, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(InvitationResend, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(InvitationResend, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(InvitationResend.Flags())
}
func initInvitationShow() {
	params := viper.New()
	var InvitationShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationShow", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Get a single invitation",
		Long:  `Get details on a single invitation. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.InvitationsApi.InvitationShow(auth, accountId, id, &localVarOptionals)

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

	InvitationsApiCmd.AddCommand(InvitationShow)

	AddFlag(InvitationShow, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(InvitationShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(InvitationShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(InvitationShow.Flags())
}
func initInvitationUpdate() {
	params := viper.New()
	var InvitationUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationUpdate", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "Update an invitation",
		Long:  `Update an existing invitation (must not be accepted yet). The &lt;code&gt;email&lt;/code&gt; cannot be updated. Developers and translators need &lt;code&gt;project_ids&lt;/code&gt; and &lt;code&gt;locale_ids&lt;/code&gt; assigned to access them. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			accountId := params.GetString(helpers.ToSnakeCase("AccountId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			invitationUpdateParameters := api.InvitationUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &invitationUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", invitationUpdateParameters)
			}
			data, api_response, err := client.InvitationsApi.InvitationUpdate(auth, accountId, id, invitationUpdateParameters, &localVarOptionals)

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

	InvitationsApiCmd.AddCommand(InvitationUpdate)

	AddFlag(InvitationUpdate, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(InvitationUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(InvitationUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(InvitationUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(InvitationUpdate.Flags())
}
func initInvitationsList() {
	params := viper.New()
	var InvitationsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("InvitationsList", strings.TrimSuffix("InvitationsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("InvitationsApi", "Api"), "s"))),
		Short: "List invitations",
		Long:  `List invitations for an account. It will also list the accessible resources like projects and locales the invited user has access to. In case nothing is shown the default access from the role is used. Access token scope must include &lt;code&gt;team.manage&lt;/code&gt;.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration(Config)
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.InvitationsListOpts{}

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

			data, api_response, err := client.InvitationsApi.InvitationsList(auth, accountId, &localVarOptionals)

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

	InvitationsApiCmd.AddCommand(InvitationsList)

	AddFlag(InvitationsList, "string", helpers.ToSnakeCase("AccountId"), "", "Account ID", true)
	AddFlag(InvitationsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(InvitationsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(InvitationsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	params.BindPFlags(InvitationsList.Flags())
}
