import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import * as d3 from 'd3';
import { Observable, Subscription, timer } from 'rxjs';

import * as gongd3 from '../../../../gongd3/src/public-api'

@Component({
  selector: 'lib-scatter',
  templateUrl: './scatter.component.html',
  standalone: true,
  styleUrls: ['./scatter.component.css']
})
export class ScatterComponent implements OnInit {

  @Input() name: string = ""
  @Input() StackName: string = ""

  checkGongd3CommitNbFromBackTimer: Observable<number> = timer(500, 500);
  checkGongd3CommitNbFromBackTimerSubscription: Subscription = new Subscription

  lastCommitNbFromBack = -1
  lastDiagramId = 0
  currTime: number = 0

  private x_serieName: string = ""
  private y_serieName: string = ""
  private text_serieName: string = ""

  private data = new Array<any>()
  private svg: any
  private margin = 50;
  private width = 750 - (this.margin * 2);
  private height = 400 - (this.margin * 2);

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private gongd3CommitNbFromBackService: gongd3.CommitNbFromBackService,
    private gongd3FrontRepoService: gongd3.FrontRepoService,
  ) { }

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
        for (let scatter of frontRepo.getFrontArray<gongd3.Scatter>(gongd3.Scatter.GONGSTRUCT_NAME)) {
          console.log("Scatter name " + scatter.Name)
          if (scatter.Name == this.name) {
            console.log("Selected Scatter name " + scatter.Name)

            // set up setting
            this.width = scatter.Width
            this.height = scatter.Heigth
            this.margin = scatter.Margin

            this.x_serieName = scatter.X!.Name
            this.y_serieName = scatter.Y!.Name
            this.text_serieName = scatter.Text!.Name

            this.data = new (Array<any>)
            let indexSerie = 0
            for (let serie of scatter.Set!) {
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
            this.drawPlot();
          }
        }
      }
    )
  }
  private createSvg(): void {
    this.svg = d3.select("figure#scatter")
      .append("svg")
      .attr("width", this.width + (this.margin * 2))
      .attr("height", this.height + (this.margin * 2))
      .append("g")
      .attr("transform", "translate(" + this.margin + "," + this.margin + ")");
  }

  private drawPlot(): void {
    // Add X axis
    const x = d3.scaleLinear()
      .domain([2009, 2017])
      .range([0, this.width]);
    this.svg.append("g")
      .attr("transform", "translate(0," + this.height + ")")
      .call(d3.axisBottom(x).tickFormat(d3.format("d")));

    // Add Y axis
    const y = d3.scaleLinear()
      .domain([0, 200000])
      .range([this.height, 0]);
    this.svg.append("g")
      .call(d3.axisLeft(y));

    // Add dots
    const dots = this.svg.append('g');
    dots.selectAll("dot")
      .data(this.data)
      .enter()
      .append("circle")
      .attr("cx", (d: any) => x(d[this.x_serieName]))
      .attr("cy", (d: any) => y(d[this.y_serieName]))
      .attr("r", 7)
      .style("opacity", .5)
      .style("fill", "#69b3a2");
    // Add labels
    dots.selectAll("text")
      .data(this.data)
      .enter()
      .append("text")
      .text((d: any) => d[this.text_serieName])
      .attr("x", (d: any) => x(d[this.x_serieName]))
      .attr("y", (d: any) => y(d[this.y_serieName]))
  }
}
