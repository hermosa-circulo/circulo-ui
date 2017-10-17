import { Component } from '@angular/core';
import 'clarity-icons';
import 'clarity-icons/shapes/all-shapes';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: [
              './app.component.css',
              '../../node_modules/clarity-ui/clarity-ui.min.css'
             ],
})
export class AppComponent {
  public sampleImagePath = "../images/main-icon.png"
  title = 'circle';
}
