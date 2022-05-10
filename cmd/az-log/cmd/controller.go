/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"errors"
	"strings"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	consts "sigs.k8s.io/azuredisk-csi-driver/pkg/azureconstants"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// volumes, nodes, requestIds, sinceTime, isFollow, isPrevious := GetFlags(cmd)
		config := getConfig()
		clientsetK8s := getKubernetesClientset(config)

		pod := getControllerPodName(clientsetK8s)
		GetLogsByAzDriverPod(clientsetK8s, pod, "azuredisk", nil, nil, nil, "2022-05-10T22:07:40Z", false, false) // 17:57:01.170816
	},
}

func init() {
	getCmd.AddCommand(controllerCmd)
}

func getControllerPodName(clientsetK8s kubernetes.Interface) string {
	// client, err := coordinationv1.NewForConfig(config)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// lease, err := client.Leases("azure-disk-csi").Get(context.Background(), "default", metav1.GetOptions{})

	lease, err := clientsetK8s.CoordinationV1().Leases(consts.DefaultAzureDiskCrdNamespace).Get(context.Background(), "default", metav1.GetOptions{}) //List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	holder := *lease.Spec.HolderIdentity
	node := strings.Split(holder, "_")[0]

	pods, err := clientsetK8s.CoreV1().Pods(consts.ReleaseNamespace).List(context.Background(), metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + node,
		LabelSelector: "app=csi-azuredisk2-controller",
	})
	if err != nil {
		panic(err.Error())
	}
	if len(pods.Items) > 1 {
		panic(errors.New("More than one controller pods were found."))
	}

	return pods.Items[0].Name
}

