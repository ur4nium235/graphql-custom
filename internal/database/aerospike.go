package database

import (
	"base/internal/utils"
	aero "github.com/aerospike/aerospike-client-go"
	"strconv"
	"strings"
	"time"
)

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 09:56
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */



func NewConnection(hostStr string) (*aero.Client, error) {
	hosts := strings.Split(hostStr,",")
	listHost := make([]*aero.Host, len(hosts))
	for index, value := range hosts {
		host := strings.Split(value, ":")
		port, _ := strconv.Atoi(host[1])
		listHost[index] = aero.NewHost(host[0], port)
	}
	var clientPolicy = aero.NewClientPolicy()
	clientPolicy.Timeout = AerospikeConnectTimeout
	return aero.NewClientWithPolicyAndHost(clientPolicy, listHost...)
}

func GetRecord(client *aero.Client, namespace, setname string, keyRaw interface{}) (*aero.Record, error) {
	key, _ := aero.NewKey(namespace, setname, keyRaw)
	policy := CreatePolicyGetDefault()
	return client.Get(policy, key)
}

func GetBatchRecord(client *aero.Client, namespace, setname string, keyRaws []string) ([]*aero.Record, error) {
	defer func() {
		if err := recover(); err != nil {
			utils.HandleError(err)
		}
	}()

	keys := make([]*aero.Key, 0)
	for _, keyCache := range keyRaws {
		key, err := aero.NewKey(namespace, setname, keyCache)
		if err == nil {
			keys = append(keys, key)
		}
	}
	var policy = aero.NewBatchPolicy()
	policy.TotalTimeout = AerospikeReadTimeout
	policy.MaxRetries = MaxRetries

	return client.BatchGet(policy, keys)
}

func CreatePolicyGetDefault() *aero.BasePolicy {
	policy := aero.NewPolicy()
	policy.TotalTimeout = time.Second
	policy.MaxRetries = MaxRetries
	policy.SleepBetweenRetries = SleepBetweenRetries
	return policy
}

func SaveRecord(client *aero.Client, namespace, setname string, keyRaw string, bins ...*aero.Bin)  {
	key, _ := aero.NewKey(namespace, setname, keyRaw)
	var policy = aero.NewWritePolicy(0, 0)
	policy.TotalTimeout = AerospikeReadTimeout
	policy.MaxRetries = MaxRetries
	policy.SleepBetweenRetries = SleepBetweenRetries
	policy.SendKey = true
	err := client.PutBins(policy, key, bins...)

	if err != nil {
		utils.HandleError(err)
	}
}

func DeleteRecord(client *aero.Client, namespace, setname string, keyRaw string) bool {
	key, _ := aero.NewKey(namespace, setname, keyRaw)
	var policy = aero.NewWritePolicy(0, 0)
	policy.TotalTimeout = AerospikeReadTimeout
	policy.MaxRetries = MaxRetries
	result, err := client.Delete(policy, key)
	if err != nil {
		utils.HandleError(err)
		return false
	}
	return result
}

