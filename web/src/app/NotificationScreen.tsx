import { NotificationItem, NotificationList } from "@/types/notification";
import dayjs from "dayjs";
import React, { useEffect, useState } from "react";
import { io } from "socket.io-client";

const NotificationScreen = () => {
  const [notificationList, setNotificationList] = useState<NotificationList>(
    []
  );
  useEffect(() => {
    const socket = io({
      path: "/api/socket/io",
      addTrailingSlash: false,
    });

    socket.on("connect", () => {
      console.log("Connect to websocket");
    });

    socket.on("place-order", (item: NotificationItem) => {
      setNotificationList((prev) => [...prev, item]);
    });

    return () => {
      socket?.disconnect();
    };
  }, []);

  return (
    <div className="mt-4">
      <h3 className="mb-6">Notification</h3>
      <div className="flex flex-col gap-4">
        {notificationList.map((noti) => (
          <div
            key={noti.message}
            className="border border-black p-10 flex justify-between"
          >
            <p>{noti.message}</p>
            <p>{dayjs(noti.createdAt).format("DD/MM/YYYY")}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default NotificationScreen;
