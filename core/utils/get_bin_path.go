package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GetBinPath(name string) (string, error) {
	if p, err := exec.LookPath(name); err == nil {
		return p, nil
	}

	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		for _, p := range []string{
			filepath.Join(exeDir, name),
			filepath.Join(exeDir, "bin", name),
		} {
			if info, err := os.Stat(p); err == nil && !info.IsDir() {
				return p, nil
			}
		}
	}

	paths := []string{
		"./" + name,
		filepath.Join(".", "bin", name),
		filepath.Join("/usr/bin", name),
		filepath.Join("/usr/local/bin", name),
	}

	for _, p := range paths {
		if info, err := os.Stat(p); err == nil && !info.IsDir() {
			return filepath.Abs(p)
		}
	}

	return "", fmt.Errorf("binary %q not found", name)
}
