import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validator, Validators} from "@angular/forms";
import {ApiService} from "../api.service";
import {CaptorRangeData} from "../models/CaptorRangeData";
import {CaptorAverageData} from "../models/CaptorAverageData";
import {Observable} from "rxjs";

@Component({
  selector: 'app-form-api',
  templateUrl: './form-api.component.html',
  styleUrls: ['./form-api.component.scss']
})
export class FormApiComponent implements OnInit {
  myForm: FormGroup;
  captorAverageData: CaptorAverageData | null;
  captorRangeData: CaptorRangeData[];

  constructor(
    private fb: FormBuilder,
    private apiService: ApiService
  ) { }

  ngOnInit(): void {
    this.myForm = this.fb.group({
      airportCode: ['', Validators.required],
      type: 'average',
      date: new Date(),
      startDate: new Date(),
      endDate: new Date(),
      nature: 'temperature'
    });
    this.apiService.rangeData.subscribe((value => this.captorRangeData = value));
  }

  /**
   * Send the form data to the API using apiService
   */
  async onSubmit(): Promise<void> {
    if (this.myForm.valid) {
      if (this.type === 'range') {
        this.captorAverageData = null;
        this.apiService.getRangeData(
          this.airportCode,
          this.nature,
          Math.round(this.startDate.getTime()/1000),
          Math.round(this.endDate.getTime()/1000)
        );
      } else {
        this.captorAverageData = null;
        this.apiService.getAverageData(
          this.airportCode,
          Math.round(this.date.getTime()/1000)
        );
      }
    }
  }

  get type() {
    return this.myForm.get('type')?.value;
  }

  get nature() {
    return this.myForm.get('nature')?.value;
  }

  get airportCode() {
    return this.myForm.get('airportCode')?.value;
  }

  get date(): Date {
    return this.myForm.get('date')?.value;
  }

  get startDate(): Date {
    return this.myForm.get('startDate')?.value;
  }

  get endDate(): Date {
    return this.myForm.get('endDate')?.value;
  }
}
