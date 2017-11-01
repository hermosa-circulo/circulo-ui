import { Directive, ElementRef, Input, ContentChild, ViewChild } from '@angular/core';
import {BrowserDomAdapter} from 'angular2/platform/browser';
import * as THREE from 'three';

import { SceneComponent } from './scene.component';
import { OrbitControlsComponent } from './controls/orbit.component';

@Directive({ selector: 'three-renderer' })
export class RendererComponent {


  @ContentChild(SceneComponent) sceneComp: SceneComponent;
  @ContentChild(OrbitControlsComponent) orbitComponent: OrbitControlsComponent;

  renderers: THREE.WebGLRenderer[] = [];
  angle: number = 0;
  get scenes() {
    return this.sceneComp.scenes;
  }

  constructor(private element: ElementRef) {
    this.dom = new BrowserDomAdapter();
  }

  ngOnChanges(changes) {
  }
 
  private dom: BrowserDomAdapter;

  ngAfterContentInit() {
    let element = this.element;
    let animate = this.animate;
    let renderers = this.renderers;
    this.scenes.forEach( function( scene ) {
      let renderer = new THREE.WebGLRenderer({ antialias: true });
      renderer.setSize(200, 200);
      renderer.setClearColor( 0xffffff, 1 );
      renderer.setPixelRatio(Math.floor(window.devicePixelRatio));
      let el = scene.userData.element;
      el.appendChild(renderer.domElement)
      element.nativeElement.appendChild(el);
      scene.children[0].rotation.y = Date.now() * 0.001;
      let camera = scene.userData.camera;
      camera.lookAt(scene.position);
      renderer.render(scene, camera);
      renderers.push( renderer );
    });
    this.animate();
  }
  
  animate() {
    for(let i = 0; i < this.renderers.length; i++) {
      let scene = this.scenes[i];
      let camera = scene.userData.camera;
      //torus.position.set(0,0,500);
      //scene.children[0].rotation.y = Date.now() * 0.001 * i;
      camera.position.x = 10 * Math.cos( this.angle );
      camera.position.y = 10 * Math.sin( this.angle );
      camera.position.z = 10 * Math.sin( this.angle );
      this.angle += 0.002;
      camera.lookAt(scene.position);
      this.renderers[i].render( scene, camera );
    }
    requestAnimationFrame(() => this.animate());
  }


}
