rows, err = db.Query("SELECT fname, lname, email FROM users WHERE UPPER(users.fname) LIKE UPPER ($1)", searchString)
