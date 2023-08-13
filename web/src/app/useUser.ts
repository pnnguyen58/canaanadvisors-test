import { User } from "@/types/user";
import { useEffect, useState } from "react";
import { getCookie } from "cookies-next";

const useUser = () => {
  const [isLoggedUser, serIsLoggedUser] = useState(false);

  useEffect(() => {
    serIsLoggedUser(!!getCookie("auth"));
  }, []);

  const [state, setState] = useState<User | null>(() => {
    try {
      const value = window?.localStorage.getItem("user");
      return value ? JSON.parse(value) : null;
    } catch (error) {
      console.log(error);
    }
  });

  const setValue = (value: User) => {
    try {
      const valueToStore = value instanceof Function ? value(state) : value;
      window?.localStorage.setItem("user", JSON.stringify(valueToStore));
      serIsLoggedUser(true);
      setState(value);
    } catch (error) {
      console.log(error);
    }
  };

  return { user: state, setUser: setValue, isLoggedUser: isLoggedUser };
};

export default useUser;
