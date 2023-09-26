package database
import(
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
  	"gorm.io/gorm"
	"log"
	"fmt"
	"os"
)



var DB *gorm.DB

func ConnectDb(){
	var err error
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error cargando variables de entorno: %v", err)
    }

	dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword +
        " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	

	if err != nil {
		log.Fatal("Failed to connecto to the database! \n", err.Error())
		// os.exit(2)
	}
	fmt.Println("DB Connected!")
}
	