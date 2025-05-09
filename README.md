# Firmware for Raspberry Pi (from official bookworm repo)

This repository holds bootloader firmware files for the Raspberry Pi, downloaded from https://archive.raspberrypi.org/debian/, for usage by the [gokrazy](https://github.com/gokrazy/gokrazy) project.

To use the files in this repository (as well as the corresponding kernel), set the `KernelPackage` and `FirmwarePackage` of your gokrazy instance's `config.json`:

```jsonc
{
    // ...
    "KernelPackage": "github.com/gokrazy-community/kernel-rpi-os-32/dist",
    "FirmwarePackage": "github.com/gokrazy-community/firmware-rpi/dist"
}
```

When building, make sure to set the appropriate `GOARCH` and `GOARM` environment variables:

```
GOARCH=arm GOARM=6 gok -i <instance-name> update
```

## How does it differ from https://github.com/gokrazy/firmware ?

https://github.com/gokrazy/firmware follows the `master` branch of https://github.com/raspberrypi/firmware

Whereas this repo follows the latest release from https://archive.raspberrypi.org/debian/dists/bullseye/main/binary-armhf/.

## Manual retrieval

```
go run cmd/retrieve/main.go
```

It will retrieve the download the latest firmware files from https://archive.raspberrypi.org/debian/ and extract them to the `dist` folder.

## Licenses

- `start*.elf`, `fixup*.dat` and `bootcode.bin` are the GPU firmwares and bootloader. Their licence is described in `dist/LICENCE.broadcom`.
- The rest of the repository is released under BSD 3-Clause License (see `LICENSE`)
