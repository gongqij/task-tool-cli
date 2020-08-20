package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"task-tool-cli/api"
	"task-tool-cli/client"
)

type List struct {
	taskID        string
	status        string
	sourceAddress string
}

func PrintResult(resp *client.TaskInfoResp, objectType string) bool {
	flag := false
	listRes := make(map[string][]List)
	logrus.Info("Total tasks Num: ", len(resp.Tasks))
	for _, v := range resp.Tasks {
		listRes[v.Info.ObjectType] = append(listRes[v.Info.ObjectType], List{
			taskID:        v.Info.TaskID,
			status:        v.Status.Status,
			sourceAddress: v.Info.SourceAddress,
		})
	}
	for k, v := range listRes {
		if objectType != "" && k != objectType {
			continue
		}
		//该objectType存在
		flag = true
		fmt.Println("---------------------------------------------------------------------------")
		if k == api.ObjectType_OBJECT_ALGO.String() {
			fmt.Printf("%s\t\t\t\t\tTotalNum:%d\n", k, len(v))
		} else {
			fmt.Printf("%s\t\t\tTotalNum:%d\n", k, len(v))
		}
		fmt.Println("taskID\t\t\t\t     status\t\t\tsource")
		for _, v := range v {
			if strings.Contains(v.sourceAddress, "cluster") {
				slice := strings.Split(v.sourceAddress, "/")
				fmt.Printf("%s%20s\t\t%s\n", v.taskID, v.status, slice[len(slice)-1])
			} else {
				fmt.Printf("%s%20s\t\t%s\n", v.taskID, v.status, v.sourceAddress)
			}
		}
	}
	return flag
}

func SetupAuth(acc, sec string) client.HTTPAuthFunc {
	if acc == "" || sec == "" {
		return nil
	}
	return func(endpoint string, url *url.URL) (string, error) {
		token, err := getToken(acc, sec, url.Scheme+"://"+url.Host)
		if err != nil {
			return "", err
		}
		return "Bearer " + token, nil
	}
}

type TokenResp struct {
	Code  int32  `json:"code"`
	Token string `json:"token"`
}

func getToken(accessKey string, secretKey string, endpoint string) (token string, errResult error) {
	jsonKey := []byte(fmt.Sprintf(`{"access_key": "%s", "secret_key": "%s"}`, accessKey, secretKey))
	url := fmt.Sprintf("%s/components/user_manager/v1/users/sign_token", endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonKey))
	if err != nil {
		return "", fmt.Errorf("invalid http request: %s", url)
	}
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: transCfg}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result := TokenResp{}
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		if err := json.Unmarshal(body, &result); err != nil {
			return "", err
		}
	}
	return result.Token, nil
}

func IsObjectTypeExist(objectType string) bool {
	switch objectType {
	case api.ObjectType_OBJECT_MULTI_PACH.String():
		fallthrough
	case api.ObjectType_OBJECT_TRAFFIC_ANOMALY_EVENT.String():
		fallthrough
	case api.ObjectType_OBJECT_TRAFFIC_CAMERA_VISION_INFO.String():
		fallthrough
	case api.ObjectType_OBJECT_TRAFFIC_AUTOMOBILE_COUNT.String():
		fallthrough
	case api.ObjectType_OBJECT_TRAFFIC_MULTI_PACH.String():
		return true
	}
	return false
}

func AllObjectType() []string {
	return []string{api.ObjectType_OBJECT_TRAFFIC_ANOMALY_EVENT.String(), api.ObjectType_OBJECT_TRAFFIC_MULTI_PACH.String(), api.ObjectType_OBJECT_MULTI_PACH.String(), api.ObjectType_OBJECT_TRAFFIC_AUTOMOBILE_COUNT.String(), api.ObjectType_OBJECT_TRAFFIC_CAMERA_VISION_INFO.String()}
}
