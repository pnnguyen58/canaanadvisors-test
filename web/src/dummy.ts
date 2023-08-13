import { Restaurant } from "./types/menu";
import { User } from "./types/user";

type UserBackend = User & {
  password: string;
};

export const dummyUser: Array<UserBackend> = [
  {
    id: "user",
    username: "user",
    password: "user",
    role: "user",
  },
  {
    id: "owner",
    username: "owner",
    password: "owner",
    role: "owner",
  },
  {
    id: "driver",
    username: "driver",
    password: "driver",
    role: "driver",
  },
];

export const dummyToken = "super-secret-token-in-the-world";

export const dummyMenu: Restaurant[] = [
  {
    id: "res1",
    name: "Restaurant 1",
    categories: [
      {
        id: "cat1",
        name: "Category 1",
        items: [
          {
            id: "item0",
            name: "Item 0",
          },
          {
            id: "item1",
            name: "Item 1",
          },
        ],
      },
    ],
  },
  {
    id: "res2",
    name: "Restaurant 2",
    categories: [
      {
        id: "cat2",
        name: "Category 1",
        items: [
          {
            id: "item2",
            name: "Item 2",
          },
          {
            id: "item3",
            name: "Item 3",
          },
          {
            id: "item4",
            name: "Item 4",
          },
        ],
      },
    ],
  },
];
