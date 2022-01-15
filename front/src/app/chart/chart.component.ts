import {Component, Input} from '@angular/core';
import {CaptorRangeData} from "../models/CaptorRangeData";
import * as moment from "moment";
import {ChartDataset} from "chart.js";


@Component({
  selector: 'app-chart',
  templateUrl: './chart.component.html',
  styleUrls: ['./chart.component.scss']
})
export class ChartComponent {
  @Input() dataset: CaptorRangeData[];

  constructor() { }

  /**
   * Returns the labels of the graph
   * @return {string[]} The labels
   */
  get labels(): string[]  {
    const labels = [];
    for (const datum of this.dataset) {
      labels.push(moment(datum.timestamp).format('DD/MM HH:mm:ss'));
    }
    return labels;
  }

  /**
   * Return the values of the graph
   * @return {number[]} The values
   */
  get values(): number[] {

    const numbers = [];
    for (const datasetElement of this.dataset) {
      numbers.push(datasetElement.value);
    }
    return numbers;
  }

  /**
   * Return the datasets to populate the graph
   * @return {{data: number[], label: string}[]}
   */
  get datasets(): ChartDataset[] {
    return [{
      data: this.values,
      label: this.dataset[0].nature
    }]
  }
}
