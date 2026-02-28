import express from "express";
import proxy from "./routes/proxy";
import analytics from "./routes/analytics";

const app = express();
app.use(express.json());

app.use("/api", proxy);
app.use("/analytics", analytics);

app.listen(3000, () => {
  console.log("Gateway running on port 3000");
});