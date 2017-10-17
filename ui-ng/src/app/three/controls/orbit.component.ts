import { Directive, Input } from '@angular/core';

import * as THREE from 'three';
import * as controls from 'three-orbit-controls';
let OrbitControls = controls(THREE);

@Directive({ selector: 'three-orbit-controls-component' })
export class OrbitControlsComponent {

  @Input() enabled: boolean = true;

  setupControls(camera, renderer) {
    OrbitControls = new OrbitControls(camera, renderer.domElement);
    OrbitControls.enabled = this.enabled;
  }

  updateControls(scene, camera) {
    OrbitControls.update();
  }
}
