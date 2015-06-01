package main

import(
	"os"
	"fmt"
	"path/filepath"
	"git.enova.com/zsyed/args"
	"git.enova.com/zsyed/utils"
)

const(
	CHSSH_DIR = "/Users/jmcconnell1/.chssh"
	ADD_USAGE = "add name PUBLIC_KEY_FILE_PATH PRIVATE_KEY_FILE_PATH"
	COMMAND_USAGE = "Need command or profile. Example [add]."
)

func main() {
	if _, err := os.Stat(CHSSH_DIR); err != nil {
		os.Mkdir(CHSSH_DIR, os.ModePerm)
	}

	a := args.New(os.Args)

	command := a.Get(1)
	switch command {
	case "add":
		Add(CHSSH_DIR, a)
	}
}

func Add(dir string, a *args.Args) {
	if a.Size() < 5 {
		utils.Exit(ADD_USAGE)
	}
	name := a.Get(2)
	pub_key_path := a.Get(3)
	private_key_path := a.Get(4)
	profileDir := filepath.Join(dir, name)

	if _, err := os.Stat(profileDir); err != nil {
		os.Mkdir(profileDir, os.ModePerm)
	}

	to_pub_key_path := ToPath(dir, name, "id_rsa.pub")
	CreateFile(to_pub_key_path)
	utils.CopyFile(pub_key_path, to_pub_key_path, pub_key_path + " - public key")

	to_private_key_path := ToPath(dir, name, "id_rsa")
	CreateFile(to_private_key_path)
	utils.CopyFile(private_key_path, to_private_key_path, private_key_path  + " - private key")
}

func ToPath(basePath, name, remaining string) string {
	return filepath.Join(basePath, name, remaining)
}


func CreateFile(filepath string) {
	file, err := os.Create(filepath)
	file.Close()
	if err != nil {
		utils.Exit(fmt.Sprintf("Could not create %s. %s", filepath, err))
	}
}
