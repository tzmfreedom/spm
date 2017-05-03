$name = "spm"
$bin_name = "spm"
$version = "0.2.1"
$github_user = "tzmfreedom"
If ($Env:PROCESSOR_ARCHITECTURE -match "64") {
  $arch = "amd64"
} Else {
  $arch = "386"
}
$archive_file = "${name}-${version}-windows-${arch}.zip"
$url = "https://github.com/${github_user}/${name}/releases/download/v${version}/${archive_file}"
$dest_dir = "$Env:APPDATA\${name}"

wget $url -OutFile "${dest_dir}\tmp.zip"
Expand-Archive -Path "${dest_dir}\tmp.zip" -DestinationPath "${dest_dir}\tmp" -Force
if (!(Test-Path -path "${dest_dir}\bin")) {  New-Item "${dest_dir}\bin" -Type Directory }
mv "${dest_dir}\tmp\${bin_name}" "${dest_dir}\bin\${bin_name}.exe" -Force
rm "${dest_dir}\tmp.zip" -Force
rm "${dest_dir}\tmp" -Force -Recurse

# ensure added to path in registry
$reg_path = [Environment]::GetEnvironmentVariable("PATH", "User")
If ($reg_path -notcontains $dest_dir) {
  [Environment]::SetEnvironmentVariable("PATH", $reg_path + ";" + $dest_dir, "User")
}
# ensure added to path for current session
If ($Env:Path -notcontains $dest_dir) {
  $Env:Path += ";" + $dest_dir
}
Write-Host "spm installed. Run 'spm help' to try it out."
