# FileSystem Simulator in Go

This project is a **Go-based filesystem simulator** that explores various real-world file handling scenarios.  
The aim is to simulate how systems behave when dealing with:

1. **Millions of small files**  
2. **Large files (hundreds of MBs to multiple GBs)**  
3. **Failure cases** – missing files, permission errors, partially deleted files, etc.

It’s designed as a learning and experimentation playground for:
- **File I/O in Go**
- **Memory management**
- **Streaming large data**
- **Hashing and validation**
- **Error handling at scale**

---

## 🚀 Project Goals

We want to **simulate and test different file system behaviors** in controlled experiments:

### Phase 1 – Millions of Small Files
- Create millions of small files.
- Iterate through them efficiently without running out of memory.
- Compute hashes for verification.

### Phase 2 – Large Files
- Generate large files (1GB+).
- Use **streaming techniques** to read & hash files without loading them fully into memory.
- Monitor Go’s **memory usage & GC behavior**.

### Phase 3 – Failure Scenarios
- Attempt to read deleted files.
- Handle permission-denied errors.
- Simulate partial file corruption.
- Retry and error logging strategies.

---
