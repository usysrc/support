#!/bin/sh

# Create top-level directories
mkdir cmd pkg internal

# Create cmd subdirectories
mkdir -p cmd/service1
mkdir -p cmd/service2

# Create internal subdirectories
mkdir -p internal/api
mkdir -p internal/database

# Create pkg subdirectories
mkdir -p pkg/config
mkdir -p pkg/logger
mkdir -p pkg/server

# Create empty files in each directory
for dir in cmd/service1 cmd/service2 internal/api internal/database pkg/config pkg/logger pkg/server; do
  touch "${dir}/.empty"
done