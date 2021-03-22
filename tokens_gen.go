// Code generated by internal/cmd/gentokens/main.go; DO NOT EDIT.

package schemalex

// TokenType describes the possible types of tokens that schemalex understands
type TokenType int

// Token represents a token
type Token struct {
	Type  TokenType
	Value string
	Pos   int
	Line  int
	Col   int
	EOF   bool
}

// NewToken creates a new token of type `t`, with value `v`
func NewToken(t TokenType, v string) *Token {
	return &Token{Type: t, Value: v}
}

// List of possible tokens
const (
	ILLEGAL TokenType = iota
	EOF
	SPACE
	IDENT
	BACKTICK_IDENT
	DOUBLE_QUOTE_IDENT
	SINGLE_QUOTE_IDENT
	NUMBER
	LPAREN        // (
	RPAREN        // )
	COMMA         // ,
	SEMICOLON     // ;
	DOT           // .
	SLASH         // /
	ASTERISK      // *
	DASH          // -
	PLUS          // +
	SINGLE_QUOTE  // '
	DOUBLE_QUOTE  // "
	EQUAL         // =
	COMMENT_IDENT // // /*   */, --, #
	ACTION
	AUTO_INCREMENT
	AVG_ROW_LENGTH
	BIGINT
	BINARY
	BIT
	BLOB
	BOOL
	BOOLEAN
	BTREE
	CASCADE
	CHAR
	CHARACTER
	CHARSET
	CHECK
	CHECKSUM
	COLLATE
	COMMENT
	COMPACT
	COMPRESSED
	CONNECTION
	CONSTRAINT
	CREATE
	CURRENT_TIMESTAMP
	DATA
	DATABASE
	DATE
	DATETIME
	DECIMAL
	DEFAULT
	DELAY_KEY_WRITE
	DELETE
	DIRECTORY
	DISK
	DOUBLE
	DROP
	DYNAMIC
	ENGINE
	ENUM
	EXISTS
	FALSE
	FIRST
	FIXED
	FLOAT
	FOREIGN
	FULL
	FULLTEXT
	GEOMETRY
	HASH
	IF
	INDEX
	INSERT_METHOD
	INT
	INTEGER
	JSON
	KEY
	KEY_BLOCK_SIZE
	LAST
	LIKE
	LONGBLOB
	LONGTEXT
	MATCH
	MAX_ROWS
	MEDIUMBLOB
	MEDIUMINT
	MEDIUMTEXT
	MEMORY
	MIN_ROWS
	NO
	NOT
	NULL
	NUMERIC
	ON
	PACK_KEYS
	PARTIAL
	PARSER
	PASSWORD
	PRIMARY
	REAL
	REDUNDANT
	REFERENCES
	RESTRICT
	ROW_FORMAT
	SET
	SIMPLE
	SMALLINT
	SPATIAL
	STATS_AUTO_RECALC
	STATS_PERSISTENT
	STATS_SAMPLE_PAGES
	STORAGE
	TABLE
	TABLESPACE
	TEMPORARY
	TEXT
	TIME
	TIMESTAMP
	TINYBLOB
	TINYINT
	TINYTEXT
	TRUE
	UNION
	UNIQUE
	UNSIGNED
	UPDATE
	USE
	USING
	VARBINARY
	VARCHAR
	YEAR
	WITH
	ZEROFILL
	ASC
	DESC
	NOW
)

