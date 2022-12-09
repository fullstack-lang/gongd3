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
			Heigth: 78.000000,
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
			},
			Fields: []*uml.Field{
				{
					Field: models.Bar{}.Name,
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
