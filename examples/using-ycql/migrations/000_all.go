// This is auto-generated file using 'gofr migrate' tool. DO NOT EDIT.
package migrations

import dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"

func All() map[string]dbmigration.Migrator {
	return map[string]dbmigration.Migrator{

		"20230116104833": K20230116104833{},
	}
}
