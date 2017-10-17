import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { SceneComponent } from './scene.component';
import { RendererComponent } from './renderer.component';

import { ClarityModule } from 'clarity-angular';
import { OrbitControlsComponent } from './controls/orbit.component';

@NgModule({
  declarations: [
    SceneComponent,
    RendererComponent,
    OrbitControlsComponent,
  ],
  imports: [
    BrowserModule,
    ClarityModule.forRoot(),
  ],
  exports: [
    SceneComponent,
    RendererComponent,
    OrbitControlsComponent,
  ]
})
export class ThreeModule { }
