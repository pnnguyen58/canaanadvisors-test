import { dummyMenu, dummyUser } from "@/dummy";
import { Restaurant } from "@/types/menu";
import { NextApiRequest, NextApiResponse } from "next";

const getMenu = async (req: NextApiRequest, res: NextApiResponse) => {
  if (req.method === "GET") {
    res.status(200).json(dummyMenu);
  }
};

export default getMenu;
