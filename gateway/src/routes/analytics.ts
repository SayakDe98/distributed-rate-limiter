import { Router } from "express";
import { redisClient } from "../config/redis";

const router = Router();

router.get("/:userId", async (req, res) => {
  const { userId } = req.params;

  const total = await redisClient.get(`analytics:${userId}:total`);
  const blocked = await redisClient.get(`analytics:${userId}:blocked`);

  res.json({
    total: total || 0,
    blocked: blocked || 0,
  });
});

export default router;