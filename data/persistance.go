/*
Copyright © 2024 Yannik Schiebelhut <yannik.schiebelhut@gmail.com>
*/
package data

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var db DB
var usr, _ = user.Current()
var dbFilePath = filepath.Join(usr.HomeDir, ".config/sotp/db.yaml")

type DB struct {
	Secrets map[string]string `yaml:"secrets"`
}

func init() {
	dbFile, err := os.ReadFile(dbFilePath)
	if err != nil {
		log.Fatal("error reading db file: ", err)
	}
	err = yaml.Unmarshal(dbFile, &db)
	if err != nil {
		log.Fatal("error unmarshalling db file: ", err)
	}
}

func GetAllSecrets() ([]string, error) {
	keys := make([]string, len(db.Secrets))
	for k := range db.Secrets {
		keys = append(keys, k)
	}
	return keys, nil
}

func LookupSecret(key string) (string, error) {
	if secret, exists := db.Secrets[key]; exists {
		return secret, nil
	}
	return "", fmt.Errorf("secret with name %s does not exist", key)
}

func AddSecret(name, secret string) error {
	if _, exists := db.Secrets[name]; exists {
		return fmt.Errorf("secret with name %s already exists", name)
	}
	db.Secrets[name] = secret
	return persistDB()
}

func RemoveSecret(name string) error {
	delete(db.Secrets, name)
	return persistDB()
}

func persistDB() error {
	out, err := yaml.Marshal(db)
	if err != nil {
		return err
	}
	err = os.WriteFile(dbFilePath, out, 0600)
	return err
}
