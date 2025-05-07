# Codewithwest Go Documentation

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Documentation Structure](#documentation-structure)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

    This project serves and codewithwest projects powerhouse housing all
    information created by the engineer and providing data to FE web
    application.

## Features

- [Client Documentation](./docs/client.md)
- [Admin Documentation](./docs/admin-user.md)
- [Project Categories Documentation](./docs/project-categories.md)
- [Projects Documentation](./docs/projects.md) In-Progress

## Prerequisites

    - Go 1.22.2
    - github.com/graphql-go/handler v0.2.4
    - github.com/joho/godotenv v1.5.1
    - github.com/rs/zerolog v1.33.0
    - go.mongodb.org/mongo-driver v1.17.2
    - golang.org/x/crypto v0.33.0
    - github.com/golang/snappy v0.0.4
    - github.com/klauspost/compress v1.16.7
    - github.com/montanaflynn/stats v0.7.1
    - github.com/xdg-go/pbkdf2 v1.0.0
    - github.com/xdg-go/scram v1.1.2
    - github.com/xdg-go/stringprep v1.0.4
    - github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78
    - golang.org/x/sync v0.11.0
    - golang.org/x/text v0.22.0
    - github.com/gorilla/mux v1.8.1
    - github.com/graphql-go/graphql v0.8.1
    - github.com/mattn/go-colorable v0.1.13
    - github.com/mattn/go-isatty v0.0.19
    - golang.org/x/sys v0.30.0

## Installation

    `git clone https://github.com/codewithwest/project_codewithwest-go  codewithwest-go`

    install docker compose if not installed

    cd codewithwest-go

    docker compose up -d

    Access resource on https://localhost:3071/graphql

## Project Structure

## Configuration

## Documentation Structure

    name

    Description:
    
    Input Parameters:
    
    Process:
    
    Sample:
    
    Possible Errors:
    
    Returns:
    
    Security Considerations:
     
    Technical Details:
    
    Performance Notes:


## Testing

    docker exec -it codewithwest-go bash
    go test ./tests
