import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongd3 from 'gongd3'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  view = 'D3 view'
  d3 = 'D3 view'
  data = 'Data view'
  diagrams = 'Diagrams view'
  meta = 'Meta view'

  views: string[] = [this.d3, this.data, this.diagrams, this.meta];

  // variable that enables pooling of selected gongstruct
  obsTimer: Observable<number> = timer(1000, 1000)
  lastSelectionDate: string = ''

  constructor(private gongdocClassshapeService: gongdoc.ClassshapeService,
    private gongstructSelectionService: gongd3.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {

    // pool the gongdoc command and check wether a gongstruct has been selected
    this.obsTimer.subscribe(
      currTime => {
        // pool all classshapes and find which one is selected
        this.gongdocClassshapeService.getClassshapes().subscribe(
          classshapes => {
            for (let classshape of classshapes) {
              if (classshape.IsSelected) {
                classshape.IsSelected = false
                // console.log("classshape " + classshape.ReferenceName + " is selected")
                this.gongdocClassshapeService.updateClassshape(classshape).subscribe(
                  classshape2 => {
                    // console.log("classshape has been unselected")
                  }
                )
                this.gongstructSelectionService.gongstructSelected(classshape.ReferenceName)
              }
            }
          }
        )
      }
    )
  }
}
