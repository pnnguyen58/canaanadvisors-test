export type Restaurant = {
  id: string;
  name: string;
  categories: MenuCategory[];
};

export type MenuCategory = {
  id: string;
  name: string;
  items: MenuItem[];
};

export type MenuItem = {
  id: string;
  name: string;
};

export type PlaceOrderData = {
  restaurantId: string;
  itemIds: string[];
};
