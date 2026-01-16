@echo off
REM NDT Logo Display - Run Latest Version
REM This script will pull the latest version and run it

setlocal

REM Get the script directory
set SCRIPT_DIR=%~dp0
cd /d "%SCRIPT_DIR%"

REM Pull latest from git repository
echo Pulling latest version from git...
cd /d "%SCRIPT_DIR%..\.."
git pull || (
    echo Warning: Git pull failed. Continuing with local version...
)

cd /d "%SCRIPT_DIR%"

REM Run the latest version
REM Version is updated by Makefile during build
echo Running NDT Logo Display v0.0.1.80...

REM Check if the versioned exe exists, if not find the latest
if exist "v0.0.1.80.exe" (
    start "" "v0.0.1.80.exe"
) else (
    echo Version v0.0.1.80.exe not found. Searching for latest version...
    for /f "delims=" %%f in ('dir /b /o-d v*.exe 2>nul') do (
        echo Found latest version: %%f
        start "" "%%f"
        goto :done
    )
    echo Error: No version files found!
    exit /b 1
)

:done
endlocal
