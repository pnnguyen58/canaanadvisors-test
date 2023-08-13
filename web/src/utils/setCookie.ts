import { NextApiResponse } from "next";
import { serialize, CookieSerializeOptions } from "cookie";

const setCookie = (
  res: NextApiResponse,
  name: string,
  value: string,
  options: CookieSerializeOptions = {}
): void => {
  const stringValue =
    typeof value === "object" ? `j:${JSON.stringify(value)}` : String(value);
  res.setHeader("Set-Cookie", serialize(name, stringValue, options));
};

export default setCookie;
