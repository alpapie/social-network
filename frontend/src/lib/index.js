// place files you want to import through the `$lib` alias in this folder.
import { log } from "console";
import { writeFileSync } from "fs";


export const  generateRandom = () =>
  Math.random().toString(36).substring(2, 15) +
  Math.random().toString(23).substring(2, 5);

  export const saveImage = async (image)=>{
    console.log("here is the image" , image)
    if (image.name == 'undefined' || image.name == '' ) return ""
    let extension = image.name.split(".").pop()
    let path = `./src/lib/images/${generateRandom()}.${extension}`
    writeFileSync(path , Buffer.from(await image.arrayBuffer()))
    return path
  }