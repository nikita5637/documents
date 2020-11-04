package apiserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()
	assert.NotNil(t, config)
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		BindAddr          string
		LogLevel          string
		MinDocsPerRequest int64
		MaxDocsPerRequest int64
		DBAddr            string
		DBPort            uint16
		DBUser            string
		DBPassword        string
		DBName            string
		MigrationsDir     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Invalid dbaddr(empty)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbaddr(invalid)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "invalid",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbport(invalid)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            0,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbuser(empty)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "",
				DBPassword:        "Password",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbuser(invalid)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "---",
				DBPassword:        "Password",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbpassword(empty)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbpassword(invalid)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        string(byte(0xa)),
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbname(empty)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid dbname(invalid)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "-",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "MinDocsPerRequest less than 1",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "db",
				MinDocsPerRequest: 0,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "MinDocsPerRequest more than 100",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "db",
				MinDocsPerRequest: 101,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "MaxDocsPerRequest less than 1",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "db",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 0,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "MaxDocsPerRequest more than 100",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "db",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 101,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "MaxDocsPerRequest less than MinDocsPerRequest",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "db",
				MinDocsPerRequest: 10,
				MaxDocsPerRequest: 1,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Valid config",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "debug",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "db",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: false,
		},
		{
			name: "Invalid loglevel",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "invalid",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "/tmp/",
			},
			wantErr: true,
		},
		{
			name: "Invalid migrations dir(empty)",
			fields: fields{
				BindAddr:          ":8080",
				LogLevel:          "invalid",
				DBAddr:            "127.0.0.1",
				DBPort:            5432,
				DBUser:            "Username",
				DBPassword:        "Password",
				DBName:            "docs",
				MinDocsPerRequest: 1,
				MaxDocsPerRequest: 10,
				MigrationsDir:     "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BindAddr:          tt.fields.BindAddr,
				LogLevel:          tt.fields.LogLevel,
				MinDocsPerRequest: tt.fields.MinDocsPerRequest,
				MaxDocsPerRequest: tt.fields.MaxDocsPerRequest,
				DBAddr:            tt.fields.DBAddr,
				DBPort:            tt.fields.DBPort,
				DBUser:            tt.fields.DBUser,
				DBPassword:        tt.fields.DBPassword,
				DBName:            tt.fields.DBName,
				MigrationsDir:     tt.fields.MigrationsDir,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
