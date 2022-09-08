// ---------------------------------------------- modules import
import { Request, Response } from "express";

// ---------------------------------------------- route handler
const handlePing = async (req: Request, res: Response) => {
  // ---------------------------------------------- response
  const response = {
    message: "Pong",
  };

  res.status(200).json(response);
};

export default handlePing;
