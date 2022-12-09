import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { BarDB } from '../bar-db'
import { BarService } from '../bar.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface barDummyElement {
}

const ELEMENT_DATA: barDummyElement[] = [
];

@Component({
	selector: 'app-bar-presentation',
	templateUrl: './bar-presentation.component.html',
	styleUrls: ['./bar-presentation.component.css'],
})
export class BarPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	bar: BarDB = new (BarDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private barService: BarService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getBar();

		// observable for changes in 
		this.barService.BarServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getBar()
				}
			}
		)
	}

	getBar(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.bar = this.frontRepo.Bars.get(id)!

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
				gongd3_go_editor: ["gongd3_go-" + "bar-detail", ID]
			}
		}]);
	}
}
