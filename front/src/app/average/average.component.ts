import {Component, Input, OnInit} from '@angular/core';
import {CaptorAverageData} from "../models/CaptorAverageData";

@Component({
  selector: 'app-average',
  templateUrl: './average.component.html',
  styleUrls: ['./average.component.scss']
})
export class AverageComponent implements OnInit {
  @Input() averageData: CaptorAverageData;

  constructor() { }

  ngOnInit(): void {
  }

}
