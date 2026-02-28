import type { Request, Response } from "express";

const { Router } = require("express");
const { redisClient } = require("../config/redis");

const router = Router();

// Use the type annotations inline
router.get("/:userId", async (req: Request, res: Response) => {
  const { userId } = req.params;

  // Redis values may be null, so convert to numbers
  const totalStr = await redisClient.get(`analytics:${userId}:total`);
  const blockedStr = await redisClient.get(`analytics:${userId}:blocked`);

  const total = totalStr ? Number(totalStr) : 0;
  const blocked = blockedStr ? Number(blockedStr) : 0;

  res.json({ total, blocked });
});

module.exports = router;