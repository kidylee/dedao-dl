package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yann0917/dedao-dl/cmd/app"
)

var ebookCmd = &cobra.Command{
	Use:     "ebook",
	Short:   "获取我的电子书架",
	Long:    `使用 dedao-dl ebook 获取我的电子书架`,
	Args:    cobra.OnlyValidArgs,
	Example: "dedao-dl ebook",
	PreRunE: AuthFunc,
	RunE: func(cmd *cobra.Command, args []string) error {
		if bookID > 0 {
			return ebookDetail(bookID)
		}
		return courseList(app.CateEbook)
	},
}

func init() {
	rootCmd.AddCommand(ebookCmd)

	ebookCmd.PersistentFlags().IntVarP(&bookID, "id", "i", 0, "电子书ID")
	// rootCmd.PersistentFlags().StringVarP(&cType, "type", "t", "bauhinia", "课程类型(all, bauhinia, ebook, compass")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}

func ebookDetail(id int) (err error) {
	detail, err := app.EbookDetail(id)
	if err != nil {
		return
	}

	out := os.Stdout
	table := tablewriter.NewWriter(out)
	fmt.Fprint(out, "书名："+detail.OperatingTitle+"\n")
	fmt.Fprint(out, "单价："+detail.Price+"\n")
	fmt.Fprint(out, "作者："+detail.BookAuthor+"\n")
	fmt.Fprint(out, "类型："+detail.ClassifyName+"\n")
	fmt.Fprint(out, "专家推荐指数："+detail.ProductScore+"\n")
	fmt.Fprint(out, "豆瓣评分："+detail.DoubanScore+"\n")
	fmt.Fprint(out, "发行日期："+detail.PublishTime+"\n")
	fmt.Fprint(out, "出版社："+detail.Press.Name+"\n")
	fmt.Fprintln(out)

	table.SetHeader([]string{"#", "ID", "章节名称"})
	table.SetAutoWrapText(false)
	// table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	// table.SetCenterSeparator("|")
	for i, p := range detail.CatalogList {
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.PlayOrder),
			p.Text,
		})
	}
	table.Render()
	return
}
