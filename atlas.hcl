variable "ATLAS_DB_DEV" {
  type = string
  default = "postgres://postgres:pwd123!@127.0.0.1:5432/x-user?sslmode=disable&search_path=dev"
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./pkg/database/models",
    "--dialect", "postgres", // | postgres | sqlite
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = var.ATLAS_DB_DEV
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}