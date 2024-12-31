package parameters

import (
	"container/list"
	"fmt"
	"github.com/Keith1039/Capstone_Test/db"
	"log"
	"sort"
)

type QueryWriter struct {
	AllRelations     map[string]map[string]map[string]string
	LevelMap         map[string]int
	pkMap            map[string]int
	TableOrderQueue  *list.List // queue
	InsertQueryQueue *list.List // queue
	DeleteQueryQueue *list.List // queue
}

func createTable(tableName string) table {
	t := table{TableName: tableName}
	columnMap := db.GetColumnMap(tableName)
	columns := make([]column, 0, len(columnMap))
	i := 0
	for columnName, dataType := range columnMap {
		c := column{ColumnName: columnName, Type: dataType, Parser: getColumnParser(dataType)}
		columns[i] = c
		i++
	}
	t.Columns = columns
	return t
}

func appendValues(colStringPtr *string, valStringPtr *string, newColumn string, newVal string) {
	if *colStringPtr == "(" {
		*colStringPtr = *colStringPtr + newColumn
	} else {
		*colStringPtr = *colStringPtr + "," + newColumn
	}

	if *valStringPtr == "(" {
		*valStringPtr = *valStringPtr + newVal
	} else {
		*valStringPtr = *valStringPtr + "," + newVal
	}

}

func (qw *QueryWriter) CreateTableOrder() {
	l := list.New()
	tnames := make([]string, len(qw.LevelMap))
	for key := range qw.LevelMap {
		tnames = append(tnames, key)
	}
	// sort in descending order
	sort.SliceStable(tnames, func(i, j int) bool {
		return qw.LevelMap[tnames[i]] > qw.LevelMap[tnames[j]]
	})
	for _, tname := range tnames {
		l.PushBack(tname) // push to the back of the queue
	}
	qw.TableOrderQueue = l // set the queue
}

func (qw *QueryWriter) ProcessTables() {
	for qw.TableOrderQueue.Len() > 0 {
		qw.ProcessTable()
	}
}

func (qw *QueryWriter) ProcessTable() {
	//var writer SQLWriter
	colString := "("
	colValString := "("
	tableName := qw.TableOrderQueue.Front().Value.(string)
	t := createTable(tableName)
	for _, col := range t.Columns {
		colVal, err := col.Parser.ParseColumn()
		if err != nil {
			log.Fatal(err)
		}
		appendValues(&colString, &colValString, col.ColumnName, colVal)
	}
	colString = colString + ")"
	colValString = colValString + ")"
	query := fmt.Sprintf("INSERT INTO %s %s VALUES %s", t.TableName, colString, colValString)
	qw.InsertQueryQueue.PushBack(query)
	qw.TableOrderQueue.Remove(qw.TableOrderQueue.Front()) // remove the first in the queue
}
