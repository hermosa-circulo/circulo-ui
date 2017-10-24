import { ElementRef, Component, Input, Directive, ContentChild, ContentChildren } from '@angular/core';
import {BrowserDomAdapter} from 'angular2/platform/browser';
import * as THREE from 'three';

import * as objloader from 'three-obj-loader';
const OBJLoader = new objloader(THREE);

import * as jsonloader from 'three-json-loader';
const JSONLoader = new jsonloader(THREE);

@Directive({ selector: 'three-scene' })
 
export class SceneComponent {

  scenes: THREE.Scene[] = [];
  constructor(private element: ElementRef) {
    this.dom = new BrowserDomAdapter();
  }

  private dom: BrowserDomAdapter;
  ngAfterContentInit() {
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
      camera.position.set( -90,0,0);
      camera.lookAt(scene.position);
      scene.userData.camera = camera;
      var X,Y,Z;
      for (X = 0; X < 5; X++) {
      for (Y = 0; Y < 5; Y++) {
      for (Z = 0; Z < 5; Z++) {
        let loader = new THREE.OBJLoader();
        //console.log(typeof loader.load);
        loader.load(
           '/api/cube/'+X*10+'-'+Y*10+'-'+Z*10,
           function(object) {
             scene.add(object);
             //object.rotateX( 45 * Math.PI / 180 );
           }
        );
      } 
      }
      }
      //let geometry = new THREE.BoxGeometry( 10, 10, 10 );
      //let material = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
      //let cube = new THREE.Mesh( geometry, material );

      scene.add(light); 
      this.scenes.push( scene );
    }
  }
}
