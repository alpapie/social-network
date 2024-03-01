import axios from "axios";
import { GetCookies } from "../routes/db";
import { PUBLIC_BACKEND_URL } from '$env/static/public';

export async function makeRequest(endpoint, method, data = {}, headers = {}, cookies) {
  try {
    
    let url = `${PUBLIC_BACKEND_URL}/${endpoint}`;
    if (cookies){
      headers.Cookie=GetCookies(cookies)
    }
    
    const config = { withCredentials: true ,method, headers };
    if (method !== "GET") {
      config.data = data;
    }
    const response = await axios(url,config);
    
    return response;
  } catch (err) {
    return err?.response
  }
}