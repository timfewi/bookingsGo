@echo off
echo Building the Go project...
go build -o bookingsApp.exe ./cmd/web/.
if %ERRORLEVEL% neq 0 (
  echo Build failed, exiting.
  exit /b %ERRORLEVEL%
)

echo Running the application...
bookingsApp.exe
