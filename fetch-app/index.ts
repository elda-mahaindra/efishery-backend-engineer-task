// ---------------------------------------------- modules import
import cors from "cors";
import express from "express";

import App from "./app";

new App()
  .withMiddleware(express.urlencoded({ extended: true }))
  .withMiddleware(express.json())
  .withMiddleware(cors())
  .start();
