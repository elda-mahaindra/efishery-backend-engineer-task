// ---------------------------------------------- modules import
import cors from "cors";
import express from "express";

import App from "./app";
import PingRouterExtender from "./routers/ping";

new App()
  .withMiddleware(express.urlencoded({ extended: true }))
  .withMiddleware(express.json())
  .withMiddleware(cors())
  .extendRouter(new PingRouterExtender())
  .start();
