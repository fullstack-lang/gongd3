import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { BarsTableComponent } from './bars-table/bars-table.component'
import { BarDetailComponent } from './bar-detail/bar-detail.component'

import { KeysTableComponent } from './keys-table/keys-table.component'
import { KeyDetailComponent } from './key-detail/key-detail.component'

import { PiesTableComponent } from './pies-table/pies-table.component'
import { PieDetailComponent } from './pie-detail/pie-detail.component'

import { ScattersTableComponent } from './scatters-table/scatters-table.component'
import { ScatterDetailComponent } from './scatter-detail/scatter-detail.component'

import { SeriesTableComponent } from './series-table/series-table.component'
import { SerieDetailComponent } from './serie-detail/serie-detail.component'

import { ValuesTableComponent } from './values-table/values-table.component'
import { ValueDetailComponent } from './value-detail/value-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongd3_go-bars', component: BarsTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-bar-adder', component: BarDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-bar-adder/:id/:originStruct/:originStructFieldName', component: BarDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-bar-detail/:id', component: BarDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-keys', component: KeysTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-key-adder', component: KeyDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-key-adder/:id/:originStruct/:originStructFieldName', component: KeyDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-key-detail/:id', component: KeyDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-pies', component: PiesTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-pie-adder', component: PieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-pie-adder/:id/:originStruct/:originStructFieldName', component: PieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-pie-detail/:id', component: PieDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-scatters', component: ScattersTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-scatter-adder', component: ScatterDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-scatter-adder/:id/:originStruct/:originStructFieldName', component: ScatterDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-scatter-detail/:id', component: ScatterDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-series', component: SeriesTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-serie-adder', component: SerieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-serie-adder/:id/:originStruct/:originStructFieldName', component: SerieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-serie-detail/:id', component: SerieDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-values', component: ValuesTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-value-adder', component: ValueDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-value-adder/:id/:originStruct/:originStructFieldName', component: ValueDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-value-detail/:id', component: ValueDetailComponent, outlet: 'gongd3_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
