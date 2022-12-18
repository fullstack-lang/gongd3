import { Component, ElementRef, OnInit } from '@angular/core';

import * as d3 from 'd3';
import * as d3_selection from 'd3-selection'
import * as topojson_client from 'topojson-client';
import * as topojson_specification from 'topojson-specification'
import * as geojson from 'geojson';

// see
// https://medium.com/geekculture/advanced-map-visualization-and-cartography-with-d3-js-geojson-topojson-2de786ece0c3

import { TopographyService } from '../topography-service'
import { registerLocaleData } from '@angular/common';

const us_map = 'assets/counties-albers-10m.json'
const fr_map = 'assets/fr-departments.json'

@Component({
  selector: 'lib-geojson',
  templateUrl: './geojson.component.html',
  styleUrls: ['./geojson.component.css']
})
export class GeojsonComponent implements OnInit {

  svg!: d3_selection.Selection<d3.BaseType, unknown, HTMLElement, any>
  projection!: d3.GeoIdentityTransform;
  topoFeatureStates!: geojson.FeatureCollection;
  geoPathGenerator!: d3.GeoPath;

  drawStateLines: boolean = true
  drawCountiesLines: boolean = false

  constructor(
    private topographyService: TopographyService,
    private el: ElementRef
  ) { }

  ngOnInit(

  ): void {
    this.initialMap();
  }

  initialMap(): void {
    let source = fr_map

    this.topographyService.getTopographyData(source).subscribe((topography: topojson_specification.Topology) => {
      this.draw(source, topography);
    });
  }

  draw(source: string, topography: topojson_specification.Topology): void {
    const { width, height } = this.getMapContainerWidthAndHeight();


    if (source == us_map) {
      this.topoFeatureStates = topojson_client.feature(
        topography,
        topography.objects['states']
      ) as geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties>
    }
    if (source == fr_map) {
      this.topoFeatureStates = topojson_client.feature(
        topography,
        topography.objects['FRA_adm2']
      ) as geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties>
    }


    this.projection = d3
      .geoIdentity()
      .fitSize([width, height], this.topoFeatureStates);

    this.geoPathGenerator = d3.geoPath(this.projection);

    // render svg
    this.svg = d3
      .select('svg')
      .attr('width', width + 50)
      .attr('height', height);

    if (source == us_map) {
      this.renderNationFeaturesWithShadow(topography);
      // this.renderCountiesFeatures(topography);
      // this.renderStateFeaures(topography);
      // this.renderStatesFeatures2(topography);
    }
    if (source == fr_map) {
      this.renderFrance(topography, width, height)
    }

    // resize event
    d3.select(window).on('resize', this.resizeMap);
  }

  renderFrance(topography: topojson_specification.Topology, width: number, height: number): void {
    const defs = this.svg.select('defs');

    var projection = d3.geoConicConformal() // Lambert-93
      .center([2.716666667, 45.983333333]) // Center on France carte Lambert93_0
      .parallels([42.38333333, 49.58333333])
      .rotate([0, 0]) // [0,0]
      .center([2, 47]) //[0,47]
      .scale(3000) //original
      .translate([width / 2, height / 2]);

    let featureCollection = topojson_client.feature(topography, topography.objects['FRA_adm2'])

    // Sets the current projection to the specified projection
    this.geoPathGenerator = d3.geoPath().projection(projection)

    defs
      .append('path')
      .datum(featureCollection)
      .attr('id', 'FRA_adm2')
      .attr('d', this.geoPathGenerator);

    this.svg
      .append('use')
      .attr('xlink:href', '#FRA_adm2')
      .attr('fill-opacity', 0.2)
      .attr("fill", "#e6f7ff")        //color of country: gris, #b8b8b8 bleu, #ccffff #80d4ff
      .style("stroke", "#80d4ff")            //original, border color #4da6ff)
      .style("stroke-width", 0.5) //border size
  }

