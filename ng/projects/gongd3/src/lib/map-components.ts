// insertion point sub template for components imports 
  import { BarsTableComponent } from './bars-table/bars-table.component'
  import { BarSortingComponent } from './bar-sorting/bar-sorting.component'
  import { KeysTableComponent } from './keys-table/keys-table.component'
  import { KeySortingComponent } from './key-sorting/key-sorting.component'
  import { PiesTableComponent } from './pies-table/pies-table.component'
  import { PieSortingComponent } from './pie-sorting/pie-sorting.component'
  import { SeriesTableComponent } from './series-table/series-table.component'
  import { SerieSortingComponent } from './serie-sorting/serie-sorting.component'
  import { ValuesTableComponent } from './values-table/values-table.component'
  import { ValueSortingComponent } from './value-sorting/value-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfBarsComponents: Map<string, any> = new Map([["BarsTableComponent", BarsTableComponent],])
  export const MapOfBarSortingComponents: Map<string, any> = new Map([["BarSortingComponent", BarSortingComponent],])
  export const MapOfKeysComponents: Map<string, any> = new Map([["KeysTableComponent", KeysTableComponent],])
  export const MapOfKeySortingComponents: Map<string, any> = new Map([["KeySortingComponent", KeySortingComponent],])
  export const MapOfPiesComponents: Map<string, any> = new Map([["PiesTableComponent", PiesTableComponent],])
  export const MapOfPieSortingComponents: Map<string, any> = new Map([["PieSortingComponent", PieSortingComponent],])
  export const MapOfSeriesComponents: Map<string, any> = new Map([["SeriesTableComponent", SeriesTableComponent],])
  export const MapOfSerieSortingComponents: Map<string, any> = new Map([["SerieSortingComponent", SerieSortingComponent],])
  export const MapOfValuesComponents: Map<string, any> = new Map([["ValuesTableComponent", ValuesTableComponent],])
  export const MapOfValueSortingComponents: Map<string, any> = new Map([["ValueSortingComponent", ValueSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Bar", MapOfBarsComponents],
      ["Key", MapOfKeysComponents],
      ["Pie", MapOfPiesComponents],
      ["Serie", MapOfSeriesComponents],
      ["Value", MapOfValuesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Bar", MapOfBarSortingComponents],
      ["Key", MapOfKeySortingComponents],
      ["Pie", MapOfPieSortingComponents],
      ["Serie", MapOfSerieSortingComponents],
      ["Value", MapOfValueSortingComponents],
    ]
  )
