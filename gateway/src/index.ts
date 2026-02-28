const express = require("express");
const proxy = require("./routes/proxy");
const analytics = require("./routes/analytics");

const app = express();
app.use(express.json());

app.use("/api", proxy);
app.use("/analytics", analytics);

app.listen(3000, () => {
  console.log("Gateway running on port 3000");
});