import React, { FormEvent, FC } from "react";
import axios from "axios";
import { User } from "@/types/user";

type Props = {
  onSuccess?: (data: User) => void;
};

const LoginForm: FC<Props> = ({ onSuccess }) => {
  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const data = Object.fromEntries(new FormData(e.target as HTMLFormElement));
    try {
      const result = await axios.post<User>("/api/auth/login", data);
      if (onSuccess) onSuccess(result.data);
    } catch (error) {
      alert("Login fail!");
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="flex flex-col gap-4 max-w-[528px] mx-auto pt-20"
    >
      <label>
        <p className="mb-2 ml-1 font-bold text-xs text-slate-700">Username</p>
        <input
          name="username"
          className="focus:shadow-soft-primary-outline text-sm leading-5.6 ease-soft block w-full appearance-none rounded-lg border border-solid border-gray-300 bg-white bg-clip-padding px-3 py-2 font-normal text-gray-700 transition-all focus:border-gray-700 focus:outline-none focus:transition-shadow"
        />
      </label>
      <label>
        <p className="mb-2 ml-1 font-bold text-xs text-slate-700">Password</p>
        <input
          type="password"
          name="password"
          className="focus:shadow-soft-primary-outline text-sm leading-5.6 ease-soft block w-full appearance-none rounded-lg border border-solid border-gray-300 bg-white bg-clip-padding px-3 py-2 font-normal text-gray-700 transition-all focus:border-gray-700 focus:outline-none focus:transition-shadow"
        />
      </label>
      <button
        type="submit"
        className="inline-block w-full px-6 py-3 mt-2 mb-0 font-bold text-center text-white uppercase align-middle transition-all bg-transparent border-0 rounded-lg cursor-pointer shadow-soft-md bg-x-25 bg-150 leading-pro text-xs ease-soft-in tracking-tight-soft bg-gradient-to-tl from-blue-600 to-cyan-400 hover:scale-102 hover:shadow-soft-xs active:opacity-85"
      >
        Login
      </button>
    </form>
  );
};

export default LoginForm;
