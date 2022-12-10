import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { BarsTableComponent } from './bars-table/bars-table.component'
import { BarSortingComponent } from './bar-sorting/bar-sorting.component'
import { BarDetailComponent } from './bar-detail/bar-detail.component'
import { BarPresentationComponent } from './bar-presentation/bar-presentation.component'

import { KeysTableComponent } from './keys-table/keys-table.component'
import { KeySortingComponent } from './key-sorting/key-sorting.component'
import { KeyDetailComponent } from './key-detail/key-detail.component'
import { KeyPresentationComponent } from './key-presentation/key-presentation.component'

import { PiesTableComponent } from './pies-table/pies-table.component'
import { PieSortingComponent } from './pie-sorting/pie-sorting.component'
import { PieDetailComponent } from './pie-detail/pie-detail.component'
import { PiePresentationComponent } from './pie-presentation/pie-presentation.component'

import { ScattersTableComponent } from './scatters-table/scatters-table.component'
import { ScatterSortingComponent } from './scatter-sorting/scatter-sorting.component'
import { ScatterDetailComponent } from './scatter-detail/scatter-detail.component'
import { ScatterPresentationComponent } from './scatter-presentation/scatter-presentation.component'

import { SeriesTableComponent } from './series-table/series-table.component'
import { SerieSortingComponent } from './serie-sorting/serie-sorting.component'
import { SerieDetailComponent } from './serie-detail/serie-detail.component'
import { SeriePresentationComponent } from './serie-presentation/serie-presentation.component'

import { ValuesTableComponent } from './values-table/values-table.component'
import { ValueSortingComponent } from './value-sorting/value-sorting.component'
import { ValueDetailComponent } from './value-detail/value-detail.component'
import { ValuePresentationComponent } from './value-presentation/value-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		BarsTableComponent,
		BarSortingComponent,
		BarDetailComponent,
		BarPresentationComponent,

		KeysTableComponent,
		KeySortingComponent,
		KeyDetailComponent,
		KeyPresentationComponent,

		PiesTableComponent,
		PieSortingComponent,
		PieDetailComponent,
		PiePresentationComponent,

		ScattersTableComponent,
		ScatterSortingComponent,
		ScatterDetailComponent,
		ScatterPresentationComponent,

		SeriesTableComponent,
		SerieSortingComponent,
		SerieDetailComponent,
		SeriePresentationComponent,

		ValuesTableComponent,
		ValueSortingComponent,
		ValueDetailComponent,
		ValuePresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		BarsTableComponent,
		BarSortingComponent,
		BarDetailComponent,
		BarPresentationComponent,

		KeysTableComponent,
		KeySortingComponent,
		KeyDetailComponent,
		KeyPresentationComponent,

		PiesTableComponent,
		PieSortingComponent,
		PieDetailComponent,
		PiePresentationComponent,

		ScattersTableComponent,
		ScatterSortingComponent,
		ScatterDetailComponent,
		ScatterPresentationComponent,

		SeriesTableComponent,
		SerieSortingComponent,
		SerieDetailComponent,
		SeriePresentationComponent,

		ValuesTableComponent,
		ValueSortingComponent,
		ValueDetailComponent,
		ValuePresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class Gongd3Module { }
