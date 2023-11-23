package container

import (
	"fmt"
	"go-chat/internal/model"
	"go-chat/pkg/config"

	"log"
	"net"

	goSQLDriver "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase(env *config.Environment) (*gorm.DB, error) {
	var dsn string

	if env.SSH_TUNNEL {
		// Konfigurasi SSH
		sshConfig := &ssh.ClientConfig{
			User:            env.SSH_USERNAME,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth: []ssh.AuthMethod{
				ssh.Password(env.SSH_PASSWORD),
			},
		}

		sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", env.SSH_HOSTNAME, env.SSH_PORT), sshConfig)
		if err != nil {
			log.Fatalf("Failed to establish SSH connection: %v", err)
			return nil, err
		}

		// Konfigurasi MySQL dengan SSH tunnel
		goSQLDriver.RegisterDial("mysql+ssh", func(addr string) (net.Conn, error) {
			return sshConn.Dial("tcp", addr)
		})

		dsn = fmt.Sprintf("%s:%s@mysql+ssh(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", env.MYSQL_USERNAME, env.MYSQL_PASSWORD, env.MYSQL_HOST, env.MYSQL_PORT, env.MYSQL_DB_NAME, env.MYSQL_DB_LOC)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", env.MYSQL_USERNAME, env.MYSQL_PASSWORD, env.MYSQL_HOST, env.MYSQL_PORT, env.MYSQL_DB_NAME, env.MYSQL_DB_LOC)
	}

	db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
		return nil, err
	}

	// Auto Migrate tabel sesuai dengan model Anda
	if env.MYSQL_DB_MIGRATE {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			log.Fatalf("Failed to auto migrate database: %v", err)
			return nil, err
		}
	}

	return db, nil
}
