import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { HelloDB } from '../hello-db'
import { HelloService } from '../hello.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface helloDummyElement {
}

const ELEMENT_DATA: helloDummyElement[] = [
];

@Component({
	selector: 'app-hello-presentation',
	templateUrl: './hello-presentation.component.html',
	styleUrls: ['./hello-presentation.component.css'],
})
export class HelloPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	hello: HelloDB = new (HelloDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private helloService: HelloService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getHello();

		// observable for changes in 
		this.helloService.HelloServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getHello()
				}
			}
		)
	}

	getHello(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.hello = this.frontRepo.Hellos.get(id)!

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
				gongd3_go_editor: ["gongd3_go-" + "hello-detail", ID]
			}
		}]);
	}
}
