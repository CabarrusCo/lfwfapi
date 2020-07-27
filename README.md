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

### Usage
---

To get started with the Laserfiche API for Go, simply call it in a import
```
import "github.com/CabarrusCo/lfwfapi
```

