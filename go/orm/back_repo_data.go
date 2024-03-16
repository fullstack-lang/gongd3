// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	BarAPIs []*BarAPI

	KeyAPIs []*KeyAPI

	PieAPIs []*PieAPI

	ScatterAPIs []*ScatterAPI

	SerieAPIs []*SerieAPI

	ValueAPIs []*ValueAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, barDB := range backRepo.BackRepoBar.Map_BarDBID_BarDB {

		var barAPI BarAPI
		barAPI.ID = barDB.ID
		barAPI.BarPointersEncoding = barDB.BarPointersEncoding
		barDB.CopyBasicFieldsToBar_WOP(&barAPI.Bar_WOP)

		backRepoData.BarAPIs = append(backRepoData.BarAPIs, &barAPI)
	}

	for _, keyDB := range backRepo.BackRepoKey.Map_KeyDBID_KeyDB {

		var keyAPI KeyAPI
		keyAPI.ID = keyDB.ID
		keyAPI.KeyPointersEncoding = keyDB.KeyPointersEncoding
		keyDB.CopyBasicFieldsToKey_WOP(&keyAPI.Key_WOP)

		backRepoData.KeyAPIs = append(backRepoData.KeyAPIs, &keyAPI)
	}

	for _, pieDB := range backRepo.BackRepoPie.Map_PieDBID_PieDB {

		var pieAPI PieAPI
		pieAPI.ID = pieDB.ID
		pieAPI.PiePointersEncoding = pieDB.PiePointersEncoding
		pieDB.CopyBasicFieldsToPie_WOP(&pieAPI.Pie_WOP)

		backRepoData.PieAPIs = append(backRepoData.PieAPIs, &pieAPI)
	}

	for _, scatterDB := range backRepo.BackRepoScatter.Map_ScatterDBID_ScatterDB {

		var scatterAPI ScatterAPI
		scatterAPI.ID = scatterDB.ID
		scatterAPI.ScatterPointersEncoding = scatterDB.ScatterPointersEncoding
		scatterDB.CopyBasicFieldsToScatter_WOP(&scatterAPI.Scatter_WOP)

		backRepoData.ScatterAPIs = append(backRepoData.ScatterAPIs, &scatterAPI)
	}

	for _, serieDB := range backRepo.BackRepoSerie.Map_SerieDBID_SerieDB {

		var serieAPI SerieAPI
		serieAPI.ID = serieDB.ID
		serieAPI.SeriePointersEncoding = serieDB.SeriePointersEncoding
		serieDB.CopyBasicFieldsToSerie_WOP(&serieAPI.Serie_WOP)

		backRepoData.SerieAPIs = append(backRepoData.SerieAPIs, &serieAPI)
	}

	for _, valueDB := range backRepo.BackRepoValue.Map_ValueDBID_ValueDB {

		var valueAPI ValueAPI
		valueAPI.ID = valueDB.ID
		valueAPI.ValuePointersEncoding = valueDB.ValuePointersEncoding
		valueDB.CopyBasicFieldsToValue_WOP(&valueAPI.Value_WOP)

		backRepoData.ValueAPIs = append(backRepoData.ValueAPIs, &valueAPI)
	}

}