  renderNationFeaturesWithShadow(topography: topojson_specification.Topology): void {
    const defs = this.svg.select('defs');
    defs
      .append('path')
      .datum(topojson_client.feature(topography, topography.objects['nation']))
      .attr('id', 'nation')
      .attr('d', this.geoPathGenerator);

    this.svg
      .append('use')
      .attr('xlink:href', '#nation')
      .attr('fill-opacity', 0.2)
    // .attr('filter', 'url(#blur)'); // bring shadows

    // extra touch (counties in grid)
    if (this.drawStateLines) {
      let mesh = topojson_client.mesh(
        topography,
        topography.objects['states'] as any,
        this.compareIds
      )
      let path = this.geoPathGenerator(mesh)

      this.svg
        .append('path')
        .attr('fill', 'none')
        .attr('stroke', '#777')
        .attr('stroke-width', 0.35)
        .attr(
          'd',
          path
        );
    }

    // extra touch (counties in grid)
    if (this.drawCountiesLines) {
      this.svg
        .append('path')
        .attr('fill', 'none')
        .attr('stroke', '#777')
        .attr('stroke-width', 0.35)
        .attr(
          'd',
          this.geoPathGenerator(
            topojson_client.mesh(
              topography,
              topography.objects['counties'] as any,
              this.compareIds
            )
          )
        );
    }
    // end extra touch
  }

  // This function takes in two arguments, a and b, which are of type any. 
  // The function uses a bitwise operator (the | symbol) to divide the id property of 
  // both a and b by 1000, and then uses the bitwise OR operator to return the result
  // of this operation as an integer.
  //
  // Finally, the function uses the === operator to check
  // if the result for a is equal to the result for b, and returns this value.
  compareIds(a: topojson_specification.GeometryObjectA, b: topojson_specification.GeometryObjectA): boolean {
    let a_id = (a.id! as number) / 1000
    let b_id = b.id! as number

    let a_id_red = a_id | 0
    let b_id_red = (b_id / 1000) | 0

    if (a_id_red != 0) {
      console.log("a_id != 0 :" + a_id_red)
    }

    // tslint:disable-next-line:no-bitwise
    return a_id_red === b_id_red
  }

  renderCountiesFeatures(topography: topojson_specification.Topology): void {

    let mapFeatures: geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties> =
      topojson_client.feature(topography, topography.objects['counties']) as
      geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties>

    this.svg
      .append('g')
      .attr('class', 'county')
      .attr('fill', '#fff')
      .selectAll('path')
      .data(
        mapFeatures.features
      )
      .join('path')
      .attr('id', (d: any) => {
        return d.id;
      })
      .attr('d', this.geoPathGenerator);
  }

  renderStateFeaures(topography: topojson_specification.Topology): void {

    let mapFeatures: geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties> =
      topojson_client.feature(topography, topography.objects['states']) as
      geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties>

    console.log("Let's inspect the svg")

    let pathString = this.geoPathGenerator.toString()
    this.svg
      .append('g')
      .attr('class', 'state')
      .attr('fill', 'none')
      .attr('stroke', '#BDBDBD')
      .attr('stroke-width', '0.7')
      .selectAll('path.state')
      .data(
        mapFeatures as any
      )
      .join('path')
      .attr('id', (d: any) => {
        return d.id;
      })
      .attr('d', this.geoPathGenerator.toString());

    console.log("Let's inspect the svg")
  }

  resizeMap = () => {
    const { width, height } = this.getMapContainerWidthAndHeight();

    this.svg.attr('width', width + 50).attr('height', height);

    // update projection
    this.projection.fitSize([width, height], this.topoFeatureStates);

    // resize the map
    this.svg.selectAll('path').attr('d', this.geoPathGenerator.toString());
  };

  getMapContainerWidthAndHeight = (): { width: number; height: number } => {
    const mapContainerEl = this.el.nativeElement.querySelector(
      '#map'
    ) as HTMLDivElement;
    const width = mapContainerEl.clientWidth - 50;
    const height = (width / 960) * 600;

    return { width, height };
  };
}
