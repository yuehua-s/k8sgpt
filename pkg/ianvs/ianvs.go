package ianvs

import (
	"fmt"
	ianvsclient "git.woa.com/ianvs/ianvs-sdk/pkg/client"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"time"
)

func GetClusterClient(clsid string) (*kubernetes.Clientset, error) {
	ianvsClient := ianvsclient.NewClient(
		"ianvs-for-cls.tke.woa.com",
		"1b09517f2a20bffe4df296858fa66cd4",
		"eb9dec621c31f1e9cba741f8a9624d41",
	)

	restConfig, err := ianvsClient.GetClientConfig("yuehuazhang", clsid)
	if err != nil {
		logrus.Errorf(fmt.Sprintf("Failed to get restconfig from cls %s.ERROR:%s", clsid, err.Error()))
		return nil, err
	}
	restConfig.Timeout = 5 * time.Second
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logrus.Errorf(fmt.Sprintf("Failed to get kubeclient from cls %s.ERROR:%s", clsid, err.Error()))
		return nil, err
	}

	return kubeClient, nil
}
