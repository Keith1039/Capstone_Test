# DBVG (Work in progress)

Database validator and generator for Postgres. Use as a CLI or import as a library

__[CLI Documentation](cmd/README.md)__

__[Go Documentation]()__

## Main Offering
This project provides tools to help a fledgling database mature quickly. This allows the developer to side step
hours of planning. This tool is intended for use in a new PERSONAL project or for helping in create a proof of concept. 
Please don't use this in prod...

### Validation
The validation provided by DBVG is cycle aversion as well as cycle resolution. As databases grow, it becomes incredibly easy to 
accidentally create cyclic relationships between tables. This can be averted with proper planning but in cases
where time is limited, this cannot be done.

### Data Generation
As a database grows it also becomes harder to generate data for it as you manage the complex web of relationships. 
While real data is better for testing, in a product that has yet to be released this isn't feasible.
Oftentimes, the solution is to create a script that executes the set of queries needed to generate 
database entries. This solution however, is inflexible and requires constant maintenance.

This problem is further compounded when you're actively making changes to the database schema. 
This causes even more time loss. With this library, you can generate test data on the fly with no worries. 
Let the code handle the hard work.

## Basic Usage (As a library)

### Verify if your database schema has cyclical relationships

``` go
package main

import (
	"database/sql"
	"fmt"
	"github.com/Keith1039/dbvg/graph"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // open the database connection
	if err != nil {
		log.Fatal(err)
	}

	ord := graph.NewOrdering(db) // get a new ordering struct
	cycles := ord.GetCycles()    // get a linked list of cycles
	// loop through the list
	node := cycles.Front()
	for node != nil {
		fmt.Println(node.Value.(string)) // print out the cycles
		node = node.Next()
	}

	defer db.Close()
}
```
Sample output:
```
b --> d --> e --> b
a --> b --> c --> a
```

### Remove All cyclical relationships
``` go
package main

import (
	"database/sql"
	database "github.com/Keith1039/dbvg/db"
	"github.com/Keith1039/dbvg/graph"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // open the database connection
	if err != nil {
		log.Fatal(err)
	}

	ord := graph.NewOrdering(db)                      // get a new ordering struct
	suggestions := ord.GetSuggestionQueries()         // get a linked list of the suggestion queries
	err = database.RunQueriesVerbose(db, suggestions) // runs the suggestion queries and prints them
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
```

sample output:
```
Query 1: ALTER TABLE e DROP COLUMN bref;
Query 2: ALTER TABLE e DROP COLUMN bref2;
Query 3: CREATE TABLE IF NOT EXISTS b_e(
         b_bkey_ref INT4,
         b_bkey2_ref INT4,
         e_ekey_ref INT4,
        FOREIGN KEY (b_bkey_ref, b_bkey2_ref) REFERENCES b(bkey, bkey2),
        FOREIGN KEY (e_ekey_ref) REFERENCES e(ekey),
        PRIMARY KEY (b_bkey_ref, b_bkey2_ref, e_ekey_ref)
)
Query 4: ALTER TABLE a DROP COLUMN bref;
Query 5: ALTER TABLE a DROP COLUMN bref2;
Query 6: CREATE TABLE IF NOT EXISTS b_a(
         b_bkey_ref INT4,
         b_bkey2_ref INT4,
         a_akey_ref INT4,
        FOREIGN KEY (b_bkey_ref, b_bkey2_ref) REFERENCES b(bkey, bkey2),
        FOREIGN KEY (a_akey_ref) REFERENCES a(akey),
        PRIMARY KEY (b_bkey_ref, b_bkey2_ref, a_akey_ref)
)
```

### Generate X amounts of entries for a table
```go
package main

import (
	"database/sql"
	database "github.com/Keith1039/dbvg/db"
	"github.com/Keith1039/dbvg/parameters"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")) // open the database connection
	if err != nil {
		log.Fatal(err)
	}

	writer, err := parameters.NewQueryWriterFor(db, "b")  // create a new query writer for table "b"
	if err != nil {
		log.Fatal(err)
	}
	writer.GenerateEntries(1)  // functional equivalent of writer.GenerateEntry() 

	err = database.RunQueriesVerbose(db, writer.InsertQueryQueue) // run the insert queries
	if err != nil {
		log.Fatal(err)
	}
	err = database.RunQueries(db, writer.DeleteQueryQueue) // run the deletion queries for cleanup (optional)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
```
sample output:
```
Query 1: INSERT INTO e (ekey) VALUES (0);
Query 2: INSERT INTO d (dkey, eref) VALUES (0, 0);
Query 3: INSERT INTO a (akey) VALUES (0);
Query 4: INSERT INTO c (aref, ckey) VALUES (0, 0);
Query 5: INSERT INTO b (bkey, bkey2, cref, dref) VALUES (0, 0, 0, 0);
Query 1: DELETE FROM b WHERE bkey=0 AND bkey2=0 AND cref=0 AND dref=0;
Query 2: DELETE FROM c WHERE aref=0 AND ckey=0;
Query 3: DELETE FROM a WHERE akey=0;
Query 4: DELETE FROM d WHERE dkey=0 AND eref=0;
Query 5: DELETE FROM e WHERE ekey=0;
```
*Note*: The `QueryWriter` struct cannot be used if a cycle exists in the path for the given table.
It is recommended to always resolve cycles before generating data. below is the result of using the above
code on a schema that has cycles.
```
2025/02/18 15:27:55 error, the following cycles have been detected in the database schema: b --> d --> e --> b | b --> c --> a --> b
exit status 1
```
