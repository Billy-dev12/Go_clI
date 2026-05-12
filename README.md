# 🚀 Bill CLI - Asisten Developer CLI Terkeren

**Bill CLI** adalah command-line tool yang powerful dan user-friendly untuk developer, khususnya yang bekerja dengan Laravel. Tool ini mengotomasi workflow development, git operations, dan system setup dengan fitur interaktif yang memudahkan pekerjaan sehari-hari.

[![GitHub](https://img.shields.io/badge/GitHub-Billy--dev12%2FGo_clI-blue?logo=github)](https://github.com/Billy-dev12/Go_clI)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.22.0+-blue?logo=go)](https://golang.org)
[![Version](https://img.shields.io/badge/version-1.1.2-brightgreen)](go.mod)

---

## ✨ Fitur Utama

### 🐙 **GitHub & Git Integration**
- **Smart Push** - Push otomatis dengan init repo, commit, dan sync
- **Branch Management** - List, switch, dan manage branches dengan mudah
- **Reset Configuration** - Bersihkan konfigurasi git lokal dengan satu command
- Support khusus untuk pengguna Arch Linux

### 🔨 **Laravel Development Tools**
- **Project Creation** - Buat project Laravel baru secara instant
- **Environment Check** - Diagnose status dan dependencies PHP/Composer
- **Auto-Setup** - Setup otomatis untuk project hasil clone
- **Development Environment** - Setup lengkap dev env (PHP, Composer, MySQL, dll)
- **Development Server** - Run artisan serve dengan custom port
- **Cleanup Tools** - Bersihkan cache, logs, dan temp files Laravel

### 📦 **Build & Compilation**
- **Multi-Platform Build** - Compile ke Windows, Linux, dan macOS
- **Interactive Build** - Mode interaktif untuk build dengan pilihan platform
- **Custom Naming** - Rename binary output sesuai kebutuhan

### 💻 **System & Utilities**
- **Interactive Mode (REPL)** - Shell mode interaktif dengan command history
- **System Info** - Tampilkan informasi environment dan dependencies
- **Auto-Install** - Install otomatis ke PATH system
- **Help System** - Bantuan lengkap untuk setiap command

---

## 📥 Instalasi

### 🚀 Quick Install (Recommended)

#### 🐧 Linux / macOS
```bash
curl -fsSL https://raw.githubusercontent.com/Billy-dev12/Go_clI/main/install.sh | bash
```

#### 🪟 Windows (PowerShell)
Buka PowerShell sebagai Administrator dan jalankan:
```powershell
irm 'https://raw.githubusercontent.com/Billy-dev12/Go_clI/main/install.ps1' | iex
```

### 📦 Build dari Source

**Prasyarat:**
- Go 1.22 atau lebih baru
- Git

**Langkah-langkah:**
```bash
# Clone repository
git clone https://github.com/Billy-dev12/Go_clI.git
cd Go_clI

# Build binary
go build -o bill

# Install ke PATH
sudo mv bill /usr/local/bin/
# atau gunakan
./bill install
```

### ✅ Verifikasi Instalasi
```bash
bill help
# atau
bill --help
```

---

## 🚀 Quick Start

### Mode Interaktif (Recommended)
Jalankan tanpa arguments untuk memasuki interactive mode:
```bash
bill
```

Anda akan masuk ke shell interaktif dimana bisa mengetik commands tanpa prefix `bill`:
```
🚀 BILL CLI 1.1.2 - Asisten CLI Terkeren kamu!
bill> buat laravel myproject
bill> cek laravel
bill> ser 8080
bill> exit
```

### Mode Command (Direct)
Jalankan langsung dari terminal:
```bash
bill [command] [options]
```

---

## 📖 Command Reference

### 🐙 GitHub & Git Commands

#### **bill push** 
Smart push ke GitHub dengan auto-init, commit, dan sync
```bash
bill push
bill push "commit message"
```

#### **bill branch**
Manage git branches
```bash
bill branch              # List semua branches
bill branch [nama]       # Switch ke branch tertentu
```

#### **bill arch push**
Khusus pengguna Arch Linux dengan smart push
```bash
bill arch push
bill arch push "pesan commit"
```

#### **bill push delete**
Reset dan bersihkan konfigurasi git
```bash
bill push delete         # Remove .git dan reset config
```

---

### 🔨 Laravel Development Commands

#### **bill buat laravel [nama]**
Buat project Laravel baru dengan struktur lengkap
```bash
bill buat laravel blog
bill buat laravel ecommerce-app
```

#### **bill cek laravel**
Check status environment dan dependencies
```bash
bill cek laravel
```
Outputnya:
- ✅ PHP version
- ✅ Composer status
- ✅ Database connectivity
- ✅ Required extensions

#### **bill setup laravel**
Auto-setup untuk project hasil clone
```bash
bill setup laravel
```
Proses:
1. Check dependencies
2. Install composer packages
3. Generate `.env` file
4. Generate app key
5. Run migrations (optional)

#### **bill setup lingkungan laravel**
Setup development environment lengkap
```bash
bill setup lingkungan laravel
```
Setup:
- PHP & PHP extensions
- Composer
- MySQL/MariaDB
- Node.js & npm (untuk frontend tooling)

#### **bill ser [port]**
Run Laravel development server (artisan serve)
```bash
bill ser                 # Default port 8000
bill ser 3000           # Custom port 3000
bill ser 5173           # Frontend server
```

#### **bill cleanup**
Bersihkan cache dan log files Laravel
```bash
bill cleanup
```
Membersihkan:
- `storage/logs/`
- `bootstrap/cache/`
- `storage/app/`
- Temp files

---

### 📦 Build & Compilation Commands

#### **bill go build [nama]**
Build project Go secara interaktif
```bash
bill go build
bill go build myapp
```

Mode interaktif akan menanyakan:
- Platform target (Windows/Linux/macOS)
- Architecture (amd64/arm64)
- Output filename
- Optimization level

#### **bill build-windows [nama]**
Build untuk Windows (amd64)
```bash
bill build-windows bill.exe
bill build-windows myapp
```

#### **bill build-linux [nama]**
Build untuk Linux (amd64)
```bash
bill build-linux bill
bill build-linux myapp
```

#### **bill build-mac [nama]**
Build untuk macOS (amd64/arm64)
```bash
bill build-mac bill
bill build-mac myapp
```

---

### 💻 System & Utility Commands

#### **bill info**
Tampilkan informasi environment system
```bash
bill info
```
Menampilkan:
- OS & Architecture
- Go version
- PHP version
- Node.js version
- Git version
- Installed tools status

#### **bill install**
Install Bill CLI ke system PATH
```bash
bill install
```

#### **bill help**
Tampilkan help menu lengkap
```bash
bill help
```

#### **bill exit / bill keluar**
Keluar dari interactive mode
```bash
bill exit
# atau
bill keluar
```

---

## 💡 Contoh Penggunaan

### Scenario 1: Buat Project Laravel Baru
```bash
# Buat project
bill buat laravel myshop

# Enter project directory
cd myshop

# Run setup
bill setup laravel

# Start development server
bill ser 8000

# Di terminal lain, jalankan build tools
npm run dev
```

### Scenario 2: Clone Project Dari GitHub
```bash
# Clone repo
git clone https://github.com/user/project.git
cd project

# Auto setup dengan Bill CLI
bill setup laravel

# Jalankan server
bill ser
```

### Scenario 3: Push Code ke GitHub
```bash
# Make changes
# ... edit files ...

# Push dengan Bill CLI
bill push "add new feature"

# atau di interactive mode
bill
# > push "add new feature"
```

### Scenario 4: Build untuk Multiple Platforms
```bash
# Interactive build
bill go build myapp

# Atau direct build
bill build-windows myapp.exe  # → myapp.exe
bill build-linux myapp        # → myapp
bill build-mac myapp          # → myapp (universal/arm64)
```

### Scenario 5: Check Environment
```bash
# Check Laravel dev environment
bill cek laravel

# Check system info
bill info

# Interactive mode
bill
# > info
# > cek laravel
```

---

## 🎯 Keunggulan Bill CLI

✅ **User-Friendly** - Interface yang intuitif dan mudah dipelajari

✅ **Automation** - Otomasi workflow repetitif, hemat waktu

✅ **Interactive Mode** - Shell interaktif untuk workflow yang lebih smoothly

✅ **Multi-Platform** - Berjalan di Windows, Linux, dan macOS

✅ **Laravel Focused** - Built specifically untuk Laravel developers

✅ **Git Integration** - Smart git operations tanpa ribet

✅ **System Agnostic** - Bekerja dengan berbagai setup development

✅ **Open Source** - Kode terbuka, transparan, dan aman

✅ **Actively Maintained** - Regular updates dan improvements

✅ **Fast** - Built dengan Go, super cepat dan efficient

---

## 🛠️ Development

### Requirements
- Go 1.22+
- Git

### Setup Development Environment
```bash
# Clone repo
git clone https://github.com/Billy-dev12/Go_clI.git
cd Go_clI

# Install dependencies (jika ada)
go mod download

# Build
go build -o bill

# Test
./bill help
```

### Project Structure
```
GO_ps/
├── main.go                    # Entry point
├── go.mod                     # Module definition
├── commands/
│   └── kode/
│       └── pusat.go          # Command router/dispatcher
└── commands/fitur/
    ├── build/                 # Build commands
    │   └── builder.go
    ├── github/                # Git & GitHub integration
    │   ├── branch.go
    │   └── push.go
    ├── help/                  # Help & info
    │   └── help.go
    ├── laravel/               # Laravel automation
    │   ├── automation.go
    │   ├── buat.go
    │   ├── cek.go
    │   └── setup.go
    ├── system/                # System utilities
    │   ├── doctor.go
    │   └── setup_wizard.go
    └── version/               # Version management
        └── version.go
```

### Adding New Command
1. Create new feature folder di `commands/fitur/`
2. Implement command logic
3. Register di `commands/kode/pusat.go` (switch case)
4. Update help menu di `help/help.go`

---

## 🤝 Contributing

Kontribusi sangat diterima! Ikuti langkah-langkah berikut:

1. **Fork** repository
2. **Create feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit changes** (`git commit -m 'Add some amazing feature'`)
4. **Push to branch** (`git push origin feature/amazing-feature`)
5. **Open Pull Request**

### Report Bugs
Temukan bug? Buka [GitHub Issues](https://github.com/Billy-dev12/Go_clI/issues) dengan:
- Deskripsi jelas tentang bug
- Steps untuk reproduce
- Expected vs actual behavior
- OS dan versi Go

### Suggest Features
Ada ide fitur? Buat issue dengan label `enhancement`

---

## 📝 License

Project ini dilisensikan di bawah **MIT License** - lihat [LICENSE](LICENSE) file untuk detail.

---

## 🙏 Acknowledments

- Terima kasih kepada Laravel community
- Go standard library
- Inspirasi dari tools seperti Artisan, npm CLI, dan git

---

## 📞 Support & Resources

- 💬 **GitHub Discussions** - Tanya jawab dan diskusi
- 🐛 **GitHub Issues** - Report bugs dan request features  
- 📜 **Documentation** - Lengkap di repo
- 🌐 **Website** - [github.com/Billy-dev12/Go_clI](https://github.com/Billy-dev12/Go_clI)

---

## 👨‍💻 Author

**Billahi Robby** (Bill-dev12)
- GitHub: [@Billy-dev12](https://github.com/Billy-dev12)
- Project: [Go_clI](https://github.com/Billy-dev12/Go_clI)

---

<div align="center">

**Dibuat dengan ❤️ untuk Developer Indonesia**

Jika project ini membantu, berikan ⭐ di [GitHub](https://github.com/Billy-dev12/Go_clI)!

</div>
