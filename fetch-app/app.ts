// ---------------------------------------------- modules import
import dotenv from "dotenv";
import express from "express";

export interface RouterExtender {
  extend(router: express.Router): void;
}

class App {
  private routeExtenders: RouterExtender[];
  private requestHandlerMiddlewares: express.RequestHandler[];
  public app: express.Application;

  constructor() {
    this.app = express();
    this.routeExtenders = [];
    this.requestHandlerMiddlewares = [];
  }

  private implementMiddlewares() {
    this.requestHandlerMiddlewares.forEach((rhm) => this.app.use(rhm));
  }

  private routes(): void {
    const router = express.Router();

    this.routeExtenders.forEach((re) => re.extend(router));

    this.app.get(
      "/route-test",
      (req: express.Request, res: express.Response) => {
        res.status(200).json({ message: "PONG" });
      }
    );

    this.app.use("/", router);

    this.app.use(
      (
        req: express.Request,
        res: express.Response,
        next: express.NextFunction
      ) => {
        const error = new Error("Request not recognized.");
        const response = { message: error.message };

        res.status(404).json(response);
      }
    );

    this.app.use(
      (
        err: any,
        req: express.Request,
        res: express.Response,
        next: express.NextFunction
      ) => {
        const response = { message: err.message, err };

        res.status(500).json(response);
      }
    );
  }

  public extendRouter(routerExtender: RouterExtender): this {
    this.routeExtenders.push(routerExtender);

    return this;
  }

  public withMiddleware(middleware: express.RequestHandler): this {
    this.requestHandlerMiddlewares.push(middleware);

    return this;
  }

  public start() {
    dotenv.config();

    const port = process.env.PORT;

    if (!port) throw new Error("Failed to read environments variables.");

    this.implementMiddlewares();
    this.routes();

    this.app.listen(port, () => {
      console.log(`server started at http://localhost:${port}`);
    });
  }
}

export default App;
