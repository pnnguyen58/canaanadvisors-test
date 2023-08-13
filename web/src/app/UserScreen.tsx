"use client";

import { MenuItem, PlaceOrderData, Restaurant } from "@/types/menu";
import React, { useEffect, useMemo, useState } from "react";
import axios from "axios";

const UserScreen = () => {
  const [menu, setMenu] = useState<Restaurant[]>([]);
  const [cart, setCart] = useState<Record<string, MenuItem[]> | null>(null);

  useEffect(() => {
    const getMenu = async () => {
      const result = await axios.get<Restaurant[]>("/api/menu");
      setMenu(result.data);
    };
    getMenu();
  }, []);

  const handleAddItem = (item: MenuItem, resId: string) => {
    setCart((prev) =>
      prev
        ? {
            ...prev,
            [resId]: [...(prev[resId] || []), item],
          }
        : { [resId]: [item] }
    );
  };

  const handlePlaceOrder = async (data: PlaceOrderData) => {
    await axios.post("/api/socket/place-order", data);
    setCart((prev) => {
      if (prev) {
        return Object.fromEntries(
          Object.entries(prev).filter(([key]) => key !== data.restaurantId)
        );
      }
      return null;
    });
  };

  return (
    <div>
      <div className="flex flex-col">
        {menu.map((res) => (
          <div key={res.id}>
            <p>{res.name}</p>
            {res.categories.map((cat) => (
              <div key={cat.id} className="ml-10">
                <p>{cat.name}</p>
                {cat.items.map((item) => (
                  <div key={item.id} className="ml-20 flex gap-2 items-center">
                    <p>{item.name}</p>
                    <button onClick={() => handleAddItem(item, res.id)}>
                      +
                    </button>
                  </div>
                ))}
              </div>
            ))}
          </div>
        ))}
      </div>
      {cart && Object.keys(cart).length > 0 && (
        <div>
          <h3 className="mb-6">Cart</h3>
          <div className="flex flex-col gap-2">
            {Object.entries(cart).map(([resId, items]) => (
              <div className="border border-black p-10" key={resId}>
                <div className="flex justify-between items-end">
                  <div>
                    <p className="mb-4">
                      {menu.find((res) => res.id === resId)?.name}
                    </p>
                    <ul>
                      {items.map((item) => (
                        <li key={item.id}>{item.name}</li>
                      ))}
                    </ul>
                  </div>
                  <div>
                    <button
                      onClick={() =>
                        handlePlaceOrder({
                          restaurantId: resId,
                          itemIds: items.map(({ id }) => id),
                        })
                      }
                    >
                      Place order
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default UserScreen;
