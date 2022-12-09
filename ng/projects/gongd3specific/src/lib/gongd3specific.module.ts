import { NgModule } from '@angular/core';
import { Gongd3specificComponent } from './gongd3specific.component';
import { BarComponent } from './bar/bar.component';
import { PieComponent } from './pie/pie.component';
import { ScatterComponent } from './scatter/scatter.component';



@NgModule({
  declarations: [
    Gongd3specificComponent,
    BarComponent,
    PieComponent,
    ScatterComponent
  ],
  imports: [
  ],
  exports: [
    Gongd3specificComponent,
    BarComponent,
    PieComponent,
    ScatterComponent
  ]
})
export class Gongd3specificModule { }
