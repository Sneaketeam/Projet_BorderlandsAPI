@echo off
echo ---------------------------------------------
echo   NETTOYAGE ET DEMARRAGE DE XAMPP (GIOVANNI)
echo ---------------------------------------------


cd /d "C:\xampp\mysql\data"


if exist "ib_logfile0" del "ib_logfile0"
if exist "ib_logfile1" del "ib_logfile1"
if exist "aria_log.*" del "aria_log.*"


echo.
echo [OK] Fichiers logs supprimes avec succes.
echo [OK] Fichiers aria_log supprimes.
echo.


echo Lancement du panneau XAMPP...
start C:\xampp\xampp-control.exe


exit