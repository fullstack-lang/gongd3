import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ScatterDB } from '../scatter-db'
import { ScatterService } from '../scatter.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface scatterDummyElement {
}

const ELEMENT_DATA: scatterDummyElement[] = [
];

@Component({
	selector: 'app-scatter-presentation',
	templateUrl: './scatter-presentation.component.html',
	styleUrls: ['./scatter-presentation.component.css'],
})
export class ScatterPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	scatter: ScatterDB = new (ScatterDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private scatterService: ScatterService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getScatter();

		// observable for changes in 
		this.scatterService.ScatterServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getScatter()
				}
			}
		)
	}

	getScatter(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.scatter = this.frontRepo.Scatters.get(id)!

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
				gongd3_go_editor: ["gongd3_go-" + "scatter-detail", ID]
			}
		}]);
	}
}
