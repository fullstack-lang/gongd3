import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { SerieDB } from '../serie-db'
import { SerieService } from '../serie.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface serieDummyElement {
}

const ELEMENT_DATA: serieDummyElement[] = [
];

@Component({
	selector: 'app-serie-presentation',
	templateUrl: './serie-presentation.component.html',
	styleUrls: ['./serie-presentation.component.css'],
})
export class SeriePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	serie: SerieDB = new (SerieDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private serieService: SerieService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSerie();

		// observable for changes in 
		this.serieService.SerieServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSerie()
				}
			}
		)
	}

	getSerie(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.serie = this.frontRepo.Series.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				gongd3_go_presentation: ["gongd3_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				gongd3_go_editor: ["gongd3_go-" + "serie-detail", ID]
			}
		}]);
	}
}
