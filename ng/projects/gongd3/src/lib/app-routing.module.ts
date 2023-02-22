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
	{ path: 'gongd3_go-bars/:GONG__StackPath', component: BarsTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-bar-adder/:GONG__StackPath', component: BarDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-bar-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: BarDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-bar-detail/:id/:GONG__StackPath', component: BarDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-keys/:GONG__StackPath', component: KeysTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-key-adder/:GONG__StackPath', component: KeyDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-key-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: KeyDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-key-detail/:id/:GONG__StackPath', component: KeyDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-pies/:GONG__StackPath', component: PiesTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-pie-adder/:GONG__StackPath', component: PieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-pie-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: PieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-pie-detail/:id/:GONG__StackPath', component: PieDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-scatters/:GONG__StackPath', component: ScattersTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-scatter-adder/:GONG__StackPath', component: ScatterDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-scatter-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: ScatterDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-scatter-detail/:id/:GONG__StackPath', component: ScatterDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-series/:GONG__StackPath', component: SeriesTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-serie-adder/:GONG__StackPath', component: SerieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-serie-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: SerieDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-serie-detail/:id/:GONG__StackPath', component: SerieDetailComponent, outlet: 'gongd3_go_editor' },

	{ path: 'gongd3_go-values/:GONG__StackPath', component: ValuesTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-value-adder/:GONG__StackPath', component: ValueDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-value-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: ValueDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-value-detail/:id/:GONG__StackPath', component: ValueDetailComponent, outlet: 'gongd3_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
