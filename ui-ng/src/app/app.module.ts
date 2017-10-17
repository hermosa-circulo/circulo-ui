import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { ClarityModule } from 'clarity-angular';

import { DocsComponent } from './docs/docs.component';
import { AboutComponent } from './about/about.component';
import { ThreeComponent } from './three/three.component';

import { ThreeModule } from './three/three.module';
import { DocsModule } from './docs/docs.module';

const appRoutes: Routes = [
  {
    path: 'three',
    component: ThreeComponent
  },
  {
    path: 'about',
    component: AboutComponent
  },
  {
    path: 'docs',
    component: DocsComponent
  },
];

@NgModule({
  declarations: [
    AppComponent,
    DocsComponent,
    AboutComponent,
    ThreeComponent,
  ],
  imports: [
    BrowserModule,
    ThreeModule,
    DocsModule,
    ClarityModule.forRoot(),
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true }
    ),
  ],
  providers: [],
  bootstrap: [AppComponent]
})

export class AppModule { }


