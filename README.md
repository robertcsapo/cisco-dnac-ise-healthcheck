# Cisco DNA Center ISE Health Check

*Cisco DNA-C Communication Health Check from Console*

Sometimes Cisco DNA-C fails to communicate with Cisco ISE with Error on Application level.  
This script helps you to verify that the communication paths between Cisco DNA-C and Cisco ISE is working (non-Application level)

## Demo

What visual, if shown, clearly articulates the impact of what you have created?  In as concise a visualization as possible (code sample, CLI output, animated GIF, or screenshot) show what your project makes possible.

## Requirements

* Cisco DNA Center Console access (_ssh maglev@<dnac-ip-address> -p 2222_)
* Internet Connectivity on DNA Center (Cloud Enabled)
  - DNA Center downloads this container image from Public Docker Registry (docker.io)

## Usage

Login to Cisco DNA-C
* ```ssh maglev@<dnac-ip-address> -p 2222```

Run script
* ```docker run -it --rm robertcsapo/cisco-dnac-ise-healthcheck && docker rmi robertcsapo/cisco-dnac-ise-healthcheck```

Enter Cisco ISE FQDN (then hit ENTER)
* ```Enter host (FQDN): cisco-ise.example.com```

### Automation

Run script with host flag
* ```docker run -it --rm robertcsapo/cisco-dnac-ise-healthcheck -host cisco-ise.example.com && docker rmi robertcsapo/cisco-dnac-ise-healthcheck```

## Installation

If you need to manually build the script on the host.  
(Important that you predownload golang:1.9.2-alpine3.7 from docker.io)  
```docker pull golang:1.9.2-alpine3.7```

### Build
* ```docker build -t robertcsapo/cisco-dnac-ise-healthcheck .```

## Technologies & Frameworks Used

**Cisco Products & Services:**

- Cisco DNA Center
- Cisco ISE

**Tools & Frameworks:**

- Golang (1.9.2)


## Authors & Maintainers

- Robert Csapo <rcsapo@cisco.com>

## License

This project is licensed to you under the terms of the [Cisco Sample
Code License](./LICENSE).
