package ast

import "github.com/daniil-kzn/graphql-go-tools/pkg/lexer/position"

type SchemaExtension struct {
	ExtendLiteral position.Position
	SchemaDefinition
}
