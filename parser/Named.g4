/*
Copyright 2021 taichi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
 */
/*
 https://github.com/tshprecher/antlr_psql/blob/master/antlr4/PostgreSQLLexer.g4
 https://github.com/antlr/grammars-v4/blob/master/sql/sqlite/SQLiteLexer.g4
 */
grammar Named;

@header {
// Copyright 2021 taichi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
}

/*
 * Parser Rules
 */

parse: stmt (SEMI+ stmt)* SEMI*;

stmt: (inExpr | anyStmtParts | parameter)+;

inExpr: IN OPEN_PAREN parameter (COMMA parameter)* CLOSE_PAREN;

parameter: namedParamter | qmarkParameter | decParameter | staticParameter;

namedParamter: (AT | COLON) IDENTIFIER (DOT IDENTIFIER)*;

qmarkParameter: QMARK;

decParameter: (DOLLAR | AT P) DIGIT+;

staticParameter: STRING;

anyStmtParts:
  IDENTIFIER (DOT (IDENTIFIER | STAR))?
  | OPEN_PAREN
  | CLOSE_PAREN
  | COMMA
  | STAR
  | ANY_SYMBOL
  | NUMBER;

/*
 * Lexer Rules
 */

// Skip
SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);
BLOCK_COMMENT: '/*' .*? '*/' -> channel(HIDDEN);
LINE_COMMENT: '--' .*? '\n' -> channel(HIDDEN);

fragment HEX_DIGIT: [0-9a-fA-F];
DIGIT: [0-9];

IN: [iI] [nN];

fragment DQUOTA_STRING: '"' ('\\' . | '""' | ~('"' | '\\'))* '"';
fragment SQUOTA_STRING: '\'' ('\\' . | '\'\'' | ~('\'' | '\\'))* '\'';

STRING: DQUOTA_STRING | SQUOTA_STRING;
NUMBER: ((DIGIT+ (DOT DIGIT*)?) | (DOT DIGIT+)) ([eE] [-+]? DIGIT+)? | '0x' HEX_DIGIT+;

IDENTIFIER: LETTER (LETTER | UNICODE_DIGIT)*;

fragment LETTER: UNICODE_LETTER | '_';
fragment UNICODE_LETTER: [\p{L}];
fragment UNICODE_DIGIT: [\p{Nd}];

OPEN_PAREN: '(';
CLOSE_PAREN: ')';
QMARK: '?';
COMMA: ',';
AT: '@';
DOLLAR: '$';
COLON: ':';
SEMI: ';';
DOT: '.';
STAR: '*';
P: 'p' | 'P';

ANY_SYMBOL:
  '['
  | ']'
  | '&'
  | '&&'
  | '&<'
  | '@@'
  | '@>'
  | '!'
  | '!!'
  | '!='
  | '^'
  | '='
  | '=>'
  | '>'
  | '>='
  | '>>'
  | '#'
  | '#='
  | '#>'
  | '#>>'
  | '##'
  | '->'
  | '->>'
  | '-|-'
  | '<'
  | '<='
  | '<@'
  | '<^'
  | '<>'
  | '<->'
  | '<<'
  | '<<='
  | '<?>'
  | '-'
  | '%'
  | '|'
  | '||'
  | '||/'
  | '|/'
  | '+'
  | '?&'
  | '?#'
  | '?-'
  | '?|'
  | '/'
  | '~'
  | '~='
  | '~>=~'
  | '~>~'
  | '~<=~'
  | '~<~'
  | '~*'
  | '~~';
