import { Directive, ElementRef, Component, Renderer } from '@angular/core';
import { Http, Response ,Headers, RequestOptions} from '@angular/http';

@Component({
  templateUrl: './threeintro.component.html',
  styleUrls: [
              '../../../node_modules/clarity-ui/clarity-ui.min.css'
  ],
})


@Directive({ selector: '[appHighlight]' })

export class ThreeIntroComponent {
    constructor(private elementRef:ElementRef, private renderer: Renderer) {
    }
    creategenesmessage;
    alertcreategenesmessage;
    CreateGenes() {
        var url = '/api/createpop'
        var request = new XMLHttpRequest();
        request.open('POST', url, false);
        request.send();
        if (request.status == 200 ) {
            this.creategenesmessage = request.responseText;
        } else {
            this.alertcreategenesmessage = "Canot Get Response";
        } 
        //let tmp = document.createElement('div');
        //let el = this.elementRef.nativeElement.querySelector('.createmessage');
        //tmp.innerHTML = '';
        //el.appendChild(tmp);
        //this.creategenesmessage = "tue";
    }
}
