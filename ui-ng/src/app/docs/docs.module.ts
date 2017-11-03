import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { RouterModule, Routes } from '@angular/router';
import { ClarityModule } from 'clarity-angular';
 
import { MarkdownModule } from 'angular2-markdown';
import { IntroComponent } from './intro.component';
import { DatabaseComponent } from './database.component';

@NgModule({
  declarations: [
    IntroComponent,
    DatabaseComponent,
  ],
  imports: [
    BrowserModule,
    MarkdownModule.forRoot(),
    ClarityModule.forRoot(),
  ],
  providers: [],
  exports: [
    MarkdownModule,
  ]
})

export class DocsModule { }
