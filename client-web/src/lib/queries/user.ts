import { useQuery } from "@tanstack/react-query";
import { getUser } from "../apis/getUser";
import { getQueryClient } from "../../app/providers";

/** Retrieves app user details from DB */
export const useGetUser = () =>
  useQuery({
    queryKey: ["user"],
    queryFn: getUser,
  });

export const invalidateUser = () =>
  getQueryClient().invalidateQueries({
    queryKey: ["user"],
  });