var keywordIdentMap = map[string]TokenType{
	"ACTION":             ACTION,
	"AUTO_INCREMENT":     AUTO_INCREMENT,
	"AVG_ROW_LENGTH":     AVG_ROW_LENGTH,
	"BIGINT":             BIGINT,
	"BINARY":             BINARY,
	"BIT":                BIT,
	"BLOB":               BLOB,
	"BOOL":               BOOL,
	"BOOLEAN":            BOOLEAN,
	"BTREE":              BTREE,
	"CASCADE":            CASCADE,
	"CHAR":               CHAR,
	"CHARACTER":          CHARACTER,
	"CHARSET":            CHARSET,
	"CHECK":              CHECK,
	"CHECKSUM":           CHECKSUM,
	"COLLATE":            COLLATE,
	"COMMENT":            COMMENT,
	"COMPACT":            COMPACT,
	"COMPRESSED":         COMPRESSED,
	"CONNECTION":         CONNECTION,
	"CONSTRAINT":         CONSTRAINT,
	"CREATE":             CREATE,
	"CURRENT_TIMESTAMP":  CURRENT_TIMESTAMP,
	"DATA":               DATA,
	"DATABASE":           DATABASE,
	"DATE":               DATE,
	"DATETIME":           DATETIME,
	"DECIMAL":            DECIMAL,
	"DEFAULT":            DEFAULT,
	"DELAY_KEY_WRITE":    DELAY_KEY_WRITE,
	"DELETE":             DELETE,
	"DIRECTORY":          DIRECTORY,
	"DISK":               DISK,
	"DOUBLE":             DOUBLE,
	"DROP":               DROP,
	"DYNAMIC":            DYNAMIC,
	"ENGINE":             ENGINE,
	"ENUM":               ENUM,
	"EXISTS":             EXISTS,
	"FALSE":              FALSE,
	"FIRST":              FIRST,
	"FIXED":              FIXED,
	"FLOAT":              FLOAT,
	"FOREIGN":            FOREIGN,
	"FULL":               FULL,
	"FULLTEXT":           FULLTEXT,
	"GEOMETRY":           GEOMETRY,
	"HASH":               HASH,
	"IF":                 IF,
	"INDEX":              INDEX,
	"INSERT_METHOD":      INSERT_METHOD,
	"INT":                INT,
	"INTEGER":            INTEGER,
	"JSON":               JSON,
	"KEY":                KEY,
	"KEY_BLOCK_SIZE":     KEY_BLOCK_SIZE,
	"LAST":               LAST,
	"LIKE":               LIKE,
	"LONGBLOB":           LONGBLOB,
	"LONGTEXT":           LONGTEXT,
	"MATCH":              MATCH,
	"MAX_ROWS":           MAX_ROWS,
	"MEDIUMBLOB":         MEDIUMBLOB,
	"MEDIUMINT":          MEDIUMINT,
	"MEDIUMTEXT":         MEDIUMTEXT,
	"MEMORY":             MEMORY,
	"MIN_ROWS":           MIN_ROWS,
	"NO":                 NO,
	"NOT":                NOT,
	"NULL":               NULL,
	"NUMERIC":            NUMERIC,
	"ON":                 ON,
	"PACK_KEYS":          PACK_KEYS,
	"PARTIAL":            PARTIAL,
	"PARSER":             PARSER,
	"PASSWORD":           PASSWORD,
	"PRIMARY":            PRIMARY,
	"REAL":               REAL,
	"REDUNDANT":          REDUNDANT,
	"REFERENCES":         REFERENCES,
	"RESTRICT":           RESTRICT,
	"ROW_FORMAT":         ROW_FORMAT,
	"SET":                SET,
	"SIMPLE":             SIMPLE,
	"SMALLINT":           SMALLINT,
	"SPATIAL":            SPATIAL,
	"STATS_AUTO_RECALC":  STATS_AUTO_RECALC,
	"STATS_PERSISTENT":   STATS_PERSISTENT,
	"STATS_SAMPLE_PAGES": STATS_SAMPLE_PAGES,
	"STORAGE":            STORAGE,
	"TABLE":              TABLE,
	"TABLESPACE":         TABLESPACE,
	"TEMPORARY":          TEMPORARY,
	"TEXT":               TEXT,
	"TIME":               TIME,
	"TIMESTAMP":          TIMESTAMP,
	"TINYBLOB":           TINYBLOB,
	"TINYINT":            TINYINT,
	"TINYTEXT":           TINYTEXT,
	"TRUE":               TRUE,
	"UNION":              UNION,
	"UNIQUE":             UNIQUE,
	"UNSIGNED":           UNSIGNED,
	"UPDATE":             UPDATE,
	"USE":                USE,
	"USING":              USING,
	"VARBINARY":          VARBINARY,
	"VARCHAR":            VARCHAR,
	"YEAR":               YEAR,
	"WITH":               WITH,
	"ZEROFILL":           ZEROFILL,
	"ASC":                ASC,
	"DESC":               DESC,
	"NOW":                NOW,
}

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case SPACE:
		return "SPACE"
	case IDENT:
		return "IDENT"
	case BACKTICK_IDENT:
		return "BACKTICK_IDENT"
	case DOUBLE_QUOTE_IDENT:
		return "DOUBLE_QUOTE_IDENT"
	case SINGLE_QUOTE_IDENT:
		return "SINGLE_QUOTE_IDENT"
	case NUMBER:
		return "NUMBER"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case COMMA:
		return "COMMA"
	case SEMICOLON:
		return "SEMICOLON"
	case DOT:
		return "DOT"
	case SLASH:
		return "SLASH"
	case ASTERISK:
		return "ASTERISK"
	case DASH:
		return "DASH"
	case PLUS:
		return "PLUS"
	case SINGLE_QUOTE:
		return "SINGLE_QUOTE"
	case DOUBLE_QUOTE:
		return "DOUBLE_QUOTE"
	case EQUAL:
		return "EQUAL"
	case COMMENT_IDENT:
		return "COMMENT_IDENT"
	case ACTION:
		return "ACTION"
	case AUTO_INCREMENT:
		return "AUTO_INCREMENT"
	case AVG_ROW_LENGTH:
		return "AVG_ROW_LENGTH"
	case BIGINT:
		return "BIGINT"
	case BINARY:
		return "BINARY"
	case BIT:
		return "BIT"
	case BLOB:
		return "BLOB"
	case BOOL:
		return "BOOL"
	case BOOLEAN:
		return "BOOLEAN"
	case BTREE:
		return "BTREE"
	case CASCADE:
		return "CASCADE"
	case CHAR:
		return "CHAR"
	case CHARACTER:
		return "CHARACTER"
	case CHARSET:
		return "CHARSET"
	case CHECK:
		return "CHECK"
	case CHECKSUM:
		return "CHECKSUM"
	case COLLATE:
		return "COLLATE"
	case COMMENT:
		return "COMMENT"
	case COMPACT:
		return "COMPACT"
	case COMPRESSED:
		return "COMPRESSED"
	case CONNECTION:
		return "CONNECTION"
	case CONSTRAINT:
		return "CONSTRAINT"
	case CREATE:
		return "CREATE"
	case CURRENT_TIMESTAMP:
		return "CURRENT_TIMESTAMP"
	case DATA:
		return "DATA"
	case DATABASE:
		return "DATABASE"
	case DATE:
		return "DATE"
	case DATETIME:
		return "DATETIME"
	case DECIMAL:
		return "DECIMAL"
	case DEFAULT:
		return "DEFAULT"
	case DELAY_KEY_WRITE:
		return "DELAY_KEY_WRITE"
	case DELETE:
		return "DELETE"
	case DIRECTORY:
		return "DIRECTORY"
	case DISK:
		return "DISK"
	case DOUBLE:
		return "DOUBLE"
	case DROP:
		return "DROP"
	case DYNAMIC:
		return "DYNAMIC"
	case ENGINE:
		return "ENGINE"
	case ENUM:
		return "ENUM"
	case EXISTS:
		return "EXISTS"
	case FALSE:
		return "FALSE"
	case FIRST:
		return "FIRST"
	case FIXED:
		return "FIXED"
	case FLOAT:
		return "FLOAT"
	case FOREIGN:
		return "FOREIGN"
	case FULL:
		return "FULL"
	case FULLTEXT:
		return "FULLTEXT"
	case GEOMETRY:
		return "GEOMETRY"
	case HASH:
		return "HASH"
	case IF:
		return "IF"
	case INDEX:
		return "INDEX"
	case INSERT_METHOD:
		return "INSERT_METHOD"
	case INT:
		return "INT"
	case INTEGER:
		return "INTEGER"
	case JSON:
		return "JSON"
	case KEY:
		return "KEY"
	case KEY_BLOCK_SIZE:
		return "KEY_BLOCK_SIZE"
	case LAST:
		return "LAST"
	case LIKE:
		return "LIKE"
	case LONGBLOB:
		return "LONGBLOB"
	case LONGTEXT:
		return "LONGTEXT"
	case MATCH:
		return "MATCH"
	case MAX_ROWS:
		return "MAX_ROWS"
	case MEDIUMBLOB:
		return "MEDIUMBLOB"
	case MEDIUMINT:
		return "MEDIUMINT"
	case MEDIUMTEXT:
		return "MEDIUMTEXT"
	case MEMORY:
		return "MEMORY"
	case MIN_ROWS:
		return "MIN_ROWS"
	case NO:
		return "NO"
	case NOT:
		return "NOT"
	case NULL:
		return "NULL"
	case NUMERIC:
		return "NUMERIC"
	case ON:
		return "ON"
	case PACK_KEYS:
		return "PACK_KEYS"
	case PARTIAL:
		return "PARTIAL"
	case PARSER:
		return "PARSER"
	case PASSWORD:
		return "PASSWORD"
	case PRIMARY:
		return "PRIMARY"
	case REAL:
		return "REAL"
	case REDUNDANT:
		return "REDUNDANT"
	case REFERENCES:
		return "REFERENCES"
	case RESTRICT:
		return "RESTRICT"
	case ROW_FORMAT:
		return "ROW_FORMAT"
	case SET:
		return "SET"
	case SIMPLE:
		return "SIMPLE"
	case SMALLINT:
		return "SMALLINT"
	case SPATIAL:
		return "SPATIAL"
	case STATS_AUTO_RECALC:
		return "STATS_AUTO_RECALC"
	case STATS_PERSISTENT:
		return "STATS_PERSISTENT"
	case STATS_SAMPLE_PAGES:
		return "STATS_SAMPLE_PAGES"
	case STORAGE:
		return "STORAGE"
	case TABLE:
		return "TABLE"
	case TABLESPACE:
		return "TABLESPACE"
	case TEMPORARY:
		return "TEMPORARY"
	case TEXT:
		return "TEXT"
	case TIME:
		return "TIME"
	case TIMESTAMP:
		return "TIMESTAMP"
	case TINYBLOB:
		return "TINYBLOB"
	case TINYINT:
		return "TINYINT"
	case TINYTEXT:
		return "TINYTEXT"
	case TRUE:
		return "TRUE"
	case UNION:
		return "UNION"
	case UNIQUE:
		return "UNIQUE"
	case UNSIGNED:
		return "UNSIGNED"
	case UPDATE:
		return "UPDATE"
	case USE:
		return "USE"
	case USING:
		return "USING"
	case VARBINARY:
		return "VARBINARY"
	case VARCHAR:
		return "VARCHAR"
	case YEAR:
		return "YEAR"
	case WITH:
		return "WITH"
	case ZEROFILL:
		return "ZEROFILL"
	case ASC:
		return "ASC"
	case DESC:
		return "DESC"
	case NOW:
		return "NOW"
	}
	return "(invalid)"
}
