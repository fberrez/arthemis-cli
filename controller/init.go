package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/arthemis-minecraft/arthemis-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Initialize imports all required plugins (which are listed in the
// config file) and clone running config files from the given
// git repository.
func Initialize(cmd *cobra.Command, args []string) {
	// target is the path where plugins will be installed
	target := cmd.Flag("target").Value.String()
	var config config.Config

	// Unmarshalls config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Used to check if the target exists or not
	_, err := os.Stat(target)

	// Checks if target exist
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist. Making it...\n", target)
		if err := os.Mkdir(target, os.ModePerm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Get informations about the target
	fi, err := os.Stat(target)
	mode := fi.Mode()

	// Checks if target is a directory
	if !mode.IsDir() {
		fmt.Println("target must be a directory")
		os.Exit(1)
	}

	// Get `force` flag value
	force, err := strconv.ParseBool(cmd.Flag("force").Value.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// For each plugins, it downloads it and saves it in target
	for _, plugin := range config.Plugins {
		filename := fmt.Sprintf("%s%s", plugin.Name, ".jar")
		filepath := path.Join(target, filename)

		_, err := os.Stat(filepath)
		// If filepath already exists, it checks if the `force` flag
		// is activated or not. If it is not, it does nothing.
		// Else it replaces the existent file.
		if !os.IsNotExist(err) {
			msg := fmt.Sprintf("`%s` already exists.", filepath)
			if !force {
				fmt.Printf("%s Do nothing.\n", msg)
				continue
			}

			fmt.Printf("%s Replacing it.\n", msg)
		}

		fmt.Printf("Downloading %s into `%s`\n", plugin.Name, target)

		if err := downloadFile(filepath, plugin.Url); err != nil {
			fmt.Printf("an error occured while download %s: %s\n", plugin.Name, err)
		}
	}

}

// downloadFile downloads a url to a local file.
func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
