package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"kubemanager/LoadFiles"

	"github.com/wonderivan/logger"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Pod pod

type pod struct{}

type PodsResp struct {
	Total int          `json:"total"`
	Items []corev1.Pod `json:"items"`
}

type PodsNp struct {
	Namespace string
	PodNum    int
}

func (p *pod) GetPods(filterName, namespace string, limit, page int) (podsResp *PodsResp, err error) {
	podlist, err := K8s.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取Pod列表失败！" + err.Error())
		return nil, errors.New("获取Pod列表失败！" + err.Error())
	}
	selectableData := &dataSelector{
		GenericDataList: p.toCells(podlist.Items),
		DataSelect: &DataSelectQuery{
			Filter:   &FilterQuery{Name: filterName},
			Paginate: &PaginateQuery{Limit: limit, Page: page},
		},
	}

	filterd := selectableData.Filter()
	total := len(filterd.GenericDataList)

	data := filterd.Sort().Paginate()

	pods := p.formCells(data.GenericDataList)

	return &PodsResp{
		Total: total,
		Items: pods,
	}, nil
}

func (p *pod) GetPodDetail(podName, namespace string) (pod *corev1.Pod, err error) {
	pod, err = K8s.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取Pod详情失败！" + err.Error())
		return nil, errors.New("获取Pod详情失败！" + err.Error())
	}
	return pod, nil
}

func (p *pod) DeletePod(podName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除Pod失败！" + err.Error())
		return errors.New("删除Pod失败！" + err.Error())
	}
	return nil
}

func (p *pod) UpdatePod(namespace, content string) (err error) {
	var pod = &corev1.Pod{}
	err = json.Unmarshal([]byte(content), pod)
	if err != nil {
		logger.Error("反序列化Pod失败！" + err.Error())
		return errors.New("反序列化Pod失败！" + err.Error())
	}
	_, err = K8s.ClientSet.CoreV1().Pods(namespace).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		logger.Error("更新Pod失败！" + err.Error())
		return errors.New("更新Pod失败！" + err.Error())
	}
	return nil
}

func (p *pod) GetPodContainers(podName, namespace string) (containers []string, err error) {
	pod, err := p.GetPodDetail(podName, namespace)
	if err != nil {
		return nil, err
	}
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return containers, nil
}

func (p *pod) GetPodLog(containerName, podName, namespace string) (log string, err error) {
	lineLimit := int64(LoadFiles.ReadPodLogTailLines())
	option := &corev1.PodLogOptions{
		Container: containerName,
		TailLines: &lineLimit,
	}
	req := K8s.ClientSet.CoreV1().Pods(namespace).GetLogs(podName, option)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		logger.Error("获取Pod日志失败！" + err.Error())
		return "", errors.New("获取Pod日志失败！" + err.Error())
	}
	defer podLogs.Close()
	buff := new(bytes.Buffer)
	_, err = io.Copy(buff, podLogs)
	if err != nil {
		logger.Error("复制Pod日志失败！" + err.Error())
		return "", errors.New("复制Pod日志失败！" + err.Error())
	}
	return buff.String(), nil
}

func (p *pod) GetPodNumPreNp() (podsNps []*PodsNp, err error) {
	namespaceList, err := K8s.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, namespace := range namespaceList.Items {
		podList, err := K8s.ClientSet.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		podsNp := &PodsNp{
			Namespace: namespace.Name,
			PodNum:    len(podList.Items),
		}
		podsNps = append(podsNps, podsNp)
	}
	return podsNps, nil
}

func (p *pod) toCells(pods []corev1.Pod) []DataCell {
	cells := make([]DataCell, len(pods))
	for i := range pods {
		cells[i] = podCell(pods[i])
	}
	return cells
}

func (p *pod) formCells(cells []DataCell) []corev1.Pod {
	pods := make([]corev1.Pod, len(cells))
	for i := range cells {
		pods[i] = corev1.Pod(cells[i].(podCell))
	}
	return pods
}
