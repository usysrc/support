#!/bin/sh

# Create top-level directories
mkdir cmd pkg internal migrations

# Create cmd subdirectories
mkdir -p cmd/support

# Create internal subdirectories
mkdir -p internal/models
mkdir -p internal/api
mkdir -p internal/database
mkdir -p internal/services

# Create pkg subdirectories
mkdir -p pkg/config
mkdir -p pkg/logger
mkdir -p pkg/server

# Create empty files in each directory
for dir in cmd/myapp internal/models internal/api internal/database internal/services pkg/config pkg/logger pkg/server; do
  touch "${dir}/.empty"
done