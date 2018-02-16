package facade

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/markdown-preview/options"
	"github.com/spiegel-im-spiegel/markdown-preview/proc"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

//newVersionCmd returns cobra.Command instance for version sub-command
func newProcCmd() *cobra.Command {
	procCmd := &cobra.Command{
		Use:   "proc [flags] [markdown file]",
		Short: "Processing Markdown",
		Long:  "Processing Markdown",
		RunE: func(cmd *cobra.Command, args []string) error {
			opt := options.New()
			f, err := cmd.Flags().GetBool("github")
			if err != nil {
				return err
			}
			if f {
				opt.SetProcType(options.GitHub)
			}
			f, err = cmd.Flags().GetBool("line-break")
			if err != nil {
				return err
			}
			if f {
				opt.AddExtensions(blackfriday.HardLineBreak)
			}
			f, err = cmd.Flags().GetBool("page")
			if err != nil {
				return err
			}
			if f {
				opt.EnablePageFlag()
			}
			f, err = cmd.Flags().GetBool("sanitize")
			if err != nil {
				return err
			}
			if f {
				opt.EnableSanitizing()
			}
			css, err := cmd.Flags().GetString("css")
			if err != nil {
				return err
			}
			if len(css) > 0 {
				opt.SetCSS(css)
			}

			reader := cui.Reader()
			if len(args) > 0 {
				file, err := os.OpenFile(args[0], os.O_RDONLY, 0400) //args[0] is maybe file path
				if err != nil {
					return err
				}
				defer file.Close()
				reader = file
			}
			html, err := proc.Run(reader, opt)
			if err != nil {
				return err
			}
			cui.WriteFrom(html)
			return nil
		},
	}

	procCmd.Flags().StringP("css", "c", "", "CSS file URL (with --page option)")
	procCmd.Flags().BoolP("github", "g", false, "use GitHub Markdown API")
	procCmd.Flags().BoolP("line-break", "l", false, "translate newlines into line breaks")
	procCmd.Flags().BoolP("page", "p", false, "generate a complete HTML page")
	procCmd.Flags().BoolP("sanitize", "s", false, "sanitize untrusted content")

	return procCmd
}
