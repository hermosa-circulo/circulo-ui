import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { ClarityModule } from 'clarity-angular';
 
import { MarkdownModule } from 'angular2-markdown';

@NgModule({
  declarations: [
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
