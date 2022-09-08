export const median = (numbers: number[]) => {
  const sorted = numbers.sort((a, b) => (a > b ? 1 : -1));

  const mid = Math.floor(sorted.length / 2);

  if (sorted.length % 2) return sorted[mid];

  return (sorted[mid - 1] + sorted[mid]) / 2;
};
