package database

import "time"

/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 10:38
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

const (
	AerospikeConnectTimeout = 5 * time.Second
	AerospikeReadTimeout    = 5 * time.Second
	MaxRetries              = 3
	SleepBetweenRetries     = 10 * time.Millisecond
)