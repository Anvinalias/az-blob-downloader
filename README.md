# az-blob-downloader
A command-line tool, developed with Go programming language, designed to streamline the process of downloading and organizing offline upgrade files from Azure Blob Storage. 
It automates the retrieval of application upgrade steps and generates corresponding metadata files (`uploadedversion.txt`) for each step.

## **Project Status: In Progress**

### ✅ Completed
1. Decrypt connection details and connect to Azure storage account
2. List containers based on request pattern
3. Parse requests from a text file
4. Find the shortest path for requests using a graph and BFS algorithm
5. Download blobs into separate directories based on steps
6. Create `uploadversion.txt` based on the downloaded files
7. Create `instructions.txt` on how to setup and run project
8. Create a log file

### 🛠️ Pending
1. Update the Blob Storage connection to use a SAS token instead of a connection string, and remove the decrypt package from the codebase.
2. Optional download path in config
3. Refactor entire codebase for error handling, comments, logging, best practices, etc.
4. Update Readme.md

### 🌟 Optional Enhancements
1. Support another download scenario: download blobs into separate directories based on application, controlled by a flag  
2. Introduce concurrency for downloading blobs using goroutines

## Features

- 🔐 Secure Connection via Config Decryption  
  Reads encrypted connection configuration and decrypts credentials at runtime (pending migration to SAS token authentication).

- 📦 Selective Blob Filtering  
  Filters blobs from Azure Blob Storage based on application name and version range defined in a plain-text request file.

- 🔀 Shortest Path Version Resolution  
  Uses a graph traversal (BFS) algorithm to find the shortest chain of application updates between requested versions.

- 📁 Organized Output  
  Downloads versioned files into structured folders (set1, set2, ...) for each valid upgrade path.

- 📄 Auto-Generated Metadata  
  - uploadversion.txt: A required metadata file that records the application and version details of each step in a JSON format

- 📝 Activity Logging  
  Logs key events and errors to help monitor behavior during execution.

## 🧾 Installation & Setup

Follow these steps to set up the az-blob-downloader project on your machine.

### ✅ Prerequisites

- Go 1.24 (or higher) installed on your system  
- Azure Storage account with the required blob containers and access permissions  
- Git (for cloning the repository)

### 1. Clone the repository

```
git clone https://github.com/Anvinalias/az-blob-downloader.git
cd az-blob-downloader
```

### 2. Install Go module dependencies

```
go mod tidy
```

### 3. Build the project

```
go build -o offline-files.exe ./cmd
```

### 4. Prepare the configuration

- Copy the example config file:
  ```
  cp config.example.yaml config.yaml
  ```
- Open `config.yaml` in a text editor.
- Update only the following key:
  - `downloadPath`: Set this to the folder where you want offline files to be downloaded.
  - **Note:** Use double backslashes (`\\`) or a single forward slash (`/`) in your path.
    - Examples:
      ```
      downloadPath: "D:\\OfflineFiles\\Downloads"
      ```
      or
      ```
      downloadPath: "D:/OfflineFiles/Downloads"
      ```
- Save the file.

### 5. Prepare the request file

- Copy the example request file:
  ```
  cp request.example.txt request.txt
  ```
- Open `request.txt` in a text editor.
- Add your offline file requests in the following format, one per line:
  ```
  applicationName-fromVersion-toVersion
  ```
  - Example:
    ```
    exampleapplication-1.0.0.0-4.0.0.10
    ```
- Save the file.

### 6. Run the application

```
./offline-files.exe
```

- You can run this from the command prompt or by double-clicking the executable.

---

## After Running

- Wait for the task to complete.
- Check the generated log file for status updates or any errors.
- Your offline files will be downloaded to the folder specified in downloadPath.
- The application will also generate uploadversion.txt and a log file for each run.


