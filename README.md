# ecms-go-filter (work-in-progress)
Filter package for filter/transforming values.

## Docs summary
- bytes 
    - [-] - SubSequences - Untested.
    - [-] - SliceSubSequences - Untested.
- string
    - [x] - LowerCase
    - [x] - Trim
    - [x] - XmlEntities
    - [-] - GetStripHtmlTags (tested/completed with caveats (doesnt' support self closing tags, and tags with attributes containing json content)).
    - [x] - GetStripPopulatedHtmlAttribs (currently doesn't support boolean html attribs (hence name))
    - [-] - GetPatternFilter (there is in implementation but it is untested)
    - [x] - GetBoolFilter

## License
Apache 2.0 License
