// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { ValueDB } from '../value-db'
import { ValueService } from '../value.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-valuestable',
  templateUrl: './values-table.component.html',
  styleUrls: ['./values-table.component.css'],
})
export class ValuesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Value instances
  selection: SelectionModel<ValueDB> = new (SelectionModel)
  initialSelection = new Array<ValueDB>()

  // the data source for the table
  values: ValueDB[] = []
  matTableDataSource: MatTableDataSource<ValueDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.values
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (valueDB: ValueDB, property: string) => {
      switch (property) {
        case 'ID':
          return valueDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return valueDB.Name;

        case 'Serie_Values':
          if (this.frontRepo.Series.get(valueDB.Serie_ValuesDBID.Int64) != undefined) {
            return this.frontRepo.Series.get(valueDB.Serie_ValuesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (valueDB: ValueDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the valueDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += valueDB.Name.toLowerCase()
      if (valueDB.Serie_ValuesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Series.get(valueDB.Serie_ValuesDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private valueService: ValueService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of value instances
    public dialogRef: MatDialogRef<ValuesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.valueService.ValueServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getValues()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Serie_Values",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Serie_Values",
      ]
      this.selection = new SelectionModel<ValueDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getValues()
    this.matTableDataSource = new MatTableDataSource(this.values)
  }

  getValues(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.values = this.frontRepo.Values_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let value of this.values) {
            let ID = this.dialogData.ID
            let revPointer = value[this.dialogData.ReversePointer as keyof ValueDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(value)
            }
            this.selection = new SelectionModel<ValueDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ValueDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as ValueDB[]
          for (let associationInstance of sourceField) {
            let value = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ValueDB
            this.initialSelection.push(value)
          }

          this.selection = new SelectionModel<ValueDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.values
      }
    )
  }

  // newValue initiate a new value
  // create a new Value objet
  newValue() {
  }

  deleteValue(valueID: number, value: ValueDB) {
    // list of values is truncated of value before the delete
    this.values = this.values.filter(h => h !== value);

    this.valueService.deleteValue(valueID).subscribe(
      value => {
        this.valueService.ValueServiceChanged.next("delete")
      }
    );
  }

  editValue(valueID: number, value: ValueDB) {

  }

  // display value in router
  displayValueInRouter(valueID: number) {
    this.router.navigate(["gongd3_go-" + "value-display", valueID])
  }

  // set editor outlet
  setEditorRouterOutlet(valueID: number) {
    this.router.navigate([{
      outlets: {
        gongd3_go_editor: ["gongd3_go-" + "value-detail", valueID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(valueID: number) {
    this.router.navigate([{
      outlets: {
        gongd3_go_presentation: ["gongd3_go-" + "value-presentation", valueID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.values.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.values.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<ValueDB>()

      // reset all initial selection of value that belong to value
      for (let value of this.initialSelection) {
        let index = value[this.dialogData.ReversePointer as keyof ValueDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(value)

      }

      // from selection, set value that belong to value
      for (let value of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = value[this.dialogData.ReversePointer as keyof ValueDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(value)
      }


      // update all value (only update selection & initial selection)
      for (let value of toUpdate) {
        this.valueService.updateValue(value)
          .subscribe(value => {
            this.valueService.ValueServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ValueDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedValue = new Set<number>()
      for (let value of this.initialSelection) {
        if (this.selection.selected.includes(value)) {
          // console.log("value " + value.Name + " is still selected")
        } else {
          console.log("value " + value.Name + " has been unselected")
          unselectedValue.add(value.ID)
          console.log("is unselected " + unselectedValue.has(value.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let value = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ValueDB
      if (unselectedValue.has(value.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<ValueDB>) = new Array<ValueDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          value => {
            if (!this.initialSelection.includes(value)) {
              // console.log("value " + value.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + value.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = value.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = value.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("value " + value.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<ValueDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
