// Code generated by openapi-cli-generator. DO NOT EDIT.
// See https://github.com/danielgtaylor/openapi-cli-generator

package main

import (
	"fmt"
	"strings"
	

	"github.com/danielgtaylor/openapi-cli-generator/cli"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/h2non/gentleman.v2"
)







var openapi.YmlSubcommand bool

func openapi.YmlServers() []map[string]string {
	return []map[string]string{
		
			map[string]string{
				"description": "",
				"url": "https://petstore3.swagger.io",
			},
		
	}
}


	// Openapi.YmlCreateMessage Create Message
	func Openapi.YmlCreateMessage(params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
		handlerPath := "create-message"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/messages"

		req := cli.Client.Post().URL(url)

		

		
			if body != "" {
				req = req.AddHeader("Content-Type", "application/json").BodyString(body)
			}
		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded map[string]interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after.(map[string]interface{})
		}

		return resp, decoded, nil
	}

	// Openapi.YmlListMessages List messages
	func Openapi.YmlListMessages(params *viper.Viper) (*gentleman.Response, interface{}, error) {
		handlerPath := "list-messages"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/messages"

		req := cli.Client.Get().URL(url)

		
			paramInbox := params.GetString("inbox")
			if paramInbox != "" {
					req = req.AddQuery("inbox", fmt.Sprintf("%v", paramInbox))
			}
			paramStatus := params.GetString("status")
			if paramStatus != "" {
					req = req.AddQuery("status", fmt.Sprintf("%v", paramStatus))
			}

		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after
		}

		return resp, decoded, nil
	}

	// Openapi.YmlDeleteMessage Deletes a message
	func Openapi.YmlDeleteMessage(paramMessageid string, params *viper.Viper) (*gentleman.Response, interface{}, error) {
		handlerPath := "delete-message"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/messages/{messageId}"
				url = strings.Replace(url, "{messageId}", paramMessageid, 1)

		req := cli.Client.Delete().URL(url)

		
			
		
			paramApiKey := params.GetString("api-key")
			if paramApiKey != "" {
					req = req.AddHeader("api_key", fmt.Sprintf("%v", paramApiKey))
			}

		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after
		}

		return resp, decoded, nil
	}

	// Openapi.YmlGetMessage Get message
	func Openapi.YmlGetMessage(paramMessageid string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
		handlerPath := "get-message"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/messages/{messageId}"
				url = strings.Replace(url, "{messageId}", paramMessageid, 1)

		req := cli.Client.Get().URL(url)

		
			
		

		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded map[string]interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after.(map[string]interface{})
		}

		return resp, decoded, nil
	}

	// Openapi.YmlCancelMessage Cancel a message
	func Openapi.YmlCancelMessage(paramMessageid string, params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
		handlerPath := "cancel-message"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/messages/{messageId}/cancel"
				url = strings.Replace(url, "{messageId}", paramMessageid, 1)

		req := cli.Client.Post().URL(url)

		
			
		

		
			if body != "" {
				req = req.AddHeader("Content-Type", "").BodyString(body)
			}
		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded map[string]interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after.(map[string]interface{})
		}

		return resp, decoded, nil
	}

	// Openapi.YmlSendMessage Sends a message
	func Openapi.YmlSendMessage(paramMessageid string, params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
		handlerPath := "send-message"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/messages/{messageId}/send"
				url = strings.Replace(url, "{messageId}", paramMessageid, 1)

		req := cli.Client.Post().URL(url)

		
			
		

		
			if body != "" {
				req = req.AddHeader("Content-Type", "").BodyString(body)
			}
		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded map[string]interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after.(map[string]interface{})
		}

		return resp, decoded, nil
	}

	// Openapi.YmlListNetworks List networks
	func Openapi.YmlListNetworks(params *viper.Viper) (*gentleman.Response, interface{}, error) {
		handlerPath := "list-networks"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/networks"

		req := cli.Client.Get().URL(url)

		
			paramCountryCode := params.GetString("country-code")
			if paramCountryCode != "" {
					req = req.AddQuery("country_code", fmt.Sprintf("%v", paramCountryCode))
			}

		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after
		}

		return resp, decoded, nil
	}

	// Openapi.YmlGetNetwork Get network
	func Openapi.YmlGetNetwork(paramNetworkid string, params *viper.Viper) (*gentleman.Response, map[string]interface{}, error) {
		handlerPath := "get-network"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/networks/{networkId}"
				url = strings.Replace(url, "{networkId}", paramNetworkid, 1)

		req := cli.Client.Get().URL(url)

		
			
		
			paramCountryCode := params.GetInt64("country-code")
			if paramCountryCode != 0 {
					req = req.AddQuery("country_code", fmt.Sprintf("%v", paramCountryCode))
			}

		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded map[string]interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after.(map[string]interface{})
		}

		return resp, decoded, nil
	}

	// Openapi.YmlGetPricing List network rates
	func Openapi.YmlGetPricing(paramNetworkid string, params *viper.Viper) (*gentleman.Response, interface{}, error) {
		handlerPath := "get-pricing"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/networks/{networkId}/pricing"
				url = strings.Replace(url, "{networkId}", paramNetworkid, 1)

		req := cli.Client.Get().URL(url)

		
			
		

		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after
		}

		return resp, decoded, nil
	}

	// Openapi.YmlCreatePricing Create network price
	func Openapi.YmlCreatePricing(paramNetworkid string, params *viper.Viper, body string) (*gentleman.Response, map[string]interface{}, error) {
		handlerPath := "create-pricing"
		if openapi.YmlSubcommand {
			handlerPath = "openapi.yml " + handlerPath
		}

		server := viper.GetString("server")
		if server == "" {
			server = openapi.YmlServers()[viper.GetInt("server-index")]["url"]
		}

		url := server+"/v1/sms/networks/{networkId}/pricing"
				url = strings.Replace(url, "{networkId}", paramNetworkid, 1)

		req := cli.Client.Put().URL(url)

		
			
		

		
			if body != "" {
				req = req.AddHeader("Content-Type", "").BodyString(body)
			}
		

		cli.HandleBefore(handlerPath, params, req)

		resp, err := req.Do()
		if err != nil {
			return nil, nil, errors.Wrap(err, "Request failed")
		}

		var decoded map[string]interface{}

		if resp.StatusCode < 400 {
			if err := cli.UnmarshalResponse(resp, &decoded); err != nil {
				return nil, nil, errors.Wrap(err, "Unmarshalling response failed")
			}
		} else {
			return nil, nil, errors.Errorf("HTTP %d: %s", resp.StatusCode, resp.String())
		}

		after := cli.HandleAfter(handlerPath, params, resp, decoded)
		if after != nil {
			decoded = after.(map[string]interface{})
		}

		return resp, decoded, nil
	}




