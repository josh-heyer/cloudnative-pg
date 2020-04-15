/*
This file is part of Cloud Native PostgreSQL.

Copyright (C) 2019-2020 2ndQuadrant Italia SRL. Exclusively licensed to 2ndQuadrant Limited.
*/

package postgres

// PromoteAndWait promotes this instance, and wait 60 seconds for it to happen
func (instance *Instance) PromoteAndWait() error {
	superUserDB, err := instance.GetSuperUserDB()
	if err != nil {
		return err
	}

	_, err = superUserDB.Exec("SELECT pg_promote(true)")
	return err
}
