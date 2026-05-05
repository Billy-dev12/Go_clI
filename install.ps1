# =================================================================
#  Bill CLI Installer for Windows
#  Repository: https://github.com/Billy-dev12/Go_clI
# =================================================================

$ErrorActionPreference = "Stop"

$REPO = "Billy-dev12/Go_clI"
$BINARY_NAME = "bill.exe"
$INSTALL_DIR = Join-Path $env:LOCALAPPDATA "bill-tool"

# Create install directory
if (!(Test-Path $INSTALL_DIR)) {
    New-Item -ItemType Directory -Path $INSTALL_DIR | Out-Null
}

Write-Host "🚀 Finding latest version of Bill CLI..." -ForegroundColor Cyan

# Get latest release tag from GitHub API
$releaseInfo = Invoke-RestMethod -Uri "https://api.github.com/repos/$REPO/releases/latest"
$latestVersion = $releaseInfo.tag_name

if (-not $latestVersion) {
    Write-Host "❌ Could not find latest release." -ForegroundColor Red
    exit 1
}

Write-Host "📦 Latest version: $latestVersion" -ForegroundColor Green

# Determine architecture
$arch = "amd64" # Default for most Windows users
if ($env:PROCESSOR_ARCHITECTURE -eq "ARM64") { $arch = "arm64" }

$downloadUrl = "https://github.com/$REPO/releases/download/$latestVersion/bill-windows-$arch.exe"
$targetPath = Join-Path $INSTALL_DIR $BINARY_NAME

Write-Host "📥 Downloading binary..." -ForegroundColor Cyan
Invoke-WebRequest -Uri $downloadUrl -OutFile $targetPath

Write-Host "⚙️  Updating PATH environment variable..." -ForegroundColor Cyan

# Update User PATH
$oldPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($oldPath -notlike "*$INSTALL_DIR*") {
    $newPath = "$INSTALL_DIR;" + $oldPath.TrimEnd(';')
    [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "✅ PATH updated successfully." -ForegroundColor Green
} else {
    Write-Host "ℹ️  PATH already contains $INSTALL_DIR." -ForegroundColor Gray
}

Write-Host "" -ForegroundColor White
Write-Host "✅ Bill CLI $latestVersion has been installed!" -ForegroundColor Green
Write-Host "---------------------------------------------------------" -ForegroundColor Yellow
Write-Host "💡 IMPORTANT: PLEASE RESTART YOUR TERMINAL (OR VS CODE)" -ForegroundColor White
Write-Host "💡 Type 'bill help' to get started." -ForegroundColor White
Write-Host "---------------------------------------------------------" -ForegroundColor Yellow
