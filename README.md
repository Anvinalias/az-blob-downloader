# az-blob-downloader
A Go tool to download files from Azure Blob Storage containers for specific internal offline file retrieval scenarios.

## **Project Status: In Progress**

### **--Completed--**
1. Encrypt and store connection details
2. Decrypt connection details and connect to Azure storage account
3. List container names
4. List blobs inside a specific container
5. Download sample blobs
6. List containers based on request pattern
7. Parse requests from a text file
8. Find the shortest path for requests using a graph and BFS algorithm
9. Sort and arrange graph nodes in descending order
10. Download blobs based on the shortest path into a single directory
11. Download blobs into separate directories based on steps
12. Create `uploadversion.txt` based on the downloaded files
13. Validate `request.txt` for structure, spelling, etc.
14. Create `instructions.txt` on how to setup and run project
15. Create a log file

### **--Pending--**
16. Refactor the entire codebase for error handling, comments, logging, best practices, etc.
17. Update Readme.md

### **--Optional Enhancements--**
18. Support another download scenario: download blobs into separate directories based on application, controlled by a flag  
19. Introduce concurrency for downloading blobs using goroutines


