package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func k8sClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetPods(c *gin.Context) {
	client, err := k8sClient()
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, "Kubernetes client error.")
		return
	}
	pods, err := client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, "Get pods error.")
		return
	}
	res := fmt.Sprintf("There are %d pods in the cluster.", len(pods.Items))
	c.String(http.StatusOK, res)
}

func Router(route *gin.Engine) {
	router := route.Group("/")

	router.GET("/pods", GetPods)
}

func main() {
	app := gin.Default()
	Router(app)
	app.Run(":8000")
}
