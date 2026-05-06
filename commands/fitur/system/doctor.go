package system

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func CheckEverything() {
	fmt.Println("\n🩺 BILL CLI DOCTOR - Mengecek kesehatan sistem kamu...")
	fmt.Println(strings.Repeat("-", 50))

	// 1. OS & Architecture
	fmt.Printf("💻 OS: %s | Arch: %s\n", runtime.GOOS, runtime.GOARCH)

	// 2. Git
	checkTool("Git", "git", "--version")

	// 3. PHP (Core for Laravel features)
	checkTool("PHP", "php", "-v")

	// 4. Composer
	checkTool("Composer", "composer", "--version")

	// 5. Node.js
	checkTool("Node.js", "node", "-v")

	// 6. Go (Compiler)
	checkTool("Go", "go", "version")

	// 7. MySQL / MariaDB
	checkTool("MySQL", "mysql", "--version")

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("✅ Pengecekan selesai!")
}

func checkTool(name, cmdName, versionArg string) {
	path, err := exec.LookPath(cmdName)
	if err == nil {
		version := "Unknown"
		out, err := exec.Command(cmdName, versionArg).Output()
		if err == nil {
			lines := strings.Split(string(out), "\n")
			if len(lines) > 0 {
				version = strings.TrimSpace(lines[0])
			}
		}
		fmt.Printf("✅ %-10s: Terdeteksi (%s) -> %s\n", name, path, version)
	} else {
		fmt.Printf("❌ %-10s: Tidak ditemukan di PATH.\n", name)
	}
}
