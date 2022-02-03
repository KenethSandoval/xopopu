package pkg

import "os"

func CreateDir(dir string) (bool, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
