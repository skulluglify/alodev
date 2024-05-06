#!powershell

$WorkDir = (Get-Location).Path
$BuildDir = Join-Path -Path $WorkDir -ChildPath "build"

Set-Location -Path $PSScriptRoot

if (Test-Path -Path $BuildDir) {
    Remove-Item -Path $BuildDir -Recurse -Force
}

Set-Location -Path $WorkDir
