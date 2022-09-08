import { IResource } from "./resource";

interface IAggregate {
  week: number;
  min: number;
  max: number;
  avg: number;
  median: number;
  resources: IResource[];
}

interface IAggregateWithValues extends IAggregate {
  values: number[];
}

export interface IAggregationResult {
  area_provinsi: string;
  aggregates: IAggregate[];
}

export interface IAggregationResultWithValues {
  area_provinsi: string;
  aggregates: IAggregateWithValues[];
}
