@echo off
echo Building the Go project...
go run  ./cmd/web/.
if %ERRORLEVEL% neq 0 (
  echo Build failed, exiting.
  exit /b %ERRORLEVEL%
)

echo Running the application...

