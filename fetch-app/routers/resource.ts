// ---------------------------------------------- modules import
import express from "express";

import { RouterExtender } from "../app";
import * as ROUTES from "../constants/routes";
import handleReadAll from "../handlers/resources/read_all";
import handleReadPriceAggregation from "../handlers/resources/read_price_aggregation";
import verifyAuth from "../middlewares/auth";

// ---------------------------------------------- the routes
class ResourceRouterExtender implements RouterExtender {
  extend(router: express.Router) {
    router.route(`${ROUTES.RESOURCES}/`).all(verifyAuth).get(handleReadAll);

    router
      .route(`${ROUTES.RESOURCES}/aggregates/price`)
      .all(verifyAuth)
      .get(handleReadPriceAggregation);
  }
}

export default ResourceRouterExtender;
