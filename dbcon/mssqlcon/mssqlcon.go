package mssqlcon

import (
	"context"
	"database/sql"
	lg "log"
	"sync"

	"github.com/FenixAra/go-util/log"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/hgfischer/mysql"
)

//Pool of database connection
var once sync.Once
var ConnPool *sql.DB

type DBConn struct {
	conn          *sql.DB
	tx            *sql.Tx
	isTransaction bool
	l             *log.Logger
}
type IDBConn interface {
	Init(*log.Logger)
	GetQueryer() Queryer
	ExecuteInTransaction(func() error) error
	// rollbackTransaction(tx *sql.Tx)
}

// Interface to abstract the queryer(dbconnection or transaction)
type Queryer interface {
	Exec(sql string, arguments ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	Query(sql string, args ...interface{}) (*sql.Rows, error)
	QueryRow(sql string, args ...interface{}) *sql.Row
	Prepare(sql string) (*sql.Stmt, error)
}

func New(l *log.Logger) *DBConn {
	return &DBConn{
		l:    l,
		conn: ConnPool,
	}
}

func NewDBConn(l *log.Logger, conn *sql.DB) *DBConn {
	return &DBConn{
		l:    l,
		conn: conn,
	}
}

// Initialize the DB connection and assign the existing db connection
func (db *DBConn) Init(l *log.Logger) {
	db.conn = ConnPool
	db.l = l
}

func (db *DBConn) GetQueryer() Queryer {
	if db.isTransaction {
		return db.tx
	} else {
		return db.conn
	}
}

// ExecuteInTransaction executes the given function in DB transaction, i.e. It commits
// only if there is not error otherwise it is rolledback.
func (db *DBConn) ExecuteInTransaction(f func() error) (err error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	db.tx = tx
	db.isTransaction = true

	defer func() {
		if r := recover(); r != nil {
			db.l.Fatal("Recovered in function ", r)
			db.rollbackTransaction(tx)
		}
		db.isTransaction = false
	}()

	err = f()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	err = tx.Commit()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	return nil
}

func (db *DBConn) rollbackTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		db.l.Error("Error While rollback, Err: ", err)
	}
}

func MSSqlInit(url string) {
	if ConnPool == nil {
		once.Do(func() {
			var err error
			// Create connection pool
			ConnPool, err = sql.Open("mysql", url)
			if err != nil {
				lg.Println("Error creating connection pool: %+v", err)
			}
			ConnPool.SetMaxOpenConns(15)
			ctx := context.Background()
			err = ConnPool.PingContext(ctx)
			if err != nil {
				lg.Println("Unable to ping to DB. Err: %+v", err)
				return
			}
			lg.Println("Connected to database successfully!")
			CreateTable(ConnPool)
		})
	}
}

func CreateTable(db *sql.DB) {
	AboutPrivacyTableCreation(db)
	BehaviourChangeTechniquesTableCreation(db)    //BCT
	BehaviourChangeInterventionsTableCreation(db) //BCN
	PatientEngagementReminderTableCreation(db)
	BenefitTherapy(db)

}

func MSSqlConnClose() {
	if ConnPool != nil {
		ConnPool.Close()
	}
}

