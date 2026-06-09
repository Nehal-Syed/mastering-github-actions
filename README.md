# 🚀 Mastering GitHub Actions with Go

[![CI](https://github.com/Nehal-Syed/mastering-github-actions/actions/workflows/ci.yml/badge.svg)](https://github.com/Nehal-Syed/mastering-github-actions/actions/workflows/ci.yml)
[![CD](https://github.com/Nehal-Syed/mastering-github-actions/actions/workflows/cd.yml/badge.svg)](https://github.com/Nehal-Syed/mastering-github-actions/actions/workflows/cd.yml)
[![Security](https://github.com/Nehal-Syed/mastering-github-actions/actions/workflows/security.yml/badge.svg)](https://github.com/Nehal-Syed/mastering-github-actions/actions/workflows/security.yml)

A hands-on project designed to master **GitHub Actions**, **CI/CD pipelines**, **automated testing**, and **security scanning** using a production-style **Go web application**.

The primary goal of this repository is to demonstrate modern DevOps practices and GitHub Actions workflows while building and testing a real-world Go application.

---

## 🎯 Project Objectives

This project showcases how to:

* Build scalable CI/CD pipelines with GitHub Actions
* Automate testing and code quality checks
* Run integration tests using service containers
* Integrate security scanning into development workflows
* Deploy applications through environment-based pipelines
* Implement branch protection and deployment gates
* Structure a Go application using clean architecture principles

---

# ⚙️ GitHub Actions Workflows

## 1️⃣ Continuous Integration Pipeline

**Workflow:** `.github/workflows/ci.yml`

### Trigger Events

* Push
* Pull Request
* Manual Dispatch

### Jobs

| Job               | Purpose                                      |
| ----------------- | -------------------------------------------- |
| code-quality      | Runs `gofmt` and `go vet`                    |
| unit-tests        | Executes Go unit tests with coverage         |
| integration-tests | Runs tests against a MySQL service container |
| build-test        | Verifies application compiles successfully   |
| security-scan     | Performs GoSec vulnerability scanning        |
| final-status      | Pipeline summary and status reporting        |

### CI Flow

```text
Push / PR
    │
    ▼
┌─────────────┐
│ Code Quality│
└──────┬──────┘
       ▼
┌─────────────┐
│ Unit Tests  │
└──────┬──────┘
       ▼
┌─────────────────┐
│ Integration Test│
│    (MySQL)      │
└──────┬──────────┘
       ▼
┌─────────────┐
│ Build Check │
└──────┬──────┘
       ▼
┌─────────────┐
│GoSec Scan   │
└──────┬──────┘
       ▼
┌─────────────┐
│Final Status │
└─────────────┘
```

---

## 2️⃣ Continuous Deployment Pipeline

**Workflow:** `.github/workflows/cd.yml`

### Trigger Events

* Git Tags (`v*`)
* Manual Dispatch

### Jobs

| Job               | Purpose                                      |
| ----------------- | -------------------------------------------- |
| deploy-staging    | Deploys application to staging               |
| deploy-production | Deploys to production after staging succeeds |

### Deployment Flow

```text
Tag Release (v1.0.0)
          │
          ▼
 ┌─────────────────┐
 │ Deploy Staging  │
 └────────┬────────┘
          ▼
 ┌─────────────────┐
 │Deploy Production│
 └─────────────────┘
```

---

## 3️⃣ Security Pipeline

**Workflow:** `.github/workflows/security.yml`

### Trigger Events

* Daily Schedule
* Push to Main Branch
* Manual Dispatch

### Jobs

| Job           | Purpose                          |
| ------------- | -------------------------------- |
| security-scan | Automated vulnerability scanning |

### Features

* Daily automated security checks
* Continuous vulnerability monitoring
* Security reports as workflow artifacts
* Shift-left security practices

---

# 🧠 GitHub Actions Concepts Demonstrated

| Concept                 | Implementation                |
| ----------------------- | ----------------------------- |
| Multi-job workflows     | CI pipeline                   |
| Job dependencies        | Production depends on staging |
| Service containers      | MySQL integration testing     |
| Scheduled workflows     | Daily security scans          |
| Workflow dispatch       | Manual execution              |
| Environment deployments | Staging & Production          |
| Security automation     | GoSec scanning                |
| Artifact generation     | Security reports              |
| Status checks           | Branch protection readiness   |
| Tag-based releases      | Production deployment         |

---

# 🛠️ Technology Stack

| Component | Technology                      |
| --------- | ------------------------------- |
| Language  | Go 1.25.1                       |
| Router    | Gorilla Mux                     |
| Database  | MySQL 8.0                       |
| Frontend  | HTML5, CSS3, Vanilla JavaScript |
| CI/CD     | GitHub Actions                  |
| Security  | GoSec                           |
| Testing   | Go Test Framework               |

---

# 📁 Project Structure

```text
mastering-github-actions/
│
├── .github/
│   └── workflows/
│       ├── ci.yml
│       ├── cd.yml
│       └── security.yml
│
├── backend/
│   ├── cmd/
│   │   └── main.go
│   │
│   ├── internal/
│   │   ├── config/
│   │   ├── database/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── models/
│   │   ├── router/
│   │   └── services/
│   │
│   └── go.mod
│
└── frontend/
```

---

# 🚀 Running Locally

## Clone Repository

```bash
git clone https://github.com/Nehal-Syed/mastering-github-actions.git

cd mastering-github-actions
```

## Run Backend

```bash
cd backend

go mod tidy

go run cmd/main.go
```

## Test Endpoints

### Health Check

```bash
curl http://localhost:8080/health
```

### Users Endpoint

```bash
curl http://localhost:8080/api/users
```

---

# 🧪 Running Tests

### Unit Tests

```bash
go test ./...
```

### Test Coverage

```bash
go test ./... -cover
```

### Verbose Output

```bash
go test ./... -v
```

---

# 🔒 Security Scanning

Run GoSec locally:

```bash
gosec ./...
```

The same vulnerability checks are automatically executed in:

* CI Pipeline
* Security Pipeline
* Scheduled Daily Scans

---

# 📊 Pipeline Dashboard

After each workflow execution, GitHub Actions provides:

* ✅ Real-time job execution logs
* 📈 Test coverage reports
* 🔒 Security scan results
* 🚀 Deployment status tracking
* 📋 Workflow run history
* 🛠️ Build and test artifacts

Navigate to the **Actions** tab in GitHub to monitor pipeline activity.

---

# 🎓 Key Learning Outcomes

Through this project, I gained practical experience with:

### CI/CD Pipeline Design

* Workflow orchestration
* Multi-stage pipelines
* Job dependencies

### Automated Testing

* Unit testing
* Integration testing
* Database service containers

### Security Automation

* Vulnerability scanning
* Continuous security monitoring
* Security-first development practices

### Deployment Strategies

* Staging environments
* Production environments
* Release management with tags

### Go Application Architecture

* Clean project structure
* Modular design
* Separation of concerns

---
