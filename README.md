# Go Laserfiche Workflow API Client

### About Cabarrus County
---
Cabarrus is an ever-growing county in the southcentral area of North Carolina. Cabarrus is part of the Charlotte/Concord/Gastonia NC-SC Metropolitan Statistical Area and has a population of about 210,000. Cabarrus is known for its rich stock car racing history and is home to Reed Gold Mine, the site of the first documented commercial gold find in the United States.

### About our team
---
The Business & Location Innovative Services (BLIS) team for Cabarrus County consists of five members:

+ Joseph Battinelli - Team Supervisor
+ Mark McIntyre - Software Developer
+ Landon Patterson - Software Developer
+ Brittany Yoder - Software Developer
+ Marci Jones - Software Developer

Our team is responsible for software development and support for the [County](https://www.cabarruscounty.us/departments/information-technology). We work under the direction of the Chief Information Officer.

### About
---
The LFWFAPI is a minimal Laserfiche Workflow API wrapper written in Go. It currently has API functionality built in to

1. Grab all workflows from the current workflow server
2. Grab all parameters of a workflow
3. Start a workflow, with or without parameters passed to it.

### Installation
---
```
go get -u github.com/CabarrusCo/lfwfapi
```

To get started with the Laserfiche API for Go, simply call it in a import
```
import "github.com/CabarrusCo/lfwfapi"
```

### Spinning up a Client
---
To spin up a basic client, call the lfwfapi.NewClient method, pass your base workflow URL to the NewClient as such

```
wfClient := lfwfapi.NewClient("http://WORKFLOWURLHERE")
```

### Using NTLM Auth
---
The package comes with no NTLM auth support baked in. If your workflow API uses NTLM auth, you can create your own Http Client that has NTLM auth and pass it to the created Workflow API client.
A good package for handling NTLM auth can be found at https://github.com/vadimi/go-http-ntlm. An example is provided below

```
client := http.Client{
	Transport: &httpntlm.NtlmTransport{
		Domain:   "DOMAIN",
		User:     "USERNAME",
		Password: "PASSWORD",
		// Configure RoundTripper if necessary, otherwise DefaultTransport is used
		RoundTripper: &http.Transport{
			// provide tls config
			TLSClientConfig: &tls.Config{},
			// other properties RoundTripper, see http.DefaultTransport
		},
	},
}
wfClient := lfwfapi.NewClient("http://WORKFLOWURLHERE")
wfClient.HttpClient = &client
```

### Retrieving all workflows on a server
---
To retrieve all workflows on a server, simply use the GetAllWorkflows method.

```
grabAllWorkflows, err := wfClient.GetAllWorkflows()
if err != nil {
	log.Println(err)
	return
}

fmt.Printf("%+v", grabAllWorkflows)
  
```

### Retrieving all workflow parameters
---
To retrieve all workflow parameters on a particular workflow, just pass in the name of the workflow to the method GetWorkflowParameters.

```
grabWorkflowParams, err := wfClient.GetWorkflowParameters("WORKFLOW NAME HERE")
if err != nil {
	log.Println(err)
	return
}

fmt.Printf("%+v", grabWorkflowParams)

for _, v := range grabWorkflowParams {
	fmt.Println()
	fmt.Println("--------------")
	fmt.Printf("Workflow Token Name: %s Workflow Token Tags: %s\n", v.TokenName, v.TokenTags)
}
```

### Running a Workflow
---
To run a workflow, call the StartWorkflow method. You can also pass parameters to this method by using []lfwfapi.Parameter{}.

```
wfParams := []lfwfapi.Parameter{
	{
		Name:  "Message",
		Value: "Hello World!",
	},
	{
		Name:  "Message Two",
		Value: "I AM ANOTHER PARAMETER!",
	},
}

runWorkflow, err := wfClient.StartWorkflow("WORKFLOW NAME HERE", wfParams)
if err != nil {
	log.Println(err)
	return
}

fmt.Println(runWorkflow)
```

Additionally, if you just want to start a workflow with no parameters, pass nil to the StartWorkflow method
```

runWorkflow, err := wfClient.StartWorkflow("WORKFLOW NAME HERE", nil)
if err != nil {
	log.Println(err)
	return
}

fmt.Println(runWorkflow)
```
### More methods coming
Here in this release we covered the main Three Methods. Retrieve all workflows, retrieve workflow parameters, and starting a workflow with parameters. We will add more minor methods in upcoming releases such as Queue workflow.
