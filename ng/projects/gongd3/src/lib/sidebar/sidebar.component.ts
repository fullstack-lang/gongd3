import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { BarService } from '../bar.service'
import { getBarUniqueID } from '../front-repo.service'
import { KeyService } from '../key.service'
import { getKeyUniqueID } from '../front-repo.service'
import { PieService } from '../pie.service'
import { getPieUniqueID } from '../front-repo.service'
import { SerieService } from '../serie.service'
import { getSerieUniqueID } from '../front-repo.service'
import { ValueService } from '../value.service'
import { getValueUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongd3-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbFromBackService: CommitNbFromBackService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private barService: BarService,
    private keyService: KeyService,
    private pieService: PieService,
    private serieService: SerieService,
    private valueService: ValueService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.barService.BarServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.keyService.KeyServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.pieService.PieServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.serieService.SerieServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.valueService.ValueServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Bar part of the mat tree
      */
      let barGongNodeStruct: GongNode = {
        name: "Bar",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Bar",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(barGongNodeStruct)

      this.frontRepo.Bars_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Bars_array.forEach(
        barDB => {
          let barGongNodeInstance: GongNode = {
            name: barDB.Name,
            type: GongNodeType.INSTANCE,
            id: barDB.ID,
            uniqueIdPerStack: getBarUniqueID(barDB.ID),
            structName: "Bar",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          barGongNodeStruct.children!.push(barGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association X
          */
          let XGongNodeAssociation: GongNode = {
            name: "(Key) X",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: barDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Bar",
            associationField: "X",
            associatedStructName: "Key",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          barGongNodeInstance.children!.push(XGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation X
            */
          if (barDB.X != undefined) {
            let barGongNodeInstance_X: GongNode = {
              name: barDB.X.Name,
              type: GongNodeType.INSTANCE,
              id: barDB.X.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getBarUniqueID(barDB.ID)
                + 5 * getKeyUniqueID(barDB.X.ID),
              structName: "Key",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            XGongNodeAssociation.children.push(barGongNodeInstance_X)
          }

          /**
          * let append a node for the association Y
          */
          let YGongNodeAssociation: GongNode = {
            name: "(Key) Y",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: barDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Bar",
            associationField: "Y",
            associatedStructName: "Key",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          barGongNodeInstance.children!.push(YGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Y
            */
          if (barDB.Y != undefined) {
            let barGongNodeInstance_Y: GongNode = {
              name: barDB.Y.Name,
              type: GongNodeType.INSTANCE,
              id: barDB.Y.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getBarUniqueID(barDB.ID)
                + 5 * getKeyUniqueID(barDB.Y.ID),
              structName: "Key",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            YGongNodeAssociation.children.push(barGongNodeInstance_Y)
          }

          /**
          * let append a node for the slide of pointer Set
          */
          let SetGongNodeAssociation: GongNode = {
            name: "(Serie) Set",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: barDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Bar",
            associationField: "Set",
            associatedStructName: "Serie",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          barGongNodeInstance.children.push(SetGongNodeAssociation)

          barDB.Set?.forEach(serieDB => {
            let serieNode: GongNode = {
              name: serieDB.Name,
              type: GongNodeType.INSTANCE,
              id: serieDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getBarUniqueID(barDB.ID)
                + 11 * getSerieUniqueID(serieDB.ID),
              structName: "Serie",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            SetGongNodeAssociation.children.push(serieNode)
          })

        }
      )

      /**
      * fill up the Key part of the mat tree
      */
      let keyGongNodeStruct: GongNode = {
        name: "Key",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Key",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(keyGongNodeStruct)

      this.frontRepo.Keys_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Keys_array.forEach(
        keyDB => {
          let keyGongNodeInstance: GongNode = {
            name: keyDB.Name,
            type: GongNodeType.INSTANCE,
            id: keyDB.ID,
            uniqueIdPerStack: getKeyUniqueID(keyDB.ID),
            structName: "Key",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          keyGongNodeStruct.children!.push(keyGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Pie part of the mat tree
      */
      let pieGongNodeStruct: GongNode = {
        name: "Pie",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Pie",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(pieGongNodeStruct)

      this.frontRepo.Pies_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Pies_array.forEach(
        pieDB => {
          let pieGongNodeInstance: GongNode = {
            name: pieDB.Name,
            type: GongNodeType.INSTANCE,
            id: pieDB.ID,
            uniqueIdPerStack: getPieUniqueID(pieDB.ID),
            structName: "Pie",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          pieGongNodeStruct.children!.push(pieGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association X
          */
          let XGongNodeAssociation: GongNode = {
            name: "(Key) X",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: pieDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Pie",
            associationField: "X",
            associatedStructName: "Key",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          pieGongNodeInstance.children!.push(XGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation X
            */
          if (pieDB.X != undefined) {
            let pieGongNodeInstance_X: GongNode = {
              name: pieDB.X.Name,
              type: GongNodeType.INSTANCE,
              id: pieDB.X.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getPieUniqueID(pieDB.ID)
                + 5 * getKeyUniqueID(pieDB.X.ID),
              structName: "Key",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            XGongNodeAssociation.children.push(pieGongNodeInstance_X)
          }

          /**
          * let append a node for the association Y
          */
          let YGongNodeAssociation: GongNode = {
            name: "(Key) Y",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: pieDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Pie",
            associationField: "Y",
            associatedStructName: "Key",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          pieGongNodeInstance.children!.push(YGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Y
            */
          if (pieDB.Y != undefined) {
            let pieGongNodeInstance_Y: GongNode = {
              name: pieDB.Y.Name,
              type: GongNodeType.INSTANCE,
              id: pieDB.Y.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getPieUniqueID(pieDB.ID)
                + 5 * getKeyUniqueID(pieDB.Y.ID),
              structName: "Key",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            YGongNodeAssociation.children.push(pieGongNodeInstance_Y)
          }

          /**
          * let append a node for the slide of pointer Set
          */
          let SetGongNodeAssociation: GongNode = {
            name: "(Serie) Set",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: pieDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Pie",
            associationField: "Set",
            associatedStructName: "Serie",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          pieGongNodeInstance.children.push(SetGongNodeAssociation)

          pieDB.Set?.forEach(serieDB => {
            let serieNode: GongNode = {
              name: serieDB.Name,
              type: GongNodeType.INSTANCE,
              id: serieDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getPieUniqueID(pieDB.ID)
                + 11 * getSerieUniqueID(serieDB.ID),
              structName: "Serie",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            SetGongNodeAssociation.children.push(serieNode)
          })

        }
      )

      /**
      * fill up the Serie part of the mat tree
      */
      let serieGongNodeStruct: GongNode = {
        name: "Serie",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Serie",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(serieGongNodeStruct)

      this.frontRepo.Series_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Series_array.forEach(
        serieDB => {
          let serieGongNodeInstance: GongNode = {
            name: serieDB.Name,
            type: GongNodeType.INSTANCE,
            id: serieDB.ID,
            uniqueIdPerStack: getSerieUniqueID(serieDB.ID),
            structName: "Serie",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          serieGongNodeStruct.children!.push(serieGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Key
          */
          let KeyGongNodeAssociation: GongNode = {
            name: "(Key) Key",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: serieDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Serie",
            associationField: "Key",
            associatedStructName: "Key",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          serieGongNodeInstance.children!.push(KeyGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Key
            */
          if (serieDB.Key != undefined) {
            let serieGongNodeInstance_Key: GongNode = {
              name: serieDB.Key.Name,
              type: GongNodeType.INSTANCE,
              id: serieDB.Key.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getSerieUniqueID(serieDB.ID)
                + 5 * getKeyUniqueID(serieDB.Key.ID),
              structName: "Key",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            KeyGongNodeAssociation.children.push(serieGongNodeInstance_Key)
          }

          /**
          * let append a node for the slide of pointer Values
          */
          let ValuesGongNodeAssociation: GongNode = {
            name: "(Value) Values",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: serieDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Serie",
            associationField: "Values",
            associatedStructName: "Value",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          serieGongNodeInstance.children.push(ValuesGongNodeAssociation)

          serieDB.Values?.forEach(valueDB => {
            let valueNode: GongNode = {
              name: valueDB.Name,
              type: GongNodeType.INSTANCE,
              id: valueDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getSerieUniqueID(serieDB.ID)
                + 11 * getValueUniqueID(valueDB.ID),
              structName: "Value",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ValuesGongNodeAssociation.children.push(valueNode)
          })

        }
      )

      /**
      * fill up the Value part of the mat tree
      */
      let valueGongNodeStruct: GongNode = {
        name: "Value",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Value",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(valueGongNodeStruct)

      this.frontRepo.Values_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Values_array.forEach(
        valueDB => {
          let valueGongNodeInstance: GongNode = {
            name: valueDB.Name,
            type: GongNodeType.INSTANCE,
            id: valueDB.ID,
            uniqueIdPerStack: getValueUniqueID(valueDB.ID),
            structName: "Value",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          valueGongNodeStruct.children!.push(valueGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    });

    // fetch the number of commits
    this.commitNbFromBackService.getCommitNbFromBack().subscribe(
      commitNbFromBack => {
        this.commitNbFromBack = commitNbFromBack
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        gongd3_go_table: ["gongd3_go-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          gongd3_go_table: ["gongd3_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          gongd3_go_presentation: ["gongd3_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        gongd3_go_editor: ["gongd3_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        gongd3_go_editor: ["gongd3_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}
