package jsval_test

import "github.com/lestrrat/go-jsval"

func JSValFoo() *jsval.JSVal {
	V := jsval.New()
	V.SetRoot(
		jsval.Object().
			AdditionalProperties(
				jsval.NilConstraint,
			).
			AddProp(
				`$schema`,
				jsval.String().Format("uri"),
			).
			AddProp(
				`additionalItems`,
				jsval.Any().
					Add(
						jsval.Boolean(),
					).
					Add(
						jsval.Reference(V).RefersTo(`#`),
					),
			).
			AddProp(
				`additionalProperties`,
				jsval.Any().
					Add(
						jsval.Boolean(),
					).
					Add(
						jsval.Reference(V).RefersTo(`#`),
					),
			).
			AddProp(
				`allOf`,
				jsval.Reference(V).RefersTo(`#/definitions/schemaArray`),
			).
			AddProp(
				`anyOf`,
				jsval.Reference(V).RefersTo(`#/definitions/schemaArray`),
			).
			AddProp(
				`default`,
				jsval.NilConstraint,
			).
			AddProp(
				`definitions`,
				jsval.Object().
					AdditionalProperties(
						jsval.Reference(V).RefersTo(`#`),
					),
			).
			AddProp(
				`dependencies`,
				jsval.Object().
					AdditionalProperties(
						jsval.Any().
							Add(
								jsval.Reference(V).RefersTo(`#`),
							).
							Add(
								jsval.Reference(V).RefersTo(`#/definitions/stringArray`),
							),
					),
			).
			AddProp(
				`description`,
				jsval.String(),
			).
			AddProp(
				`enum`,
				jsval.Array().MinItems(1).MaxItems(0).UniqueItems(),
			).
			AddProp(
				`exclusiveMaximum`,
				jsval.Boolean().Default(false),
			).
			AddProp(
				`exclusiveMinimum`,
				jsval.Boolean().Default(false),
			).
			AddProp(
				`id`,
				jsval.String().Format("uri"),
			).
			AddProp(
				`items`,
				jsval.Any().
					Add(
						jsval.Reference(V).RefersTo(`#`),
					).
					Add(
						jsval.Reference(V).RefersTo(`#/definitions/schemaArray`),
					),
			).
			AddProp(
				`maxItems`,
				jsval.Reference(V).RefersTo(`#/definitions/positiveInteger`),
			).
			AddProp(
				`maxLength`,
				jsval.Reference(V).RefersTo(`#/definitions/positiveInteger`),
			).
			AddProp(
				`maxProperties`,
				jsval.Reference(V).RefersTo(`#/definitions/positiveInteger`),
			).
			AddProp(
				`maximum`,
				jsval.Number(),
			).
			AddProp(
				`minItems`,
				jsval.Reference(V).RefersTo(`#/definitions/positiveIntegerDefault0`),
			).
			AddProp(
				`minLength`,
				jsval.Reference(V).RefersTo(`#/definitions/positiveIntegerDefault0`),
			).
			AddProp(
				`minProperties`,
				jsval.Reference(V).RefersTo(`#/definitions/positiveIntegerDefault0`),
			).
			AddProp(
				`minimum`,
				jsval.Number(),
			).
			AddProp(
				`multipleOf`,
				jsval.Number().Minimum(0.000000).ExclusiveMinimum(true),
			).
			AddProp(
				`not`,
				jsval.Reference(V).RefersTo(`#`),
			).
			AddProp(
				`oneOf`,
				jsval.Reference(V).RefersTo(`#/definitions/schemaArray`),
			).
			AddProp(
				`pattern`,
				jsval.String().Format("regex"),
			).
			AddProp(
				`patternProperties`,
				jsval.Object().
					AdditionalProperties(
						jsval.Reference(V).RefersTo(`#`),
					),
			).
			AddProp(
				`properties`,
				jsval.Object().
					AdditionalProperties(
						jsval.Reference(V).RefersTo(`#`),
					),
			).
			AddProp(
				`required`,
				jsval.Reference(V).RefersTo(`#/definitions/stringArray`),
			).
			AddProp(
				`title`,
				jsval.String(),
			).
			AddProp(
				`type`,
				jsval.Any().
					Add(
						jsval.Reference(V).RefersTo(`#/definitions/simpleTypes`),
					).
					Add(
						jsval.Array().MinItems(1).MaxItems(0).UniqueItems(),
					),
			).
			AddProp(
				`uniqueItems`,
				jsval.Boolean().Default(false),
			).
			PropDependency("exclusiveMinimum", "minimum").
			PropDependency("exclusiveMaximum", "maximum"),
	)

	V.SetReference(
		`#`,
		V.Root(),
	)
	V.SetReference(
		`#/definitions/positiveInteger`,
		jsval.Integer().Minimum(0),
	)
	V.SetReference(
		`#/definitions/positiveIntegerDefault0`,
		jsval.All().
			Add(
				jsval.Reference(V).RefersTo(`#/definitions/positiveInteger`),
			).
			Add(
				jsval.NilConstraint,
			),
	)
	V.SetReference(
		`#/definitions/schemaArray`,
		jsval.Array().MinItems(1).MaxItems(0),
	)
	V.SetReference(
		`#/definitions/simpleTypes`,
		jsval.String().Enum([]interface{}{"array", "boolean", "integer", "null", "number", "object", "string"},),
	)
	V.SetReference(
		`#/definitions/stringArray`,
		jsval.Array().MinItems(1).MaxItems(0).UniqueItems(),
	)
	return V
}
