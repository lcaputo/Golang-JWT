package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DATABASE_URL = "postgres://admin:password@localhost:5432/postgres?sslmode=disable"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})
	if error != nil {
		log.Fatal("failed to connect database")
	} else {
		log.Println("database connected")
	}
}

//
//type key int
//
//var (
//	dbKey key
//	conn  *gorm.DB
//	once  sync.Once
//)
//
//func logMode() logger.LogLevel {
//	value := config.FindEnvOrDefault("DB_LOGS", "silent")
//	switch value {
//	case "info":
//		return logger.Info
//	case "error":
//		return logger.Error
//	case "warn":
//		return logger.Warn
//	}
//	return logger.Silent
//}
//
//func PrepareConnection() (err error) {
//	once.Do(func() {
//		host := config.FindEnvOrDefault("DB_HOST", "localhost")
//		port := config.FindEnvOrDefault("DB_PORT", "5432")
//		user := config.FindEnvOrDefault("DB_USER", "admin")
//		password := config.FindEnvOrDefault("DB_PASSWORD", "password")
//		dbname := config.FindEnvOrDefault("DB_NAME", "postgres")
//
//		const layer = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
//		dsn := fmt.Sprintf(layer, host, user, password, dbname, port)
//
//		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
//			SkipDefaultTransaction: true,
//			Logger:                 logger.Default.LogMode(logMode()),
//			NamingStrategy: schema.NamingStrategy{
//				SingularTable: true,
//			},
//		})
//	})
//	return err
//}
//
//func GormMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		ctx := WithConnection(c.Request().Context())
//		c.SetRequest(c.Request().WithContext(ctx))
//		return next(c)
//	}
//}
//
//func WithConnection(ctx context.Context) context.Context {
//	return context.WithValue(ctx, dbKey, conn.WithContext(ctx))
//}
//
//func Conn(ctx context.Context) *gorm.DB {
//	value := ctx.Value(dbKey)
//	if value == nil {
//		panic("connection value not found with dbKey")
//	}
//	conn, ok := value.(*gorm.DB)
//	if !ok {
//		panic("connection invalid type")
//	}
//	return conn
//}
//func Transaction(ctx context.Context, f func(tx *gorm.DB) error) (err error) {
//	tx := Conn(ctx).Begin()
//	defer func() {
//		if r := recover(); r != nil {
//			if err := tx.Rollback(); err != nil {
//				log.Println("rollback error")
//			}
//			err = errors.New(fmt.Sprint(r))
//		}
//	}()
//	if err := f(tx); err != nil {
//		return err
//	}
//	if err := tx.Commit().Error; err != nil {
//		return errors.New("commit error")
//	}
//	return nil
//}
