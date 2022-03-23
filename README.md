# Firmware for Raspberry Pi, for usage in gokrazy

Usage

```
GOARCH=arm ./gokr-packer \
		-kernel_package=github.com/oliverpool/kernel-rpi-os-32/dist \
		-firmware_package=github.com/oliverpool/firmware-rpi/dist \
		github.com/gokrazy/hello
```

## Manual retrieval

```
go run cmd/retrieve/main.go
```

It will retrieve the firmware from https://github.com/raspberrypi/firmware to the `dist` folder.

## Licenses

- `start*.elf`, `fixup*.dat` and `bootcode.bin` are the GPU firmwares and bootloader. Their licence is described in `dist/LICENCE.broadcom`.
- The rest of the repository is released under BSD 3-Clause License (see `LICENSE`)
