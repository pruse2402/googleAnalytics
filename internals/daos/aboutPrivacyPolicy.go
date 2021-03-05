package daos

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"fmt"

	"github.com/FenixAra/go-util/log"
)

type AboutPrivacyPolicy struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewAboutPrivacyPolicy(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *AboutPrivacyPolicy {
	return &AboutPrivacyPolicy{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

type AboutPrivacyPolicyDao interface {
	GetLatestAboutPrivacyPolicy() (*dtos.AboutPrivacyPolicy, error)
	AboutPrivacyPolicyInfoList(versionCode int64) (*[]dtos.AboutPrivacyPolicyInfo, error)
}

func (p *AboutPrivacyPolicy) GetLatestAboutPrivacyPolicy() (*dtos.AboutPrivacyPolicy, error) {
	pPolicy := dtos.AboutPrivacyPolicy{}
	row := p.dbConnMSSQL.GetQueryer().QueryRow("SELECT about_privacy_policy_id,version_code,version_name,updated_at FROM ac_about_privacy_policy WHERE version_code = (SELECT MAX(version_code) FROM ac_about_privacy_policy)")
	err := row.Scan(&pPolicy.AboutPrivacyPolicyID, &pPolicy.VersionCode, &pPolicy.VersionName, &pPolicy.UpdatedAt)
	if err != nil {
		p.l.Error("Error fetching about_privacy_policy - ", err)
		return nil, err
	}
	return &pPolicy, nil
}

func (a *AboutPrivacyPolicy) AboutPrivacyPolicyInfoList(versionCode int64) (*[]dtos.AboutPrivacyPolicyInfo, error) {
	list := []dtos.AboutPrivacyPolicyInfo{}
	rows, err := a.dbConnMSSQL.GetQueryer().Query(fmt.Sprintf("SELECT id,version_code, version_name, sequence_no,content_type, message_info from ac_about_privacy_policy_info where version_code='%d' order by sequence_no", versionCode))
	if err != nil {
		a.l.Error("Error AboutPrivacyPolicyInfoList ", versionCode, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		info := dtos.AboutPrivacyPolicyInfo{}
		err := rows.Scan(&info.ID,
			&info.VersionCode,
			&info.VersionName,
			&info.SequenceNO,
			&info.ContentType,
			&info.MessageInfo)
		if err != nil {
			a.l.Error("Error AboutPrivacyPolicyInfoList - ", versionCode, err)
			return nil, err
		}
		list = append(list, info)
	}
	return &list, nil
}
