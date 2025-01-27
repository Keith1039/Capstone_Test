/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package generate

import (
	"container/list"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Keith1039/Capstone_Test/graph"
	"github.com/jimsmart/schema"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	dirPath string
	table   string
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "generates a template in a specific folder for a specific group of tables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := InitDB()
		if err != nil {
			log.Fatal(err)
		}
		ordering := graph.NewOrdering(db)

		tableOrder, err := ordering.GetOrder(table)
		if err != nil {
			log.Fatal(err)
		}
		templates := makeTemplates(db, tableOrder)
		jsonString, err := json.MarshalIndent(templates, "", "  ")
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err = os.Mkdir(dirPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = os.WriteFile(fmt.Sprintf("%s_path_template.json", table), jsonString, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {

	templateCmd.Flags().StringVarP(&dirPath, "dir", "", "", "relative path of a directory to place the template file in, if the path doesn't exist it will make the folder")
	templateCmd.Flags().StringVarP(&table, "table", "", "", "the name of the table we want an entry for")

	err := templateCmd.MarkFlagRequired("dir")
	if err != nil {
		log.Fatal(err)
	}
	err = templateCmd.MarkFlagRequired("table")
	if err != nil {
		log.Fatal(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func makeTemplates(db *sql.DB, l *list.List) map[string]map[string]map[string]string {
	m := make(map[string]map[string]map[string]string)

	node := l.Front()
	for node != nil {
		tableName := node.Value.(string)
		m[table] = makeTemplate(db, tableName)
		node = node.Next()
	}
	return m
}

func makeTemplate(db *sql.DB, tableName string) map[string]map[string]string {
	m := make(map[string]map[string]string)
	cols, err := schema.ColumnTypes(db, "", tableName)
	if err != nil {
		log.Fatal(err)
	}
	for _, col := range cols {
		m[col.Name()] = map[string]string{"Code": "", "Value": ""}
	}
	return m
}
