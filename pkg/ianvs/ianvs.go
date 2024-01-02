package ianvs

import (
	"fmt"
	ianvsclient "git.woa.com/ianvs/ianvs-sdk/pkg/client"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"time"
)

func GetClusterClient(clsid string) (*kubernetes.Clientset, error) {
	ianvscli := ianvsclient.NewClient(
		"ianvs-for-cls.tke.woa.com",
		"9c50924be5333464c2b50cc08d915bc2",
		"9b1d669726c4bd6d76e7ae50213db254",
	)

	restConfig, err := ianvscli.GetClientConfig("yuehuazhang", clsid)
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
