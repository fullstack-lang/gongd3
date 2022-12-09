import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { KeyDB } from '../key-db'
import { KeyService } from '../key.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface keyDummyElement {
}

const ELEMENT_DATA: keyDummyElement[] = [
];

@Component({
	selector: 'app-key-presentation',
	templateUrl: './key-presentation.component.html',
	styleUrls: ['./key-presentation.component.css'],
})
export class KeyPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	key: KeyDB = new (KeyDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private keyService: KeyService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getKey();

		// observable for changes in 
		this.keyService.KeyServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getKey()
				}
			}
		)
	}

	getKey(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.key = this.frontRepo.Keys.get(id)!

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
				gongd3_go_editor: ["gongd3_go-" + "key-detail", ID]
			}
		}]);
	}
}
