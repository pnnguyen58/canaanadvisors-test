"use client";

import { useEffect, useState } from "react";
import LoginForm from "./LoginForm";
import { User } from "@/types/user";
import useUser from "./useUser";
import { useRouter } from "next/navigation";
import UserScreen from "./UserScreen";
import NotificationScreen from "./NotificationScreen";

export default function Home() {
  const [loading, setLoading] = useState(true);
  const { user, setUser, isLoggedUser } = useUser();
  const router = useRouter();

  useEffect(() => {
    setLoading(false);
  }, []);

  const handleSuccessLogin = (user: User) => {
    setUser(user);
    router.refresh();
  };

  if (loading) return <div>Loading ...</div>;
  if (!isLoggedUser) return <LoginForm onSuccess={handleSuccessLogin} />;
  if (user?.role === "user") return <UserScreen />;
  return <NotificationScreen />;
}
