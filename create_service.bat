nssm install "PowerControl Daemon" "pc_daemon.exe"
nssm set "PowerControl Daemon" AppDirectory %CD%
nssm start "PowerControl Daemon"
del "%~f0"