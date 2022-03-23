# Firmare for Raspberry Pi, for usage in gokrazy

Usage

```
GOARCH=arm ./gokr-packer \
		-kernel_package=github.com/oliverpool/kernel-rpi-os-32/dist \
		-firmware_package=github.com/oliverpool/firmware-rpi/dist \
		-serial_console=disabled \
		github.com/gokrazy/hello
```

## Manual retrieval

```
go run cmd/retrieve/main.go
```

It will retrieve the firmware from https://github.com/raspberrypi/firmware to the `dist` folder.
