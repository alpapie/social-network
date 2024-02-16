
class ServerStorage {
  constructor() {
    this.data = {};
  }

  // Méthode pour obtenir une valeur par clé
  get(key) {
    return this.data[key] || '';
  }

  // Méthode pour enregistrer une valeur par clé
  set(key, value) {
    this.data[key] = value;
  }

  // Méthode pour mettre à jour une valeur par clé
  update(key, value) {
    this.data[key] = value;
  }

  // Méthode pour supprimer une valeur par clé
  delete(key) {
    delete this.data[key];
  }
}

// Créez une instance de ServerStorage
export const localStorageObj = new ServerStorage();

export function DB(type = "", key = "", data = null) {
  try{
    let storedData;
    if (type === "get") {
      storedData = JSON.parse(localStorageObj.getItem(key) || "false");
      return data ? storedData[data] : storedData;
    }
  
    if (type === "set") {
      localStorageObj.set(key, JSON.stringify(data));
      return;
    }
  
    if (type === "update") {
      data = typeof data === "string" ? JSON.parse(data) : data;
  
      storedData = JSON.parse(localStorageObj.get(key) || "{}");
      const updatedData = { ...storedData, ...data };
      localStorageObj.set(key, JSON.stringify(updatedData));
      return;
    }
  
    if (type === "remove") {
      storedData = JSON.parse(localStorageObj.getItem(key) || "{}");
      delete storedData[data];
      localStorageObj.delete(key);
      return;
    }
  }catch(err){
    console.error("localsorage",err)
  }
}

  export function GetCookies(cookies){
    // console.log("alpaie",cookies);
    return "sessionId="+ cookies?.get("sessionId")
  }