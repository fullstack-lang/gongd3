import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongd3 from 'gongd3'

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { GongtableModule } from 'gongtable'
import { GongtablespecificModule } from 'gongtablespecific'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {


  d3 = 'D3 view'
  probe = 'Probe view'

  view = this.d3

  views: string[] = [this.d3, this.probe];


  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "gongd3"
  StackType = "github.com/fullstack-lang/gongd3/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
