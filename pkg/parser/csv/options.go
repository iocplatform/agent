package csv

// csvOptions configures a CSV parser.
type csvOptions struct {
	separator string
	fields    []string
	skipLines int

	tagMap       map[string][]string
	tagField     string
	tagSeparator string
	dateField    string
	dateFormat   string
	typeField    string
}

// Option configures how set up the connection
type Option func(*csvOptions)

// WithSeparator returns a CsvOption which sets the separator for input splitting
func WithSeparator(separator string) Option {
	return func(o *csvOptions) {
		o.separator = separator
	}
}

// WithFields returns o CsvOption which sets the fields attribute
func WithFields(fields ...string) Option {
	return func(o *csvOptions) {
		o.fields = fields
	}
}

// WithSkipLines returns o CsvOption which sets the skip lines attribute.
func WithSkipLines(sl int) Option {
	return func(o *csvOptions) {
		o.skipLines = sl
	}
}

// WithTagMap returns a CsvOption which set the tag map for translation
func WithTagMap(maps map[string][]string) Option {
	return func(o *csvOptions) {
		o.tagMap = maps
	}
}

// WithTagField returns a CsvOption which set the tag field name
func WithTagField(field string) Option {
	return func(o *csvOptions) {
		o.tagField = field
	}
}

// WithTagSeparator returns a CsvOption which set the tag separator
func WithTagSeparator(separator string) Option {
	return func(o *csvOptions) {
		o.tagSeparator = separator
	}
}

// WithDateField returns a CsvOption which set the date field name
func WithDateField(field string) Option {
	return func(o *csvOptions) {
		o.dateField = field
	}
}

// WithDateFormat returns a CsvOption which set the date format
func WithDateFormat(format string) Option {
	return func(o *csvOptions) {
		o.dateFormat = format
	}
}

// WithTypeField returns a CsvOption which set the type field name
func WithTypeField(field string) Option {
	return func(o *csvOptions) {
		o.typeField = field
	}
}
