package cmd

import (
	"fmt"
	"os"
	"text/template"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cluster string

var getCmd = &cobra.Command{
	Use:    "get",
	Short:  "Get config",
	Long:   `Get config when given a config type. For example: get kubeconfig --cluster pocdev-mkt01-aws-us-east-1`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get command")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

var kubeconfigCmd = &cobra.Command{
	Use:     "kubeconfig",
	Aliases: []string{"kubeconfig"},
	Short:   "Get kubeconfig",
	Long: `Get kubeconfig for requested cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		generateKubeconfig(Cluster)
	},
}

func init() {
	getCmd.AddCommand(kubeconfigCmd)

	kubeconfigCmd.PersistentFlags().StringVar(&Cluster, "cluster", DefaultCluster, "Name of cluster. Use 'list clusters' to see available options")
}

func generateKubeconfig(cluster string) {

	clusterPublicKey, clusterUrl := clusterIdentifiers(cluster)

	data := KubeconfigTemplateData{
		ClusterPublicKey: clusterPublicKey,
		ClusterURL: clusterUrl,
		ClusterName: cluster,
		AccessToken: viper.GetString("AccessToken"),
		ClientID: viper.GetString("LoginClientId"),
		IDToken: viper.GetString("IdToken"),
		IDPIssuerURL: viper.GetString("IdpIssuerUrl"),
		RefreshToken: viper.GetString("RefreshToken"),
	}

	kubeconfigTpl, err := template.New("yamlTemplate").Parse(kubefonfigYAML)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	err = kubeconfigTpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return
	}

}

func clusterIdentifiers(cluster string) (clusterPublicKey, clusterUrl string) {
	for _, clusters := range clustersList {
		if clusters.clusterName == cluster {
			return clusters.base64CertificateAuthorityData, clusters.clusterEndpoint
		}
	}
	return "Not Found", "Not Found"
}
