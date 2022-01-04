import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, Validator, Validators} from "@angular/forms";
import {ApiService} from "../api.service";
import {CaptorRangeData} from "../models/CaptorRangeData";
import {CaptorAverageData} from "../models/CaptorAverageData";


@Component({
  selector: 'app-form-api',
  templateUrl: './form-api.component.html',
  styleUrls: ['./form-api.component.scss']
})
export class FormApiComponent implements OnInit {

  myForm: FormGroup;
  captorAverageData: CaptorAverageData | null;
  captorRangeData: CaptorRangeData[] | null;

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
  }

  onSubmit(): void {
    if (this.myForm.valid) {
      if (this.type === 'range') {
        this.captorAverageData = null;
        this.captorRangeData = this.apiService.getRangeData(
          this.airportCode,
          this.nature,
          this.startDate,
          this.endDate
        );
      } else {
        this.captorAverageData = null;
        this.captorAverageData = this.apiService.getAverageData(
          this.airportCode,
          this.date
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

  get date() {
    return this.myForm.get('date')?.value;
  }

  get startDate() {
    return this.myForm.get('startDate')?.value;
  }

  get endDate() {
    return this.myForm.get('endDate')?.value;
  }
}
