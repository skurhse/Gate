package repository

import (
	"github.com/mrNobody95/Gate/pkg/gormext"
	log "github.com/sirupsen/logrus"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

var (
	Assets     *AssetRepository
	Brokerages *BrokerageRepository
)

func LoadRepositories(dsn string) {
	db, err := gormext.Open(gormext.Mariadb, dsn)
	if err != nil {
		log.WithError(err).Fatal("can not load repository configs")
	}
	Assets = &AssetRepository{db: db}
	Brokerages = &BrokerageRepository{db: db}
}