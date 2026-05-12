package github

import (
	"bill/commands/fitur/system"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func HandleBranch(args []string) {
	if len(args) == 0 {
		listBranches()
		return
	}

	targetBranch := args[0]
	switchBranch(targetBranch)
}

func listBranches() {
	fmt.Println("\n🌿 DAFTAR BRANCH:")
	out, err := exec.Command("git", "branch").Output()
	if err != nil {
		fmt.Printf("❌ Gagal mengambil daftar branch: %v\n", err)
		return
	}
	fmt.Println(string(out))

	currentBranchRaw, _ := exec.Command("git", "branch", "--show-current").Output()
	currentBranch := strings.TrimSpace(string(currentBranchRaw))

	fmt.Printf("\n💡 Branch saat ini: %s\n", currentBranch)
	fmt.Println("Gunakan 'bill branch [nama]' untuk pindah branch.")
}

func switchBranch(name string) {
	fmt.Printf("🔄 Mencoba pindah ke branch '%s'...\n", name)

	// Cek apakah branch ada
	cmd := exec.Command("git", "checkout", name)
	out, err := cmd.CombinedOutput()

	if err != nil {
		if strings.Contains(string(out), "did not match any file") {
			fmt.Printf("⚠️  Branch '%s' tidak ditemukan.\n", name)
			answer := system.ReadInput(fmt.Sprintf("🤔 Mau buat branch baru '%s'? (y/n): ", name))
			if strings.ToLower(answer) == "y" {
				createBranch(name)
			}
		} else {
			fmt.Printf("❌ Gagal pindah branch: %s\n", string(out))
		}
	} else {
		fmt.Printf("✅ Berhasil pindah ke branch '%s'.\n", name)
	}
}

func createBranch(name string) {
	fmt.Printf("🌱 Membuat branch baru '%s'...\n", name)
	cmd := exec.Command("git", "checkout", "-b", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ Gagal membuat branch: %v\n", err)
	} else {
		fmt.Printf("✅ Branch '%s' berhasil dibuat dan aktif.\n", name)
	}
}
