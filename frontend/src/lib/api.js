import axios from "axios";

import { GetCookies } from "../db";

export async function makeRequest(endpoint, method, data = {}, headers = {},cookies) {
  try {
    
<<<<<<< HEAD
    let url = `http://localhost:8080/server/${endpoint}`;
    if (cookies){
      headers.Cookie=GetCookies(cookies)
    }
    
    const config = { withCredentials: true ,method, headers };
    if (method !== "GET") {
=======
  let url = `http://localhost:8080/server/${endpoint}`;
  if (cookies){
    headers.Cookie=GetCookies(cookies)
  }
  const config = { withCredentials: true ,method, headers };
    if (method !== "get") {
>>>>>>> mirak
      config.data = data;
    }
    const response = await axios(url,config);
    
    return response;
  } catch (err) {
<<<<<<< HEAD
    console.log(err);
=======
>>>>>>> mirak
    return err?.response
  }
}