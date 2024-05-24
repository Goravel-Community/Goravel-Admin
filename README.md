
# An admin panel for Goravel inspired by Backpack for Laravel

# Installation

```bash
# publish configs
go run . artisan vendor:publish --package=github.com/goravel-community/goravel-admin --tag=goravel-admin-config
# publish migrations
go run . artisan vendor:publish --package=github.com/goravel-community/goravel-admin --tag=goravel-admin-migrations
# migrate
go run . artisan migrate
```
