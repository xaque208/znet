package arpwatch

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	junos "github.com/scottdware/go-junos"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	macsTable = "macs"
)

// ArpWatch is a Junos ARP table reader for prometheus export.
type ArpWatch struct {
	Hosts       []string
	Auth        *junos.AuthMethod
	redisClient *redis.Client
}

// NewArpWatch loads configuration to create a new ArpWatch object for return.
func NewArpWatch() *ArpWatch {
	hosts := viper.GetStringSlice("junos.hosts")
	if len(hosts) == 0 {
		log.Error("List of hosts required")
		return &ArpWatch{}
	}

	log.Debugf("Arpwatch for: %+v", hosts)

	auth := &junos.AuthMethod{
		Username:   viper.GetString("junos.username"),
		PrivateKey: viper.GetString("junos.keyfile"),
	}

	aw := &ArpWatch{
		Hosts: hosts,
		Auth:  auth,
	}

	aw.redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", viper.GetString("redis.host")),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return aw
}

// Update reaches out to each device to update the arp table held in redis.
func (a *ArpWatch) Update() {
	// redisClient := NewRedisClient()
	// defer redisClient.Close()
	log.Debug("Arpwatch updating")

	for _, h := range a.Hosts {
		session, err := junos.NewSession(h, a.Auth)
		if err != nil {
			log.Error(err)
			continue
		}
		defer session.Close()

		views, err := session.View("arp")
		if err != nil {
			log.Error(err)
			continue
		}

		for _, arp := range views.Arp.Entries {
			result, err := a.redisClient.SIsMember(macsTable, arp.MACAddress).Result()
			if err != nil {
				log.Error(err)
				continue
			}

			if !result {
				log.Infof("New MACAddress seen: %+v", arp.MACAddress)
				_, err := a.redisClient.SAdd(macsTable, arp.MACAddress).Result()
				if err != nil {
					log.Error(err)
					continue
				}
			}

			keyName := fmt.Sprintf("mac:%s", arp.MACAddress)
			a.redisClient.HSet(keyName, "mac", arp.MACAddress)
			a.redisClient.HSet(keyName, "ip", arp.IPAddress)
			a.redisClient.Expire(keyName, 900*time.Second)
		}

	}

}
