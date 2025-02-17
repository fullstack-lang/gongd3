import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import * as d3 from 'd3';
import { Observable, Subscription, timer } from 'rxjs';

import * as gongd3 from '../../../../gongd3/src/public-api'


@Component({
  selector: 'lib-pie',
  templateUrl: './pie.component.html',
  standalone: true,
  styleUrls: ['./pie.component.css']
})
export class PieComponent implements OnInit {

  @Input() name: string = ""
  @Input() StackName: string = ""

  checkGongd3CommitNbFromBackTimer: Observable<number> = timer(500, 500);
  checkGongd3CommitNbFromBackTimerSubscription: Subscription = new Subscription

  lastCommitNbFromBack = -1
  lastDiagramId = 0
  currTime: number = 0

  private x_serieName: string = ""
  private y_serieName: string = ""

  private data = new Array<any>()
  private svg: any;
  private margin = 50;
  private width = 750;
  private height = 600;
  // The radius of the pie chart is half the smallest side
  private radius = Math.min(this.width, this.height) / 2 - this.margin;
  private colors: any;

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
        for (let pie of frontRepo.getFrontArray<gongd3.Pie>(gongd3.Bar.GONGSTRUCT_NAME)) {
          console.log("Pie name " + pie.Name)
          if (pie.Name == this.name) {
            console.log("Selected Pie name " + pie.Name)

            // set up setting
            this.width = pie.Width
            this.height = pie.Heigth
            this.margin = pie.Margin

            this.x_serieName = pie.X!.Name
            this.y_serieName = pie.Y!.Name

            this.data = new (Array<any>)
            let indexSerie = 0
            for (let serie of pie.Set!) {
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
            this.createColors();
            this.drawChart();
          }
        }
      }
    )
  }


  private createSvg(): void {
    this.svg = d3.select("figure#pie")
      .append("svg")
      .attr("width", this.width)
      .attr("height", this.height)
      .append("g")
      .attr(
        "transform",
        "translate(" + this.width / 2 + "," + this.height / 2 + ")"
      );
  }
  private createColors(): void {
    this.colors = d3.scaleOrdinal()
      .domain(this.data.map(

        d => d[this.y_serieName].toString()

      ))
      .range(["#c7d3ec", "#a5b8db", "#879cc4", "#677795", "#5a6782"]);
  }

  private drawChart(): void {
    // Compute the position of each group on the pie:
    const pie = d3.pie<any>().value((d: any) => Number(d[this.y_serieName]));

    // Build the pie chart
    let pieD3 = pie(this.data)
    this.svg
      .selectAll('pieces')
      .data(pieD3)
      .enter()
      .append('path')
      .attr('d', d3.arc()
        .innerRadius(0)
        .outerRadius(this.radius)
      )
      .attr('fill', (d: any, i: any) => (this.colors(i)))
      .attr("stroke", "#121926")
      .style("stroke-width", "1px");

    // Add labels
    const labelLocation = d3.arc()
      .innerRadius(100)
      .outerRadius(this.radius);

    this.svg
      .selectAll('pieces')
      .data(pieD3)
      .enter()
      .append('text')
      .text((d: any) => d.data[this.x_serieName])
      .attr("transform", (d: any) => "translate(" + labelLocation.centroid(d) + ")")
      .style("text-anchor", "middle")
      .style("font-size", 15);
  }

}
