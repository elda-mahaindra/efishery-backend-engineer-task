// ---------------------------------------------- modules import
import axios, { AxiosResponse } from "axios";
import dayjs from "dayjs";
import weekOfYear from "dayjs/plugin/weekOfYear";
import { Request, Response } from "express";

import { ErrorEnum } from "../../enums/error";
import { RoleEnum } from "../../enums/role";
import { average } from "../../helpers/average";
import { median } from "../../helpers/median";
import { IRequestExtended } from "../../middlewares/auth";
import {
  IAggregationResult,
  IAggregationResultWithValues,
} from "../../models/aggregation_result";
import { ICurrencyRate } from "../../models/currency_rate";
import { IPayload } from "../../models/payload";
import { IResource, IResourceData } from "../../models/resource";

// ---------------------------------------------- route handler
const handleReadSizeAggregation = async (req: Request, res: Response) => {
  dayjs.extend(weekOfYear);

  try {
    const { payload } = req as IRequestExtended;

    // ---------------------------------------------- authorization by role
    const isAuthorized = ((payload: IPayload) =>
      payload.role === RoleEnum.ADMIN)(payload);

    if (!isAuthorized) throw new Error(ErrorEnum.UNAUTHORIZED_ACCESS);

    // ---------------------------------------------- fetching
    const resourceUrl =
      "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list";

    const resourceData = (
      await axios.get<IResourceData[], AxiosResponse<IResourceData[]>>(
        resourceUrl
      )
    ).data;

    const currencyRateUrl =
      "https://api.exchangerate.host/latest?base=USD&symbols=IDR";

    const currencyRate = (
      await axios.get<ICurrencyRate, AxiosResponse<ICurrencyRate>>(
        currencyRateUrl
      )
    ).data;

    const resources: IResource[] = resourceData.map((r) => {
      const { rates } = currencyRate;

      return {
        ...r,
        price_in_usd: r.price ? parseInt(r.price) * rates["IDR"] : null,
      };
    });

    const sorted = resources.sort((a, b) =>
      a.tgl_parsed && b.tgl_parsed && a.tgl_parsed > b.tgl_parsed ? 1 : -1
    );

    const aggregationResults: IAggregationResult[] = sorted
      .reduce((results, resource): IAggregationResultWithValues[] => {
        const { area_provinsi, size, tgl_parsed } = resource;

        if (!area_provinsi || !tgl_parsed || !size) return results;

        if (!results.length) {
          return [
            {
              area_provinsi,
              aggregates: [
                {
                  week: dayjs(tgl_parsed).week(),
                  min: parseInt(size),
                  max: parseInt(size),
                  avg: parseInt(size),
                  median: parseInt(size),
                  values: [parseInt(size)],
                  resources: [resource],
                },
              ],
            },
          ];
        }

        const foundByAreaProvinsi = results.find(
          (r) => r.area_provinsi === area_provinsi
        );

        if (foundByAreaProvinsi) {
          return results.map((result) => {
            if (result.area_provinsi === foundByAreaProvinsi.area_provinsi) {
              const foundByWeek = foundByAreaProvinsi.aggregates.find(
                (a) => a.week === dayjs(tgl_parsed).week()
              );

              if (foundByWeek) {
                return {
                  ...result,
                  aggregates: result.aggregates.map((a) => {
                    if (a.week === foundByWeek.week) {
                      return {
                        ...a,
                        min: a.min > parseInt(size) ? parseInt(size) : a.min,
                        max: a.min > parseInt(size) ? a.min : parseInt(size),
                        avg: average([...a.values, parseInt(size)]),
                        median: median([...a.values, parseInt(size)]),
                        values: [...a.values, parseInt(size)],
                        resources: [...a.resources, resource],
                      };
                    }

                    return a;
                  }),
                };
              }

              return {
                ...result,
                aggregates: [
                  ...result.aggregates,
                  {
                    week: dayjs(tgl_parsed).week(),
                    min: parseInt(size),
                    max: parseInt(size),
                    avg: parseInt(size),
                    median: parseInt(size),
                    values: [parseInt(size)],
                    resources: [resource],
                  },
                ],
              };
            }

            return result;
          });
        } else {
          return [
            ...results,
            {
              area_provinsi,
              aggregates: [
                {
                  week: dayjs(tgl_parsed).week(),
                  min: parseInt(size),
                  max: parseInt(size),
                  avg: parseInt(size),
                  median: parseInt(size),
                  values: [parseInt(size)],
                  resources: [resource],
                },
              ],
            },
          ];
        }

        return [];
      }, [] as IAggregationResultWithValues[])
      .map((result) => ({
        ...result,
        aggregates: result.aggregates.map((aggregate) => {
          const { week, min, max, avg, median, resources } = aggregate;

          return { week, min, max, avg, median, resources };
        }),
      }));

    // ---------------------------------------------- response
    const response = {
      aggregation_results: aggregationResults,
    };

    res.status(200).json(response);
  } catch (error: any) {
    if (
      error instanceof Error &&
      error.message === ErrorEnum.UNAUTHORIZED_ACCESS
    ) {
      const response = {
        message: "you don't have authorization to access this endpoint.",
        error,
      };

      res.status(403).json(response);
    } else {
      console.log("ERROR 500: ", error);

      const response = {
        message: "something went wrong, please try again later.",
        error,
      };

      res.status(500).json(response);
    }
  }
};

export default handleReadSizeAggregation;
