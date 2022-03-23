package main

import (
	"archive/tar"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/blakesmith/ar"
	"github.com/ulikunitz/xz"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)

		os.Exit(1)
	}
}

const baseURL = "https://archive.raspberrypi.org/debian/"
const packagesURL = baseURL + "dists/buster/main/binary-armhf/Packages"

func run() error {
	dstFolder := filepath.Join(".", "dist")
	os.RemoveAll(dstFolder) // ignore any error
	if err := os.MkdirAll(dstFolder, 0755); err != nil {
		return err
	}

	log.Println("checking:", packagesURL)
	bootLoaderURL := ""
	bootLoaderPrefix := "Filename: pool/main/r/raspberrypi-firmware/raspberrypi-bootloader_"
	version := ""
	versionPrefix := "Version: "
	err := scanOnlineTextFile(packagesURL, func(s string) bool {
		if strings.HasPrefix(s, versionPrefix) {
			version = s[len(versionPrefix):]
		}
		if strings.HasPrefix(s, bootLoaderPrefix) {
			bootLoaderURL = baseURL + s[len("Filename: "):]
			return true
		}
		return false
	})
	if bootLoaderURL == "" {
		if err != nil {
			return err
		}
		return errors.New("could not find bootloader URL in package list")
	}

	fmt.Println(version)
	log.Println("downloading:", bootLoaderURL)
	resp, err := http.Get(bootLoaderURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = extractFirmwareFiles(resp.Body, dstFolder)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dstFolder, "placeholder.go"), []byte(`package dist

// empty package so we can use the go tool with this repository
`), 0755); err != nil {
		return err
	}

	return nil
}

func extractFirmwareFiles(debSrc io.Reader, dstFolder string) error {

	ar := ar.NewReader(debSrc)
	var dataReader io.Reader
	for {
		header, err := ar.Next()
		if err != nil {
			return err
		}
		if header.Name == "data.tar.xz" {
			dataReader = ar
			break
		}
	}

	r, err := xz.NewReader(dataReader)
	if err != nil {
		return err
	}

	// Create a tar Reader
	tr := tar.NewReader(r)
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return err
		}
		switch hdr.Typeflag {
		// case tar.TypeDir:
		case tar.TypeReg, tar.TypeRegA:
			bootPrefix := "./boot/"
			if !strings.HasPrefix(hdr.Name, bootPrefix) {
				continue
			}
			name := hdr.Name[len(bootPrefix):]
			// write a file
			log.Println("extracting: " + name)
			err = writeFile(tr, filepath.Join(dstFolder, name))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func writeFile(r io.Reader, dst string) error {
	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}

func scanOnlineTextFile(url string, stopScanning func(string) bool) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if stopScanning(scanner.Text()) {
			break
		}
	}
	return scanner.Err()
}
