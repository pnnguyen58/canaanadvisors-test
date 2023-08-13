import { dummyMenu } from "@/dummy";
import { PlaceOrderData } from "@/types/menu";
import { NextApiResponseServerIO } from "@/types/next";
import { NextApiRequest } from "next";

const chat = async (req: NextApiRequest, res: NextApiResponseServerIO) => {
  if (req.method === "POST") {
    // get message
    const order = req.body as PlaceOrderData;

    const restaurantName = dummyMenu.find(
      ({ id }) => order.restaurantId === id
    )?.name;
    const message = `A new order of ${restaurantName} has just been placed`;

    // dispatch to channel "message"
    res?.socket?.server?.io?.emit("place-order", {
      message,
      createdAt: new Date().toJSON(),
    });

    // return message
    res.status(201).json(order);
  }
};

export default chat;
