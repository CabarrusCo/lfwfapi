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
The LFWFAPI is a Laserfiche Workflow API wrapper written in Go. It currently has API functionality built in to

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

### Making a new client
---
There are two ways to make a new API client, the first way is with NTLM auth, the second is without NTLM auth.

To use NTLM auth, spin up your client using the lfwfapi Credentials struct.

```
loginInfo := lfwfapi.Credentials{Username: "xxxxx", Password: "xxxxxxx"}
wfClient := lfwfapi.NewClient("http://WORKFLOWBASEURLHERE", &loginInfo) // Pass loginInfo
```
It is up to you store your credentials and retrieve them securely.

If you are not using any auth on your API, simply spin up the client with nil in the Credentials.

```
wfClient := lfwfapi.NewClient("http://WORKFLOWBASEURLHERE", nil) // Pass loginInfo
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
