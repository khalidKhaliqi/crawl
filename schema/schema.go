// Code generated by go generate gen.go; DO NOT EDIT.

//go:generate go run gen.go

package schema

var bq = []schemaItem{
	{
		Name: "Address",
		Type: "RECORD",
		Mode: "NULLABLE",
		Fields: []schemaItem{
			{
				Name: "Full",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Scheme",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Opaque",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Host",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Path",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Query",
				Type: "STRING",
				Mode: "NULLABLE",
			},
		},
	},
	{
		Name: "Depth",
		Type: "INT64",
		Mode: "REQUIRED",
	},
	{
		Name: "BodyTextHash",
		Type: "STRING",
		Mode: "NULLABLE",
	},
	{
		Name: "Description",
		Type: "STRING",
		Mode: "NULLABLE",
	},
	{
		Name: "Title",
		Type: "STRING",
		Mode: "NULLABLE",
	},
	{
		Name: "H1",
		Type: "STRING",
		Mode: "NULLABLE",
	},
	{
		Name: "Robots",
		Type: "STRING",
		Mode: "NULLABLE",
	},
	{
		Name: "Canonical",
		Type: "RECORD",
		Mode: "NULLABLE",
		Fields: []schemaItem{
			{
				Name: "Address",
				Type: "RECORD",
				Mode: "NULLABLE",
				Fields: []schemaItem{
					{
						Name: "Full",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Scheme",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Opaque",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Host",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Path",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Query",
						Type: "STRING",
						Mode: "NULLABLE",
					},
				},
			},
			{
				Name: "Href",
				Type: "STRING",
				Mode: "NULLABLE",
			},
		},
	},
	{
		Name: "Links",
		Type: "RECORD",
		Mode: "REPEATED",
		Fields: []schemaItem{
			{
				Name: "Address",
				Type: "RECORD",
				Mode: "NULLABLE",
				Fields: []schemaItem{
					{
						Name: "Full",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Scheme",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Opaque",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Host",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Path",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Query",
						Type: "STRING",
						Mode: "NULLABLE",
					},
				},
			},
			{
				Name: "Anchor",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Href",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Nofollow",
				Type: "BOOL",
				Mode: "NULLABLE",
			},
		},
	},
	{
		Name: "Hreflang",
		Type: "RECORD",
		Mode: "REPEATED",
		Fields: []schemaItem{
			{
				Name: "Address",
				Type: "RECORD",
				Mode: "NULLABLE",
				Fields: []schemaItem{
					{
						Name: "Full",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Scheme",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Opaque",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Host",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Path",
						Type: "STRING",
						Mode: "NULLABLE",
					},
					{
						Name: "Query",
						Type: "STRING",
						Mode: "NULLABLE",
					},
				},
			},
			{
				Name: "Href",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Hreflang",
				Type: "STRING",
				Mode: "NULLABLE",
			},
		},
	},
	{
		Name: "Status",
		Type: "STRING",
		Mode: "NULLABLE",
	},
	{
		Name: "StatusCode",
		Type: "INT64",
		Mode: "NULLABLE",
	},
	{
		Name: "Proto",
		Type: "STRING",
		Mode: "NULLABLE",
	},
	{
		Name: "ProtoMajor",
		Type: "INT64",
		Mode: "NULLABLE",
	},
	{
		Name: "ProtoMinor",
		Type: "INT64",
		Mode: "NULLABLE",
	},
	{
		Name: "Header",
		Type: "RECORD",
		Mode: "REPEATED",
		Fields: []schemaItem{
			{
				Name: "K",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "V",
				Type: "STRING",
				Mode: "NULLABLE",
			},
		},
	},
	{
		Name: "ResolvesTo",
		Type: "RECORD",
		Mode: "NULLABLE",
		Fields: []schemaItem{
			{
				Name: "Full",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Scheme",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Opaque",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Host",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Path",
				Type: "STRING",
				Mode: "NULLABLE",
			},
			{
				Name: "Query",
				Type: "STRING",
				Mode: "NULLABLE",
			},
		},
	},
}
