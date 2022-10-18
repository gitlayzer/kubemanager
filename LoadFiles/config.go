package LoadFiles

import (
	"flag"
)

var (
	KubeConfig      string
	Address         string
	Admin           string
	Password        string
	PodLogTailLines int
)

func SetFlags() {
	flag.StringVar(&KubeConfig, "kubeconfig", "kubeconfig", "Set kubeconfig file")
	flag.StringVar(&Address, "listen", "0.0.0.0:9090", "Set address")
	flag.StringVar(&Admin, "user", "admin", "Set admin")
	flag.StringVar(&Password, "pass", "123456", "Set password")
	flag.IntVar(&PodLogTailLines, "podlogtaillines", 2000, "Set pod log tail lines")
}

func ReadKubeConfigFile() string {
	flag.Parse()
	return KubeConfig
}

func ReadAddress() string {
	flag.Parse()
	return Address
}

func ReadAdmin() string {
	flag.Parse()
	return Admin
}

func ReadPassword() string {
	flag.Parse()
	return Password
}

func ReadPodLogTailLines() int {
	flag.Parse()
	return PodLogTailLines
}
