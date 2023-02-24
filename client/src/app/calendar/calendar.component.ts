import { Component, ViewChild, OnInit, AfterViewInit, ViewEncapsulation } from '@angular/core';
import { SchedulerComponent, SchedulerResource, SchedulerViews } from 'smart-webcomponents-angular/scheduler';

@Component({
    selector: 'app-root',
    templateUrl: './calendar.component.html',
    styleUrls: ['./calendar.component.css'],
    entryComponents: [CalendarComponent],
    encapsulation: ViewEncapsulation.None
})

export class CalendarComponent implements AfterViewInit, OnInit {
    @ViewChild('scheduler', { read: SchedulerComponent, static: false }) scheduler!: SchedulerComponent;

    dataSource: SchedulerResource[] = this.getData();

    currentTimeIndicator: boolean = true;

    shadeUntilCurrentTime: boolean = true;

    view: string = 'day';

    views: string[] = ['day', 'week', 'month', 'timelineDay', 'timelineWeek', 'timelineMonth'];

    firstDayOfWeek: number = 1;

    ngOnInit(): void {
        // onInit code.
    }

    ngAfterViewInit(): void {
        // afterViewInit code.
    }

    getData() {
        const today = new Date(),
            todayDate = today.getDate(),
            currentYear = today.getFullYear(),
            currentMonth = today.getMonth(),
            data = [
                {
                    label: 'Example Event',
                    dateStart: new Date(currentYear, currentMonth, todayDate, 9, 0),
                    dateEnd: new Date(currentYear, currentMonth, todayDate, 10, 30),
                    backgroundColor: '#E67C73'
                },
            ]

        return data
    }
}
