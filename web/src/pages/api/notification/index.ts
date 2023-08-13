import { Restaurant } from "@/types/menu";
import { NextApiRequest, NextApiResponse } from "next";

/**
 * Api to handle get history notification from server
 */
const getNotification = async (req: NextApiRequest, res: NextApiResponse) => {
  if (req.method === "GET") {
    res.status(200).json({});
  }
};

export default getNotification;
