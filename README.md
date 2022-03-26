# Firmware for Raspberry Pi (from official bullseye repo)

This repository holds bootloader firmware files for the Raspberry Pi, downloaded from https://archive.raspberrypi.org/debian/, for usage by the [gokrazy](https://github.com/gokrazy/gokrazy) project.

To use the files in this repository, adjust the `-firmware_package`
of `gokr-packer`:

```
GOARCH=arm gokr-packer \
    -firmware_package=github.com/gokrazy-community/firmware-rpi/dist \
    github.com/gokrazy/hello
```

## How does it differ from https://github.com/gokrazy/firmware ?

https://github.com/gokrazy/firmware follows the `master` branch of https://github.com/raspberrypi/firmware

Whereas this repo follows the latest release from https://archive.raspberrypi.org/debian/ (bullseye as of now).

## Manual retrieval

```
go run cmd/retrieve/main.go
```

It will retrieve the download the latest firmware files from https://archive.raspberrypi.org/debian/ and extract them to the `dist` folder.

## Licenses

- `start*.elf`, `fixup*.dat` and `bootcode.bin` are the GPU firmwares and bootloader. Their licence is described in `dist/LICENCE.broadcom`.
- The rest of the repository is released under BSD 3-Clause License (see `LICENSE`)
