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
	initJobComplete()
	initJobCreate()
	initJobDelete()
	initJobKeysCreate()
	initJobKeysDelete()
	initJobReopen()
	initJobShow()
	initJobStart()
	initJobUpdate()
	initJobsList()

	rootCmd.AddCommand(JobsApiCmd)
}

var JobsApiCmd = &cobra.Command{
	Use:   helpers.ToSnakeCase("Jobs"),
	Short: "Jobs API",
}

func initJobComplete() {
	params := viper.New()
	var JobComplete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobComplete", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Complete a job",
		Long:  `Mark a job as completed.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobCompleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobCompleteParameters := api.JobCompleteParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobCompleteParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobCompleteParameters)
			}
			data, api_response, err := client.JobsApi.JobComplete(auth, projectId, id, jobCompleteParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobComplete)

	AddFlag(JobComplete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobComplete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobComplete, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobComplete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobComplete.Flags())
}
func initJobCreate() {
	params := viper.New()
	var JobCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobCreate", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Create a job",
		Long:  `Create a new job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			jobCreateParameters := api.JobCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobCreateParameters)
			}
			data, api_response, err := client.JobsApi.JobCreate(auth, projectId, jobCreateParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobCreate)

	AddFlag(JobCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobCreate.Flags())
}
func initJobDelete() {
	params := viper.New()
	var JobDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobDelete", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Delete a job",
		Long:  `Delete an existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.JobsApi.JobDelete(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobDelete)

	AddFlag(JobDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(JobDelete.Flags())
}
func initJobKeysCreate() {
	params := viper.New()
	var JobKeysCreate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobKeysCreate", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Add keys to job",
		Long:  `Add multiple keys to a existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobKeysCreateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobKeysCreateParameters := api.JobKeysCreateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobKeysCreateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobKeysCreateParameters)
			}
			data, api_response, err := client.JobsApi.JobKeysCreate(auth, projectId, id, jobKeysCreateParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobKeysCreate)

	AddFlag(JobKeysCreate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobKeysCreate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobKeysCreate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobKeysCreate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobKeysCreate.Flags())
}
func initJobKeysDelete() {
	params := viper.New()
	var JobKeysDelete = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobKeysDelete", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Remove keys from job",
		Long:  `Remove multiple keys from existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobKeysDeleteOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.JobsApi.JobKeysDelete(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobKeysDelete)

	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobKeysDelete, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(JobKeysDelete, "string", "data", "d", "payload in JSON format", false)

	params.BindPFlags(JobKeysDelete.Flags())
}
func initJobReopen() {
	params := viper.New()
	var JobReopen = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobReopen", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Reopen a job",
		Long:  `Mark a job as uncompleted.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobReopenOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobReopenParameters := api.JobReopenParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobReopenParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobReopenParameters)
			}
			data, api_response, err := client.JobsApi.JobReopen(auth, projectId, id, jobReopenParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobReopen)

	AddFlag(JobReopen, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobReopen, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobReopen, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobReopen, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobReopen.Flags())
}
func initJobShow() {
	params := viper.New()
	var JobShow = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobShow", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Get a single job",
		Long:  `Get details on a single job for a given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobShowOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			data, api_response, err := client.JobsApi.JobShow(auth, projectId, id, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobShow)

	AddFlag(JobShow, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobShow, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobShow, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobShow, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	params.BindPFlags(JobShow.Flags())
}
func initJobStart() {
	params := viper.New()
	var JobStart = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobStart", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Start a job",
		Long:  `Starts an existing job in state draft.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobStartOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobStartParameters := api.JobStartParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobStartParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobStartParameters)
			}
			data, api_response, err := client.JobsApi.JobStart(auth, projectId, id, jobStartParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobStart)

	AddFlag(JobStart, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobStart, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobStart, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobStart, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobStart.Flags())
}
func initJobUpdate() {
	params := viper.New()
	var JobUpdate = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobUpdate", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "Update a job",
		Long:  `Update an existing job.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobUpdateOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))
			id := params.GetString(helpers.ToSnakeCase("Id"))

			jobUpdateParameters := api.JobUpdateParameters{}
			if err := json.Unmarshal([]byte(params.GetString("data")), &jobUpdateParameters); err != nil {
				HandleError(err)
			}
			if Config.Debug {
				fmt.Printf("%+v\n", jobUpdateParameters)
			}
			data, api_response, err := client.JobsApi.JobUpdate(auth, projectId, id, jobUpdateParameters, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobUpdate)

	AddFlag(JobUpdate, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobUpdate, "string", helpers.ToSnakeCase("Id"), "", "ID", true)
	AddFlag(JobUpdate, "string", "data", "d", "payload in JSON format", true)

	AddFlag(JobUpdate, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	params.BindPFlags(JobUpdate.Flags())
}
func initJobsList() {
	params := viper.New()
	var JobsList = &cobra.Command{
		// this weird approach is due to mustache template limitations
		Use:   helpers.ToSnakeCase(strings.TrimPrefix(strings.TrimPrefix("JobsList", strings.TrimSuffix("JobsApi", "Api")), strings.TrimSuffix(strings.TrimSuffix("JobsApi", "Api"), "s"))),
		Short: "List jobs",
		Long:  `List all jobs for the given project.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{
				Key:    Config.Credentials.Token,
				Prefix: "token",
			})

			cfg := api.NewConfiguration()
			client := api.NewAPIClient(cfg)

			localVarOptionals := api.JobsListOpts{}

			if params.IsSet(helpers.ToSnakeCase("xPhraseAppOTP")) {
				localVarOptionals.XPhraseAppOTP = optional.NewString(params.GetString(helpers.ToSnakeCase("XPhraseAppOTP")))
			}
			if params.IsSet(helpers.ToSnakeCase("page")) {
				localVarOptionals.Page = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("Page")))
			}
			if params.IsSet(helpers.ToSnakeCase("perPage")) {
				localVarOptionals.PerPage = optional.NewInt32(params.GetInt32(helpers.ToSnakeCase("PerPage")))
			}
			if params.IsSet(helpers.ToSnakeCase("branch")) {
				localVarOptionals.Branch = optional.NewString(params.GetString(helpers.ToSnakeCase("Branch")))
			}
			if params.IsSet(helpers.ToSnakeCase("ownedBy")) {
				localVarOptionals.OwnedBy = optional.NewString(params.GetString(helpers.ToSnakeCase("OwnedBy")))
			}
			if params.IsSet(helpers.ToSnakeCase("assignedTo")) {
				localVarOptionals.AssignedTo = optional.NewString(params.GetString(helpers.ToSnakeCase("AssignedTo")))
			}
			if params.IsSet(helpers.ToSnakeCase("state")) {
				localVarOptionals.State = optional.NewString(params.GetString(helpers.ToSnakeCase("State")))
			}

			projectId := params.GetString(helpers.ToSnakeCase("ProjectId"))

			data, api_response, err := client.JobsApi.JobsList(auth, projectId, &localVarOptionals)

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

	JobsApiCmd.AddCommand(JobsList)

	AddFlag(JobsList, "string", helpers.ToSnakeCase("ProjectId"), "", "Project ID", true)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("XPhraseAppOTP"), "", "Two-Factor-Authentication token (optional)", false)
	AddFlag(JobsList, "int32", helpers.ToSnakeCase("Page"), "", "Page number", false)
	AddFlag(JobsList, "int32", helpers.ToSnakeCase("PerPage"), "", "allows you to specify a page size up to 100 items, 10 by default", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("Branch"), "", "specify the branch to use", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("OwnedBy"), "", "filter by user owning job", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("AssignedTo"), "", "filter by user assigned to job", false)
	AddFlag(JobsList, "string", helpers.ToSnakeCase("State"), "", "filter by state of job Valid states are <code>draft</code>, <code>in_progress</code>, <code>completed</code>", false)
	params.BindPFlags(JobsList.Flags())
}
