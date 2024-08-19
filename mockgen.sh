# Repository
mockery --disable-version-string --name=DBRepository --dir ./internal/repository --output ./mocks/repository
mockery --disable-version-string --name=RedisRepository --dir ./internal/repository --output ./mocks/repository

# Service
mockery --disable-version-string --name=Service --dir ./internal/service --output ./mocks/service