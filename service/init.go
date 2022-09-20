package service

import (
	"kubemanager/config"

	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	ClientSet *kubernetes.Clientset
}

func (k *k8s) Init() {
	conf, err := clientcmd.BuildConfigFromFlags("", config.Kubeconfig)
	if err != nil {

		panic("获取k8s配置失败" + err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("创建k8sClientset失败！" + err.Error())
	} else {
		logger.Info("创建k8sClientset成功！")
	}

	k.ClientSet = clientSet
}
