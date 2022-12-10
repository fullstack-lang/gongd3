import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { PieDB } from '../pie-db'
import { PieService } from '../pie.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface pieDummyElement {
}

const ELEMENT_DATA: pieDummyElement[] = [
];

@Component({
	selector: 'app-pie-presentation',
	templateUrl: './pie-presentation.component.html',
	styleUrls: ['./pie-presentation.component.css'],
})
export class PiePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	pie: PieDB = new (PieDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private pieService: PieService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getPie();

		// observable for changes in 
		this.pieService.PieServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getPie()
				}
			}
		)
	}

	getPie(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.pie = this.frontRepo.Pies.get(id)!

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
				gongd3_go_editor: ["gongd3_go-" + "pie-detail", ID]
			}
		}]);
	}
}
