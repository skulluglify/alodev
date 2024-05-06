#!powershell

$WorkDir = (Get-Location).Path
$BuildDir = Join-Path -Path $WorkDir -ChildPath "build"

Set-Location -Path $PSScriptRoot

if (Test-Path -Path $BuildDir) {
    Remove-Item -Path $BuildDir -Recurse -Force
}

New-Item -ItemType Directory -Path $BuildDir -Force | Out-Null

Set-Location -Path $WorkDir

go build -o $BuildDir\app.exe .
