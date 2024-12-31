package parameters

import (
	"container/list"
	"fmt"
	"github.com/Keith1039/Capstone_Test/db"
	"log"
	"sort"
)

type QueryWriter struct {
	AllRelations    map[string]map[string]map[string]string
	LevelMap        map[string]int
	pkMap           map[string]int
	TableOrderQueue *list.List // queue
	QueryQueue      *list.List // queue
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

func (qw *QueryWriter) CreateTableOrder() {
	l := list.New()
	tnames := make([]string, len(qw.LevelMap))
	for key, _ := range qw.LevelMap {
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
	tableName := qw.TableOrderQueue.Front().Value.(string)
	t := createTable(tableName)
	for _, col := range t.Columns {
		colVal, err := col.Parser.ParseColumn()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(colVal)
	}
	qw.TableOrderQueue.Remove(qw.QueryQueue.Front()) // remove the first in the queue
}
