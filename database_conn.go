// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

// // OpenWithOptions -  Options must be converted into pg.Options{}, if not - use default options
// func openWithOptions(user, database, password string, data []byte) (IDatabase, error) {
// 	opts := pg.Options{}
// 	err := json.Unmarshal(data, &opts)
// 	if err != nil {
// 		//connect with default options
// 		pgDB := pg.Connect(defaultOptions(user, database, password))
// 		return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil
// 	}

// 	pgDB := pg.Connect(&opts)
// 	return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil
// }

// // openWithDefaultOpts -  Options must be converted into pg.Options{}, if not - use default options
// func openWithDefaultOpts(user, database, password string) (IDatabase, error) {
// 	pgDB := pg.Connect(defaultOptions(user, database, password))
// 	return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil
// }

// // Close closes the database client
// func (d *Database) Close() {
// 	err := d.DB.Close()
// 	if err != nil {
// 		d.Error("Database.Close", "Can not Close DB", err)
// 		return
// 	}

// 	d.Info("Database.Close", "Closed")
// }

// // defaultOptions sets the default options.
// // Note: this is called (first) even if you open a connection with parameters.
// func defaultOptions(user, database, password string) *pg.Options {
// 	opts := &pg.Options{
// 		User:     user,
// 		Database: database,
// 		Password: password,
// 		TLSConfig: &tls.Config{
// 			InsecureSkipVerify: true,
// 		},

// 		MaxRetries:      1,
// 		MinRetryBackoff: -1,

// 		DialTimeout:  30 * time.Second,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 10 * time.Second,

// 		PoolSize:           10,
// 		MaxConnAge:         10 * time.Second,
// 		PoolTimeout:        30 * time.Second,
// 		IdleTimeout:        10 * time.Second,
// 		IdleCheckFrequency: 100 * time.Millisecond,
// 	}

// 	//LoadCertificate(opts)

// 	return opts
// }
