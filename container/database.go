package container

import (
	"fmt"
	"go-chat/config"
	"go-chat/internal/model"
	"log"
	"net"

	goSQLDriver "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(env *config.Environment) (*gorm.DB, error) {
	var dsn string

	if env.SSHTunnel {
		// Konfigurasi SSH
		sshConfig := &ssh.ClientConfig{
			User:            env.SSHUsername,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Auth: []ssh.AuthMethod{
				ssh.Password(env.SSHPasssword),
			},
		}

		sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", env.SSHHostname, env.SSHPort), sshConfig)
		if err != nil {
			log.Fatalf("Failed to establish SSH connection: %v", err)
			return nil, err
		}

		// Konfigurasi MySQL dengan SSH tunnel
		goSQLDriver.RegisterDial("mysql+ssh", func(addr string) (net.Conn, error) {
			return sshConn.Dial("tcp", addr)
		})

		dsn = fmt.Sprintf("%s:%s@mysql+ssh(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	}

	db, err := gorm.Open(mysqlDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
		return nil, err
	}

	// Auto Migrate tabel sesuai dengan model Anda
	if env.DBMigrate {
		err := db.AutoMigrate(&model.User{})
		if err != nil {
			log.Fatalf("Failed to auto migrate database: %v", err)
			return nil, err
		}
	}

	return db, nil
}
