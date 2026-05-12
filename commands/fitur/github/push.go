package github

import (
	"bill/commands/fitur/system"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	GithubToken string `json:"github_token"`
}

func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".bill_config.json")
}

func loadConfig() Config {
	var cfg Config
	file, err := os.Open(getConfigPath())
	if err != nil {
		return cfg
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&cfg)
	return cfg
}

func saveConfig(cfg Config) error {
	dir := filepath.Dir(getConfigPath())
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	file, err := os.Create(getConfigPath())
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(cfg)
}

func setupGithubToken() string {
	fmt.Println("\n🔑 Halo! Sepertinya ini pertama kali kamu pakai Bill.")
	fmt.Println("Biar aku bisa bantu push ke GitHub, aku butuh Personal Access Token (PAT) kamu.")
	fmt.Println("\nLangkah-langkah ambil token:")
	fmt.Println("1. Buka: https://github.com/settings/tokens")
	fmt.Println("2. Klik 'Generate new token (classic)'")
	fmt.Println("3. Kasih nama (misal: 'Bill CLI')")
	fmt.Println("4. Centang bagian 'repo' (paling atas)")
	fmt.Println("5. Klik 'Generate token' di bawah")
	fmt.Println("6. Copy token-nya dan paste di sini.")

	token := system.ReadInput("\nInput GitHub Token: ")
	token = strings.TrimSpace(token)

	if token != "" {
		cfg := loadConfig()
		cfg.GithubToken = token
		saveConfig(cfg)
		fmt.Println("✅ Token berhasil disimpan secara global!")
	}
	return token
}

