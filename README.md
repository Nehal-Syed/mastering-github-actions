# User CRUD Application with Advanced CI/CD

![CI Pipeline](https://github.com/yourusername/go-user-crud/actions/workflows/ci.yml/badge.svg)
![CD Pipeline](https://github.com/yourusername/go-user-crud/actions/workflows/cd.yml/badge.svg)
![Security Scan](https://github.com/yourusername/go-user-crud/actions/workflows/security.yml/badge.svg)
![Code Coverage](https://codecov.io/gh/yourusername/go-user-crud/branch/main/graph/badge.svg)

## 🚀 DevOps Features Implemented

### GitHub Actions Advanced Concepts
- ✅ Multi-stage Docker builds with caching
- ✅ Matrix testing across Go versions
- ✅ Integration tests with MySQL service containers
- ✅ Parallel job execution with dependencies
- ✅ Security scanning (Trivy, GoSec, Dependabot)
- ✅ Performance testing (k6 load tests)
- ✅ Automated Docker Hub pushes
- ✅ Staging & Production environments
- ✅ Slack notifications
- ✅ Scheduled security scans

### Application Features
- ✅ RESTful API with Go
- ✅ MySQL database integration
- ✅ Frontend HTML/CSS/JS interface
- ✅ Full CRUD operations
- ✅ Rate limiting & security headers
- ✅ Docker & Docker Compose
- ✅ Health checks & graceful shutdown

## 🛠️ Tech Stack
- **Backend**: Go 1.21, Gorilla Mux
- **Database**: MySQL 8.0
- **Frontend**: HTML5, CSS3, Vanilla JS
- **Container**: Docker, Docker Compose
- **CI/CD**: GitHub Actions
- **Security**: Trivy, GoSec, gitleaks
- **Testing**: Go tests, k6, Lighthouse

## 📦 Local Development

```bash
# Clone repository
git clone https://github.com/yourusername/go-user-crud.git

# Start with Docker Compose
cd backend && docker-compose up -d

# Or run locally
make build
make run