func openapi.YmlRegister(subcommand bool) {
	root := cli.Root

	if subcommand {
		root = &cobra.Command{
			Use: "openapi.yml",
			Short: "Swagger Petstore - OpenAPI 3.0",
			Long: cli.Markdown("This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about\nSwagger at [https://swagger.io](https://swagger.io). In the third iteration of the pet store, we've switched to the design first approach!\nYou can now help us improve the API whether it's by making changes to the definition itself or to the code.\nThat way, with time, we can improve the API in general, and expose some of the new features in OAS3.\n\n_If you're looking for the Swagger 2.0/OAS 2.0 version of Petstore, then click [here](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml). Alternatively, you can load via the `Edit > Load Petstore OAS 2.0` menu option!_\n\nSome useful links:\n- [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)\n- [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)"),
		}
		openapi.YmlSubcommand = true
	} else {
		cli.Root.Short = "Swagger Petstore - OpenAPI 3.0"
		cli.Root.Long = cli.Markdown("This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about\nSwagger at [https://swagger.io](https://swagger.io). In the third iteration of the pet store, we've switched to the design first approach!\nYou can now help us improve the API whether it's by making changes to the definition itself or to the code.\nThat way, with time, we can improve the API in general, and expose some of the new features in OAS3.\n\n_If you're looking for the Swagger 2.0/OAS 2.0 version of Petstore, then click [here](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml). Alternatively, you can load via the `Edit > Load Petstore OAS 2.0` menu option!_\n\nSome useful links:\n- [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)\n- [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)")
	}

	

	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "create-message",
				Short: "Create Message",
				Long: cli.Markdown("Update an existing pet by Id\n## Request Schema (application/json)\n\nproperties:\n  callback_url:\n    type: string\n  from:\n    type: string\n  network_id:\n    type: integer\n  to:\n    items:\n      type: string\n    type: array\ntype: object\n"),
				Example: examples,
				Args: cobra.MinimumNArgs(0),
				Run: func(cmd *cobra.Command, args []string) {
					body, err := cli.GetBody("application/json", args[0:])
					if err != nil {
						log.Fatal().Err(err).Msg("Unable to get body")
					}

					_, decoded, err := Openapi.YmlCreateMessage(params, body)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "list-messages",
				Short: "List messages",
				Long: cli.Markdown("Update an existing pet by Id"),
				Example: examples,
				Args: cobra.MinimumNArgs(0),
				Run: func(cmd *cobra.Command, args []string) {

					_, decoded, err := Openapi.YmlListMessages(params)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			
			cmd.Flags().String("inbox", "", "")
			cmd.Flags().String("status", "", "")

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "delete-message messageid",
				Short: "Deletes a message",
				Long: cli.Markdown("delete a message"),
				Example: examples,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {

					_, decoded, err := Openapi.YmlDeleteMessage(args[0], params)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			
			cmd.Flags().String("api-key", "", "")

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "get-message messageid",
				Short: "Get message",
				Long: cli.Markdown("Returns a single pet"),
				Example: examples,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {

					_, decoded, err := Openapi.YmlGetMessage(args[0], params)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "cancel-message messageid",
				Short: "Cancel a message",
				Long: cli.Markdown("Returns a single pet"),
				Example: examples,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					body, err := cli.GetBody("", args[1:])
					if err != nil {
						log.Fatal().Err(err).Msg("Unable to get body")
					}

					_, decoded, err := Openapi.YmlCancelMessage(args[0], params, body)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "send-message messageid",
				Short: "Sends a message",
				Long: cli.Markdown("Returns a single pet"),
				Example: examples,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					body, err := cli.GetBody("", args[1:])
					if err != nil {
						log.Fatal().Err(err).Msg("Unable to get body")
					}

					_, decoded, err := Openapi.YmlSendMessage(args[0], params, body)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "list-networks",
				Short: "List networks",
				Long: cli.Markdown("Returns a single pet"),
				Example: examples,
				Args: cobra.MinimumNArgs(0),
				Run: func(cmd *cobra.Command, args []string) {

					_, decoded, err := Openapi.YmlListNetworks(params)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			
			cmd.Flags().String("country-code", "", "ID of pet to return")

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "get-network networkid",
				Short: "Get network",
				Long: cli.Markdown("Returns a single pet"),
				Example: examples,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {

					_, decoded, err := Openapi.YmlGetNetwork(args[0], params)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			
			cmd.Flags().Int64("country-code", 0, "ID of pet to return")

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "get-pricing networkid",
				Short: "List network rates",
				Long: cli.Markdown("Returns a single pet"),
				Example: examples,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {

					_, decoded, err := Openapi.YmlGetPricing(args[0], params)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
		func () {
			params := viper.New()

			var examples string

			

			cmd := &cobra.Command{
				Use: "create-pricing networkid",
				Short: "Create network price",
				Long: cli.Markdown("Returns a single pet"),
				Example: examples,
				Args: cobra.MinimumNArgs(1),
				Run: func(cmd *cobra.Command, args []string) {
					body, err := cli.GetBody("", args[1:])
					if err != nil {
						log.Fatal().Err(err).Msg("Unable to get body")
					}

					_, decoded, err := Openapi.YmlCreatePricing(args[0], params, body)
					if err != nil {
						log.Fatal().Err(err).Msg("Error calling operation")
					}

					if err := cli.Formatter.Format(decoded); err != nil {
						log.Fatal().Err(err).Msg("Formatting failed")
					}

					
				},
			}
			root.AddCommand(cmd)

			

	cli.SetCustomFlags(cmd)

	if cmd.Flags().HasFlags() {
		params.BindPFlags(cmd.Flags())
	}

		}()
	
}