func createGithubRepo(token string, repoName string, description string) (string, error) {
	apiURL := "https://api.github.com/user/repos"
	data := map[string]interface{}{
		"name":        repoName,
		"private":     false,
		"description": description,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 422 {
		return "", fmt.Errorf("nama_digunakan")
	}

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Gagal (HTTP %d): %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result["html_url"].(string), nil
}

func PushToGithub(args []string) {
	// Pastikan konfigurasi Git (email & username) sudah ada
	ensureGitConfig()

	// 0. Cek argumen 'delete' untuk reset Git
	if len(args) > 0 && args[0] == "delete" {
		confirm := system.ReadInput("⚠️  Kamu yakin ingin menghapus folder .git? Ini akan mereset konfigurasi Git di folder ini. (y/n): ")
		if confirm == "y" || confirm == "Y" {
			err := os.RemoveAll(".git")
			if err != nil {
				fmt.Printf("❌ Gagal menghapus folder .git: %v\n", err)
			} else {
				fmt.Println("✅ Folder .git berhasil dihapus. Konfigurasi Git telah di-reset.")
			}
		} else {
			fmt.Println("❌ Proses pembatalan penghapusan.")
		}
		return
	}

	// 1. Context Discovery: Cek folder .git
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("📂 Folder .git tidak ditemukan. Bill akan bantu buatkan repositori baru di GitHub.")

		cfg := loadConfig()
		token := cfg.GithubToken
		if token == "" {
			token = setupGithubToken()
		}

		if token == "" {
			fmt.Println("❌ Error: Token dibutuhkan untuk membuat repositori.")
			return
		}

		// Ambil nama folder sebagai default nama repo
		dir, _ := os.Getwd()
		defaultRepoName := filepath.Base(dir)
		repoName := system.ReadInput(fmt.Sprintf("📦 Nama repositori (default: %s): ", defaultRepoName))
		if repoName == "" {
			repoName = defaultRepoName
		}

		repoDesc := system.ReadInput(fmt.Sprintf("📝 Deskripsi repositori (default: %s): ", repoName))
		if repoDesc == "" {
			repoDesc = repoName
		}

		var repoURL string
		for {
			url, err := createGithubRepo(token, repoName, repoDesc)
			if err == nil {
				repoURL = url
				break
			}

			if err.Error() == "nama_digunakan" {
				fmt.Printf("❌ Nama '%s' sudah digunakan di GitHub kamu.\n", repoName)
				newName := system.ReadInput("📝 Masukkan nama lain: ")
				repoName = strings.TrimSpace(newName)
				if repoName == "" {
					fmt.Println("❌ Error: Nama tidak boleh kosong.")
					return
				}
			} else {
				fmt.Printf("❌ Error API: %v\n", err)
				return
			}
		}

		commitMsg := system.ReadInput("📝 Masukkan Pesan Commit awal (default: Initial commit): ")
		if commitMsg == "" {
			commitMsg = "Initial commit"
		}

		fmt.Println("📂 Inisialisasi Git repository baru...")
		runGitCommand("git", "init")

		// Gunakan token dalam remote URL agar tidak nanya password
		// Format: https://<token>@github.com/<user>/<repo>.git
		// repoURL biasanya https://github.com/<user>/<repo>
		repoPath := strings.TrimPrefix(repoURL, "https://github.com/")
		remoteWithToken := fmt.Sprintf("https://%s@github.com/%s.git", token, repoPath)

		runGitCommand("git", "remote", "add", "origin", remoteWithToken)
		runGitCommand("git", "branch", "-M", "main")

		fmt.Println("📝 Menambahkan file dan membuat commit...")
		runGitCommand("git", "add", ".")
		runGitCommand("git", "commit", "-m", commitMsg)

		fmt.Println("⬆️  Sedang push ke GitHub...")
		err := runGitCommand("git", "push", "-u", "origin", "main")
		if err != nil {
			fmt.Printf("\n❌ Gagal push ke GitHub: %v\n", err)
		} else {
			fmt.Println("\n✅ Berhasil! Repositori kamu sudah online di GitHub.")
			fmt.Printf("🔗 Link Repo: %s\n", repoURL)
		}
		return
	}

	// 2. Ambil Remote URL & Branch
	remoteURLRaw, _ := exec.Command("git", "remote", "get-url", "origin").Output()
	remoteURL := strings.TrimSpace(string(remoteURLRaw))
	if remoteURL == "" {
		fmt.Println("❌ Error: Remote 'origin' tidak ditemukan. Tambahkan remote dulu dengan: git remote add origin [link]")
		return
	}

	branchRaw, _ := exec.Command("git", "branch", "--show-current").Output()
	currentBranch := strings.TrimSpace(string(branchRaw))
	if currentBranch == "" {
		currentBranch = "main" // Default jika tidak terdeteksi
	}

	// 3. Credential Check & Auto-config (The "Arch" Special)
	helperRaw, _ := exec.Command("git", "config", "credential.helper").Output()
	helper := strings.TrimSpace(string(helperRaw))
	if helper == "" {
		fmt.Println("🔐 Mengaktifkan credential helper store agar tidak perlu ngetik password terus...")
		runGitCommand("git", "config", "--global", "credential.helper", "store")
	}

	// 4. Staging Check & Auto-commit
	status, _ := exec.Command("git", "status", "--short").Output()
	if len(status) > 0 {
		fmt.Println("📝 Terdeteksi perubahan yang belum di-commit.")
		message := "Auto push from Bill"
		if len(args) > 0 {
			message = strings.Join(args, " ")
		} else {
			inputMsg := system.ReadInput(fmt.Sprintf("Masukkan pesan commit (default: %s): ", message))
			if inputMsg != "" {
				message = inputMsg
			}
		}

		fmt.Println("📦 Menambahkan file dan membuat commit...")
		runGitCommand("git", "add", ".")
		runGitCommand("git", "commit", "-m", message)
	}

	// 5. Final Push
	fmt.Printf("🚀 Pushing to GitHub (%s branch %s)...\n", remoteURL, currentBranch)
	output, err := runGitCombinedOutput("git", "push", "origin", currentBranch)

	if err != nil {
		// Cek apakah error karena push rejected (perlu pull)
		if strings.Contains(output, "rejected") || strings.Contains(output, "non-fast-forward") {
			fmt.Println("\n⚠️  Push ditolak karena ada perubahan di GitHub yang belum kamu ambil (pull).")
			answer := system.ReadInput("🤔 Mau Bill bantu lakukan pull & push ulang? (y/n): ")
			if answer == "y" || answer == "Y" {
				fmt.Println("⬇️  Sedang mengambil perubahan (pull)...")
				_, pullErr := runGitCombinedOutput("git", "pull", "origin", currentBranch)
				if pullErr != nil {
					fmt.Println("❌ Gagal melakukan pull. Mungkin ada konflik file yang harus diperbaiki manual.")
					return
				}
				fmt.Println("✅ Pull berhasil! Mencoba push kembali...")
				_, pushRetryErr := runGitCombinedOutput("git", "push", "origin", currentBranch)
				if pushRetryErr != nil {
					fmt.Printf("❌ Gagal push ulang: %v\n", pushRetryErr)
				} else {
					fmt.Println("\n✅ Selesai! Kode kamu sudah aman di GitHub setelah sinkronisasi.")
				}
				return
			}
		}

		fmt.Printf("\n❌ Gagal push ke GitHub: %v\n", err)
		fmt.Println("Pastikan token kamu benar dan punya akses ke repo ini.")
	} else {
		fmt.Println("\n✅ Selesai! Kode kamu sudah aman di GitHub.")
		fmt.Printf("🔗 Link Repo: %s\n", remoteURL)
	}
}

func ensureGitConfig() {
	// Cek user.email
	emailRaw, _ := exec.Command("git", "config", "user.email").Output()
	email := strings.TrimSpace(string(emailRaw))
	if email == "" {
		inputEmail := system.ReadInput("📧 GitHub Email belum diatur. Masukkan email GitHub kamu: ")
		if inputEmail != "" {
			runGitCommand("git", "config", "--global", "user.email", inputEmail)
			fmt.Println("✅ Email berhasil disimpan!")
		}
	}

	// Cek user.name
	nameRaw, _ := exec.Command("git", "config", "user.name").Output()
	name := strings.TrimSpace(string(nameRaw))
	if name == "" {
		inputName := system.ReadInput("👤 GitHub Username belum diatur. Masukkan username GitHub kamu: ")
		if inputName != "" {
			runGitCommand("git", "config", "--global", "user.name", inputName)
			fmt.Println("✅ Username berhasil disimpan!")
		}
	}
}

func runGitCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runGitCombinedOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	output, err := cmd.CombinedOutput()
	fmt.Print(string(output)) // Tetap tampilkan output ke user
	return string(output), err
}
