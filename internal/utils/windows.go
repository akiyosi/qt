package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func QT_MSYS2() bool {
	return (os.Getenv("QT_MSYS2") == "true" || IsMsys2QtDir() || MSYSTEM() != "") && !MSYS_DOCKER()
}

func QT_MSYS2_DIR() string {
	if dir, ok := os.LookupEnv("QT_MSYS2_DIR"); ok {
		switch QT_MSYS2_ARCH() {
		case "amd64":
			return filepath.Join(dir, "mingw64")
		case "arm64":
			return filepath.Join(dir, "clangarm64")
		default:
			return filepath.Join(dir, "mingw32")
		}
	}
	// default prefix
	prefix := "msys64"
	if QT_MSYS2_ARCH() == "386" {
		prefix = "msys32"
	}
	// choose subdir by arch
	suffix := "mingw32"
	switch QT_MSYS2_ARCH() {
	case "amd64":
		suffix = "mingw64"
	case "arm64":
		suffix = "clangarm64"
	}
	return fmt.Sprintf("%v\\%v\\%v", windowsSystemDrive(), prefix, suffix)
}

func IsMsys2QtDir() bool {
	d := os.Getenv("QT_MSYS2_DIR")
	if d == "" {
		return false
	}
	if ExistsFile(filepath.Join(d, "msys2.exe")) {
		return true
	}
	if ExistsFile(filepath.Join(d, "bin", "qmake.exe")) {
		return true
	}
	return ExistsDir(d)
}

func QT_MSYS2_ARCH() string {
	if arch, ok := os.LookupEnv("QT_MSYS2_ARCH"); ok {
		return arch
	}
	switch MSYSTEM() {
	case "CLANGARM64":
		return "arm64"
	case "MINGW64":
		return "amd64"
	case "MINGW32":
		return "386"
	}
	switch runtime.GOARCH {
	case "arm64":
		return "arm64"
	case "amd64":
		return "amd64"
	default:
		return "386"
	}
}

func QT_MSYS2_STATIC() bool {
	return os.Getenv("QT_MSYS2_STATIC") == "true"
}

func MSYSTEM() string {
	return os.Getenv("MSYSTEM")
}

func MSYS_DOCKER() bool {
	_, ok := os.LookupEnv("DOCKER_MACHINE_NAME")
	return ok
}

func windowsSystemDrive() string {
	if vol, ok := os.LookupEnv("SystemDrive"); ok {
		return vol
	}
	if vol, ok := os.LookupEnv("SystemRoot"); ok {
		return filepath.VolumeName(vol)
	}
	if vol, ok := os.LookupEnv("WinDir"); ok {
		return filepath.VolumeName(vol)
	}
	return "C:"
}

func MINGWDIR() string {
	if QT_MSVC() {
		if GOARCH() == "386" {
			return "msvc2017"
		}
		return "msvc2017_64"
	}
	if QT_VERSION_NUM() >= 5122 {
		if GOARCH() == "386" {
			return "mingw73_32"
		}
	}
	return "mingw73_64"
}

func MINGWTOOLSDIR() string {
	version := "mingw730_64"
	if QT_VERSION_NUM() >= 5122 {
		if GOARCH() == "386" {
			version = "mingw730_32"
		}
	}
	path := filepath.Join(QT_DIR(), "Tools", version, "bin")
	if !ExistsDir(path) {
		path = strings.Replace(path, version, "mingw530_32", -1)
	}
	if !ExistsDir(path) {
		path = strings.Replace(path, "mingw530_32", "mingw492_32", -1)
	}
	return path
}

func QT_MSVC() bool {
	return os.Getenv("QT_MSVC") == "true"
}

func GOVSVARSPATH() string {
	if p, ok := os.LookupEnv("GOVSVARSPATH"); ok {
		return p
	}
	bits := "64"
	if GOARCH() == "386" {
		bits = "32"
	}
	return fmt.Sprintf(`%v\Program Files (x86)\Microsoft Visual Studio\2017\BuildTools\VC\Auxiliary\Build\vcvars%v.bat`, windowsSystemDrive(), bits)
}
