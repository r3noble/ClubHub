import { Injectable } from '@angular/core';
import { SchedulerEvent } from 'smart-webcomponents-angular';

@Injectable()

export class DataService {

  constructor() { }

  GetData(): SchedulerEvent[] {

    const today = new Date();

    const year = today.getFullYear();
    const month = today.getMonth();
    const date = today.getDate();

    return [
      {
        label: 'Example Event',
        dateStart: new Date(year, month, date, 9, 0),
        dateEnd: new Date(year, month, date + 1, 10, 30),
        backgroundColor: '#39BF00',
        allDay: true
      }
    ]
  }
}
