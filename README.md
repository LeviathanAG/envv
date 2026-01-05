# envv

(First GO project after learning GO for 3 weeks after endsems.)

In my internship everyone begs for env and reply comes late sometimes from people which is annoying asl so I figured i might as well make an open source CLI tool just for the sake of it.

idc if people use it or not tbh.

i named it envv since its similar to git

it scans ur repo for .env files and stores each file in a mongo document and each repo is a separate cluster (yet to implement) rn its one cluster.

each doc is hashed and every update we check this hash for changes and if there are any we replace the doc with local file.
(yet to implement)

what is implemented is :

envv init which makes a repo hash and stores it in mongo and makes a doc of each env with its own hash and stores it in mongo. 

also have git import command which takes a folder name and imports it to local

## Prerequisites

- **Go**: v1.25.5 or higher
- **MongoDB**: A running MongoDB instance (Atlas or local)

## Installation

### Global Install (Recommended)

Install the binary into your `GOBIN` so you can use it anywhere:

```bash
go install ./src/cmd/envv
```

*Ensure your `PATH` includes your Go bin directory (e.g., `export PATH=$PATH:$(go env GOPATH)/bin`).*

### Building Manually

To build the binary in the current directory:

```bash
go build -o envv ./src/cmd/envv
```

## Configuration

The tool requires a `MONGO_URI` to connect to your database. You can provide this in two ways:

1. **Environment Variable**:

    ```bash
    export MONGO_URI="mongodb://your-mongodb-uri"
    ```

2. **Local .env File**: Create a `.env` file in the project root:

    ```env
    MONGO_URI=mongodb://your-mongodb-uri
    ```

## Usage

### 1. Initialize a Repository

Scan the current repository and upload all discovered `.env` files to MongoDB:

```bash
envv init
```

### 2. Import / Restore a .env File

Restore a specific `.env` file from MongoDB back to your local filesystem:

```bash
envv import <folder_name>
```

* Use `root` or `.` for the repository root.
- The tool handles creating missing parent directories automatically.

## Project Structure

- `src/cmd/envv`: Main application entry point and command implementations (Cobra).
- `src/environ`: Core logic for scanning, hashing, and writing `.env` files.
- `src/model`: Data structures for database storage.
- `src/store`: MongoDB interaction layer.
- `src/config`: Configuration loading and management.
