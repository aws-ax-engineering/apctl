package cmd

type ClusterConfig struct {
	clusterName 										string
	hidden 													bool
	clusterEndpoint 								string
	base64CertificateAuthorityData  string
}

type KubeconfigTemplateData struct {
	ClusterPublicKey string
	ClusterURL       string
	ClusterName      string
	AccessToken      string
	ClientID         string
	IDToken          string
	IDPIssuerURL     string
	RefreshToken     string
}

var ( 
	clustersList = []ClusterConfig{ 
		{
			clusterName: "pocdev-mkt01-aws-us-east-1",
			hidden: true,
			clusterEndpoint: "$POCDEV_CLUSTER_URL",
			base64CertificateAuthorityData: "$POCDEV_CLUSTER_PUBLIC_CERTIFICATE_AUTHORITY_DATA",
		},
		{
			clusterName: "nonprod-mkt01-aws-us-west-2",
			hidden: false,
			clusterEndpoint: "$NONPROD_CLUSTER_URL",
			base64CertificateAuthorityData: "$NONPROD_CLUSTER_PUBLIC_CERTIFICATE_AUTHORITY_DATA",
		},
		{
			clusterName: "prod-mkt01-aws-us-east-2",
			hidden: false,
			clusterEndpoint: "$PROD_CLUSTER_URL",
			base64CertificateAuthorityData: "$PROD_CLUSTER_PUBLIC_CERTIFICATE_AUTHORITY_DATA",
		},
		{
			clusterName: "mapi-i01-aws-us-east-2",
			hidden: true,
			clusterEndpoint: "$MAPI_CLUSTER_URL",
			base64CertificateAuthorityData: "$MAPI_CLUSTER_PUBLIC_CERTIFICATE_AUTHORITY_DATA",
		},
	}

	kubefonfigYAML = `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: {{ .ClusterPublicKey }}
    server: {{ .ClusterURL }}
  name: {{ .ClusterName }}
contexts:
- context:
    cluster: {{ .ClusterName }}
    user: poc-user
  name: {{ .ClusterName }}
current-context: {{ .ClusterName }}
kind: Config
preferences: {}
users:
- name: poc-user
  user:
    auth-provider:
      config:
        access-token: {{ .AccessToken }}
        client-id: {{ .ClientID }}
        id-token: {{ .IDToken }}
        idp-issuer-url: {{ .IDPIssuerURL }}
        refresh-token: {{ .RefreshToken }}
      name: oidc
`
)

const (
	LoginClientId		 						 = "$APCTL_CLIENT_ID"
	LoginScope                   = "openid offline_access profile email"
	LoginAudience                = "https://dev-apctl.us.auth0.com/api/v2/"
	IdpIssuerUrl								 = "https://dev-apctl.us.auth0.com/"

	DefaultShowHidden						 = true
	DefaultCluster							 = "nonprod-mkt01-aws-us-west-2"

	ConfigEnvDefault             = "APCTL"
	ConfigFileDefaultName        = "config"
	ConfigFileDefaultType        = "yaml"
	ConfigFileDefaultLocation    = "/.apctl" // path will begin with $HOME dir
	ConfigFileDefaultLocationMsg = "config file (default is $HOME/.apctl/config.yaml)"
)
