@echo off

SET SETUP=0
SET SERVER=0

:ARGS
IF "%1"=="" GOTO EXEC
IF "%1"=="SETUP" SET SETUP=1
IF "%1"=="SERVER" SET SERVER=1
SHIFT

GOTO ARGS

:EXEC

if "%SERVER%"=="1" (
    echo "building server module"
    SET GOOS=windows
    SET GOARCH=amd64
    SET SETUP=1

    pushd .\web
    go build
    popd
)

if "%SETUP%"=="1" (
    echo "setting up environment"
    SET GOOS=js
    SET GOARCH=wasm
    echo "copying wasm executor"
    cp "c:\Go\misc\wasm\wasm_exec.js" ".\web\wasm_exec.js"
)

echo "building module..."
pushd .\module

go build -o ..\web\main.wasm

popd
