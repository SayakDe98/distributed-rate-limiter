import axios from "axios";

export async function checkRateLimit(userId: string) {
  const res = await axios.post("http://rate-limiter:8080/check", {
    user_id: userId,
  });

  return res.data;
}