// insertion point sub template for components imports 
  import { CountrysTableComponent } from './countrys-table/countrys-table.component'
  import { CountrySortingComponent } from './country-sorting/country-sorting.component'
  import { HellosTableComponent } from './hellos-table/hellos-table.component'
  import { HelloSortingComponent } from './hello-sorting/hello-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfCountrysComponents: Map<string, any> = new Map([["CountrysTableComponent", CountrysTableComponent],])
  export const MapOfCountrySortingComponents: Map<string, any> = new Map([["CountrySortingComponent", CountrySortingComponent],])
  export const MapOfHellosComponents: Map<string, any> = new Map([["HellosTableComponent", HellosTableComponent],])
  export const MapOfHelloSortingComponents: Map<string, any> = new Map([["HelloSortingComponent", HelloSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Country", MapOfCountrysComponents],
      ["Hello", MapOfHellosComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Country", MapOfCountrySortingComponents],
      ["Hello", MapOfHelloSortingComponents],
    ]
  )
