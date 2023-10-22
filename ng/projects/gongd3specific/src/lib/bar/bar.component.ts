import { Component, Input, OnInit } from '@angular/core';
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

  @Input() name: string = ""
  @Input() StackName: string = ""

  checkGongd3CommitNbFromBackTimer: Observable<number> = timer(500, 500);
  checkGongd3CommitNbFromBackTimerSubscription: Subscription = new Subscription

  lastCommitNbFromBack = -1
  lastDiagramId = 0
  currTime: number = 0

  private data = new (Array<any>)

  private svg: any;

  private bar = new (gongd3.BarDB)

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

        this.gongd3CommitNbFromBackService.getCommitNbFromBack(500, this.StackName).subscribe(
          commitNbFromBack => {
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
    this.gongd3FrontRepoService.pull(this.StackName).subscribe(
      frontRepo => {
        for (let bar of frontRepo.Bars_array) {
          console.log("Bar name " + bar.Name)
          if (bar.Name == this.name) {
            this.bar = bar
            console.log("Selected Bar name " + bar.Name)

            this.data = new (Array<any>)
            let indexSerie = 0
            for (let serie of bar.Set!) {
              let indexValue = 0

              for (let value of serie.Values!) {
                var obj: any
                // when parsing the first serie, creates the object
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
      .attr("width", this.bar.Width + (this.bar.Margin * 2))
      .attr("height", this.bar.Heigth + (this.bar.Margin * 2))
      .append("g")
      .attr("transform", "translate(" + this.bar.Margin + "," + this.bar.Margin + ")");
  }

  private drawBars(data: any[]): void {

    let xDomain: any[]
    if (this.bar.AutoDomainX) {
      xDomain = data.map(d => d[this.bar.X!.Name])
    } else {
      xDomain = [this.bar.XMin, this.bar.XMax]
    }

    let yDomain: any[]
    if (this.bar.AutoDomainY) {
      yDomain = data.map(d => parseInt(d[this.bar.Y!.Name]))
    } else {
      yDomain = [this.bar.YMin, this.bar.YMax]
    }

    // Create the X-axis band scale
    const x = d3.scaleBand()
      .range([0, this.bar.Width])
      .domain(data.map(d => d[this.bar.X!.Name]))
      .padding(0.2)

    // Draw the X-axis on the DOM
    this.svg.append("g")
      .attr("transform", "translate(0," + this.bar.Heigth + ")")
      .call(d3.axisBottom(x))
      .selectAll("text")
      .attr("transform", "translate(-10,0)rotate(-45)")
      .style("text-anchor", "end");

    // Create the Y-axis band scale
    const y = d3.scaleLinear()
      .domain(yDomain)
      .range([this.bar.Heigth, 0]);

    // Draw the Y-axis on the DOM
    this.svg.append("g")
      .call(d3.axisLeft(y));

    this.svg.append("text")
      .attr("transform", "rotate(-90)")
      .attr("x", -(this.bar.Heigth / 2))
      .attr("y", -50)
      .attr("text-anchor", "middle")
      .text("Y-axis label")

    // Create and fill the bars
    this.svg.selectAll("bars")
      .data(data)
      .enter()
      .append("rect")
      .attr("x", (d: any) => x(d[this.bar.X!.Name]))
      .attr("y", (d: any) => y(d[this.bar.Y!.Name]))
      .attr("width", x.bandwidth())
      .attr("height", (d: any) => this.bar.Heigth - y(d[this.bar.Y!.Name]))
      .attr("fill", "#d04a35")
  }
}
