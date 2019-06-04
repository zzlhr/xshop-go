package util

import log "github.com/Sirupsen/logrus"

func CheckErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
