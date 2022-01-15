import { Injectable } from '@angular/core';
import {CaptorRangeData} from "./models/CaptorRangeData";
import {CaptorAverageData} from "./models/CaptorAverageData";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {BehaviorSubject, Observable} from "rxjs";
import {ResponseCaptorRangeData} from "./models/ResponseCaptorRangeData";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private readonly URL_API = 'http://localhost:8080/';
  private readonly OPTIONS = {
    headers: new HttpHeaders({
      "Access-Control-Allow-Headers": "*",
      'Access-Control-Allow-Origin': '*'
    })
  };
  private _rangeData: BehaviorSubject<CaptorRangeData[]> = new BehaviorSubject<CaptorRangeData[]>([]);
  public readonly rangeData: Observable<CaptorRangeData[]> = this._rangeData.asObservable();
  private _averageData: BehaviorSubject<CaptorRangeData[]> = new BehaviorSubject<CaptorRangeData[]>([]);
  public readonly averageData: Observable<CaptorRangeData[]> = this._rangeData.asObservable();

  constructor(private http: HttpClient) {

  }

  getRangeData (airportCode: string, type: string, startDate: number, endDate: number) {
    this.http.get<string>(this.URL_API + 'GetMesureFromTypeInRange/'+airportCode+'/'+type+'/'+startDate+'/'+endDate+'', this.OPTIONS)
    .subscribe((response) => {
      const data = JSON.parse(response)[0] as ResponseCaptorRangeData;
      const result = [];
      console.log(data);
      for (const datum of data.DataPoints) {
        result.push({
          airportId: data.Labels.airport_id,
          nature: type,
          timestamp: datum.Timestamp*1000,
          value: datum.Value
        } as CaptorRangeData)
      }
      this._rangeData.next(result);
    })
  }

  getAverageData (airportCode: string, date: number) {
    this.http.get<string>(this.URL_API + 'AverageForDay/'+airportCode,
      this.OPTIONS)
      .subscribe((response) => {
        const data = JSON.parse(response)[0];
        console.log(data);
      })

  }
}
