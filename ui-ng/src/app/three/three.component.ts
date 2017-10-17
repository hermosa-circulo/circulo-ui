import { Input, Component, HostListener } from '@angular/core';

@Component({
  selector: 'three',
  templateUrl: './three.component.html',
  styleUrls: [
              '../../../node_modules/clarity-ui/clarity-ui.min.css'
  ],
})
export class ThreeComponent {

  @Input() ngModel: any;
  @Input() image: any;

  @Input() height: number;
  @Input() width: number;

  ngOnInit() {
    this.resetWidthHeight();
  }

  ngOnChanges(changes) {
    if(changes.ngModel && changes.ngModel.currentValue) {
      console.log('changes', changes);
    }
  }

  @HostListener('window:resize')
  @HostListener('window:vrdisplaypresentchange')
  resetWidthHeight() {
    this.height = window.innerHeight;
    this.width = window.innerWidth;
    console.log('window resize', this.height, this.width);
  }

}
