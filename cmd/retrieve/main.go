package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)

		os.Exit(1)
	}
}

const baseURL = "https://raw.githubusercontent.com/raspberrypi/firmware/c6d56567ff6ef17fd85159770f22abcf2c5953ed/boot/"

var files = []string{
	"LICENCE.broadcom",
	"bootcode.bin",
	"fixup.dat",
	"fixup4.dat",
	"fixup4cd.dat",
	"fixup4db.dat",
	"fixup4x.dat",
	"fixup_cd.dat",
	"fixup_db.dat",
	"fixup_x.dat",
	"start.elf",
	"start4.elf",
	"start4cd.elf",
	"start4db.elf",
	"start4x.elf",
	"start_cd.elf",
	"start_db.elf",
	"start_x.elf",
}

func run() error {
	dstFolder := filepath.Join(".", "dist")
	os.RemoveAll(dstFolder) // ignore any error
	if err := os.MkdirAll(dstFolder, 0755); err != nil {
		return err
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	ch := make(chan string, len(files))
	for _, file := range files {
		ch <- file
	}
	close(ch)

	group, ctx := errgroup.WithContext(ctx)
	const workers = 4
	for i := 0; i < workers; i++ {
		group.Go(func() error {
			for file := range ch {
				err := downloadFile(baseURL+file, filepath.Join(dstFolder, file))
				if err != nil {
					return err
				}
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(dstFolder, "placeholder.go"), []byte(`package dist

// empty package so we can use the go tool with this repository
`), 0755); err != nil {
		return err
	}

	return nil
}

func downloadFile(src, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(src)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return out.Close()
}
