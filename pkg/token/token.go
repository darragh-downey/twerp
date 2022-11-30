package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifiers + literals
  IDENT = "IDENT" // add, foobar, x, y, etc.
  INT = "INT" // 132456

  // Operators
  ASSIGN = "="
  PLUS = "+"

  //Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)

