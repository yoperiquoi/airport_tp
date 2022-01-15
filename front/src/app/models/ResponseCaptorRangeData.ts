export interface ResponseCaptorRangeData {
  Name: string
  Labels: {
    airport_id: string,
    sensor_type: string,
    unit: string
  }
  DataPoints: {
    Timestamp: number,
    Value: number
  }[]
}
