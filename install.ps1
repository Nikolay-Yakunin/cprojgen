$currentUser = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
$isAdmin = $currentUser.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

if (-not $isAdmin) { 
    Write-Host "Please run as administrator"
    exit
}

$binaryName = "cprojgen.exe"
$destDir = "$env:ProgramFiles\cprojgen"

if (-not (Test-Path $destDir)) {
    New-Item -ItemType Directory -Force -Path $destDir | Out-Null
}

Write-Host "Building the binary..."
make build-windows

if (-not (Test-Path "build\windows\$binaryName")) {
    Write-Host "Failed to build the binary."
    exit 1
}

Write-Host "Copying binary to $destDir..."
Copy-Item "build\windows\$binaryName" "$destDir\" -Force

if ($?) {
    Write-Host "Binary successfully installed to $destDir"

    $currentPath = [System.Environment]::GetEnvironmentVariable("PATH", "Machine")
    if (-not $currentPath.Contains($destDir)) {
        Write-Host "Adding $destDir to PATH..."
        [System.Environment]::SetEnvironmentVariable("PATH", "$currentPath;$destDir", "Machine")
        Write-Host "Restart your terminal or system to apply changes."
    } else {
        Write-Host "$destDir is already in PATH"
    }
} else {
    Write-Host "Failed to copy binary to $destDir"
    exit 1
}
