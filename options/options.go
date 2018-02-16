package options

import blackfriday "gopkg.in/russross/blackfriday.v2"

//Processor is type of Processor
type Processor int

const (
	//Unknown type
	Unknown Processor = iota
	//Blackfriday is russross/blackfriday.v2 package
	Blackfriday
	//GitHub Markdown REST API
	GitHub
)

//Options is russross/blackfriday's options
type Options struct {
	procType       Processor
	sanitize       bool
	extentionFlags blackfriday.Extensions
	htmlFlags      blackfriday.HTMLFlags
	pageFlag       bool
	css            string
	title          string
}

//New returns instance of Options
func New() *Options {
	extFlags := blackfriday.CommonExtensions //NoIntraEmphasis | Tables | FencedCode | Autolink | Strikethrough | SpaceHeadings | HeadingIDs | BackslashLineBreak | DefinitionLists
	extFlags |= blackfriday.Footnotes        //Pandoc-style footnotes
	extFlags |= blackfriday.HeadingIDs       //specify heading IDs  with {#id}
	extFlags |= blackfriday.Titleblock       //Titleblock ala pandoc
	extFlags |= blackfriday.DefinitionLists  //Render definition lists

	htmlFlags := blackfriday.HTMLFlagsNone
	htmlFlags |= blackfriday.FootnoteReturnLinks     //Generate a link at the end of a footnote to return to the source
	htmlFlags |= blackfriday.Smartypants             //Enable smart punctuation substitutions
	htmlFlags |= blackfriday.SmartypantsFractions    // Enable smart fractions (with Smartypants)
	htmlFlags |= blackfriday.SmartypantsDashes       // Enable smart dashes (with Smartypants)
	htmlFlags |= blackfriday.SmartypantsLatexDashes  // Enable LaTeX-style dashes (with Smartypants)
	htmlFlags |= blackfriday.SmartypantsAngledQuotes // Enable angled double quotes (with Smartypants) for double quotes rendering
	htmlFlags |= blackfriday.SmartypantsQuotesNBSP   // Enable French guillemets Â» (with Smartypants)

	return &Options{procType: Blackfriday, sanitize: false, extentionFlags: extFlags, htmlFlags: htmlFlags, css: "", title: ""}
}

//SetProcType sets Processor type
func (o *Options) SetProcType(p Processor) {
	o.procType = p
}

//ProcType returns Processor type
func (o *Options) ProcType() Processor {
	return o.procType
}

//EnableSanitizing sets sanitize flag
func (o *Options) EnableSanitizing() {
	o.sanitize = true
}

//Sanitizing returns sanitize flag
func (o *Options) Sanitizing() bool {
	return o.sanitize
}

//AddExtensions adds blackfriday.Extensions flag
func (o *Options) AddExtensions(e blackfriday.Extensions) {
	o.extentionFlags |= e
}

//Extensions returns blackfriday.Extensions flag
func (o *Options) Extensions() blackfriday.Extensions {
	return o.extentionFlags
}

//AddHTMLFlags adds blackfriday.HTMLFlags flag
func (o *Options) AddHTMLFlags(f blackfriday.HTMLFlags) {
	if (blackfriday.CompletePage & f) == 0 {
		o.htmlFlags |= f
	}
}

//HTMLFlags returns blackfriday.HTMLFlags flag
func (o *Options) HTMLFlags() blackfriday.HTMLFlags {
	return o.htmlFlags
}

//EnablePageFlag sets page rendering flag
func (o *Options) EnablePageFlag() {
	o.pageFlag = true
}

//PageFlag returns page rendering flag
func (o *Options) PageFlag() bool {
	return o.pageFlag
}

//SetCSS sets URL of CSS
func (o *Options) SetCSS(css string) {
	o.css = css
}

//CSS returns URL of CSS
func (o *Options) CSS() string {
	return o.css
}

//SetTitle sets title string
func (o *Options) SetTitle(title string) {
	o.title = title
}

//Title returns title string
func (o *Options) Title() string {
	return o.title
}
