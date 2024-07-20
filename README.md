
# An admin panel for Goravel inspired by Backpack for Laravel

# Installation

```bash
# install goravel-admin
go get -u github.com/goravel-community/goravel-admin
# publish configs
go run . artisan vendor:publish --package=github.com/goravel-community/goravel-admin --tag=goravel-admin-config
# publish migrations
go run . artisan vendor:publish --package=github.com/goravel-community/goravel-admin --tag=goravel-admin-migrations
# migrate
go run . artisan migrate
```

### Checklist

✅ make login via session

✅ install and configure templ

✅ frontend for login

⬜ logged in template

⬜ change email and name page

