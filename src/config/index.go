package config

import "os"
import "errors"
import "fmt"
import "github.com/joho/godotenv"
import "sync"


type Config struct {
    MongoURI string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() error {
    var err error
    once.Do(func(){
        _ = godotenv.Load()
        mongoURI := os.Getenv("MONGO_URI")
        if mongoURI == "" {
            err = errors.New("MONGO_URI not set. Set the variable in .env")
            return
        }
        fmt.Println("MONGO_URI:", mongoURI)
        cfg = &Config{
            MongoURI: mongoURI,
        }

    })
    return err


}

func Get() Config {
	if cfg == nil {
		panic("config not loaded: call config.Load() first")
	}
	return *cfg
}


/* This loads the mongoDb uri from the parent .env file once and makes it available for use across the application */
/* load gets called and stores the uri in the cfg variable */
/* Get returns the config struct with the mongo uri */


// TODO : add support mongo-URI flag so that init can be called with custom uris if devs wants to create a loader for diff repos with different conntection strings.