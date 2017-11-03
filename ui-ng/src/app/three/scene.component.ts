import { ElementRef, Component, Input, Directive, ContentChild, ContentChildren } from '@angular/core';
import {BrowserDomAdapter} from 'angular2/platform/browser';
import * as THREE from 'three';

import * as objloader from 'three-obj-loader';
const OBJLoader = new objloader(THREE);
import { Http, Response ,Headers, RequestOptions} from '@angular/http';

import * as jsonloader from 'three-json-loader';
const JSONLoader = new jsonloader(THREE);

@Directive({ selector: 'three-scene' })

 
export class SceneComponent {

  scenes: THREE.Scene[] = [];
  constructor(private element: ElementRef,private http: Http) {
    this.dom = new BrowserDomAdapter();
  }

  private dom: BrowserDomAdapter;
  public resdata;
  ngAfterContentInit() {
    let query = new Array("","","")
    for (var i = 0; i < 3; i++) {   
      let scene = new THREE.Scene();

      let el = this.dom.createElement("div");
      el.id = "list-item" + i;
      el.className = "list-item";
      let el2 = this.dom.createElement("div");
      el2.className = "scene";
      el.appendChild(el2)
      this.element.nativeElement.appendChild(el);
      scene.userData.element = el.querySelector( ".scene" );
      console.log(scene.userData.element);

      let light = new THREE.DirectionalLight( 0xffffff, 0.5 );
      light.position.set(0,250,0);
      let camera = new THREE.PerspectiveCamera(75,1,0.1,10000);
      camera.position.set( 100,0,0);
      camera.lookAt(scene.position);
      scene.userData.camera = camera;
      
      let send_data = { Point: i  };
      let trans_data = JSON.stringify(send_data);
      let gene
      var url = '/api/picgene'
      var request = new XMLHttpRequest();
      request.open('POST', url, false);
      request.setRequestHeader("Content-type", "application/json");
      request.send(trans_data);
      if (request.status == 200 ) {
        gene = request.responseText;
      } else {
        gene = "0000000000000001001001000100000000000000000000000000000000000000"
      } 
  
      let bit = gene.split('');
      let arealen = 10;
      for (var x = 0; x < arealen; x++){
        for (var y = 0; y < arealen; y++){
          for (var z = 0; z < arealen; z++){
            if ( bit[(x*arealen*arealen)+(y*arealen)+z] == "1" ) {
              query[i] = query[i] + x +"~"+ y +"~"+ z +"/" 
            }
          }
        }
      }
      query[i] = query[i].slice( 0, -1 ) ;

      let loader = new THREE.OBJLoader();
      loader.load(
         '/api/cube/'+query[i],
         function(object) {
           scene.add(object);
         }
      );
      scene.add(light); 
      this.scenes.push( scene );
    }
  }
}
