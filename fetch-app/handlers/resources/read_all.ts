// ---------------------------------------------- modules import
import axios, { AxiosResponse } from "axios";
import { Request, Response } from "express";

import { ErrorEnum } from "../../enums/error";
import { RoleEnum } from "../../enums/role";
import { IRequestExtended } from "../../middlewares/auth";
import { ICurrencyRate } from "../../models/currency_rate";
import { IPayload } from "../../models/payload";
import { IResource, IResourceData } from "../../models/resource";

// ---------------------------------------------- route handler
const handleReadAll = async (req: Request, res: Response) => {
  try {
    const { payload } = req as IRequestExtended;

    // ---------------------------------------------- authorization by role
    const isAuthorized = ((payload: IPayload) =>
      payload.role === RoleEnum.SUPER ||
      payload.role === RoleEnum.ADMIN ||
      payload.role === RoleEnum.USER)(payload);

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

    // ---------------------------------------------- response
    const response = {
      resources,
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

export default handleReadAll;
