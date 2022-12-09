import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { CountrysTableComponent } from './countrys-table/countrys-table.component'
import { CountryDetailComponent } from './country-detail/country-detail.component'
import { CountryPresentationComponent } from './country-presentation/country-presentation.component'

import { HellosTableComponent } from './hellos-table/hellos-table.component'
import { HelloDetailComponent } from './hello-detail/hello-detail.component'
import { HelloPresentationComponent } from './hello-presentation/hello-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongd3_go-countrys', component: CountrysTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-country-adder', component: CountryDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-country-adder/:id/:originStruct/:originStructFieldName', component: CountryDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-country-detail/:id', component: CountryDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-country-presentation/:id', component: CountryPresentationComponent, outlet: 'gongd3_go_presentation' },
	{ path: 'gongd3_go-country-presentation-special/:id', component: CountryPresentationComponent, outlet: 'gongd3_gocountrypres' },

	{ path: 'gongd3_go-hellos', component: HellosTableComponent, outlet: 'gongd3_go_table' },
	{ path: 'gongd3_go-hello-adder', component: HelloDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-hello-adder/:id/:originStruct/:originStructFieldName', component: HelloDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-hello-detail/:id', component: HelloDetailComponent, outlet: 'gongd3_go_editor' },
	{ path: 'gongd3_go-hello-presentation/:id', component: HelloPresentationComponent, outlet: 'gongd3_go_presentation' },
	{ path: 'gongd3_go-hello-presentation-special/:id', component: HelloPresentationComponent, outlet: 'gongd3_gohellopres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
