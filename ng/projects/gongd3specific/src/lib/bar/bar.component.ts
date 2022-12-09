import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import * as d3 from 'd3';
import { Observable, Subscription, timer } from 'rxjs';

import * as gongd3 from 'gongd3'
import { keyBy } from 'lodash';
import { index } from 'd3';

interface LooseObject {
  [key: string]: any
}

@Component({
  selector: 'lib-bar',
  templateUrl: './bar.component.html',
  styleUrls: ['./bar.component.css']
})
export class BarComponent implements OnInit {

  checkGongd3CommitNbFromBackTimer: Observable<number> = timer(500, 500);
  checkGongd3CommitNbFromBackTimerSubscription: Subscription = new Subscription

  lastCommitNbFromBack = -1
  lastDiagramId = 0
  currTime: number = 0

  private data = [
    { "Framework": "Vue", "Stars": "166443", "Released": "2014" },
    { "Framework": "React", "Stars": "150793", "Released": "2013" },
    { "Framework": "Angular", "Stars": "62342", "Released": "2016" },
    { "Framework": "Backbone", "Stars": "27647", "Released": "2010" },
    { "Framework": "Ember", "Stars": "21471", "Released": "2011" },
  ];
  private svg: any;
  private margin = 50;
  private width = 750 - (this.margin * 2);
  private height = 400 - (this.margin * 2);

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private gongd3CommitNbFromBackService: gongd3.CommitNbFromBackService,
    private gongd3FrontRepoService: gongd3.FrontRepoService,
  ) { }

  ngOnDestroy() {
    // console.log("on destroy")
    this.checkGongd3CommitNbFromBackTimerSubscription.unsubscribe()
  }
  ngOnInit(): void {

    // check loop for refresh from the back repo
    this.checkGongd3CommitNbFromBackTimerSubscription = this.checkGongd3CommitNbFromBackTimer.subscribe(
      currTime => {
        this.currTime = currTime

        this.gongd3CommitNbFromBackService.getCommitNbFromBack().subscribe(
          commitNbFromBack => {

            // const id = +this.route.snapshot.paramMap.get('id')!;

            // console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
            // console.log("last diagram id " + this.lastDiagramId + " new: " + id)
            // console.log("last drawn diagram id " + this.idOfDrawnClassDiagram + " new: " + id)

            // condition for refresh
            if (this.lastCommitNbFromBack < commitNbFromBack) {

              console.log("gongd3: last commit nb " + this.lastCommitNbFromBack + " new: " + commitNbFromBack)
              this.lastCommitNbFromBack = commitNbFromBack
              // this.lastDiagramId = id
              this.redraw()
            }
          }
        )
      }
    )
  }


  private redraw(): void {
    this.gongd3FrontRepoService.pull().subscribe(
      frontRepo => {
        for (let bar of frontRepo.Bars_array) {
          console.log("Bar name " + bar.Name)
          if (bar.Name == "Stars per Framework") {
            console.log("Selected Bar name " + bar.Name)
            this.data = []
            let indexSerie = 0
            for (let serie of bar.Set!) {
              let indexValue = 0

              // sort serie.Values according to the index
              serie.Values!.sort((t1, t2) => {
                if (t1.Serie_ValuesDBID_Index.Int64 > t2.Serie_ValuesDBID_Index.Int64) {
                  return 1;
                }
                if (t1.Serie_ValuesDBID_Index.Int64 < t2.Serie_ValuesDBID_Index.Int64) {
                  return -1;
                }
                return 0;
              });

              for (let value of serie.Values!) {
                var obj: any
                if (indexSerie == 0) {
                  obj = {}
                  obj[serie.Key!.Name] = value.Name
                  this.data.push(obj)
                } else {
                  obj = this.data[indexValue]
                  obj[serie.Key!.Name] = value.Name
                }
                indexValue++
              }
              indexSerie = indexSerie + 1
            }

            this.createSvg();
            this.drawBars(this.data);
          }
        }
      }
    )
  }


  private createSvg(): void {
    this.svg = d3.select("figure#bar")
      .append("svg")
      .attr("width", this.width + (this.margin * 2))
      .attr("height", this.height + (this.margin * 2))
      .append("g")
      .attr("transform", "translate(" + this.margin + "," + this.margin + ")");
  }

  private drawBars(data: any[]): void {
    // Create the X-axis band scale
    const x = d3.scaleBand()
      .range([0, this.width])
      .domain(data.map(d => d.Framework))
      .padding(0.2);

    // Draw the X-axis on the DOM
    this.svg.append("g")
      .attr("transform", "translate(0," + this.height + ")")
      .call(d3.axisBottom(x))
      .selectAll("text")
      .attr("transform", "translate(-10,0)rotate(-45)")
      .style("text-anchor", "end");

    // Create the Y-axis band scale
    const y = d3.scaleLinear()
      .domain([0, 200000])
      .range([this.height, 0]);

    // Draw the Y-axis on the DOM
    this.svg.append("g")
      .call(d3.axisLeft(y));

    // Create and fill the bars
    this.svg.selectAll("bars")
      .data(data)
      .enter()
      .append("rect")
      .attr("x", (d: any) => x(d.Framework))
      .attr("y", (d: any) => y(d.Stars))
      .attr("width", x.bandwidth())
      .attr("height", (d: any) => this.height - y(d.Stars))
      .attr("fill", "#d04a35");
  }

}
