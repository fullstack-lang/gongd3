package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"gongd3/go/models"
)

var NewDiagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			ReferencedGong: &(models.Bar{}),
			Position: &uml.Position{
				X: 40.000000,
				Y: 190.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
			Links: []*uml.Link{
				{
					Field: models.Bar{}.Set,
					Middlevertice: &uml.Vertice{
						X: 375.000000,
						Y: 234.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Bar{}.X,
					Middlevertice: &uml.Vertice{
						X: 138.000000,
						Y: 415.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Bar{}.Y,
					Middlevertice: &uml.Vertice{
						X: 250.000000,
						Y: 329.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Bar{}.Heigth,
				},
				{
					Field: models.Bar{}.Margin,
				},
				{
					Field: models.Bar{}.Name,
				},
				{
					Field: models.Bar{}.Width,
				},
			},
		},
		{
			ReferencedGong: &(models.Key{}),
			Position: &uml.Position{
				X: 440.000000,
				Y: 370.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Key{}.Name,
				},
			},
		},
		{
			ReferencedGong: &(models.Pie{}),
			Position: &uml.Position{
				X: 270.000000,
				Y: 590.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
			Links: []*uml.Link{
				{
					Field: models.Pie{}.Set,
					Middlevertice: &uml.Vertice{
						X: 775.000000,
						Y: 351.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Pie{}.X,
					Middlevertice: &uml.Vertice{
						X: 505.000000,
						Y: 498.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Pie{}.Y,
					Middlevertice: &uml.Vertice{
						X: 355.000000,
						Y: 508.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Pie{}.Heigth,
				},
				{
					Field: models.Pie{}.Margin,
				},
				{
					Field: models.Pie{}.Name,
				},
				{
					Field: models.Pie{}.Width,
				},
			},
		},
		{
			ReferencedGong: &(models.Scatter{}),
			Position: &uml.Position{
				X: 800.000000,
				Y: 610.000000,
			},
			Width:  240.000000,
			Heigth: 123.000000,
			Links: []*uml.Link{
				{
					Field: models.Scatter{}.Set,
					Middlevertice: &uml.Vertice{
						X: 1085.000000,
						Y: 384.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Scatter{}.Text,
					Middlevertice: &uml.Vertice{
						X: 995.000000,
						Y: 484.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Scatter{}.X,
					Middlevertice: &uml.Vertice{
						X: 885.000000,
						Y: 541.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Scatter{}.Y,
					Middlevertice: &uml.Vertice{
						X: 635.000000,
						Y: 621.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Scatter{}.Heigth,
				},
				{
					Field: models.Scatter{}.Margin,
				},
				{
					Field: models.Scatter{}.Name,
				},
				{
					Field: models.Scatter{}.Width,
				},
			},
		},
		{
			ReferencedGong: &(models.Serie{}),
			Position: &uml.Position{
				X: 440.000000,
				Y: 190.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Serie{}.Key,
					Middlevertice: &uml.Vertice{
						X: 563.000000,
						Y: 293.000000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Serie{}.Values,
					Middlevertice: &uml.Vertice{
						X: 713.000000,
						Y: 226.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Serie{}.Name,
				},
			},
		},
		{
			ReferencedGong: &(models.Value{}),
			Position: &uml.Position{
				X: 810.000000,
				Y: 190.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Value{}.Name,
				},
			},
		},
	},
	Notes: []*uml.NoteShape{
	},
}
