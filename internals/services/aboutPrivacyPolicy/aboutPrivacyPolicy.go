package aboutPrivacyPolicy

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"cyberliver/go-alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type AboutPrivacyPolicy struct {
	l                  *log.Logger
	dbConnMSSQL        *mssqlcon.DBConn
	aboutPrivacyPolicy daos.AboutPrivacyPolicyDao
}

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *AboutPrivacyPolicy {
	return &AboutPrivacyPolicy{
		l:                  l,
		dbConnMSSQL:        dbConnMSSQL,
		aboutPrivacyPolicy: daos.NewAboutPrivacyPolicy(l, dbConnMSSQL),
	}
}

func (a *AboutPrivacyPolicy) AboutPrivacyPolicy() (*dtos.AboutPrivacyPolicyResponse, error) {
	pPolicy := dtos.AboutPrivacyPolicyResponse{}
	latestAboutPrivacyPolicy, err := a.aboutPrivacyPolicy.GetLatestAboutPrivacyPolicy()
	if err != nil {
		a.l.Error("GetLatestAboutPrivacyPolicy Error - ", err)
		return nil, err
	}
	a.l.Debug("latestVersionCode-", latestAboutPrivacyPolicy.VersionCode)
	list, errL := a.aboutPrivacyPolicy.AboutPrivacyPolicyInfoList(latestAboutPrivacyPolicy.VersionCode)
	if errL != nil {
		a.l.Error("AboutPrivacyPolicyInfoList Error - ", errL)
		return nil, errL
	}
	pPolicy.AboutPrivacyPolicyID = latestAboutPrivacyPolicy.AboutPrivacyPolicyID
	pPolicy.VersionCode = latestAboutPrivacyPolicy.VersionCode
	pPolicy.VersionName = latestAboutPrivacyPolicy.VersionName
	pPolicy.UpdatedAt = latestAboutPrivacyPolicy.UpdatedAt
	pPolicy.AboutPrivacyPolicyInfo = list
	return &pPolicy, nil
}
