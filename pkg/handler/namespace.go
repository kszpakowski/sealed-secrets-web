package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func (h *Handler) Namespaces(c *gin.Context) {

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	overrides := clientcmd.ConfigOverrides{}
	clientConfig := clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, &overrides, os.Stdin)

	conf, _ := clientConfig.ClientConfig()
	// if err != nil {
	// return nil, nil, err
	// }

	restClient, _ := corev1.NewForConfig(conf)
	// if err != nil {
	// return nil, nil, err
	// }

	list, _ := restClient.Namespaces().List(v1.ListOptions{})
	namespaces := list.Items
	c.JSON(http.StatusOK, namespaces)
}
