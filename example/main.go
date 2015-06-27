package main

import (
    "fmt"

    "github.com/microcosm-cc/bluemonday"
)



func main() {
    p := bluemonday.UGCPolicy()
    html := p.Sanitize(
        `<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
    )
    
    //p.Sanitize(string) string
    //p.SanitizeBytes([]byte) []byte
    //p.SanitizeReader(io.Reader) bytes.Buffer

    // Require URLs to be parseable by net/url.Parse and either:
    //   mailto: http:// or https://
    //p.AllowStandardURLs()
    // We only allow <p> and <a href="">
    //p.AllowAttrs("href").OnElements("a")
    //p.AllowElements("p")
    //p.AllowStandardAttributes()
    // Permits the "img" element and it's standard attributes
    //p.AllowImages()
    // Permits ordered and unordered lists, and also definition lists
    //p.AllowLists()
    // Permits HTML tables and all applicable elements and non-styling attributes
    //p.AllowTables()
    //p.RequireParseableURLs(true)
    //p.AllowRelativeURLs(true)
    //p.AllowURLSchemes("mailto", "http", "https")
    //p.AllowAttrs("href").Matching(regexp.MustCompile(`(?i)mailto|https?`)).OnElements("a")
    //And you can take any existing policy and extend it:
    //p := bluemonday.UGCPolicy()
    //p.AllowElements("fieldset", "select", "option")
    //An additional complexity regarding links is the data URI as defined in RFC2397. 
    //The data URI allows for images to be served inline using this format:
    //<img src="data:image/webp;base64,UklGRh4AAABXRUJQVlA4TBEAAAAvAAAAAAfQ//73v/+BiOh/AAA=">
    //We have provided a helper to verify the mimetype followed by base64 content of data URIs links:
    //p.AllowDataURIImages()


    // Output:
    // <a href="http://www.google.com" rel="nofollow">Google</a>
    fmt.Println(html)
}