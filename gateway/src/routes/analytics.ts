// src/routes/analytics.ts

// Type-only import for TypeScript
import type { Request, Response } from "express";

// Runtime import using CommonJS
const { Router } = require("express");
const { redisClient } = require("../config/redis");

const router = Router();

router.get("/:userId", async (req: Request, res: Response) => {
  const { userId } = req.params;

  const total = await redisClient.get(`analytics:${userId}:total`);
  const blocked = await redisClient.get(`analytics:${userId}:blocked`);

  res.json({
    total: total || 0,
    blocked: blocked || 0,
  });
});

module.exports = router;