package cmd

import (
	"cmdLine/internal/sql2struct"
	"log"

	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2StructCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}
		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2StructCmd)

	sql2StructCmd.Flags().StringVarP(&username, "username", "", "", "请出入数据库账户")
	sql2StructCmd.Flags().StringVarP(&password, "password", "", "", "请出入数据库账户密码")
	sql2StructCmd.Flags().StringVarP(&host, "host", "", "", "请出入主机")
	sql2StructCmd.Flags().StringVarP(&charset, "charset", "", "", "请出入数据库字符集")
	sql2StructCmd.Flags().StringVarP(&dbType, "dbType", "", "", "请出入数据库类型")
	sql2StructCmd.Flags().StringVarP(&dbName, "dbName", "", "", "请出入数据库名")
	sql2StructCmd.Flags().StringVarP(&tableName, "tableName", "", "", "请出入数据库表名")
}
