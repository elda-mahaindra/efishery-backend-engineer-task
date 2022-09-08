// ---------------------------------------------- modules import
import express from "express";

import { RouterExtender } from "../app";
import * as ROUTES from "../constants/routes";
import handlePing from "../handlers/ping/ping";

// ---------------------------------------------- the routes
class PingRouterExtender implements RouterExtender {
  extend(router: express.Router) {
    router.route(`${ROUTES.PING}/`).get(handlePing);
  }
}

export default PingRouterExtender;
