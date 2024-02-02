import axios from "axios";
import { SERVER_URL } from "$env/static/private";
import { redirect } from "@sveltejs/kit";

export async function makeRequest(endpoint, method, data = {}, headers = {}) {
    let url = `${SERVER_URL}/sever/${endpoint}/`;
  const config = { method, headers };
  if (method !== "GET") {
    const formData = new FormData();
    for (const [key, value] of Object.entries(data)) {
      if (key === "file") {
        formData.append(key, value, value.name);
        config.headers["Content-Type"] = "multipart/form-data";
      } else {
        formData.append(key, value);
      }
    }
    
    config.data = formData;
  }
  // else {
  //   const params = new URLSearchParams({ api_key: api_token, ...data });
  //   url += `?${params.toString()}`;
  // }

  try {
    const response = await axios(url, config);
    return response;
  } catch (error) {
    console.error("Error occurred during request: ", error);
    redirect(301,'/error')
    return
  }
}