func AboutPrivacyTableCreation(db *sql.DB) {
	aboutPrivacyTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS 
		ac_about_privacy_policy (about_privacy_policy_id int unsigned NOT NULL AUTO_INCREMENT, 
		version_code int unsigned NOT NULL, 
		version_name varchar(255) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT versionx_about_privacy_policy UNIQUE (version_code),
		PRIMARY KEY (about_privacy_policy_id));`)
	if err != nil {
		lg.Println(err.Error())
	}
	_, err = aboutPrivacyTable.Exec()
	if err != nil {
		lg.Println(err.Error())
	}
	aboutPrivacyInfoTable, errN := db.Prepare(`CREATE TABLE IF NOT EXISTS 
		ac_about_privacy_policy_info (id int unsigned NOT NULL AUTO_INCREMENT, 
	    about_privacy_policy_id int unsigned NOT NULL,  
		CONSTRAINT fk_ac_about_privacy_policy FOREIGN KEY (about_privacy_policy_id) REFERENCES ac_about_privacy_policy(about_privacy_policy_id),
		version_code int unsigned NOT NULL, 
		version_name varchar(255) NOT NULL,
		sequence_no int unsigned NOT NULL, 
		content_type varchar(1000) NOT NULL,
		message_info BLOB,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT versionx_about_privacy_policy_info UNIQUE (version_code,about_privacy_policy_id,sequence_no),
		PRIMARY KEY (id));`)
	if errN != nil {
		lg.Println(errN.Error())
	}
	_, err = aboutPrivacyInfoTable.Exec()
	if err != nil {
		lg.Println(err.Error())
	}
}

func BehaviourChangeTechniquesTableCreation(db *sql.DB) {
	ac_bct, err := db.Prepare(`CREATE TABLE IF NOT EXISTS ac_behaviour_change_techniques (behaviour_change_id int unsigned NOT NULL AUTO_INCREMENT, 
		bct_taxonomy_id varchar(100) NOT NULL, 
		bct_taxonomy varchar(255) NOT NULL,
		bct_id varchar(100) NOT NULL,
		bct_description varchar(1000) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT bct_idx_behaviour_change UNIQUE (bct_id),
		PRIMARY KEY (behaviour_change_id));`)
	if err != nil {
		lg.Println(err.Error())
	}
	_, err = ac_bct.Exec()
	if err != nil {
		lg.Println("Error -", err.Error())
	}

}

func BehaviourChangeInterventionsTableCreation(db *sql.DB) {
	ac_bct, err := db.Prepare(`CREATE TABLE IF NOT EXISTS ac_behaviour_change_notifications (behaviour_notification_id int unsigned NOT NULL AUTO_INCREMENT, 
		bcn_category varchar(100) NOT NULL, 
		bcn_category_desc varchar(255),
		bcn_group varchar(100),
		bcn_group_description varchar(255),
		bcn_trigger_event varchar(255),
		app_route varchar(255),
		bcn_id varchar(100) NOT NULL,
		bcn_message varchar(2555),
		bct_1 varchar(100),
		bct_2 varchar(100),
		bct_3 varchar(100),
		bct_4 varchar(100),
		alcochange_theme BOOLEAN DEFAULT FALSE,
		frames_feedback BOOLEAN DEFAULT FALSE,
		frames_responsibility BOOLEAN DEFAULT FALSE,
		frames_advice BOOLEAN DEFAULT FALSE,
		frames_menu BOOLEAN DEFAULT FALSE,
		frames_empathy BOOLEAN DEFAULT FALSE,
		frames_support_and_selfefficacy BOOLEAN DEFAULT FALSE,
		develop_discrepancy BOOLEAN DEFAULT FALSE,
		assessment BOOLEAN DEFAULT FALSE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT bcn_idx_behaviour_change UNIQUE (bcn_id),
		PRIMARY KEY (behaviour_notification_id));`)
	if err != nil {
		lg.Println(err.Error())
	}
	_, err = ac_bct.Exec()
	if err != nil {
		lg.Println("Error -", err.Error())
	}

}

func PatientEngagementReminderTableCreation(db *sql.DB) {
	acIt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS ac_intervention_type (id int unsigned NOT NULL AUTO_INCREMENT,
		intervention_type_id int unsigned NOT NULL UNIQUE,
		intervention_type varchar(255) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id, intervention_type_id));`)
	if err != nil {
		lg.Println(err.Error())
	}
	_, err = acIt.Exec()
	if err != nil {
		lg.Println("Error -", err.Error())
	}

	patientEngagementReminderTable, errN := db.Prepare(`CREATE TABLE IF NOT EXISTS
		ac_patient_engagement_reminder (id int unsigned NOT NULL AUTO_INCREMENT,
		user_id bigint unsigned NOT NULL,
		CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES user(user_id),
		user_uuid varchar(100) NOT NULL,
	    intervention_type_id int unsigned NOT NULL,
		CONSTRAINT fk_ac_intervention_type FOREIGN KEY (intervention_type_id) REFERENCES ac_intervention_type(intervention_type_id),
		notification_id int unsigned NOT NULL,
		patient_engagement_time varchar(100) NOT NULL,
		message_shown varchar(1000),
		user_action varchar(100) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		CONSTRAINT reminderx_patient_engagement UNIQUE (notification_id,patient_engagement_time),
		PRIMARY KEY (id));`)
	if errN != nil {
		lg.Println(errN.Error())
	}
	_, err = patientEngagementReminderTable.Exec()
	if err != nil {
		lg.Println(err.Error())
	}

}

func BenefitTherapy(db *sql.DB) {

	ac_bt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS ac_benefit_therapy (id int unsigned NOT NULL AUTO_INCREMENT, 
		message varchar(255) NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id));`)
	if err != nil {
		lg.Println(err.Error())
	}
	_, err = ac_bt.Exec()
	if err != nil {
		lg.Println("Error -", err.Error())
	}
}
