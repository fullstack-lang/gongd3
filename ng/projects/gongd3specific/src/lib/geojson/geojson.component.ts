import { Component, ElementRef, OnInit } from '@angular/core';

import * as d3 from 'd3';
import * as d3_selection from 'd3-selection'
import * as topojson_client from 'topojson-client';
import * as topojson_specification from 'topojson-specification'
import * as geojson from 'geojson';

// see
// https://medium.com/geekculture/advanced-map-visualization-and-cartography-with-d3-js-geojson-topojson-2de786ece0c3

import { TopographyService } from '../topography-service'

@Component({
  selector: 'lib-geojson',
  templateUrl: './geojson.component.html',
  styleUrls: ['./geojson.component.css']
})
export class GeojsonComponent implements OnInit {

  svg!: d3_selection.Selection<d3.BaseType, unknown, HTMLElement, any>
  projection!: d3.GeoIdentityTransform;
  topoFeatureStates!: geojson.FeatureCollection;
  path!: d3.GeoPath;

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
    this.topographyService.getTopographyData().subscribe((topography: topojson_specification.Topology) => {
      this.draw(topography);
    });
  }

  draw(topography: topojson_specification.Topology): void {
    const { width, height } = this.getMapContainerWidthAndHeight();

    this.topoFeatureStates = topojson_client.feature(
      topography,
      topography.objects['states']
    ) as geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties>
    this.projection = d3
      .geoIdentity()
      .fitSize([width, height], this.topoFeatureStates);

    this.path = d3.geoPath(this.projection);

    // render svg
    this.svg = d3
      .select('svg')
      .attr('width', width + 50)
      .attr('height', height);

    this.renderNationFeaturesWithShadow(topography);
    // this.renderCountiesFeatures(topography);
    // this.renderStateFeaures(topography);
    // this.renderStatesFeatures2(topography);

    // resize event
    d3.select(window).on('resize', this.resizeMap);
  }

  renderNationFeaturesWithShadow(topography: topojson_specification.Topology): void {
    const defs = this.svg.select('defs');
    defs
      .append('path')
      .datum(topojson_client.feature(topography, topography.objects['nation']))
      .attr('id', 'nation')
      .attr('d', this.path);

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
        (a: any, b: any) => {
          // tslint:disable-next-line:no-bitwise
          return ((a.id / 1000) | 0) === ((b.id / 1000) | 0);
        }
      )
      let path = this.path(mesh)

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
          this.path(
            topojson_client.mesh(
              topography,
              topography.objects['counties'] as any,
              (a: any, b: any) => {
                // tslint:disable-next-line:no-bitwise
                return ((a.id / 1000) | 0) === ((b.id / 1000) | 0);
              }
            )
          )
        );
    }
    // end extra touch
  }

  // This function takes in two arguments, a and b, which are of type any. 
  // The function uses a bitwise operator (the | symbol) to divide the id property of 
  // both a and b by 1000, and then uses the bitwise OR operator to return the result
  // of this operation as an integer. Finally, the function uses the === operator to check
  // if the result for a is equal to the result for b, and returns this value.
  toDefine(a: any, b: any): boolean {
    // tslint:disable-next-line:no-bitwise
    return ((a.id / 1000) | 0) === ((b.id / 1000) | 0);
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
      .attr('d', this.path);
  }

  renderStateFeaures(topography: topojson_specification.Topology): void {

    let mapFeatures: geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties> =
      topojson_client.feature(topography, topography.objects['states']) as
      geojson.FeatureCollection<geojson.Geometry, geojson.GeoJsonProperties>

    console.log("Let's inspect the svg")

    let pathString = this.path.toString()
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
      .attr('d', this.path.toString());

    console.log("Let's inspect the svg")
  }

  resizeMap = () => {
    const { width, height } = this.getMapContainerWidthAndHeight();

    this.svg.attr('width', width + 50).attr('height', height);

    // update projection
    this.projection.fitSize([width, height], this.topoFeatureStates);

    // resize the map
    this.svg.selectAll('path').attr('d', this.path.toString());
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
