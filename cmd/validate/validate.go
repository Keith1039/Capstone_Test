/*
Copyright © 2025 Keith Compere <KeithCompere150@gmail.com>
*/
package validate

import (
	"database/sql"
	"github.com/spf13/cobra"
	"log"
)

var (
	ConnString string
)

// validateCmd represents the validate command
var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "The palette responsible for schema validation",
	Long: `This pallete is responsible for ensuring that the database schema has no
		cycles. If a cycle is detected, a series of recommendation queries will be generated`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubCommands() {
	ValidateCmd.AddCommand(schemaCmd)
}

func init() {
	addSubCommands()
	ValidateCmd.PersistentFlags().StringVarP(&ConnString, "database", "", "", "url to connect to the database with")

	if err := ValidateCmd.MarkPersistentFlagRequired("database"); err != nil {
		log.Fatal(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", ConnString)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
