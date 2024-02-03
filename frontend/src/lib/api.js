import axios from "axios";
import { SERVER_URL } from "$env/static/private";
import { redirect } from "@sveltejs/kit";
export async function makeRequest(endpoint, method, data = {}, headers = {}) {
    let url = `http://localhost:8080/server/${endpoint}`;
    console.log(url)
  const config = { method, headers };
  if (method !== "GET") {

    config.data = data;
  }
  
  try {
    const response = await axios(url, config);
    return response;
  } catch (error) {
    return error
  }
}