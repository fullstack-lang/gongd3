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
import { KeyDB } from '../key-db'
import { KeyService } from '../key.service'

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
  selector: 'app-keystable',
  templateUrl: './keys-table.component.html',
  styleUrls: ['./keys-table.component.css'],
})
export class KeysTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Key instances
  selection: SelectionModel<KeyDB> = new (SelectionModel)
  initialSelection = new Array<KeyDB>()

  // the data source for the table
  keys: KeyDB[] = []
  matTableDataSource: MatTableDataSource<KeyDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.keys
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
    this.matTableDataSource.sortingDataAccessor = (keyDB: KeyDB, property: string) => {
      switch (property) {
        case 'ID':
          return keyDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return keyDB.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (keyDB: KeyDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the keyDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += keyDB.Name.toLowerCase()

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
    private keyService: KeyService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of key instances
    public dialogRef: MatDialogRef<KeysTableComponent>,
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
    this.keyService.KeyServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getKeys()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
      ]
      this.selection = new SelectionModel<KeyDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getKeys()
    this.matTableDataSource = new MatTableDataSource(this.keys)
  }

  getKeys(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.keys = this.frontRepo.Keys_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let key of this.keys) {
            let ID = this.dialogData.ID
            let revPointer = key[this.dialogData.ReversePointer as keyof KeyDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(key)
            }
            this.selection = new SelectionModel<KeyDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, KeyDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as KeyDB[]
          for (let associationInstance of sourceField) {
            let key = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as KeyDB
            this.initialSelection.push(key)
          }

          this.selection = new SelectionModel<KeyDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.keys
      }
    )
  }

  // newKey initiate a new key
  // create a new Key objet
  newKey() {
  }

  deleteKey(keyID: number, key: KeyDB) {
    // list of keys is truncated of key before the delete
    this.keys = this.keys.filter(h => h !== key);

    this.keyService.deleteKey(keyID).subscribe(
      key => {
        this.keyService.KeyServiceChanged.next("delete")
      }
    );
  }

  editKey(keyID: number, key: KeyDB) {

  }

  // display key in router
  displayKeyInRouter(keyID: number) {
    this.router.navigate(["gongd3_go-" + "key-display", keyID])
  }

  // set editor outlet
  setEditorRouterOutlet(keyID: number) {
    this.router.navigate([{
      outlets: {
        gongd3_go_editor: ["gongd3_go-" + "key-detail", keyID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(keyID: number) {
    this.router.navigate([{
      outlets: {
        gongd3_go_presentation: ["gongd3_go-" + "key-presentation", keyID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.keys.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.keys.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<KeyDB>()

      // reset all initial selection of key that belong to key
      for (let key of this.initialSelection) {
        let index = key[this.dialogData.ReversePointer as keyof KeyDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(key)

      }

      // from selection, set key that belong to key
      for (let key of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = key[this.dialogData.ReversePointer as keyof KeyDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(key)
      }


      // update all key (only update selection & initial selection)
      for (let key of toUpdate) {
        this.keyService.updateKey(key)
          .subscribe(key => {
            this.keyService.KeyServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, KeyDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedKey = new Set<number>()
      for (let key of this.initialSelection) {
        if (this.selection.selected.includes(key)) {
          // console.log("key " + key.Name + " is still selected")
        } else {
          console.log("key " + key.Name + " has been unselected")
          unselectedKey.add(key.ID)
          console.log("is unselected " + unselectedKey.has(key.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let key = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as KeyDB
      if (unselectedKey.has(key.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<KeyDB>) = new Array<KeyDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          key => {
            if (!this.initialSelection.includes(key)) {
              // console.log("key " + key.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + key.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = key.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = key.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("key " + key.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<KeyDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
