package cmd

import (
	"cmdLine/internal"
	"log"

	"github.com/spf13/cobra"
)

const (
	MODE_UPPER                         = iota + 1 // 单词全部转为大写
	MODE_LOWER                                    // 单词转化为小写
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE            // 转化为大驼峰
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE            // 转化为小驼峰
	MODE_CAMELCASE_TO_UNDERSCORE                  // 驼峰转化为下划线
)

var desc = `
	该子命令支持各种单词转换，模式如下：
	1：全部转化为大写
	2: 全部转化为小写
	3: 转化为大驼峰
	4: 转化为小驼峰
	5: 驼峰转化为下划线
`

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "支持多种单词格式转换",
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = internal.ToUpper(str)
		case MODE_LOWER:
			content = internal.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = internal.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = internal.UnderscoreToLowerCamelcase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = internal.CamelCaseToUnderScore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词的内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
