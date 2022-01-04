import {Component, Input, OnInit} from '@angular/core';
import {} from 'ng2-charts';
import {CaptorRangeData} from "../models/CaptorRangeData";
import * as moment from "moment";
import {ChartDataset} from "chart.js";


@Component({
  selector: 'app-chart',
  templateUrl: './chart.component.html',
  styleUrls: ['./chart.component.scss']
})
export class ChartComponent implements OnInit {
  @Input() dataset: CaptorRangeData[];
  @Input() startDate: string
  @Input() endDate: string

  constructor() { }

  ngOnInit(): void {

  }

  get labels(): any  {
    const days = [];
    const momentStartDate = moment(this.startDate);
    const momentEndDate = moment(this.endDate);
    const diff = momentEndDate.diff(momentStartDate, 'days');
    days.push(momentStartDate.format('DD/MM'));
    console.log(diff)
    for (let i = 0; i < diff; i++) {
      const moment = momentStartDate.add(1, 'days')
      days.push(moment.format('DD/MM'));    }
    return days;
  }

  get values(): any {
    const numbers = [];
    for (const datasetElement of this.dataset) {
      numbers.push(datasetElement.value);
    }
    return numbers;
  }

  get datasets(): ChartDataset[] {
    return [{
      data: this.values,
      label: 'Temperatures'
    }]
  }
}
