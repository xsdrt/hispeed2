package hispeed2

import (
	"fmt"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type HiSpeed2 struct {
	AppName string
	Debug   bool
	Version string
}

func (h *HiSpeed2) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := h.Init(pathConfig)
	if err != nil {
		return err
	}

	err = h.checkDotEnv(rootPath) // Check the root path of the application (or TestApp during development)...
	if err != nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (h *HiSpeed2) Init(p initPaths) error {
	root := p.rootPath //holds the full root path to the web app...
	for _, path := range p.folderNames {
		// create the folder if it doesn't exist...
		err := h.CreateDirIfNotExist(root + "/" + path) // creates the dir if not exists...
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *HiSpeed2) checkDotEnv(path string) error {
	err := h.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path)) // look into the root lvl of app to see if the env file exist, if not return an err...
	if err != nil {
		return err
	}
	return nil
}
