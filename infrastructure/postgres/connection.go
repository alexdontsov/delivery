package postgres

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/sarulabs/di"
	otgorm "github.com/smacker/opentracing-gorm"
	"github.com/spf13/viper"
)

func buildPostgresConnection() func(container di.Container) (i interface{}, err error) {
	return func(ctn di.Container) (i interface{}, err error) {
		viper.AutomaticEnv()

		config := pg_sugar.FromEnv()

		conn, err := gorm.Open("postgres", config.FormatGORM())
		if err != nil {
			return nil, err
		}

		otgorm.AddGormCallbacks(conn)

		return conn, err
	}
}

func closePostgresConnection(_ context.Context) func(obj interface{}) error {
	return func(obj interface{}) error {
		if err := obj.(*gorm.DB).Close(); err != nil {
			return err
		}

		return nil
	}
}
