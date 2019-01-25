package resolvers

import (
    "api/db"
)

type Resolvers struct {
  JobDefinition JobDefinition
}

func New(db *db.Connection) Resolvers {
  return Resolvers{
    JobDefinition: JobDefinition{db: db},
  }
}
