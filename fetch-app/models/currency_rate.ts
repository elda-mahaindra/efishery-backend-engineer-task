export interface ICurrencyRate {
  base: string;
  rates: { [key: string]: number };
}
