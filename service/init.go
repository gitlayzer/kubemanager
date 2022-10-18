package service

import (
	"kubemanager/LoadFiles"

	"github.com/wonderivan/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var K8s k8s

type k8s struct {
	ClientSet *kubernetes.Clientset
}

func (k *k8s) Init() {
	conf, err := clientcmd.BuildConfigFromFlags("", LoadFiles.ReadKubeConfigFile())
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic("创建k8sClientset失败！" + err.Error())
	} else {
		logger.Info("创建k8sClientset成功！")
	}

	k.ClientSet = clientSet
}
