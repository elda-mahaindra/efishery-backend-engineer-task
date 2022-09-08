export const average = (numbers: number[]) =>
  numbers.reduce((sum, n) => sum + n) / numbers.length;
