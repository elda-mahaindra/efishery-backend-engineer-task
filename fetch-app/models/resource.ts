export interface IResourceData {
  uuid: string | null;
  komoditas: string | null;
  area_provinsi: string | null;
  area_kota: string | null;
  size: string | null;
  price: string | null;
  tgl_parsed: string | null;
  timestamp: string | null;
}

export interface IResource extends IResourceData {
  price_in_usd: number | null;
}
