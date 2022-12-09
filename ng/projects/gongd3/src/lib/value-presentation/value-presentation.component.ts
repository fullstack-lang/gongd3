import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ValueDB } from '../value-db'
import { ValueService } from '../value.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface valueDummyElement {
}

const ELEMENT_DATA: valueDummyElement[] = [
];

@Component({
	selector: 'app-value-presentation',
	templateUrl: './value-presentation.component.html',
	styleUrls: ['./value-presentation.component.css'],
})
export class ValuePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	value: ValueDB = new (ValueDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private valueService: ValueService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getValue();

		// observable for changes in 
		this.valueService.ValueServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getValue()
				}
			}
		)
	}

	getValue(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.value = this.frontRepo.Values.get(id)!

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
				gongd3_go_editor: ["gongd3_go-" + "value-detail", ID]
			}
		}]);
	}
}
