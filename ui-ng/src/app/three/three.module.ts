import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { SceneComponent } from './scene.component';
import { RendererComponent } from './renderer.component';
import { ThreeIntroComponent } from './threeintro.component';

import { ClarityModule } from 'clarity-angular';
import { OrbitControlsComponent } from './controls/orbit.component';
import { ThreejsComponent } from './threejs.component';

@NgModule({
  declarations: [
    SceneComponent,
    RendererComponent,
    OrbitControlsComponent,
    ThreeIntroComponent,
    ThreejsComponent,
  ],
  imports: [
    BrowserModule,
    ClarityModule.forRoot(),
  ],
  exports: [
    SceneComponent,
    RendererComponent,
    OrbitControlsComponent,
    ThreeIntroComponent,
    ThreejsComponent,
  ]
})
export class ThreeModule { }
