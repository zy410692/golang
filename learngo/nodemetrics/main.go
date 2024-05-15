package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/client/clientset/versioned"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type DingTalkClient struct {
	WebhookURL string
	Secret     string
}

type DingDingMessage struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (c *DingTalkClient) SendMessage(content string) error {
	msg := DingDingMessage{
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}

	bytesData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(bytesData)

	// Generate sign
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	strToSign := timestamp + "\n" + c.Secret
	hmac256 := hmac.New(sha256.New, []byte(c.Secret))
	hmac256.Write([]byte(strToSign))
	data := hmac256.Sum(nil)
	sign := base64.StdEncoding.EncodeToString(data)

	// Append sign to url
	webhookURL, err := url.Parse(c.WebhookURL)
	if err != nil {
		return err
	}
	query := webhookURL.Query()
	query.Set("timestamp", timestamp)
	query.Set("sign", sign)
	webhookURL.RawQuery = query.Encode()

	req, err := http.NewRequest("POST", webhookURL.String(), reader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("dingtalk request failed with status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	var kubeconfig *string
	if home := filepath.Dir("/Users/zhangyi/"); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, _ := kubernetes.NewForConfig(config)
	metricsClient, _ := versioned.NewForConfig(config)
	client := DingTalkClient{
		WebhookURL: "https://oapi.dingtalk.com/robot/send?access_token=70a8d7519e57a80316d7fe7e7b8f35e6deab2a5e137914c6fa7243ac46a218fa",
		Secret:     "SEC9467d356a809e7a183571231eb0cea0647a77fd9c2588316aacc5bc7ef20b67f",
	}

	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	for _, node := range nodes.Items {
		nodeMetrics, _ := metricsClient.MetricsV1beta1().NodeMetricses().Get(context.TODO(), node.ObjectMeta.Name, metav1.GetOptions{})
		cpuUsageQuantity, ok := nodeMetrics.Usage[corev1.ResourceCPU]
		if !ok {
			log.Println("error")
		}
		cpuUsage := cpuUsageQuantity.MilliValue()
		cpuCapacityQuantity, ok := node.Status.Capacity[corev1.ResourceCPU]
		if !ok {
			log.Println("error")
		}
		cpuCapacity := cpuCapacityQuantity.MilliValue()
		cpuPercentage := (float64(cpuUsage) / float64(cpuCapacity)) * 100

		memCapacityQuantity, ok := node.Status.Capacity[corev1.ResourceMemory]
		if !ok {
			log.Println("error")
		}
		memCapacity := memCapacityQuantity.Value()
		memUsageQuantity, ok := nodeMetrics.Usage[corev1.ResourceMemory]
		if !ok {
			log.Println("error")

		}
		memUsage := memUsageQuantity.Value()

		memPercentage := (float64(memUsage) / float64(memCapacity)) * 100
		client.SendMessage(fmt.Sprintf("Node: %s, CPU Usage: %.2f%%, Memory Usage: %.2f%%\n", node.ObjectMeta.Name, cpuPercentage, memPercentage))
		fmt.Printf("Node: %s, CPU Usage: %.2f%%, Memory Usage: %.2f%%\n", node.ObjectMeta.Name, cpuPercentage, memPercentage)
	}
}
