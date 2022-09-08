// ---------------------------------------------- modules import
import dayjs from "dayjs";
import dotenv from "dotenv";
import { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";

import { ErrorEnum } from "../enums/error";
import { IPayload } from "../models/payload";

export interface IRequestExtended extends Request {
  rawBody: any;
  payload: IPayload;
}

// ---------------------------------------------- jwt verification
const extractToken = (req: Request): string =>
  (req.headers["authorization"] as string).replace("Bearer ", "");

const verifyAuth = (req: Request, res: Response, next: NextFunction) => {
  if (
    // req.hasOwnProperty("headers") &&
    req.headers.hasOwnProperty("authorization")
  ) {
    dotenv.config();

    try {
      const secret = process.env.SECRET;

      if (!secret) throw new Error(ErrorEnum.UNABLE_TO_READ_ENV);

      const decoded = jwt.verify(extractToken(req), secret) as IPayload;
      const expiredAt = dayjs(decoded.expired_at);

      // because jwt signed in golang didn't expired in node, wee need to manually verify the expiration time
      if (!dayjs().isBefore(expiredAt))
        throw new Error(ErrorEnum.TOKEN_EXPIRED);

      Object.assign(req, { payload: { ...decoded } });

      next();
    } catch (error: any) {
      if (error instanceof Error && error.message === ErrorEnum.TOKEN_EXPIRED) {
        res.status(401).json({ message: "Your token has expired." });
      } else if (error.name === "TokenExpiredError") {
        res.status(401).json({ message: "Your token has expired.", error });
      } else {
        res
          .status(403)
          .json({ message: "error while verifying token.", error });
      }
    }
  } else {
    res.status(401).json({ message: "no token found." });
  }
};

export default verifyAuth;
