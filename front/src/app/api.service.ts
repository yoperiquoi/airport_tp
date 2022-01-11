import { Injectable } from '@angular/core';
import {CaptorRangeData} from "./models/CaptorRangeData";
import {CaptorAverageData} from "./models/CaptorAverageData";

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor() { }

  getRangeData (airportCode: string, type: string, startDate: string, endDate: string) {
    return [
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 },
      {captorId: 1, airportId: 'aa', value: 12, nature: 'temperature', timestamp: 1239974 }
    ] as CaptorRangeData[]
  }

  getAverageData (airportCode: string, date: string) {
    return {
      windAverage: 14,
      temperatureAverage: 17,
      pressureAverage: 19
    } as CaptorAverageData
  }
}
