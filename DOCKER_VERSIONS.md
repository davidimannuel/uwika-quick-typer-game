# Docker Image Versions

Dokumentasi version untuk semua Docker images yang digunakan dalam project ini.

## Production Images

| Service | Image | Version | Notes |
|---------|-------|---------|-------|
| PostgreSQL | `postgres` | `16-alpine` | Stable PostgreSQL 16 dengan Alpine Linux |
| Migrate | `migrate/migrate` | `v4.17.0` | Database migration tool |
| Golang (Build) | `golang` | `1.23-alpine` | Go compiler untuk build stage |
| Alpine (Runtime) | `alpine` | `3.19` | Minimal runtime image |

## Version Details

### PostgreSQL 16-alpine
- **Full Image**: `postgres:16-alpine`
- **Base**: Alpine Linux
- **Size**: ~80MB
- **Release**: PostgreSQL 16.x latest patch
- **Why**: Stable, small size, production-ready

### Migrate v4.17.0
- **Full Image**: `migrate/migrate:v4.17.0`
- **Release Date**: Latest stable
- **Why**: 
  - Supports PostgreSQL
  - CLI-compatible
  - Well-maintained
  - [Release Notes](https://github.com/golang-migrate/migrate/releases/tag/v4.17.0)

### Golang 1.23-alpine
- **Full Image**: `golang:1.23-alpine`
- **Base**: Alpine Linux
- **Size**: ~180MB
- **Why**: Minimal build environment, matches go.mod version

### Alpine 3.19
- **Full Image**: `alpine:3.19`
- **Size**: ~7MB
- **Release**: Latest stable Alpine 3.19.x
- **Why**: 
  - Minimal attack surface
  - Small image size
  - Security updates
  - Production-ready

## Update Strategy

### When to Update

**PostgreSQL:**
- Update to latest patch version monthly
- Update major version after testing: `16-alpine` → `17-alpine`

**Migrate:**
- Update when new features needed
- Check compatibility: [Migration Guide](https://github.com/golang-migrate/migrate)

**Golang:**
- Keep in sync with `go.mod`
- Update to latest patch: `1.23.1-alpine`, `1.23.2-alpine`

**Alpine:**
- Update to latest stable: `3.19` → `3.20`
- Test thoroughly before production

### How to Update

1. **Test Locally:**
   ```bash
   # Update version in docker-compose.yml or Dockerfile
   # Test build
   docker-compose build
   
   # Test run
   docker-compose up
   ```

2. **Check Compatibility:**
   - Read release notes
   - Check breaking changes
   - Test migrations

3. **Update Documentation:**
   - Update this file
   - Update README if needed

## Version Pinning Best Practices

✅ **DO:**
- Use specific version tags (e.g., `v4.17.0`, `3.19`)
- Pin major + minor versions for stability
- Update patch versions regularly

❌ **DON'T:**
- Use `latest` tag in production
- Use floating tags (e.g., `alpine` without version)
- Update without testing

## Security

### Image Scanning

```bash
# Scan for vulnerabilities
docker scan postgres:16-alpine
docker scan alpine:3.19

# Or use trivy
trivy image postgres:16-alpine
trivy image alpine:3.19
```

### CVE Monitoring

Monitor for security updates:
- PostgreSQL: https://www.postgresql.org/support/security/
- Alpine: https://security.alpinelinux.org/
- Golang: https://go.dev/security/

## Rollback Plan

If update causes issues:

```bash
# 1. Revert version in docker-compose.yml
# 2. Rebuild
docker-compose down
docker-compose build --no-cache
docker-compose up -d

# 3. If database issues, restore from backup
```

## Current Versions Summary

```yaml
services:
  postgres:
    image: postgres:16-alpine
  
  migrate:
    image: migrate/migrate:v4.17.0
  
  api:
    build:
      context: .
      dockerfile: Dockerfile.api
    # Uses:
    # - golang:1.23-alpine (build)
    # - alpine:3.19 (runtime)
  
  admin-web:
    build:
      context: .
      dockerfile: Dockerfile.admin
    # Uses:
    # - golang:1.23-alpine (build)
    # - alpine:3.19 (runtime)
```

## Changelog

| Date | Image | Old Version | New Version | Reason |
|------|-------|-------------|-------------|--------|
| 2024-11-29 | All | `latest` | Specific versions | Production readiness |

---

**Last Updated**: 2024-11-29  
**Reviewed By**: Development Team  
**Next Review**: 2025-01-01

