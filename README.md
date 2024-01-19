# apctl

## installation

1. Go to release and download the correct build for your local workstation.  
2. unzip tar using OS appropraite method. On MacOS can double-click package.
3. Find CLI in unpacked folder. Move file to a location on $PATH.
4. On MacOS, since executable not signed by Apple Developer integration, you will need to approve use through the security system settings tab.

### usage  

See tool help.  
```bash
apctl --help
```

Authenticate. A browser window will open and you will be re-drected to authenticate into the github organization using the github identity yo uused for joining aws-engineering-poc. Follow the instructions.  
```bash
apctl login
```

List available cluster.  
```bash
apctl list clusters
nonprod-mkt01-aws-us-west-2
prod-mkt01-aws-us-east-2
```
or, set `export APCTL_DEFAULTSHOWHIDDEN=true` to also see cluster only authorized for paltform team usage.  
```bash
apctl list clusters
pocdev-mkt01-aws-us-east-1
nonprod-mkt01-aws-us-west-2
prod-mkt01-aws-us-east-2
mapi-mkt01-aws-us-east-2
```

Generate kubeconfig file for a cluster.
```bash
apctl get kubeconfig --cluster CLUSTERNAME
```
This will generate an oidc provider kubeconfig for the specified cluster and print to stdout. You can write the output to a local file with bash pipe like ` > kubeconfig` and then use `export KUBECONFIG=kubeconfig` to use the resulting for for kubectl or other commands that make use of the kubeconfig definition. You can of course also merge the config file with ~/.kube/config  

Once the kubeconfig file is generated and made available in a standard location you can interact with the cluster useing `kubectl` or any other k8s api compatible tools. Note: while the Authentication step can be performed by any github org member, authN alone does not effect your access or permissions to interact with the cluster api per se. In order to be authorized for any response from the k8s api a clusterrolebinding matching the claims in your id token must exist on the cluster.  

