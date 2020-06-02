import { Component } from '@angular/core';
import { ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AppComponent {
  title = 'Bistro';

  ngOnInit(){
    console.log('hi')
    setTimeout(
      () => {
        this.title = 'Hey, its been updated'
        console.log('hello')
      },
      2000
      )
  }
}

