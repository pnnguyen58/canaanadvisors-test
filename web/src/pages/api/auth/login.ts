import { dummyToken, dummyUser } from "@/dummy";
import { LoginRequest, User } from "@/types/user";
import setCookie from "@/utils/setCookie";
import { NextApiRequest, NextApiResponse } from "next";

const login = async (req: NextApiRequest, res: NextApiResponse) => {
  if (req.method === "POST") {
    const { username, password } = req.body as LoginRequest;
    const foundUser = dummyUser.find(
      (user) => user.username === username && user.password === password
    );
    if (!foundUser) return res.status(400).json({});
    const { password: _, ...userWithoutPassword } = foundUser;

    setCookie(res, "auth", dummyToken, {
      // httpOnly: true,
      maxAge: 3600,
      path: "/",
      sameSite: "strict",
      secure: process.env.NODE_ENV === "production",
    });

    return res.status(200).json(userWithoutPassword);
  }
};

export default login;
