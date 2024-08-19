package app

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ffauzann/CAI-account/internal/model"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/common/config/vault"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
	Cache    Cache
	App      *model.AppConfig
}

func (c *Config) Setup() {
	c.readConfig()

	err := c.Server.Logger.init()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = c.Database.prepare()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = c.Cache.prepare()
	if err != nil {
		log.Fatal(err)
		return
	}

	util.SetValidator()
}

func (c *Config) readConfig() {
	addr := os.Getenv("VAULT_ADDR")
	token := os.Getenv("VAULT_TOKEN")
	mountPath := os.Getenv("VAULT_MOUNT_PATH")

	if addr == "" || token == "" {
		c.readConfigFile("user.config.yaml")
		c.readConfigFile("auth.config.yaml")
	} else {
		c.readConfigVault(addr, token, mountPath, "user")
		c.readConfigVault(addr, token, mountPath, "auth")
	}
}

func (c *Config) readConfigFile(name string) {
	viper.SetConfigName(name)             // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./internal/app") // optionally look for config in the working directory
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		log.Fatalln(err)
		return
	}

	viper.Unmarshal(c)
}

func (c *Config) readConfigVault(addr, token, mountPath, secret string) {
	vaultConfig, err := vault.GetConfig(secret,
		vault.WithAddress(addr),
		vault.WithToken(token),
		vault.WithMountPath(mountPath),
		vault.WithKVVersion(2), //nolint
	)
	if err != nil {
		log.Println(err)
		return
	}

	b, err := json.Marshal(vaultConfig.GetAll())
	if err != nil {
		log.Println(err)
		return
	}

	err = json.Unmarshal(b, c)
	if err != nil {
		log.Println(err)
		return
	}
}
