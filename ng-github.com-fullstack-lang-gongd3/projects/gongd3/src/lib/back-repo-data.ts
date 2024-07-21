// generated code - do not edit

//insertion point for imports
import { BarAPI } from './bar-api'

import { KeyAPI } from './key-api'

import { PieAPI } from './pie-api'

import { ScatterAPI } from './scatter-api'

import { SerieAPI } from './serie-api'

import { ValueAPI } from './value-api'


export class BackRepoData {
	// insertion point for declarations
	BarAPIs = new Array<BarAPI>()

	KeyAPIs = new Array<KeyAPI>()

	PieAPIs = new Array<PieAPI>()

	ScatterAPIs = new Array<ScatterAPI>()

	SerieAPIs = new Array<SerieAPI>()

	ValueAPIs = new Array<ValueAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.BarAPIs = data?.BarAPIs || [];

		this.KeyAPIs = data?.KeyAPIs || [];

		this.PieAPIs = data?.PieAPIs || [];

		this.ScatterAPIs = data?.ScatterAPIs || [];

		this.SerieAPIs = data?.SerieAPIs || [];

		this.ValueAPIs = data?.ValueAPIs || [];

	}

}