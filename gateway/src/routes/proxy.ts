import { Router } from "express";
import { checkRateLimit } from "../services/rateLimiterClient";

const router = Router();

router.post("/", async (req, res) => {
  const { user_id } = req.body;

  const result = await checkRateLimit(user_id);

  if (!result.allowed) {
    return res.status(429).json({ message: "Rate limit exceeded" });
  }

  res.json({ message: "Request allowed" });
});

export default router;