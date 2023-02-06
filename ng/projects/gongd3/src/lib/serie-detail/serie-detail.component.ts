// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { SerieDB } from '../serie-db'
import { SerieService } from '../serie.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { BarDB } from '../bar-db'
import { PieDB } from '../pie-db'
import { ScatterDB } from '../scatter-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// SerieDetailComponent is initilizaed from different routes
// SerieDetailComponentState detail different cases 
enum SerieDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Bar_Set_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Pie_Set_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Scatter_Set_SET,
}

@Component({
	selector: 'app-serie-detail',
	templateUrl: './serie-detail.component.html',
	styleUrls: ['./serie-detail.component.css'],
})
export class SerieDetailComponent implements OnInit {

	// insertion point for declarations

	// the SerieDB of interest
	serie: SerieDB = new SerieDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: SerieDetailComponentState = SerieDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private serieService: SerieService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = SerieDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = SerieDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Set":
						// console.log("Serie" + " is instanciated with back pointer to instance " + this.id + " Bar association Set")
						this.state = SerieDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Bar_Set_SET
						break;
					case "Set":
						// console.log("Serie" + " is instanciated with back pointer to instance " + this.id + " Pie association Set")
						this.state = SerieDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Pie_Set_SET
						break;
					case "Set":
						// console.log("Serie" + " is instanciated with back pointer to instance " + this.id + " Scatter association Set")
						this.state = SerieDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Scatter_Set_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getSerie()

		// observable for changes in structs
		this.serieService.SerieServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getSerie()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getSerie(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case SerieDetailComponentState.CREATE_INSTANCE:
						this.serie = new (SerieDB)
						break;
					case SerieDetailComponentState.UPDATE_INSTANCE:
						let serie = frontRepo.Series.get(this.id)
						console.assert(serie != undefined, "missing serie with id:" + this.id)
						this.serie = serie!
						break;
					// insertion point for init of association field
					case SerieDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Bar_Set_SET:
						this.serie = new (SerieDB)
						this.serie.Bar_Set_reverse = frontRepo.Bars.get(this.id)!
						break;
					case SerieDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Pie_Set_SET:
						this.serie = new (SerieDB)
						this.serie.Pie_Set_reverse = frontRepo.Pies.get(this.id)!
						break;
					case SerieDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Scatter_Set_SET:
						this.serie = new (SerieDB)
						this.serie.Scatter_Set_reverse = frontRepo.Scatters.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.serie.KeyID == undefined) {
			this.serie.KeyID = new NullInt64
		}
		if (this.serie.Key != undefined) {
			this.serie.KeyID.Int64 = this.serie.Key.ID
			this.serie.KeyID.Valid = true
		} else {
			this.serie.KeyID.Int64 = 0
			this.serie.KeyID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.serie.Bar_Set_reverse != undefined) {
			if (this.serie.Bar_SetDBID == undefined) {
				this.serie.Bar_SetDBID = new NullInt64
			}
			this.serie.Bar_SetDBID.Int64 = this.serie.Bar_Set_reverse.ID
			this.serie.Bar_SetDBID.Valid = true
			if (this.serie.Bar_SetDBID_Index == undefined) {
				this.serie.Bar_SetDBID_Index = new NullInt64
			}
			this.serie.Bar_SetDBID_Index.Valid = true
			this.serie.Bar_Set_reverse = new BarDB // very important, otherwise, circular JSON
		}
		if (this.serie.Pie_Set_reverse != undefined) {
			if (this.serie.Pie_SetDBID == undefined) {
				this.serie.Pie_SetDBID = new NullInt64
			}
			this.serie.Pie_SetDBID.Int64 = this.serie.Pie_Set_reverse.ID
			this.serie.Pie_SetDBID.Valid = true
			if (this.serie.Pie_SetDBID_Index == undefined) {
				this.serie.Pie_SetDBID_Index = new NullInt64
			}
			this.serie.Pie_SetDBID_Index.Valid = true
			this.serie.Pie_Set_reverse = new PieDB // very important, otherwise, circular JSON
		}
		if (this.serie.Scatter_Set_reverse != undefined) {
			if (this.serie.Scatter_SetDBID == undefined) {
				this.serie.Scatter_SetDBID = new NullInt64
			}
			this.serie.Scatter_SetDBID.Int64 = this.serie.Scatter_Set_reverse.ID
			this.serie.Scatter_SetDBID.Valid = true
			if (this.serie.Scatter_SetDBID_Index == undefined) {
				this.serie.Scatter_SetDBID_Index = new NullInt64
			}
			this.serie.Scatter_SetDBID_Index.Valid = true
			this.serie.Scatter_Set_reverse = new ScatterDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case SerieDetailComponentState.UPDATE_INSTANCE:
				this.serieService.updateSerie(this.serie)
					.subscribe(serie => {
						this.serieService.SerieServiceChanged.next("update")
					});
				break;
			default:
				this.serieService.postSerie(this.serie).subscribe(serie => {
					this.serieService.SerieServiceChanged.next("post")
					this.serie = new (SerieDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.serie.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.serie.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Serie"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.serie.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.serie.Name == "") {
			this.serie.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
